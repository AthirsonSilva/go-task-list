package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" binding:"required"`
	Artist    string             `json:"artist" binding:"required"`
	Price     float64            `json:"price" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
