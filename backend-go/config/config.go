package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	var filename string = ".env"
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalln("Error reading ", filename)
	}
	return os.Getenv(key)
}
