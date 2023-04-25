package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/api"
)

func GetArtworkNpm(c *fiber.Ctx) error {
	artwork, err := api.GetTargetArtworkByIdNpm("04000975", c)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"message": "server err"})
	}

	return c.Status(200).JSON(fiber.Map{
		"title":             artwork.Title,
		"short-description": artwork.LongDesc,
		"artist-title":      artwork.ArticleMaker,
		"work-start":        0,
		"work-end":          0,
		"image-url":         artwork.ImageUrl,
		"museum":            "National Palace Museum",
		"message":           "OK",
	})
}
