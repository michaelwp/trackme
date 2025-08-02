package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/joho/godotenv"
	"github.com/michaelwp/trackme/internal/config"
	"github.com/michaelwp/trackme/internal/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// Only load the .env file in development (optional in production)
	if os.Getenv("ENV") == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file")
		}
	}
}

func main() {
	// Initialize database connection
	err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	handlers.SetupRoutes(app)

	// Serve static files from the web / static directory
	app.Use(static.New("./web"))

	// Set up a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		config.DisconnectDB()
		os.Exit(0)
	}()

	log.Printf("Server starting on port %s", os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
