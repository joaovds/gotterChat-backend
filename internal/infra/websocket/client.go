package websocket

import (
	"log"
	"time"

	"github.com/google/uuid"
	gws "github.com/gorilla/websocket"
)

type Message struct {
	ID     string `json:"id"`
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
	Data   string `json:"data"`
}

func NewMessage(roomId, userId string, data string) Message {
	return Message{
		ID:     uuid.New().String(),
		RoomID: roomId,
		UserID: userId,
		Data:   data,
	}
}

type Client struct {
	ID   string
	Hub  *Hub
	Conn *gws.Conn
	Send chan Message
}

func NewClient(hub *Hub, conn *gws.Conn) *Client {
	return &Client{
		ID:   uuid.New().String(),
		Hub:  hub,
		Conn: conn,
		Send: make(chan Message),
	}
}

const (
	maxMessageSize = 512
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
)

func (c *Client) MessageListener() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, messageData, err := c.Conn.ReadMessage()
		if err != nil {
			if gws.IsUnexpectedCloseError(err, gws.CloseGoingAway, gws.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message := NewMessage(c.Hub.ID, c.ID, string(messageData))

		c.Hub.Broadcast <- message
	}
}

func (c *Client) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(gws.CloseMessage, []byte{})
				return
			}

      if(message.UserID != c.ID) {
			err := c.Conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
			}
      }

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(gws.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
