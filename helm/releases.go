package helm

import (
	"awesomeProject/iproto"
	"awesomeProject/models"
	"awesomeProject/resources"
	"github.com/volatiletech/null"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/release"
)

func ListReleases(client *helm.Client) (*resources.ListReleasesResponse, error) {
	rawReleases, err := client.ListReleases()
	if err != nil {
		return &resources.ListReleasesResponse{Data: nil}, err
	}
	l := len(rawReleases.Releases)
	releases := make([]*models.Release, l)
	for i, item := range rawReleases.Releases {
		rls := newRelease(item)
		releases[i] = rls
	}
	return &resources.ListReleasesResponse{Data: releases}, err
}

func newRelease(r *release.Release) *models.Release {
	return &models.Release{
		Name:          r.Name,
		Namespace:     r.Namespace,
		Cluster:       "fake",
		Status:        r.Info.Status.Code.String(),
		CreatedAt:     *iproto.TimestampToTime(r.Info.FirstDeployed),
		Revision:      null.IntFrom(int(r.Version)),
		Schemaversion: null.IntFrom(1),
	}
}
