package main

import (
	"pscourse/handler"
	"pscourse/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	posh := utils.New()
	/* app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	}) */

	app.Static("/", "./html")

	app.Get("/lessons", handler.LessonListHandler)

	app.Get("/lesson/:section/:lesson", handler.LessonHandler)

	app.Post("/ps/:task", func(c *fiber.Ctx) error {
		return handler.PowershellHandler(c, posh)
	})

	type SomeStruct struct {
		Name string
		Age  uint8
	}

	app.Get("/json", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}

		return c.JSON(data)
		// => Content-Type: application/json
		// => "{"Name": "Grame", "Age": 20}"

	})

	app.Listen(":3000")
}
