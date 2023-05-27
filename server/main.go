package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Tasks []string = []string{"hello, world"}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Route: ", c.Route().Path)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		fmt.Printf("Tasks: %v\n", Tasks)
		return c.JSON(&Tasks)
	})

	log.Fatal(app.Listen(":8000"))
}
