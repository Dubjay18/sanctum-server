package app

import (
	"github.com/Dubjay18/sanctum-server/app/controller"
	"github.com/Dubjay18/sanctum-server/app/domain/dao"
	"github.com/Dubjay18/sanctum-server/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func StartServer() {
	hub := dao.NewHub()
	wsHandler := controller.NewWsController(hub)
	go hub.Run()
	config.ConnectDB()
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"

		},
		MaxAge: 12 * time.Hour,
	}))
	v1 := router.Group("/v1")
	{
		ws := v1.Group("/ws")
		{
			ws.POST("/create-room", wsHandler.CreateRoom)
		}
	}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
