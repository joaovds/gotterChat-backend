package websocket

import (
	"github.com/google/uuid"
)

type Hub struct {
	ID         string
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

func NewHub(id string) *Hub {
	if id == "" {
		id = uuid.New().String()
	}

	return &Hub{
		ID:         id,
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run(hubs map[string]*Hub) {
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

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

			if len(h.Clients) == 0 {
				delete(hubs, client.Hub.ID)
			}
		}
	}
}
