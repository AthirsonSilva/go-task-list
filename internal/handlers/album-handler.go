package handlers

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/internal/models"
	"github.com/AthirsonSilva/music-streaming-api/internal/repositories"
)

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

func FindOne(res http.ResponseWriter, req *http.Request) {
	log.Println("[FindOne] => Passing by here")
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

func Update(res http.ResponseWriter, req *http.Request) {
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

	id := api.PathVar(req, 1)
	if id == "" {
		response = api.Response{
			Message: "ID is required",
			Data:    nil,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	album := request.ToModel()
	album, err := repositories.Update(id, album)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album updated",
		Data:    album,
	}
	api.JSON(res, response, http.StatusOK)
}

func Delete(res http.ResponseWriter, req *http.Request) {
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

	err := repositories.Delete(id)
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
