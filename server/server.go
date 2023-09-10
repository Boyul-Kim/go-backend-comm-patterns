package server

import (
	"comm-design/util"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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

	return &s
}

func (s *Server) shortPoll(c *fiber.Ctx) error {
	id := util.RandomIdGenerator()
	go s.updateJob(id, 0)
	return c.SendString(id)
}

func (s *Server) checkShortPoll(c *fiber.Ctx) error {
	newCheckJob := new(CheckJob)
	err := c.BodyParser(newCheckJob)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result := s.checkJob(newCheckJob.Id)

	return c.Status(200).JSON(fiber.Map{"result": result})
}

func (s *Server) updateJob(id string, percentage int64) {
	fmt.Println("Jobs", s.Jobs)
	s.Jobs[id] = percentage
	fmt.Println("ID", id, "percentage", s.Jobs[id])
	if s.Jobs[id].(int64) == 100 {
		fmt.Println("ID", id, "COMPLETE")
		return
	}
	timer1 := time.NewTimer(3 * time.Second)
	<-timer1.C
	s.updateJob(id, percentage+10)
}

func (s *Server) checkJob(id string) int64 {
	return s.Jobs[id].(int64)
}
