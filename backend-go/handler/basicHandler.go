package handler

import (
	"github.com/gofiber/fiber/v2"
)

func BasicHelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Welcome To Art Notebook"})
}
