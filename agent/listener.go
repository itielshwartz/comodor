package main

import (
	"awesomeProject/helm"
	"awesomeProject/iproto"
	"awesomeProject/kube"
	"github.com/golang/protobuf/proto"
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
		var resp []byte
		_, rawMsg, err := c.conn.ReadMessage()
		if rawMsg == nil {
			// Happen when connection get closed
			//TODO: handle this differently
			break
		}
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("error: %v", err)
				c.sendError <- err
			}
		}
		message := &iproto.ServerToClientRequest{}
		err = proto.Unmarshal(rawMsg, message)
		if err != nil {
			c.sendError <- err

		}
		switch message.Request.(type) {
		case *iproto.ServerToClientRequest_ListPodsInNamespace:
			namespace := message.GetListPodsInNamespace()
			log.Info("Start getting pods")
			resp, err = kube.GetPods(c.kubeClient, namespace.Namespace)
			log.Info(string("Done getting pods"))
		case *iproto.ServerToClientRequest_ListCurrentHelmReleases:
			log.Info("Start getting releases")
			resp, err = helm.ListReleases(helm.GetClient())
			log.Info(string("Done getting releases"))
		default:
			print("WTF")
			log.Print(message)
		}
		if err != nil {
			log.Error(err)
			c.sendError <- err
		} else {
			c.send <- resp
		}
	}
}
