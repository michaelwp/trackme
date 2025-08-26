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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Set client options with retry and timeout configurations
	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetConnectTimeout(30 * time.Second).
		SetServerSelectionTimeout(30 * time.Second).
		SetRetryWrites(true).
		SetRetryReads(true).
		SetMaxConnIdleTime(30 * time.Second).
		SetMaxPoolSize(10).
		SetMinPoolSize(1)

	// Connect to MongoDB with retries
	var client *mongo.Client
	var err error
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		client, err = mongo.Connect(clientOptions)
		if err != nil {
			log.Printf("Failed to connect to MongoDB (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				time.Sleep(time.Duration(i+1) * 2 * time.Second)
				continue
			}
			return fmt.Errorf("failed to connect to MongoDB after %d attempts: %w", maxRetries, err)
		}
		break
	}

	// Test the connection with retries
	for i := 0; i < maxRetries; i++ {
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			log.Printf("Failed to ping MongoDB (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				time.Sleep(time.Duration(i+1) * 2 * time.Second)
				continue
			}
			return fmt.Errorf("failed to ping MongoDB after %d attempts: %w", maxRetries, err)
		}
		break
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
