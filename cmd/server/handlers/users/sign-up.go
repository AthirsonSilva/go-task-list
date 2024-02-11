package users

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
)

// @Summary Creates an user
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param user body models.UserRequest true "User request"
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Router /api/v1/users/signup [post]
func SignUp(res http.ResponseWriter, req *http.Request) {
	var request models.UserRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusBadRequest)
		return
	}

	user := request.ToModel()
	user, err := repositories.CreateUser(user)
	if err != nil {
		response = api.Response{
			Message: err.Error(),
			Data:    nil,
		}
		api.JSON(res, response, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "User created",
		Data:    user,
	}
	api.JSON(res, response, http.StatusCreated)
}
