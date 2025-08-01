package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/michaelwp/trackme/internal/models"
	"github.com/michaelwp/trackme/internal/repository"
	"time"
)

type locationRequest struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type LocationHandler interface {
	SaveLocation(c fiber.Ctx) error
	GetLocations(c fiber.Ctx) error
}

type locationHandler struct {
	repository repository.LocationRepository
}

func NewLocationHandler(locationRepository repository.LocationRepository) LocationHandler {
	return locationHandler{
		repository: locationRepository,
	}
}

func (l locationHandler) SaveLocation(c fiber.Ctx) error {
	location := new(locationRequest)

	if err := c.Bind().All(location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	targetID, err := uuid.NewUUID()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate UUID",
		})
	}

	locationModel := &models.Location{
		TargetID:  targetID.String(),
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Timestamp: time.Now(),
	}

	if err := l.repository.Create(locationModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save location",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": locationModel.ID,
	})
}

func (l locationHandler) GetLocations(c fiber.Ctx) error {
	locations, err := l.repository.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch locations",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"locations": locations,
	})
}
