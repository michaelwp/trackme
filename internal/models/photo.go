package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// Photo represents a captured photo entry
type Photo struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TargetID   string        `bson:"user_id" json:"user_id"`
	LocationID bson.ObjectID `bson:"location_id,omitempty" json:"location_id,omitempty"`
	Filename   string        `bson:"filename" json:"filename"`
	Path       string        `bson:"path" json:"path"`
	CreatedAt  time.Time     `bson:"created_at" json:"created_at"`
}
