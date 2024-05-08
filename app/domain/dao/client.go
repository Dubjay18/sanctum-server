package dao

import (
	"github.com/Dubjay18/sanctum-server/config"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string
	RoomId   string
	Username string
}

type Message struct {
	Content  string `json:"content"`
	RoomId   string `json:"room_id"`
	Username string `json:"username"`
}

func (c *Client) WriteMesage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Message:
			if !ok {
				return
			}
			c.Conn.WriteJSON(msg)
		}
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				config.LogError(err, "websocket error")
			}
			break
		}
		m := &Message{
			Content:  string(msg),
			RoomId:   c.RoomId,
			Username: c.Username,
		}
		hub.Broadcast <- m
	}
}
