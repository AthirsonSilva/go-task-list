package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// UpdateAlbumById @Summary Find all albums
//
//	@Tags		albums
//	@Accept		application/json
//	@Produce	application/json
//	@Param		album			body		models.AlbumRequest	true	"Album request"
//	@Param		id				path		string				true	"Album ID"
//	@Param		Authorization	header		string				true	"Authorization"
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Router		/api/v1/albums/{id} [put]
func UpdateAlbumById(res http.ResponseWriter, req *http.Request) {
	var request models.AlbumRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		api.Error(res, req, "Malformed request", err, http.StatusBadRequest)
		return
	}

	id := api.PathVar(req, 1)
	if id == "" {
		api.Error(res, req, "ID is required", errors.New("ID is required"), http.StatusBadRequest)
		return
	}

	album := request.ToModel()
	album, err := repositories.UpdateAlbumById(id, album)
	if err != nil {
		api.Error(res, req, "Error while updating album", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album updated",
		Data:    album,
	}
	api.JSON(res, response, http.StatusOK)
}
