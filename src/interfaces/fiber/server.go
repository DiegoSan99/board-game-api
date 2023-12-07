package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	Server *fiber.App
}

func NewFiberServer() *FiberServer {
	server := fiber.New() //check if sonic unmarshal is useful with websockets and add
	server.Use(cors.New())
	return &FiberServer{
		Server: server,
	}
}

func (f *FiberServer) Start(port string) error {
	return f.Server.Listen(fmt.Sprintf(":%v", port))
}
