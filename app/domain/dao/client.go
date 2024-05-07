package dao

import socketio "github.com/googollee/go-socket.io"

type Client struct {
	Conn     *socketio.Conn
	Message  chan *Message
	ID       string
	RoomId   int
	username string
}

type Message struct {
	Content  string `json:"content"`
	RoomId   int    `json:"room_id"`
	Username string `json:"username"`
}
