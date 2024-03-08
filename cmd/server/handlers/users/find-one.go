package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// FindOneUserById @Summary Find one user by ID
//
//	@Tags		users
//	@Produce	json
//	@Success	200	{object}	api.Response
//	@Failure	500	{object}	api.Response
//	@Failure	500	{object}	api.Exception
//	@Failure	400	{object}	api.Exception
//	@Failure	429	{object}	api.Exception
//	@Param		id	path		string	true	"User ID"
//	@Router		/api/v1/users/{id} [get]
func FindOneUserById(res http.ResponseWriter, req *http.Request) {
	logger.Info("FindOneUserById", "Finding user by ID")

	id := api.PathVar(req, 1)
	var response api.Response

	if id == "" {
		logger.Error("FindOneUserById", "ID is required")
		api.Error(res, req, "ID is required", errors.New("ID is required"), http.StatusBadRequest)
		return
	}

	user, err := repositories.FindUserById(id)
	if err != nil {
		logger.Error("FindOneUserById", err.Error())
		api.Error(res, req, "Error while finding user", err, http.StatusInternalServerError)
		return
	}

	userResponse := user.ToResponse()
	response = api.Response{
		Message: "User found",
		Data:    userResponse,
	}
	api.JSON(res, response, http.StatusOK)
}
