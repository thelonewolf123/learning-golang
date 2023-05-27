package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type data struct {
	Hello string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("route: ", c.Route().Path)
		json := &data{Hello: "world"} // Pass a pointer to data struct
		return c.JSON(json)
	})

	log.Fatal(app.Listen(":8000"))
}
