package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/michaelwp/trackme/internal/models"
	"github.com/michaelwp/trackme/internal/repository"
	"time"
)

type targetRequest struct {
	Location models.LocationInformation `json:"location" validate:"required"`
	Device   models.DeviceInformation   `json:"device" validate:"required"`
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
