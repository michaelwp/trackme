package repository

import (
	"context"
	"time"

	"github.com/michaelwp/trackme/internal/config"
	"github.com/michaelwp/trackme/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LocationRepository interface {
	Create(location *models.Target) error
	GetAll() ([]models.Target, error)
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
		return err
	}

	target.ID = result.InsertedID.(bson.ObjectID)
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
