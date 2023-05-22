package routes

import (
	"context"
	"net/http"

	"go-mongo/database"
	"go-mongo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAll(c *gin.Context) {
	var albums []models.Album

	// Get all albums from the database
	cursor, err := database.Collection.Find(context.Background(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while fetching all albums"})
	}

	// Iterate through the returned cursor
	for cursor.Next(context.Background()) {
		var album models.Album

		cursor.Decode(&album)

		albums = append(albums, album)
	}

	// Return all albums
	c.JSON(http.StatusOK, albums)
}
