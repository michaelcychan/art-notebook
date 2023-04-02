package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ArtworkSimple struct {
	Data struct {
		Title        string   `json:"title"`
		ID           int      `json:"id"`
		ArtistTitles []string `json:"artist_titles"`
		ImageId      string   `json:"image_id"`
		DateStart    int      `json:"date_start"`
		DateEnd      int      `json:"date_end"`
		Thumbnail    struct {
			ShortDesc string `json:"alt_text"`
		} `json:"thumbnail"`
	} `json:"data"`
	Config struct {
		IIIFUrl string `json:"iiif_url"`
	} `json:"config"`
}

func GetExampleArticPainting(c *fiber.Ctx) error {
	const ArtInstituteOfChicigoBaseUri = "https://api.artic.edu/api/v1/"

	exampleArtId := 238399

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(ArtInstituteOfChicigoBaseUri + "/artworks" + "/" + fmt.Sprint(exampleArtId))

	if err := a.Parse(); err != nil {
		log.Fatalf("error retrieving info from external api: %v", err)
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		log.Fatalln("errors returned from external api")
	}

	if code != http.StatusOK {
		log.Fatalf("return code from api not OK: %d", code)
	}

	var artwork ArtworkSimple
	err := json.Unmarshal(bodyBytes, &artwork)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"title": "error"})
	}

	fullImageUrl := artwork.Config.IIIFUrl + "/" + artwork.Data.ImageId + "/full/600,/0/default.jpg"

	return c.Status(200).JSON(fiber.Map{"title": artwork.Data.Title, "short-description": artwork.Data.Thumbnail.ShortDesc, "artist-title": artwork.Data.ArtistTitles, "image-url": fullImageUrl})
}
