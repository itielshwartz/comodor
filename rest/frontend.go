package rest

import (
	"awesomeProject/resources"
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	client *redis.Client
	db     *sql.DB
}

func NewUserHandler(client *redis.Client, db *sql.DB) UserHandler {
	return UserHandler{client: client, db: db}
}

func (h *UserHandler) HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	releasesHandler := resources.NewReleasesHandler(h.client, h.db)
	var e error
	var resp interface{}
	cmd := mux.Vars(r)["command"]
	switch cmd {
	case "listpods":
		resp, e = releasesHandler.HandleListReleaseRequest(r.Context(), "orgname");
	default:
		logrus.Error("Bad command!!")
	}
	if e != nil {
		logrus.WithError(e).Error("Failed to list Releases")
		//TODO: better handle logic
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		//TODO: log error
		json.NewEncoder(w).Encode(resp)
	}
}
