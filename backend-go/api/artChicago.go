// Art Institute of Chicago
// https://api.artic.edu/docs/#collections

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AIChicagoArtworkSimple struct {
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

func GetTargetArtworkChicago(artworkID int, c *fiber.Ctx) (AIChicagoArtworkSimple, error) {
	const ArtInstituteOfChicigoBaseUri = "https://api.artic.edu/api/v1"

	var artwork AIChicagoArtworkSimple

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(ArtInstituteOfChicigoBaseUri + "/artworks" + "/" + fmt.Sprint(artworkID))

	if err := a.Parse(); err != nil {
		return artwork, fmt.Errorf("error retrieving info from external api: %v", err)
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		return artwork, fmt.Errorf("errors returned from external api")
	}

	if code != http.StatusOK {
		return artwork, fmt.Errorf("return code from api not OK: %d", code)
	}

	if unmarshalErr := json.Unmarshal(bodyBytes, &artwork); unmarshalErr != nil {
		return artwork, fmt.Errorf("json unmarshal Error")
	}
	return artwork, nil
}
