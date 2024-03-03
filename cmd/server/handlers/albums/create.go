package handlers

import (
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// CreateAlbum @Summary Creates an album
//
//	@Tags		albums
//	@Accept		application/json
//	@Produce	application/json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		Authorization	header		string				true	"Authorization"
//	@Param		album			formData	models.AlbumRequest	true	"Album request"
//	@Param		file			formData	file				true	"File"
//	@Router		/api/v1/albums [post]
func CreateAlbum(res http.ResponseWriter, req *http.Request) {
	var request models.AlbumRequest
	var response api.Response

	artistInput := req.FormValue("artist")
	titleInput := req.FormValue("title")
	if artistInput == "" || titleInput == "" {
		api.Error(res, req, "Artist or title cannot be empty", nil, http.StatusBadRequest)
		return
	}

	request.Title = titleInput
	request.Artist = artistInput

	album := request.ToModel()
	photoUrl, err := api.FileUpload(req)
	if err != nil {
		api.Error(res, req, "Error while uploading photo", err, http.StatusInternalServerError)
		return
	}

	album.PhotoUrl = photoUrl
	album, err = repositories.CreateAlbum(album)
	if err != nil {
		api.Error(res, req, "Error while creating album", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album created",
		Data:    album,
	}
	api.JSON(res, response, http.StatusCreated)
}
