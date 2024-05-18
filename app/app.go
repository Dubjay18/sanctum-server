package app

import (
	"fmt"
	"github.com/Dubjay18/sanctum-server/app/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
	"time"
)

func StartServer() {
	//hub := dao.NewHub()
	//wsHandler := controller.NewWsController(hub)
	//go hub.Run()

	server := socketio.NewServer(nil)
	socketHandler := controller.NewSocketController()

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
	//router.Group("/v1")
	//router.Group("/ws")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.OnConnect(controller.SocketConnection, socketHandler.HandleConnection)
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		//server.Remove(s.ID())
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()
	//router.POST("/create-room", wsHandler.CreateRoom)
	//router.GET("/join-room/:id/:user_id/:username", wsHandler.JoinRoom)
	//router.GET("/get-rooms", wsHandler.GetRooms)
	//router.GET("/get-clients/:id", wsHandler.GetClients)
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.Run() // listen and serve on 0.0.0.0:8080
}
