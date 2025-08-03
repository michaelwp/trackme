package handlers

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/michaelwp/trackme/internal/config"
	"github.com/michaelwp/trackme/internal/repository"
)

func SetupRoutes(app *fiber.App, s3Client *s3.Client, s3Config *config.S3Config) {
	locationRepo := repository.NewLocationRepository()
	locationHandler := NewLocationHandler(locationRepo, s3Client, s3Config)

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Allow requests from any origin (for development, be more specific in production)
		AllowHeaders: []string{"Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin"},
		AllowMethods: []string{"GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS"},
	}))

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, Black hole!")
	})

	app.Post("/locations", locationHandler.SaveLocation)
	app.Get("/locations", locationHandler.GetLocations)
	app.Post("/locations/photos", locationHandler.UploadPhoto)
}
