package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Title     string             `json:"title"                           binding:"required"`
	Artist    string             `json:"artist"                          binding:"required"`
	Price     float64            `json:"price"                           binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type AlbumRequest struct {
	Title  string  `json:"title"  binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price"  binding:"required"`
}

type AlbumResponse struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	Title     string             `json:"title"                           binding:"required"`
	Artist    string             `json:"artist"                          binding:"required"`
	Price     float64            `json:"price"                           binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (a *AlbumRequest) ToModel() Album {
	var album Album

	album.Title = a.Title
	album.Artist = a.Artist
	album.Price = a.Price

	return album
}

func (a *Album) ToResponse() AlbumResponse {
	var album AlbumResponse

	album.ID = a.ID
	album.Title = a.Title
	album.Artist = a.Artist
	album.Price = a.Price
	album.CreatedAt = a.CreatedAt
	album.UpdatedAt = a.UpdatedAt

	return album
}
