package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
)

// @Summary Deletes an album
// @Tags albums
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Param id path string true "Album ID"
// @Router /api/v1/albums/{id} [delete]
func DeleteAlbumById(res http.ResponseWriter, req *http.Request) {
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

	err := repositories.DeleteAlbumById(id)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album deleted",
		Data:    nil,
	}
	api.JSON(res, response, http.StatusOK)
}
