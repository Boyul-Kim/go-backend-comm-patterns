package main

import (
	"comm-design/server"
)

func main() {
	server := server.InitServer()
	server.Fiber.Listen(":3000")
}
