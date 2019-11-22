package main

import (
	"awesomeProject/communication"
	"awesomeProject/helm"
	"awesomeProject/resources"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// The application runs listen in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) listen() {
	defer func() {
		c.conn.Close()
	}()
	for {
		log.Info(string("Waiting for msg"))
		var err error
		var respData communication.ResponseData
		_, rawMsg, err := c.conn.ReadMessage()
		if rawMsg == nil {
			// Happen when connection get closed
			//TODO: handle this differently
			break
		}
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("error: %v", err)
			}
		}
		var message communication.ServerToAgentRequest
		err = json.Unmarshal(rawMsg, &message)
		if err != nil {
			log.WithError(err).Error("Failed to marshal request")
			continue
		}
		resp := communication.NewAgentToServerResponse(nil, message.UUID, "")
		switch message.Cmd {
		case resources.ListReleasesRequestCMD:
			log.Info("Start getting releases")
			respData, err = helm.ListReleases(c.HelmClient)
			log.Info(string("Done getting releases"))
		default:
			print("WTF")
			log.Print(message)
		}
		resp.ResponseData, err = json.Marshal(respData)
		if err != nil {
			log.WithError(err).Error("Failed to marshal respData")
		}
		bytes, err := json.Marshal(resp)
		if err != nil {
			log.WithError(err).Error("Failed to marshal response")
		}
		if err != nil {
			resp.Err = err.Error()

		}

		c.send <- bytes
	}
}
