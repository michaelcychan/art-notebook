package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/michaelcychan/art-notebook/backend-go/config"
	"github.com/michaelcychan/art-notebook/backend-go/database"
	"github.com/michaelcychan/art-notebook/backend-go/router"
)

func main() {

	//database.Connect()

	app := fiber.New()
	database.Connect()

	serverPort := config.Config("PORT")
	if serverPort == "" {
		serverPort = "3001"
	}

	app.Static("/", "/")
	app.Use(cors.New())

	router.SetupRouter(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(&fiber.Map{"message": "not found"})
	})

	app.Listen(":" + serverPort)
}
