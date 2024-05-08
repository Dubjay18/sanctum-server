package main

import (
	"github.com/Dubjay18/sanctum-server/app"
	"github.com/Dubjay18/sanctum-server/config"
)

func main() {
	// Start the server
	app.StartServer()
	config.InitLog()
}
