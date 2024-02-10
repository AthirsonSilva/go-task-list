package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// @Summary Find one album by ID
// @Tags albums
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Param id path string true "Album ID"
// @Router /api/v1/albums/{id} [get]
func FindOne(res http.ResponseWriter, req *http.Request) {
	id := api.PathVar(req, 1)
	var response api.Response

	if id == "" {
		response = api.Response{
			Message: "ID is required",
			Data:    nil,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	album, err := repositories.FindById(id)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	albumResponse := album.ToResponse()
	response = api.Response{
		Message: "Album found",
		Data:    albumResponse,
	}
	api.JSON(res, response, http.StatusOK)
}
