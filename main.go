package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	students := getStudents()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/students", func(c *fiber.Ctx) error {
		if bytes, err := json.Marshal(students); err == nil {
			return c.SendString(string(bytes))
		} else {
			return fiber.NewError(599, "Network error!")
		}
	})

	app.Listen(":3000")
}

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func getStudents() []Student {
	students := []Student{
		{1, "Mary", "Mary", 22},
		{2, "Joseph", "Joseph", 23},
		{3, "John", "John", 21},
		{4, "Mark", "Mark", 20},
		{5, "Ray", "Ray", 221},
	}

	return students
}
