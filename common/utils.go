package common

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func GenUUID() string {
	uid, e := uuid.NewUUID()
	if e != nil {
		log.WithError(e).Error("Failed to get UUID")
	}
	return uid.String()
}

func TypeToName(o interface{}) string {
	return fmt.Sprintf("%T", o)
}

func GenRedisUUID(typeName string) string {
	uid, e := uuid.NewUUID()
	if e != nil {
		log.WithError(e).Error("Failed to get UUID")
	}
	return typeName + "-" + uid.String()
}

func buildKey(componentType, name string) string {
	return name + "_" + componentType
}

func AgentKey(name string) string {
	return buildKey(Agent, name)
}

func UserKey(name string) string {
	return buildKey(User, name)
}
