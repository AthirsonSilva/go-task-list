package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// @Summary Find all albums
// @Tags albums
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Router /api/v1/albums [get]
func FindAll(res http.ResponseWriter, req *http.Request) {
	albums, err := repositories.FindAll()
	var response api.Response

	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	var albumsResponse []models.AlbumResponse
	for _, album := range albums {
		response := album.ToResponse()
		albumsResponse = append(albumsResponse, response)
	}

	response = api.Response{
		Message: "Albums found",
		Data:    albumsResponse,
	}
	api.JSON(res, response, http.StatusOK)
}
