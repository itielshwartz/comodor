package iproto

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func TimestampToTime(t *timestamp.Timestamp) *time.Time {
	if t != nil {
		goTime, err := ptypes.Timestamp(t)
		if err != nil {
			return nil
		}
		return &goTime
	}
	return nil
}
