package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DatabaseConfig struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var DB *DatabaseConfig

// ConnectDB initializes the MongoDB connection
func ConnectDB() error {
	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Println("MONGODB_URI environment variable is not set")
		return fmt.Errorf("MONGODB_URI environment variable is not set")
	}

	// Get database name from environment variable
	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		log.Println("MONGODB_NAME environment variable is not set")
		return fmt.Errorf("MONGODB_NAME environment variable is not set")
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Println("Failed to connect to MongoDB:", err)
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Test the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Failed to ping MongoDB:", err)
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// Initialize global database configuration
	DB = &DatabaseConfig{
		Client:   client,
		Database: client.Database(dbName),
	}

	log.Println("Successfully connected to MongoDB!")
	return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if DB != nil && DB.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := DB.Client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			fmt.Println("Disconnected from MongoDB")
		}
	}
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	if DB == nil || DB.Database == nil {
		log.Fatal("Database not initialized. Call ConnectDB() first.")
	}
	return DB.Database.Collection(name)
}
