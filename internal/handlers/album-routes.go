package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/internal/models"
	"github.com/AthirsonSilva/music-streaming-api/internal/repositories"
	"github.com/labstack/echo/v4"
)

func FindAll(c echo.Context) error {
	albums, err := repositories.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if len(albums) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No albums found"})
	}

	var albumsResponse []models.AlbumResponse
	for _, album := range albums {
		response := album.ToResponse()
		albumsResponse = append(albumsResponse, response)
	}

	return c.JSON(http.StatusOK, albumsResponse)
}

func FindOne(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
	}

	album, err := repositories.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	albumResponse := album.ToResponse()
	return c.JSON(http.StatusOK, albumResponse)
}

func Create(c echo.Context) error {
	var request models.AlbumRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	album := request.ToModel()
	album, err := repositories.Create(album)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, album)
}

func Update(c echo.Context) error {
	var request models.AlbumRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
	}

	album := request.ToModel()
	album, err := repositories.Update(id, album)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, album)
}

func Delete(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID is required"})
	}

	err := repositories.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusNoContent, "")
}
