package kube

import (
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var timeoutInSeconds = int64(2)

func GetPods(kubeClient *kubernetes.Clientset, namespace string) ([]byte, error) {
	log.Info("start kube things")

	pods, err := kubeClient.CoreV1().Pods(namespace).List(metav1.ListOptions{TimeoutSeconds: &timeoutInSeconds})
	log.Info("Done kube things")

	if err != nil {
		return nil, err
	}
	if pods == nil {
		return nil, nil
	}
	podsBinary, marshelError := pods.Marshal()
	return podsBinary, marshelError
}
