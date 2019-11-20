package helm

import (
	"helm.sh/helm/pkg/helm"
	"sync"
)

const (
	tillerHost = "127.0.0.1:44134"
)

var doOnce sync.Once
var client *helm.Client

func GetClient() *helm.Client {
	doOnce.Do(func() {
		client = helm.NewClient(helm.Host(tillerHost), helm.ConnectTimeout(3))
	})
	return client
}