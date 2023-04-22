// Art Institute of Chicago
// https://api.artic.edu/docs/#collections

package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ArtData struct {
	Title        string   `json:"title"`
	ID           int      `json:"id"`
	ArtistTitles []string `json:"artist_titles"`
	ImageId      string   `json:"image_id"`
	DateStart    int      `json:"date_start"`
	DateEnd      int      `json:"date_end"`
	Thumbnail    struct {
		ShortDesc string `json:"alt_text"`
	} `json:"thumbnail"`
}

type AIChicagoArtworkSimple struct {
	Data   ArtData `json:"data"`
	Config struct {
		IIIFUrl string `json:"iiif_url"`
	} `json:"config"`
}

type AIChicagoPagination struct {
	Pagination struct {
		Total int `json:"total_pages"`
	} `json:"pagination"`
}

type ArtworkFromPage struct {
	Data   []ArtData `json:"data"`
	Config struct {
		IIIFUrl string `json:"iiif_url"`
	} `json:"config"`
}

const ArtInstituteOfChicigoBaseUri = "https://api.artic.edu/api/v1"

func GetTotalNumberOfArtworkChicago(c *fiber.Ctx) (int, error) {
	var pagination AIChicagoPagination

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(ArtInstituteOfChicigoBaseUri + "/artworks?limit=1")

	if err := a.Parse(); err != nil {
		return pagination.Pagination.Total, fmt.Errorf("error retrieving info from external api: %v", err)
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		return pagination.Pagination.Total, fmt.Errorf("errors returned from external api")
	}

	if code != http.StatusOK {
		return pagination.Pagination.Total, fmt.Errorf("return code from api not OK: %d", code)
	}

	if unmarshalErr := json.Unmarshal(bodyBytes, &pagination); unmarshalErr != nil {
		return pagination.Pagination.Total, fmt.Errorf("json unmarshal Error at GetTotalNumberOfArtworkChicago")
	}
	return pagination.Pagination.Total, nil
}

func GetArtworkChicagoByPageNumber(pageNumber int, c *fiber.Ctx) (AIChicagoArtworkSimple, error) {
	var artwork AIChicagoArtworkSimple
	var artworkFromPageNumber ArtworkFromPage

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)

	reqUri := ArtInstituteOfChicigoBaseUri + "/artworks?limit=1&page=" + fmt.Sprint(pageNumber)
	req.SetRequestURI(reqUri)

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

	if unmarshalErr := json.Unmarshal(bodyBytes, &artworkFromPageNumber); unmarshalErr != nil {
		return artwork, fmt.Errorf("json unmarshal Error at GetArtworkChicagoByPageNumber")
	}

	artwork.Data = artworkFromPageNumber.Data[0]
	artwork.Config = artworkFromPageNumber.Config
	return artwork, nil
}

func GetRandomArtworkChicagoByTotalPage(c *fiber.Ctx) (AIChicagoArtworkSimple, error) {
	var artwork AIChicagoArtworkSimple
	totalPage, err := GetTotalNumberOfArtworkChicago(c)
	if err != nil {
		return artwork, err
	}

	targetPage := rand.Intn(totalPage) + 1
	return GetArtworkChicagoByPageNumber(targetPage, c)
}

func GetTargetArtworkChicagoById(artworkID int, c *fiber.Ctx) (AIChicagoArtworkSimple, error) {

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
