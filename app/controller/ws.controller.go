package controller

import (
	"github.com/Dubjay18/sanctum-server/app/domain/dao"
	"github.com/Dubjay18/sanctum-server/app/domain/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsController struct {
	Hub *dao.Hub
}

func NewWsController(hub *dao.Hub) *WsController {
	return &WsController{
		Hub: hub,
	}
}

func (ws *WsController) CreateRoom(c *gin.Context) {
	// Create a room
	var req dto.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.RoomId = uuid.New().String()
	ws.Hub.Rooms[req.RoomId] = &dao.Room{
		ID:      req.RoomId,
		Name:    req.RoomName,
		Clients: make(map[string]*dao.Client)}
	c.JSON(http.StatusCreated, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
