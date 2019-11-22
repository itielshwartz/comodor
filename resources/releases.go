package resources

import (
	"awesomeProject/common"
	"awesomeProject/communication"
	"awesomeProject/db"
	"awesomeProject/models"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"time"
)

var ListReleasesRequestCMD = "ListReleasesRequest"

const WaitTimeOut = time.Second * 5

type ListReleasesRequest struct {
}

func (l ListReleasesRequest) Cmd() string { return ListReleasesRequestCMD }

type ListReleasesResponse struct {
	Data []*models.Release
}

func (l *ListReleasesResponse) ResponseMarker() {}

type ReleasesHandler struct {
	client *redis.Client
	db     *sql.DB
}

func NewReleasesHandler(client *redis.Client, db *sql.DB) *ReleasesHandler {
	return &ReleasesHandler{client: client, db: db}
}

func (r *ReleasesHandler) HandleListReleaseRequest(ctx context.Context, orgName string) (interface{}, error) {
	request := communication.NewServerToAgentRequest(&ListReleasesRequest{})
	bytes, e := json.Marshal(request)
	if e != nil {
		return nil, e
	}
	r.client.Publish(common.AgentKey(orgName), bytes)
	redisResp, err := r.client.BLPop(WaitTimeOut, request.UUID).Result()
	if err != nil {
		//TODO: better handle logic
		return nil, err
	}
	var rsp communication.AgentToServerResponse
	err = json.Unmarshal([]byte(redisResp[1]), &rsp)
	if err != nil {
		//TODO: better handle logic
		return nil, err
	}
	var rspData ListReleasesResponse
	err = json.Unmarshal(rsp.ResponseData, &rspData)
	db.SaveReleases(r.db, ctx, rspData.Data)
	return rspData.Data, err
}
