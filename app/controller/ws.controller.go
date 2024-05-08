package controller

import (
	"github.com/Dubjay18/sanctum-server/app/domain/dao"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type WsController struct {
	Hub *dao.Hub
}

func NewWsController(hub *dao.Hub) *WsController {
	return &WsController{
		Hub: hub,
	}
}

type CreateRoomRequest struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (wc *WsController) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wc.Hub.Rooms[strconv.Itoa(req.ID)] = &dao.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*dao.Client),
	}
	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wc *WsController) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID := c.Param("id")
	clientID := c.Param("user_id")
	username := c.Param("username")
	cl := &dao.Client{
		Conn:     conn,
		Message:  make(chan *dao.Message),
		ID:       clientID,
		RoomId:   roomID,
		Username: username,
	}
	m := &dao.Message{
		Content:  "joined",
		RoomId:   roomID,
		Username: username,
	}
	wc.Hub.Register <- cl
	wc.Hub.Broadcast <- m

	go cl.WriteMesage()
	cl.ReadMessage(wc.Hub)
}

type RoomResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (wc *WsController) GetRooms(c *gin.Context) {
	room := make([]*RoomResponse, 0)
	for _, r := range wc.Hub.Rooms {
		room = append(room, &RoomResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	c.JSON(http.StatusOK, room)
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (wc *WsController) GetClients(c *gin.Context) {
	clients := make([]*ClientResponse, 0)
	if _, ok := wc.Hub.Rooms[c.Param("id")]; !ok {
		clients = append(clients, &ClientResponse{
			ID:       "0",
			Username: "No clients",
		})
	} else {
		for _, cl := range wc.Hub.Rooms[c.Param("id")].Clients {
			clients = append(clients, &ClientResponse{
				ID:       cl.ID,
				Username: cl.Username,
			})
		}
	}
	c.JSON(http.StatusOK, clients)

}
