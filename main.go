package main

import (
	todoAPI "jya-todos/src/api/todo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("using fiber")
	})

    api := app.Group("/api")

    todoAPI.AddRoutes(api)

	app.Listen("0.0.0.0:7000")
}
