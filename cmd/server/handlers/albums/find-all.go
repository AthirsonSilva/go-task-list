package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
)

// @Summary Find all albums
// @Tags albums
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Param Authorization header string true "Authorization"
// @Param page query string false "page default" default 1
// @Param size query string false "size default" default 10
// @Param field query string false "field default" default created_at
// @Param direction query string false "direction default" default desc
// @Router /api/v1/albums [get]
func FindAllAlbums(res http.ResponseWriter, req *http.Request) {
	var response api.Response

	paginationInfo, errorData := api.GetPaginationInfo(req)
	if errorData.Message != "" {
		api.JSON(res, errorData, http.StatusBadRequest)
		return
	}

	albums, err := repositories.FindAllAlbums(paginationInfo)
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
