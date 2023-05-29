package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Tasks []string = []string{"hello, world"}

type ResponseType struct {
	Updated bool `json:"updated"`
}

type Todo struct {
	ID   uint   `gorm:"primaryKey"`
	Task string `json:"task"`
}

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Log initialize failed!")
	}

	db.AutoMigrate(&Todo{})
}

func addTodoTask(task string) error {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Create(&Todo{Task: task})
	return nil
}

func GetAllTodo() ([]Todo, error) {
	var todos []Todo

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Find(&todos)

	return todos, nil

}

func main() {
	app := fiber.New()

	InitDatabase()

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
		todos, err := GetAllTodo()
		if err != nil {
			return c.Status(500).SendString("Server side error, please contact support!")
		}
		return c.JSON(&todos)
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

		addTodoTask(payload.Task)

		fmt.Println(payload.Task)
		return c.JSON(&ResponseType{Updated: true})
	})

	app.Delete("/task/:id", func(c *fiber.Ctx) error {
		fmt.Printf("params: %T\n", c.Params("id"))
		index, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		fmt.Println("index => ", index)

		TasksCopy := make([]string, len(Tasks))
		copy(TasksCopy, Tasks)

		Tasks = make([]string, 0)

		for i, v := range TasksCopy {
			if index == i {
				continue
			}
			Tasks = append(Tasks, v)
		}

		return c.JSON(&ResponseType{Updated: true})
	})

	log.Fatal(app.Listen(":8000"))
}
