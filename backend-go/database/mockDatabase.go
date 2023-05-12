package database

import (
	"encoding/json"
	"log"
	"os"
)

type Database struct {
	Data []ArtData `json:"data"`
}

type ArtData struct {
	Source   string   `json:"source"`
	ID       int      `json:"id"`
	SourceId string   `json:"source-id"`
	User     string   `json:"user"`
	Tag      []string `json:"tag"`
	Note     string   `json:"note"`
}

func GetDataForUser(username string, rootDir string) ([]ArtData, error) {
	var payload []ArtData

	content, err := os.ReadFile(rootDir + "mock-data/mock-data.json")
	if err != nil {
		log.Fatal("error when opening file: ", err)
		return payload, err
	}

	var allArtData Database
	err = json.Unmarshal(content, &allArtData)
	if err != nil {
		log.Fatal("error when unmarshalling json: ", err)
		return payload, err
	}

	// do something with the payload...
	for i := range allArtData.Data {
		if allArtData.Data[i].User == username {
			payload = append(payload, allArtData.Data[i])
		}
	}

	return payload, nil
}
