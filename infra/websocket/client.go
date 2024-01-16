package websocket

import gws "github.com/gorilla/websocket"

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
