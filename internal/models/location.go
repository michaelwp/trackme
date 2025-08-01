package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// Location represents a GPS location tracking entry
type Location struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TargetID  string        `bson:"target_id" json:"target_id"`
	Latitude  float64       `bson:"latitude" json:"latitude"`
	Longitude float64       `bson:"longitude" json:"longitude"`
	Timestamp time.Time     `bson:"timestamp" json:"timestamp"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}
