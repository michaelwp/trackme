package handlers

import (
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/michaelwp/trackme/internal/config"

	"github.com/gofiber/fiber/v3"
	"github.com/michaelwp/trackme/internal/models"
	"github.com/michaelwp/trackme/internal/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type targetRequest struct {
	Location models.LocationInformation `json:"location" validate:"required"`
	Device   models.DeviceInformation   `json:"device" validate:"required"`
}

type photoUploadRequest struct {
	ID string `form:"id" validate:"required"`
}

type LocationHandler interface {
	SaveLocation(c fiber.Ctx) error
	GetLocations(c fiber.Ctx) error
	UploadPhoto(c fiber.Ctx) error
}

type locationHandler struct {
	repository repository.LocationRepository
	s3Client   *s3.Client
	s3Config   *config.S3Config
}

func NewLocationHandler(locationRepository repository.LocationRepository, s3Client *s3.Client, s3Config *config.S3Config) LocationHandler {
	return locationHandler{
		repository: locationRepository,
		s3Client:   s3Client,
		s3Config:   s3Config,
	}
}

func (l locationHandler) SaveLocation(c fiber.Ctx) error {
	target := new(targetRequest)

	if err := c.Bind().All(target); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	locationModel := &models.Target{
		Location:  target.Location,
		Device:    target.Device,
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

	req := new(photoUploadRequest)
	if err := c.Bind().All(req); err != nil {
		log.Println("error on get location id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	filename := time.Now().Format("20060102150405") + "_" + file.Filename
	path := "photos/" + filename

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
			"error": "Failsdsded to read file",
		})
	}

	if err := l.s3Config.UploadFile(l.s3Client, path, buffer); err != nil {
		log.Println("error on upload file to S3:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file to S3",
		})
	}

	s3URL := l.s3Config.GetObjectURL(path)
	photo := models.Photo{
		Name: filename,
		Path: s3URL,
	}

	if err := l.repository.UpdatePhoto(req.ID, photo); err != nil {
		log.Printf("UpdatePhoto failed for ID %s: %v\n", req.ID, err)

		// Check if it's a "no documents" error
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Location not found with the provided ID",
			})
		}

		// Check if it's an invalid ObjectID error
		if err.Error() == "encoding/hex: invalid byte" || err.Error() == "encoding/hex: odd length hex string" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid location ID format",
			})
		}

		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update photo information",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"filename": filename,
		"path":     path,
	})
}

func (l locationHandler) GetLocations(c fiber.Ctx) error {
	locations, err := l.repository.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch locations",
		})
	}

	// Debug: Log all available document IDs
	log.Printf("Available document IDs in database:")
	for i, location := range locations {
		log.Printf("  [%d] ID: %s", i, location.ID.Hex())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"locations": locations,
	})
}
