package main

import (
	"awesomeProject/websckt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	waitTimeout = time.Minute * 10
)

func main() {
	client, _ := websckt.ConnectToRedis("localhost:6379")
	agentHandler := websckt.NewWebsocketHandler(client)
	frontendHandler := websckt.NewUserHandler(client)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/v1/api", frontendHandler.HandleUserRequest)
	http.HandleFunc("/v1/ws-agent", agentHandler.HandleAgentWebsocket)
	log.Info("blabl")
	http.ListenAndServe(":8888", nil)
	log.Info("Done")
}
