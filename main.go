package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Tasks []string = []string{"hello, world"}

type ResponseType struct {
	Updated bool `json:"updated"`
}

type TaskType struct {
	Task string
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Route: ", c.Route().Path)
		fmt.Println("Method: ", c.Route().Method)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		fmt.Printf("Tasks: %v\n", Tasks)
		return c.JSON(&Tasks)
	})

	app.Post("/task/add", func(c *fiber.Ctx) error {
		payload := struct {
			Task string `json:"task"`
		}{}

		err := c.BodyParser(&payload)

		if err != nil {
			return c.SendStatus(400)
		}

		Tasks = append(Tasks, payload.Task)

		fmt.Println(payload.Task)
		return c.JSON(&ResponseType{Updated: true})
	})

	log.Fatal(app.Listen(":8000"))
}
