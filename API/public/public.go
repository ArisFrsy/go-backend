package public

import (
	"github.com/gofiber/fiber/v2"
)

func AccessibleHandler(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}