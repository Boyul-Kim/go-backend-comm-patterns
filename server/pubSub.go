package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *Server) RabbitMQPub(c *fiber.Ctx) error {
	fmt.Println("testing rabbitmq pub")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ"))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("q", q)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Testing RabbitMQ"),
		},
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"result": "pub msg successful"})
}

func (s *Server) RabbitMQSub(c *fiber.Ctx) error {
	fmt.Println("testing rabbitmq sub")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ"))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	//keep alive to receive messages
	keepAlive := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Println("Message: ", string(msg.Body))
		}
	}()
	fmt.Println("waiting for messages")
	<-keepAlive
	return c.Status(200).JSON(fiber.Map{"result": "msgs received"})
}
