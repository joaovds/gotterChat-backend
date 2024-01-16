package handlers

import (
	"log"
	"net/http"

	gws "github.com/gorilla/websocket"
	"github.com/joaovds/chat/infra/websocket"
)

type websocketHandler struct {
	hub *websocket.Hub
}

func NewWebsocketHandler() *websocketHandler {
	hub := websocket.NewHub()
	go hub.Run()

	return &websocketHandler{
		hub: hub,
	}
}

var upgrader = gws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *websocketHandler) ServeWs(w http.ResponseWriter, r *http.Request) {
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
