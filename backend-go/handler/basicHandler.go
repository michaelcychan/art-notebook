package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func BasicHelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Welcome To Art Notebook"})
}

func GetExamplePainting(c *fiber.Ctx) error {
	const ArtInstituteOfChicigoBaseUri = "https://api.artic.edu/api/v1/"

	exampleArtId := "15742"

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(ArtInstituteOfChicigoBaseUri + "/artworks" + "/" + exampleArtId)

	if err := a.Parse(); err != nil {
		log.Fatalf("error retrieving info from external api: %v", err)
	}

	code, bodyBytes, errs := a.Bytes()

	if len(errs) != 0 {
		log.Fatalln("errors returned from external api")
	}

	fmt.Printf("code: %d\n", code)

	body := string(bodyBytes)
	fmt.Printf("body: %s", body)

	return c.Status(200).SendString(body)
}
