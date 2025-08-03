package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Target struct {
	ID        bson.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"`
	Location  LocationInformation `bson:"location" json:"location"`
	Device    DeviceInformation   `bson:"device" json:"device"`
	Photo     Photo               `bson:"photo" json:"photo"`
	Timestamp time.Time           `bson:"timestamp" json:"timestamp"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
}

type DeviceInformation struct {
	Model           string `bson:"model" json:"model"`
	OperatingSystem string `bson:"operating_system" json:"operating_system"`
	Platform        string `bson:"platform" json:"platform"`
	UserAgent       string `bson:"user_agent" json:"user_agent"`
	Brand           string `bson:"brand" json:"brand"`
	Browser         string `bson:"browser" json:"browser"`
}

type LocationInformation struct {
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}

type Photo struct {
	Name string `bson:"name" json:"name"`
	Path string `bson:"path" json:"path"`
}
