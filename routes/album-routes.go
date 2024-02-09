package routes

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/models"
	"github.com/AthirsonSilva/music-streaming-api/repositories"
	"github.com/labstack/echo/v4"
)

func FindAll(c echo.Context) error {
	var albums []models.Album

	// Get all albums from the database
	albums, err := repositories.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Return all albums
	return c.JSON(http.StatusOK, albums)
}

func FindOne(c echo.Context) error {
	var album models.Album

	// Get album from the database
	album, err := repositories.FindById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Return album
	return c.JSON(http.StatusOK, album)
}

func Create(c echo.Context) error {
	var album models.Album

	// Bind the body to the album variable
	if err := c.Bind(&album); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Create album in the database
	album, err := repositories.Create(album)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Return album
	return c.JSON(http.StatusOK, album)
}

func Update(c echo.Context) error {
	var album models.Album

	// Bind the body to the album variable
	if err := c.Bind(&album); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Set ID
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
	}

	// Update album in the database
	album, err := repositories.Update(id, album)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Return album
	return c.JSON(http.StatusOK, album)
}

func Delete(c echo.Context) error {
	// Set ID
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
	}

	// Delete album from the database
	err := repositories.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Return message
	return c.JSON(http.StatusNoContent, "")
}
