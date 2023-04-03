// Metropolitan Museum of Art
// https://metmuseum.github.io/#search

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MMAArtworkSimple struct {
	Title     string `json:"title"`
	ID        int    `json:"objectID"`
	Artist    string `json:"artistDisplayName"` // there is artistDisplayAlphaSort, check that if necessary
	ImageURL  string `json:"primaryImageSmall"`
	DateStart int    `json:"objectBeginDate"`
	DateEnd   int    `json:"objectEndDate"`
}

func GetExampleMetMAArt(c *fiber.Ctx) error {
	const MetMuseumOfArtBaseUri = "https://collectionapi.metmuseum.org/public/collection/v1"

	listOfExmapleArtid := []int{437133}
	exampleArtid := listOfExmapleArtid[rand.Intn(len(listOfExmapleArtid))]

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(MetMuseumOfArtBaseUri + "/objects" + "/" + fmt.Sprint(exampleArtid))

	if err := a.Parse(); err != nil {
		log.Fatalf("error retrieving info from external api")
		return c.Status(500).JSON(fiber.Map{"message": "external api error"})
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		log.Fatalln("errors returned from external api")
		return c.Status(500).JSON(fiber.Map{"message": "external api error"})
	}

	if code != http.StatusOK {
		log.Fatalf("return code from api not OK: %d", code)
		return c.Status(500).JSON(fiber.Map{"message": "external api error"})
	}

	var artwork MMAArtworkSimple

	if unmarshalErr := json.Unmarshal(bodyBytes, &artwork); unmarshalErr != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error"})
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
