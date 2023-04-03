// Metropolitan Museum of Art
// https://metmuseum.github.io/#search

package handler

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/api"
)

func GetExampleMetMAArt(c *fiber.Ctx) error {
	listOfExmapleArtid := []int{437133}
	exampleArtid := listOfExmapleArtid[rand.Intn(len(listOfExmapleArtid))]

	artwork, err := api.GetArtworkMetropolitanMusuemArt(exampleArtid, c)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "server error"})
	}

	artistSlice := []string{artwork.Artist}

	return c.Status(200).JSON(fiber.Map{
		"title":             artwork.Title,
		"short-description": "nil",
		"artist-title":      artistSlice,
		"image-url":         artwork.ImageURL,
		"museum":            "Metropolitan Museum of Art",
		"message":           "OK",
	})

}
