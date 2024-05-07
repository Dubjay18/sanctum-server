package controller

import (
	"github.com/Dubjay18/sanctum-server/app/domain/dao"
	"github.com/gin-gonic/gin"
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

type CreateRoomRequest struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (wc *WsController) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	wc.Hub.Rooms[req.ID] = &dao.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*dao.Client),
	}
	c.JSON(http.StatusOK, req)
}
