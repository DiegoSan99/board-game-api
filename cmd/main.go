package main

import (
	"github.com/DiegoSan99/board-game-api/src/interfaces/fiber"
	"github.com/DiegoSan99/board-game-api/src/interfaces/websocket"
)

func main() {

	server := fiber.NewFiberServer()

	wsh := websocket.NewWebSocketHandler()
	wsh.SetupRoutes(server.Server)

	server.Start("3030")
}
