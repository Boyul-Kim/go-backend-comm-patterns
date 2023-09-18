package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type Server struct {
	Fiber *fiber.App
	Jobs  map[string]interface{}
}

type CheckJob struct {
	Id string `json:"id" bson:"id"`
}

func InitServer() *Server {
	app := fiber.New()
	jobs := make(map[string]interface{})

	s := Server{
		Fiber: app,
		Jobs:  jobs,
	}

	s.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("testing")
	})

	s.Fiber.Post("/poll/short", s.shortPoll)
	s.Fiber.Get("/poll/short/check", s.checkShortPoll)
	s.Fiber.Get("/sse", adaptor.HTTPHandler(handler(eventHandler)))
	s.Fiber.Get("/pubSub/pub", s.RabbitMQPub)
	s.Fiber.Get("/pubSub/sub", s.RabbitMQSub)
	return &s
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}
