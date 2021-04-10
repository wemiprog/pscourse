package handler

import (
	u "pscourse/utils"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PowershellHandler(c *fiber.Ctx, p *u.PowerShell) error {
	// t := c.Params("task")
	// u := c.Params("user")
	b := c.Body()

	stdOut, stdErr, err := p.Execute(fmt.Sprintf("%s", b))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":   "Code nicht ausführbar",
			"response": stdErr,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "Erfolgreich ausgeführt",
		"response": fmt.Sprintf("# Output\n%s", stdOut),
	})
}
