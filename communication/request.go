package communication

import (
	"awesomeProject/common"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
)

type WSRequest interface {
	//Empty interface to see all requests
	Cmd() string
}

type ServerToAgentRequest struct {
	UUID        string
	CreateAt    time.Time
	RequestData json.RawMessage
	Cmd         string
}

func NewServerToAgentRequest(requestData WSRequest) *ServerToAgentRequest {
	cmd := requestData.Cmd()
	uuid := common.GenRedisUUID(cmd)
	bytes, e := json.Marshal(requestData)
	if e != nil {
		logrus.WithError(e).Fatal("Failed to marshal request")
	}
	return &ServerToAgentRequest{
		UUID:        uuid,
		CreateAt:    time.Now(),
		RequestData: bytes,
		Cmd:         cmd,
	}
}
