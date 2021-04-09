package main

import (
	"pscourse/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	/* app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	}) */
	app.Static("/", "./html")
	app.Post("/ps/:task/:user", handler.PowershellHandler)

	app.Listen(":3000")
}
