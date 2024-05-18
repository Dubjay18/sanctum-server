package controller

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

const (
	// SocketConnection is the event name for a connection
	SocketConnection = "/"
	// SocketEventCreateRoom is the event name for creating a room
	SocketEventCreateRoom = "create_room"
	// SocketEventJoinRoom is the event name for joining a room
	SocketEventJoinRoom = "join_room"
	// SocketEventGetRooms is the event name for getting all rooms
	SocketEventGetRooms = "get_rooms"
	// SocketEventGetClients is the event name for getting all clients in a room
	SocketEventGetClients = "get_clients"
	// SocketEventSendMessage is the event name for sending a message
	SocketEventSendMessage = "send_message"
	// SocketEventReceiveMessage is the event name for receiving a message
	SocketEventReceiveMessage = "receive_message"
	// SocketEventUserJoined is the event name for a user joining a room
	SocketEventUserJoined = "user_joined"
	// SocketEventUserLeft is the event name for a user leaving a room
	SocketEventUserLeft = "user_left"
	// SocketEventError is the event name for an error
	SocketEventError = "error"
)

type SocketController struct {
}

func NewSocketController() *SocketController {
	return &SocketController{}
}

func (so *SocketController) HandleConnection(s socketio.Conn) error {
	s.SetContext("")
	fmt.Println("connected:", s.ID())
	return nil
}
