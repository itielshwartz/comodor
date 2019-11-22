package communication

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
)

type ResponseData interface {
	//Empty interface to see all requests
	ResponseMarker()
}

type AgentToServerResponse struct {
	UUID         string
	CreateAt     time.Time
	ResponseData json.RawMessage
	Err          string
}

func NewAgentToServerResponse(responseData ResponseData, uuid string, err string) *AgentToServerResponse {
	bytes, e := json.Marshal(responseData)
	if e != nil {
		logrus.WithError(e).Fatal("Failed to marshal response")
	}
	return &AgentToServerResponse{
		UUID:         uuid,
		CreateAt:     time.Now(),
		ResponseData: bytes,
		Err:          err,
	}
}
