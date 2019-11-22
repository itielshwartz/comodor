package websckt

import (
	"awesomeProject/resources"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	client *redis.Client
}

func NewUserHandler(client *redis.Client) UserHandler {
	return UserHandler{client: client}
}

func (h *UserHandler) HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, e := resources.HandleListReleaseRequest(h.client, "orgname");
	if e != nil {
		logrus.WithError(e).Error("Failed to list Releases")
		//TODO: better handle logic
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		//TODO: log error
		json.NewEncoder(w).Encode(resp)
	}
}
