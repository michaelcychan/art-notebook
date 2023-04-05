// Metropolitan Museum of Art
// https://metmuseum.github.io/#search

package handler

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/api"
)

func GetExampleMetMAArt(c *fiber.Ctx) error {
	listOfExmapleArtid := []int{438011, 437428, 437122}
	exampleArtid := listOfExmapleArtid[rand.Intn(len(listOfExmapleArtid))]

	artwork, err := api.GetArtworkMetropolitanMusuemArt(exampleArtid, c)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"message": "server error"})
	}

	artistSlice := []string{artwork.Artist}

	fmt.Println(artwork.ImageURL)

	return c.Status(200).JSON(fiber.Map{
		"title":             artwork.Title,
		"short-description": "nil",
		"artist-title":      artistSlice,
		"image-url":         artwork.ImageURL,
		"work-start":        artwork.DateStart,
		"work-end":          artwork.DateEnd,
		"museum":            "Metropolitan Museum of Art",
		"message":           "OK",
	})

}
