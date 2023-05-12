// Metropolitan Museum of Art
// https://metmuseum.github.io/#search

package api

import (
	"encoding/json"
	"fmt"
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

func GetArtworkMetropolitanMusuemArt(artworkID string, c *fiber.Ctx) (MMAArtworkSimple, error) {
	const MetMuseumOfArtBaseUri = "https://collectionapi.metmuseum.org/public/collection/v1"

	var artwork MMAArtworkSimple

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(MetMuseumOfArtBaseUri + "/objects" + "/" + artworkID)

	if err := a.Parse(); err != nil {
		return artwork, fmt.Errorf("error retrieving info from external api: %s", err)
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		return artwork, fmt.Errorf("errors returned from external api: %s", errs)
	}

	if code != http.StatusOK {
		return artwork, fmt.Errorf("return code from api not OK: %d", code)
	}

	if unmarshalErr := json.Unmarshal(bodyBytes, &artwork); unmarshalErr != nil {
		return artwork, fmt.Errorf("json unmarshal error")
	}

	return artwork, nil

}
