// Art Institute of Chicago
// https://api.artic.edu/docs/#collections

package handler

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/api"
)

func GetExampleArticPainting(c *fiber.Ctx) error {
	listOfExampleArtId := []int{15742, 238399, 150152}
	exampleArtId := listOfExampleArtId[rand.Intn(len(listOfExampleArtId))]

	artwork, err := api.GetTargetArtworkChicago(exampleArtId, c)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"message": "server error"})
	}

	fullImageUrl := artwork.Config.IIIFUrl + "/" + artwork.Data.ImageId + "/full/600,/0/default.jpg"

	return c.Status(200).JSON(fiber.Map{
		"title":             artwork.Data.Title,
		"short-description": artwork.Data.Thumbnail.ShortDesc,
		"artist-title":      artwork.Data.ArtistTitles,
		"image-url":         fullImageUrl,
		"museum":            "Art Institute of Chicago",
		"message":           "OK",
	})
}
