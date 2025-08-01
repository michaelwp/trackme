package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/michaelwp/trackme/internal/repository"
)

func SetupRoutes(app *fiber.App) {
	locationRepo := repository.NewLocationRepository()
	locationHandler := NewLocationHandler(locationRepo)

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, Black hole!")
	})

	app.Post("/locations", locationHandler.SaveLocation)
	app.Get("/locations", locationHandler.GetLocations)
}
