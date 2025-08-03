package repository

import (
	"context"
	"log"
	"time"

	"github.com/michaelwp/trackme/internal/config"
	"github.com/michaelwp/trackme/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LocationRepository interface {
	Create(location *models.Target) error
	GetAll() ([]models.Target, error)
	UpdatePhoto(id string, photo models.Photo) error
}

type locationRepository struct {
	collection *mongo.Collection
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{
		collection: config.GetCollection("locations"),
	}
}

func (r *locationRepository) Create(target *models.Target) error {
	target.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, target)
	if err != nil {
		log.Printf("Failed to insert document: %v", err)
		return err
	}

	target.ID = result.InsertedID.(bson.ObjectID)
	log.Printf("Successfully inserted document with ID: %s", target.ID.Hex())

	return nil
}

func (r *locationRepository) GetAll() ([]models.Target, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var locations []models.Target
	if err = cursor.All(ctx, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func (r *locationRepository) UpdatePhoto(id string, photo models.Photo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID format: %s, error: %v", id, err)
		return err
	}

	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{"photo": photo}},
	)

	if err != nil {
		log.Printf("MongoDB update error: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		log.Printf("No document found with ID: %s", objectID.Hex())
		return mongo.ErrNoDocuments
	}

	if result.ModifiedCount == 0 {
		log.Printf("Document found but not modified for ID: %s", objectID.Hex())
	} else {
		log.Printf("Successfully updated photo for document ID: %s", objectID.Hex())
	}

	return nil
}
