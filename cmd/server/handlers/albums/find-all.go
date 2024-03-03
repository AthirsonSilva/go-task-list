package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// FindAllAlbums @Summary Find all albums
//
//	@Tags		albums
//	@Produce	json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		Authorization	header		string	true	"Authorization"
//	@Param		page			query		int		false	"page"		default(1)
//	@Param		size			query		int		false	"size"		default(10)
//	@Param		field			query		string	false	"field"		default(created_at)
//	@Param		direction		query		int		false	"direction"	default(-1)
//	@Param		searchName		query		string	false	"searchName"
//	@Router		/api/v1/albums [get]
func FindAllAlbums(res http.ResponseWriter, req *http.Request) {
	var response api.Response

	paginationInfo, errorData := api.GetPaginationInfo(req)
	if errorData.Message != "" {
		api.Error(res, req, errorData.Message, errors.New(errorData.Message), http.StatusBadRequest)
		return
	}

	albums, err := repositories.FindAllAlbums(paginationInfo)
	if err != nil {
		api.Error(res, req, "Error while finding albums", err, http.StatusInternalServerError)
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
