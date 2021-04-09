package handler

import (
	u "pscourse/utils"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PowershellHandler(c *fiber.Ctx) error {
	// t := c.Params("task")
	// u := c.Params("user")
	b := c.Body()

	posh := u.New()

	stdOut, stdErr, err := posh.Execute(fmt.Sprintf("%s", b))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":   "Code nicht ausführbar",
			"response": fmt.Sprintf("%s", stdErr),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "Erfolgreich ausgeführt",
		"response": fmt.Sprintf("# Output\n%s", stdOut),
	})
}
