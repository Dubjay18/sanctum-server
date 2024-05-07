package app

import (
	"github.com/Dubjay18/sanctum-server/app/controller"
	"github.com/Dubjay18/sanctum-server/app/domain/dao"
	"github.com/gin-gonic/gin"
)

type SocketServer struct {
}

func StartServer() {
	hub := dao.NewHub()
	wsHandler := controller.NewWsController(hub)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/create-room", wsHandler.CreateRoom)
	router.Run() // listen and serve on 0.0.0.0:8080
}
