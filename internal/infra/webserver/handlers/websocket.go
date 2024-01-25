package handlers

import (
	"log"
	"net/http"

	gws "github.com/gorilla/websocket"
	"github.com/joaovds/chat/internal/infra/websocket"
	"github.com/joaovds/chat/pkg/validation"
)

type websocketHandler struct {
	hub *websocket.Hub
}

func NewWebsocketHandler() *websocketHandler {
	return &websocketHandler{}
}

var upgrader = gws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Hubs = make(map[string]*websocket.Hub)

func (h *websocketHandler) ServeWs(roomId string, w http.ResponseWriter, r *http.Request) {
	if !validation.IsValidUUID(roomId) {
		http.Error(w, "Invalid room id", http.StatusBadRequest)
		return
	}

	if hub, ok := Hubs[roomId]; !ok {
		h.hub = websocket.NewHub(roomId)
		Hubs[roomId] = h.hub
		go h.hub.Run(Hubs)
	} else {
		h.hub = hub
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := websocket.NewClient(h.hub, conn)
	client.Hub.Register <- client

	go client.MessageListener()
	go client.WriteMessage()
}
