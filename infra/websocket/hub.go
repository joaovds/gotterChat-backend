package websocket

import (
	"log"

	"github.com/google/uuid"
)

type Hub struct {
  ID         string
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub(id string) *Hub {
  if id != "" {
    id = uuid.New().String()
  }

	return &Hub{
    ID: id,
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}

		case client := <-h.Register:
			h.Clients[client] = true
			log.Println("Client registered")

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				log.Println("Client unregistered")
			}
		}
	}
}
