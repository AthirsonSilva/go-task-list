package users

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
)

// @Summary Find one user by ID
// @Tags users
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Param id path string true "User ID"
// @Router /api/v1/users/{id} [get]
func FindOneUserById(res http.ResponseWriter, req *http.Request) {
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

	user, err := repositories.FindUserById(id)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	userResponse := user.ToResponse()
	response = api.Response{
		Message: "User found",
		Data:    userResponse,
	}
	api.JSON(res, response, http.StatusOK)
}
