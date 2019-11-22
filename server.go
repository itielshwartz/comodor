package main

import (
	"awesomeProject/db"
	"awesomeProject/rest"
	"awesomeProject/websckt"
	"github.com/gorilla/mux"
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
	dbClient, e := db.GetClient()
	if e != nil {
		log.Fatal("Failed to connect to DB")
	}
	frontendHandler := rest.NewUserHandler(client, dbClient)
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./public")))
	r.HandleFunc("/v1/api/{command}", frontendHandler.HandleUserRequest)
	r.HandleFunc("/v1/ws-agent", agentHandler.HandleAgentWebsocket)
	log.Info("Start server")

	http.Handle("/", r)
	http.ListenAndServe(":8888", nil)
	log.Info("Done")
}
