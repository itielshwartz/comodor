//COPY: /Users/itiel/go/pkg/mod/k8s.io/helm@v2.16.1+incompatible/pkg/kube/client.go
package helm // import "k8s.io/helm/pkg/kube"

import (
	"context"
	goerrors "errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"

	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	batch "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
	cachetools "k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/validation"
	"k8s.io/kubernetes/pkg/apis/core"
)

// MissingGetHeader is added to Get's output when a resource is not found.
const MissingGetHeader = "==> MISSING\nKIND\t\tNAME\n"

// ErrNoObjectsVisited indicates that during a visit operation, no matching objects were found.
var ErrNoObjectsVisited = goerrors.New("no objects visited")

var metadataAccessor = meta.NewAccessor()

// Client represents a client capable of communicating with the Kubernetes API.
type Client struct {
	cmdutil.Factory
	Log func(string, ...interface{})
}

// New creates a new Client.
func New(getter genericclioptions.RESTClientGetter) *Client {
	if getter == nil {
		getter = genericclioptions.NewConfigFlags(true)
	}

	err := apiextv1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	return &Client{
		Factory: cmdutil.NewFactory(getter),
		Log:     nopLogger,
	}
}

var nopLogger = func(_ string, _ ...interface{}) {}

// ResourceActorFunc performs an action on a single resource.
type ResourceActorFunc func(*resource.Info) error

func (c *Client) validator() validation.Schema {
	schema, err := c.Validator(true)
	if err != nil {
		c.Log("warning: failed to load schema: %s", err)
	}
	return schema
}

// BuildUnstructured reads Kubernetes objects and returns unstructured infos.
func (c *Client) BuildUnstructured(namespace string, reader io.Reader) (Result, error) {
	var result Result

	result, err := c.NewBuilder().
		Unstructured().
		ContinueOnError().
		NamespaceParam(namespace).
		DefaultNamespace().
		Stream(reader, "").
		Flatten().
		Do().Infos()
	return result, scrubValidationError(err)
}


// Return the resource info as internal
func resourceInfoToObject(info *resource.Info, c *Client) runtime.Object {
	internalObj, err := asInternal(info)
	if err != nil {
		// If the problem is just that the resource is not registered, don't print any
		// error. This is normal for custom resources.
		if !runtime.IsNotRegisteredError(err) {
			c.Log("Warning: conversion to internal type failed: %v", err)
		}
		// Add the unstructured object in this situation. It will still get listed, just
		// with less information.
		return info.Object
	}

	return internalObj
}



// Get gets Kubernetes resources as pretty-printed string.
//
// Namespace will set the namespace.
func (c *Client) GetReleaseResources(namespace string, reader io.Reader) (map[string](map[string]runtime.Object), map[string][]v1.Pod, error) {
	// Since we don't know what order the objects come in, let's group them by the types and then sort them, so
	// that when we print them, they come out looking good (headers apply to subgroups, etc.).
	objs := make(map[string](map[string]runtime.Object))
	mux := &sync.Mutex{}

	infos, err := c.BuildUnstructured(namespace, reader)
	if err != nil {
		return nil, nil, err
	}

	var objPods = make(map[string][]v1.Pod)
	err = perform(infos, func(info *resource.Info) error {
		mux.Lock()
		defer mux.Unlock()
		c.Log("Doing get for %s: %q", info.Mapping.GroupVersionKind.Kind, info.Name)

		// Use APIVersion/Kind as grouping mechanism. I'm not sure if you can have multiple
		// versions per cluster, but this certainly won't hurt anything, so let's be safe.
		gvk := info.ResourceMapping().GroupVersionKind
		vk := gvk.Version + "/" + gvk.Kind

		// Initialize map. The main map groups resources based on version/kind
		// The second level is a simple 'Name' to 'Object', that will help sort
		// the individual resource later
		if objs[vk] == nil {
			objs[vk] = make(map[string]runtime.Object)
		}
		// Map between the resource name to the underlying info object
		objs[vk][info.Name] = resourceInfoToObject(info, c)

		//Get the relation pods
		objPods, err = c.getSelectRelationPod(info, objPods)
		if err != nil {
			c.Log("Warning: get the relation pod is failed, err:%s", err.Error())
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	//here, we will add the objPods to the objs
	for key, podItems := range objPods {
		for i := range podItems {
			pod := &core.Pod{}
			scheme.Scheme.Convert(&podItems[i], pod, nil)
			if objs[key+"(related)"] == nil {
				objs[key+"(related)"] = make(map[string]runtime.Object)
			}
			objs[key+"(related)"][pod.ObjectMeta.Name] = runtime.Object(pod)
		}
	}
	return objs,objPods, nil
}

func perform(infos Result, fn ResourceActorFunc) error {
	if len(infos) == 0 {
		return ErrNoObjectsVisited
	}

	errs := make(chan error)
	go batchPerform(infos, fn, errs)

	for range infos {
		err := <-errs
		if err != nil {
			return err
		}
	}
	return nil
}

func batchPerform(infos Result, fn ResourceActorFunc, errs chan<- error) {
	var kind string
	var wg sync.WaitGroup
	for _, info := range infos {
		currentKind := info.Object.GetObjectKind().GroupVersionKind().Kind
		if kind != currentKind {
			wg.Wait()
			kind = currentKind
		}
		wg.Add(1)
		go func(i *resource.Info) {
			errs <- fn(i)
			wg.Done()
		}(info)
	}
}

func getSelectorFromObject(obj runtime.Object) (map[string]string, bool) {
	switch typed := obj.(type) {

	case *v1.ReplicationController:
		return typed.Spec.Selector, true

	case *extv1beta1.ReplicaSet:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1.ReplicaSet:
		return typed.Spec.Selector.MatchLabels, true

	case *extv1beta1.Deployment:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1beta1.Deployment:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1beta2.Deployment:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1.Deployment:
		return typed.Spec.Selector.MatchLabels, true

	case *extv1beta1.DaemonSet:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1beta2.DaemonSet:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1.DaemonSet:
		return typed.Spec.Selector.MatchLabels, true

	case *batch.Job:
		return typed.Spec.Selector.MatchLabels, true

	case *appsv1beta1.StatefulSet:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1beta2.StatefulSet:
		return typed.Spec.Selector.MatchLabels, true
	case *appsv1.StatefulSet:
		return typed.Spec.Selector.MatchLabels, true

	default:
		return nil, false
	}
}

func (c *Client) watchUntilReady(timeout time.Duration, info *resource.Info) error {
	// Use a selector on the name of the resource. This should be unique for the
	// given version and kind
	selector, err := fields.ParseSelector(fmt.Sprintf("metadata.name=%s", info.Name))
	if err != nil {
		return err
	}
	lw := cachetools.NewListWatchFromClient(info.Client, info.Mapping.Resource.Resource, info.Namespace, selector)

	kind := info.Mapping.GroupVersionKind.Kind
	c.Log("Watching for changes to %s %s with timeout of %v", kind, info.Name, timeout)

	// What we watch for depends on the Kind.
	// - For a Job, we watch for completion.
	// - For all else, we watch until Ready.
	// In the future, we might want to add some special logic for types
	// like Ingress, Volume, etc.

	ctx, cancel := watchtools.ContextWithOptionalTimeout(context.Background(), timeout)
	defer cancel()
	_, err = watchtools.ListWatchUntil(ctx, lw, func(e watch.Event) (bool, error) {
		switch e.Type {
		case watch.Added, watch.Modified:
			// For things like a secret or a config map, this is the best indicator
			// we get. We care mostly about jobs, where what we want to see is
			// the status go into a good state. For other types, like ReplicaSet
			// we don't really do anything to support these as hooks.
			c.Log("Add/Modify event for %s: %v", info.Name, e.Type)
			if kind == "Job" {
				return c.waitForJob(e, info.Name)
			}
			return true, nil
		case watch.Deleted:
			c.Log("Deleted event for %s", info.Name)
			return true, nil
		case watch.Error:
			// Handle error and return with an error.
			c.Log("Error event for %s", info.Name)
			return true, fmt.Errorf("Failed to deploy %s", info.Name)
		default:
			return false, nil
		}
	})
	return err
}

// waitForJob is a helper that waits for a job to complete.
//
// This operates on an event returned from a watcher.
func (c *Client) waitForJob(e watch.Event, name string) (bool, error) {
	job := &batch.Job{}
	err := scheme.Scheme.Convert(e.Object, job, nil)
	if err != nil {
		return true, err
	}

	for _, c := range job.Status.Conditions {
		if c.Type == batch.JobComplete && c.Status == v1.ConditionTrue {
			return true, nil
		} else if c.Type == batch.JobFailed && c.Status == v1.ConditionTrue {
			return true, fmt.Errorf("Job failed: %s", c.Reason)
		}
	}

	c.Log("%s: Jobs active: %d, jobs failed: %d, jobs succeeded: %d", name, job.Status.Active, job.Status.Failed, job.Status.Succeeded)
	return false, nil
}

// scrubValidationError removes kubectl info from the message.
func scrubValidationError(err error) error {
	if err == nil {
		return nil
	}
	const stopValidateMessage = "if you choose to ignore these errors, turn validation off with --validate=false"

	if strings.Contains(err.Error(), stopValidateMessage) {
		return goerrors.New(strings.Replace(err.Error(), "; "+stopValidateMessage, "", -1))
	}
	return err
}

func (c *Client) watchPodUntilComplete(timeout time.Duration, info *resource.Info) error {
	lw := cachetools.NewListWatchFromClient(info.Client, info.Mapping.Resource.Resource, info.Namespace, fields.ParseSelectorOrDie(fmt.Sprintf("metadata.name=%s", info.Name)))

	c.Log("Watching pod %s for completion with timeout of %v", info.Name, timeout)
	ctx, cancel := watchtools.ContextWithOptionalTimeout(context.Background(), timeout)
	defer cancel()
	_, err := watchtools.ListWatchUntil(ctx, lw, func(e watch.Event) (bool, error) {
		return isPodComplete(e)
	})

	return err
}

// GetPodLogs takes pod name and namespace and returns the current logs (streaming is NOT enabled).
func (c *Client) GetPodLogs(name, ns string) (io.ReadCloser, error) {
	client, err := c.KubernetesClientSet()
	if err != nil {
		return nil, err
	}
	req := client.CoreV1().Pods(ns).GetLogs(name, &v1.PodLogOptions{})
	logReader, err := req.Stream()
	if err != nil {
		return nil, fmt.Errorf("error in opening log stream, got: %s", err)
	}
	return logReader, nil
}

func isPodComplete(event watch.Event) (bool, error) {
	o, ok := event.Object.(*v1.Pod)
	if !ok {
		return true, fmt.Errorf("expected a *v1.Pod, got %T", event.Object)
	}
	if event.Type == watch.Deleted {
		return false, fmt.Errorf("pod not found")
	}
	switch o.Status.Phase {
	case v1.PodFailed, v1.PodSucceeded:
		return true, nil
	}
	return false, nil
}

//get a kubernetes resources' relation pods
// kubernetes resource used select labels to relate pods
func (c *Client) getSelectRelationPod(info *resource.Info, objPods map[string][]v1.Pod) (map[string][]v1.Pod, error) {
	if info == nil {
		return objPods, nil
	}

	c.Log("get relation pod of object: %s/%s/%s", info.Namespace, info.Mapping.GroupVersionKind.Kind, info.Name)

	versioned := asVersionedOrUnstructured(info)
	selector, ok := getSelectorFromObject(versioned)
	if !ok {
		return objPods, nil
	}

	client, _ := c.KubernetesClientSet()

	pods, err := client.CoreV1().Pods(info.Namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(selector).AsSelector().String(),
	})
	if err != nil {
		return objPods, err
	}

	for _, pod := range pods.Items {
		vk := "v1/Pod"
		if !isFoundPod(objPods[vk], pod) {
			objPods[vk] = append(objPods[vk], pod)
		}
	}
	return objPods, nil
}

func isFoundPod(podItem []v1.Pod, pod v1.Pod) bool {
	for _, value := range podItem {
		if (value.Namespace == pod.Namespace) && (value.Name == pod.Name) {
			return true
		}
	}
	return false
}

func asVersionedOrUnstructured(info *resource.Info) runtime.Object {
	obj, _ := asVersioned(info)
	return obj
}

func asVersioned(info *resource.Info) (runtime.Object, error) {
	converter := runtime.ObjectConvertor(scheme.Scheme)
	groupVersioner := runtime.GroupVersioner(schema.GroupVersions(scheme.Scheme.PrioritizedVersionsAllGroups()))
	if info.Mapping != nil {
		groupVersioner = info.Mapping.GroupVersionKind.GroupVersion()
	}

	obj, err := converter.ConvertToVersion(info.Object, groupVersioner)
	if err != nil {
		return info.Object, err
	}
	return obj, nil
}

func asInternal(info *resource.Info) (runtime.Object, error) {
	groupVersioner := info.Mapping.GroupVersionKind.GroupKind().WithVersion(runtime.APIVersionInternal).GroupVersion()
	return scheme.Scheme.ConvertToVersion(info.Object, groupVersioner)
}
