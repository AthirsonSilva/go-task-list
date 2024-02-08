package routes

import (
	"go-mongo/models"
	"go-mongo/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FindAll(c echo.Context) {
	var albums []models.Album

	// Get all albums from the database
	albums, err := repositories.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	// Return all albums
	c.JSON(http.StatusOK, albums)
}

func FindOne(c echo.Context) {
	var album models.Album

	// Get album from the database
	album, err := repositories.FindById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	// Return album
	c.JSON(http.StatusOK, album)
}

func Create(c echo.Context) {
	var album models.Album

	// Bind the body to the album variable
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Create album in the database
	album, err := repositories.Create(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	// Return album
	c.JSON(http.StatusOK, album)
}

func Update(c echo.Context) {
	var album models.Album

	// Bind the body to the album variable
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Set ID
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
		return
	}

	// Update album in the database
	album, err := repositories.Update(id, album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	// Return album
	c.JSON(http.StatusOK, album)
}

func Delete(c echo.Context) {
	// Set ID
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
		return
	}

	// Delete album from the database
	err := repositories.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	// Return message
	c.Status(http.StatusNoContent)
}
