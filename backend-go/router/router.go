package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelcychan/art-notebook/backend-go/handler"
)

func SetupRouter(app *fiber.App) {
	v1Path := app.Group("/v1")

	v1Path.Get("/", handler.BasicHelloWorld)

	v1Path.Get("/get-example-painting", handler.GetExamplePainting)
}
