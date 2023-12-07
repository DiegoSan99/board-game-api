// websocket_handlers.go
package websocket

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WebSocketHandler struct {
	// Add any dependencies or services here
}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{}
}

// SetupRoutes sets up the routes for WebSocket endpoints
func (wsh *WebSocketHandler) SetupRoutes(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Get("/init", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host"))
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

}
