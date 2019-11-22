package websckt

import (
	"awesomeProject/common"
	"awesomeProject/communication"
	"encoding/json"
	redis "github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// message sent to us by the javascript client
type message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}

type WebsocketHandler struct {
	client *redis.Client
}

func NewWebsocketHandler(client *redis.Client) WebsocketHandler {
	return WebsocketHandler{client: client}
}

// validateMessage so that we know it's valid JSON and contains a Handle and
// Text
func validateMessage(data []byte) (message, error) {
	var msg message

	if err := json.Unmarshal(data, &msg); err != nil {
		return msg, errors.Wrap(err, "Unmarshaling message")
	}

	if msg.Handle == "" && msg.Text == "" {
		return msg, errors.New("Message has no Handle or Text")
	}

	return msg, nil
}

// handleWebsocket connection.
func (h *WebsocketHandler) HandleAgentWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		log.WithField("err", err).Println(m)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	// We listen for event for  all org/ single agent
	orgId := r.Header.Get(common.OrgId)
	clusterId := r.Header.Get(common.ClusterId)
	l := log.WithField(common.OrgId, orgId).WithField(common.ClusterId, clusterId)
	l.Info("new agent subscribe ")
	pubsub := h.client.Subscribe(common.AgentKey(orgId), clusterId)

	_, err = pubsub.Receive() // wait for subscription to be created and ignore the message
	if err != nil {
		log.WithError(err).Error("Fail to create websckt connection")
	}
	//Wait for messages/commands from the user to the agent
	h.handleIncomingCommandsForAgent(pubsub, ws, l)
	// Wait for receiving data from the remote agent
	h.handleIncomingMsgFromAgent(ws, l)
	l.Info("agent disconnected ")

	_ = pubsub.Close()
	_ = ws.WriteMessage(websocket.CloseMessage, []byte{})
}

func (h *WebsocketHandler) handleIncomingCommandsForAgent(pubsub *redis.PubSub, ws *websocket.Conn, l *logrus.Entry) {
	//TODO: check what happen when the user disconnects
	go func() {
		for {
			l.Info("Starting to listen to commands")
			receiveMessage, e := pubsub.ReceiveMessage()
			if e != nil {
				l.WithError(e).Error("failed to receive messages")
			}
			l.Info("Got a command and going to send it along")
			e = ws.WriteMessage(2, []byte(receiveMessage.Payload))
			if e != nil {
				l.WithError(e).Error("failed to write message")
			}
		}
	}()
}

func (h *WebsocketHandler) handleIncomingMsgFromAgent(ws *websocket.Conn, entry *logrus.Entry) {
	for {
		mt, data, err := ws.ReadMessage()
		l := log.WithFields(logrus.Fields{"mt": mt, "data": data, "err": err})
		if err != nil {
			l.WithError(err)
			if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
				l.Info("Websocket closed!")
				break
			}
			l.Error("Error reading websckt message")
			break

		}
		switch mt {
		case websocket.BinaryMessage:
			var msg communication.AgentToServerResponse
			err := json.Unmarshal(data, &msg)
			if err != nil {
				l.WithFields(logrus.Fields{"msg": msg, "err": err}).Error("Invalid Message")
				break
			}
			_, err = h.client.LPush(msg.UUID, data).Result()
			if err != nil {
				l.WithFields(logrus.Fields{"msg": msg, "err": err}).Error("Redis push problem")
				break
			}
		default:
			l.Warning("Unknown Message!")
		}
	}
}
