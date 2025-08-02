package repository

import (
	"github.com/michaelwp/trackme/internal/config"
	"github.com/michaelwp/trackme/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PhotoRepository interface {
	Create(photo *models.Photo) error
	GetPhotoByLocationID(bson.ObjectID) ([]models.Photo, error)
	GetPhotoByTargetID(targetID string) ([]models.Photo, error)
	GetPhotoByID(id string) (*models.Photo, error)
	GetPhotos() ([]models.Photo, error)
}

type photoRepository struct {
	collection *mongo.Collection
}

func NewPhotoRepository() PhotoRepository {
	return &photoRepository{
		collection: config.GetCollection("photos"),
	}
}

func (p photoRepository) Create(photo *models.Photo) error {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) GetPhotoByLocationID(id bson.ObjectID) ([]models.Photo, error) {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) GetPhotoByTargetID(targetID string) ([]models.Photo, error) {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) GetPhotoByID(id string) (*models.Photo, error) {
	//TODO implement me
	panic("implement me")
}

func (p photoRepository) GetPhotos() ([]models.Photo, error) {
	//TODO implement me
	panic("implement me")
}
