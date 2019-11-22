package main

import (
	"fmt"
	"os"

	"github.com/flant/kubedog/pkg/kube"
	"github.com/flant/kubedog/pkg/tracker"
	"github.com/flant/kubedog/pkg/tracker/pod"

)

func main() {
	_ = kube.Init(kube.InitOptions{})


	feed := pod.NewFeed()


	feed.OnStatusReport(func(p pod.PodStatus) error {
		fmt.Printf("Pod status: %#v\n", feed.GetStatus().stat)
		return nil
	})
	feed.OnFailed(func(reason string) error {
		return fmt.Errorf("pod failed: %s", reason)
	})

	err := feed.Track(
		"wishing-prawn-redis-master-0",
		"default",
		kube.Kubernetes,
		tracker.Options{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: po/mypod tracker failed: %s", err)
		os.Exit(1)
	}
}