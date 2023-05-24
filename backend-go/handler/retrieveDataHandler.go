package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/api"
	"github.com/michaelcychan/art-notebook/backend-go/database"
)

type ArtworkDataToFrontend struct {
	Title        string   `json:"title"`
	SourceID     string   `json:"source-id"`
	ShortDesc    string   `json:"short-description"`
	ArtistTitles []string `json:"artist-title"`
	DateStart    int      `json:"date_start"`
	DateEnd      int      `json:"date_end"`
	FullImageUrl string   `json:"image-url"`
	Musuem       string   `json:"museum"`
	Message      string
	Tags         []string `json:"tags"`
	Note         string   `json:"note"`
}

func ReturnSavedDataJsonForUser(c *fiber.Ctx) error {
	username := c.Query("username")
	fmt.Println(username)
	data, err := database.GetDataForUser(username, "./")

	if err != nil {
		fmt.Println("error at reading json")
		return c.Status(500).SendString("Internal Error")
	}

	var outputJson []ArtworkDataToFrontend

	for _, eachEntry := range data {

		var jsonEntry ArtworkDataToFrontend
		if eachEntry.Source == "National Palace Museum, Taipei" {
			artworkNPM, err := api.GetTargetArtworkByIdNpm(eachEntry.SourceId, c)
			if err != nil {
				log.Fatal("API error: ", err)
				break
			}
			jsonEntry.ArtistTitles = []string{artworkNPM.ArticleMaker}
			jsonEntry.FullImageUrl = artworkNPM.ImageUrl
			jsonEntry.ShortDesc = artworkNPM.LongDesc
			jsonEntry.Title = artworkNPM.Title
			jsonEntry.SourceID = artworkNPM.ID

		}
		if eachEntry.Source == "Metropolitan Museum of Art, New York" {
			artworkMetMuseumArt, err := api.GetArtworkMetropolitanMusuemArtWtihID(eachEntry.SourceId, c)
			if err != nil {
				log.Fatal("API error: ", err)
				break
			}
			jsonEntry.ArtistTitles = []string{artworkMetMuseumArt.Artist}
			jsonEntry.FullImageUrl = artworkMetMuseumArt.ImageURL
			jsonEntry.DateStart = artworkMetMuseumArt.DateStart
			jsonEntry.DateEnd = artworkMetMuseumArt.DateEnd
			jsonEntry.Title = artworkMetMuseumArt.Title
			jsonEntry.SourceID = fmt.Sprint(artworkMetMuseumArt.ID)
		}
		if eachEntry.Source == "Art Institute of Chicago, Chicago" {
			sourceID, err := strconv.Atoi(eachEntry.SourceId)
			if err != nil {
				log.Fatal("source id int to string conversion failed: ", err)
				break
			}
			artworkChicago, err := api.GetTargetArtworkChicagoById(sourceID, c)
			if err != nil {
				log.Fatal("API error: ", err)
				break
			}
			jsonEntry.FullImageUrl = artworkChicago.Config.IIIFUrl + "/" + artworkChicago.Data.ImageId + "/full/600,/0/default.jpg"
			jsonEntry.ArtistTitles = artworkChicago.Data.ArtistTitles
			jsonEntry.DateStart = artworkChicago.Data.DateStart
			jsonEntry.DateEnd = artworkChicago.Data.DateEnd
			jsonEntry.ShortDesc = artworkChicago.Data.Thumbnail.ShortDesc
			jsonEntry.SourceID = fmt.Sprint(artworkChicago.Data.ID)
			jsonEntry.Title = artworkChicago.Data.Title
		}
		jsonEntry.Musuem = eachEntry.Source
		jsonEntry.Tags = eachEntry.Tag
		jsonEntry.Note = eachEntry.Note
		outputJson = append(outputJson, jsonEntry)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": outputJson,
	})
}
