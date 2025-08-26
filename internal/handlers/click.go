package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/michaelwp/trackme/internal/services"
)

type ClickHandler struct {
	telegramService *services.TelegramService
}

func NewClickHandler() *ClickHandler {
	return &ClickHandler{
		telegramService: services.NewTelegramService(),
	}
}

func (h *ClickHandler) TrackClick(c fiber.Ctx) error {
	// Get visitor information
	ip := c.IP()
	userAgent := c.Get("User-Agent")
	referer := c.Get("Referer")
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Get the target URL from query parameter
	targetURL := c.Query("url", "")
	if targetURL == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "URL parameter is required",
		})
	}

	// Create a notification message
	message := fmt.Sprintf("🔗 Link Click Detected!\n\n"+
		"Time: %s\n"+
		"IP: %s\n"+
		"User Agent: %s\n"+
		"Referer: %s\n"+
		"Target URL: %s",
		timestamp, ip, userAgent, referer, targetURL)

	// Send Telegram notification
	go func() {
		if err := h.telegramService.SendNotification(message); err != nil {
			fmt.Printf("Error sending Telegram notification: %v\n", err)
		}
	}()

	// Redirect to the target URL
	return c.Redirect().To(targetURL)
}
