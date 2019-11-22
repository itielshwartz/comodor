package websckt

import (
	"github.com/go-redis/redis/v7"
)

const (
	// Channel name to use with websckt
	Channel = "chat"
)

func ConnectToRedis(addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

	// Output: PONG <nil>
}
