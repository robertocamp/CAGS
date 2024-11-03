package handlers

import (
	"github.com/gofiber/fiber/v2"
)


// hello function: handles routes to "/"
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Updated World!")
}
