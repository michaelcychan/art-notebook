// National Palace Museum, Taiwan 台灣故宮
// https://openapiweb.npm.gov.tw/APP_Prog/cht/overview_cht.aspx

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/config"
)

type ArtDataNpm struct {
	Title        string `json:"ArticleSubject"`
	ID           string `json:"Serial_No"`
	ArticleMaker string `json:"ArticleMaker"`
	ImageUrl     string `json:"imgUrl"`
	LongDesc     string `json:"ArticleContext"`
}

type ArtworkFromSearchId struct {
	Result   []ArtDataNpm `json:"result"`
	Status   int          `json:"status"`
	ErrorMsg string       `json:"error"`
}

const NpmBaseUri = "https://openapi.npm.gov.tw/v1/rest"

func GetTargetArtworkByIdNpm(artworkID string, c *fiber.Ctx) (ArtDataNpm, error) {
	var artwork ArtworkFromSearchId

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.Set("apiKey", config.Config("NPM_API"))
	req.SetRequestURI(NpmBaseUri + "/collection/search/" + artworkID)

	if err := a.Parse(); err != nil {
		return artwork.Result[0], fmt.Errorf("error retrieving info from external api: %v", err)
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		return artwork.Result[0], fmt.Errorf("errors returned from external api")
	}

	if code != http.StatusOK {
		return artwork.Result[0], fmt.Errorf("return code from api not OK: %d", code)
	}

	if unmarshalErr := json.Unmarshal(bodyBytes, &artwork); unmarshalErr != nil {
		return artwork.Result[0], fmt.Errorf("json unmarshal Error at API")
	}
	return artwork.Result[0], nil

}
