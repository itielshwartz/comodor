package resources

import (
	v1 "k8s.io/api/core/v1"
	"time"
)

type Ipod struct {
	name      string
	status    string
	phase     string
	restart   int
	createdAt time.Time
}

func newIpod(pod v1.Pod) *Ipod {
	status := pod.Status
	containerStatus := status.ContainerStatuses[0]
	return &Ipod{
		name:      pod.Name,
		status:    status.Message,
		phase:     string(status.Phase),
		restart:   int(containerStatus.RestartCount),
		createdAt: pod.GetCreationTimestamp().Time,
	}
}

func CreateIpodsList(list v1.PodList) []*Ipod {
	l := len(list.Items)
	Ipods := make([]*Ipod, l)
	for i, item := range list.Items {
		Ipods[i] = newIpod(item)
	}
	return Ipods
}
