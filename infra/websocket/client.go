package websocket

import (
	gws "github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Hub  *Hub
	Conn *gws.Conn
	Send chan []byte
}

func NewClient(hub *Hub, conn *gws.Conn) *Client {
	return &Client{
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
}

const (
	maxMessageSize = 512
)

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if gws.IsUnexpectedCloseError(err, gws.CloseGoingAway, gws.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.Hub.Broadcast <- message
	}
}
