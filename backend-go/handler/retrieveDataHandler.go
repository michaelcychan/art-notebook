package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/database"
)

func ReturnSavedDataJsonForUser(c *fiber.Ctx) error {
	username := c.Query("username")
	data, err := database.GetDataForUser(username, "./")

	if err != nil {
		return c.Status(500).SendString("Internal Error")
	}

	return c.Status(200).JSON(fiber.Map{
		"data": data,
	})
}
