package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// @Summary Creates an album
// @Tags albums
// @Accept  application/json
// @Produce  application/json
// @Param album body models.AlbumRequest true "Album request"
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Router /api/v1/albums [post]
func Create(res http.ResponseWriter, req *http.Request) {
	var request models.AlbumRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	album := request.ToModel()
	album, err := repositories.Create(album)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album created",
		Data:    album,
	}
	api.JSON(res, response, http.StatusCreated)
}
