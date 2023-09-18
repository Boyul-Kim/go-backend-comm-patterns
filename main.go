package main

import (
	"comm-design/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	initEnv()
	server := server.InitServer()
	server.Fiber.Listen(":3000")
}

func initEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
		return err
	}
	return nil
}
