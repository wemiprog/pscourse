package handler

import (
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

func LessonHandler(c *fiber.Ctx) error {

	textPath := c.Query("text")
	textContent, _ := ioutil.ReadFile(textPath)

	return c.JSON(fiber.Map{
		"text": string(textContent),
	})
}
