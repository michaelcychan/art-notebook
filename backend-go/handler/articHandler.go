// Art Institute of Chicago
// https://api.artic.edu/docs/#collections

package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/api"
)

func GetExampleArticPainting(c *fiber.Ctx) error {
	artwork, err := api.GetRandomArtworkChicagoByTotalPage(c)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"message": "server error"})
	}

	fullImageUrl := artwork.Config.IIIFUrl + "/" + artwork.Data.ImageId + "/full/600,/0/default.jpg"

	return c.Status(200).JSON(fiber.Map{
		"title":             artwork.Data.Title,
		"short-description": artwork.Data.Thumbnail.ShortDesc,
		"artist-title":      artwork.Data.ArtistTitles,
		"work-start":        artwork.Data.DateStart,
		"work-end":          artwork.Data.DateEnd,
		"image-url":         fullImageUrl,
		"museum":            "Art Institute of Chicago",
		"message":           "OK",
	})
}
