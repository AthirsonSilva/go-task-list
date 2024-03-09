package handlers

import (
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"net/http"
)

// GetLoggedInUser @Summary Check if user is logged in
//
//	@Tags		users
//	@Produce	json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		Authorization	header		string	true	"Authorization"
//	@Router		/api/v1/users/current-user [get]
func GetLoggedInUser(res http.ResponseWriter, req *http.Request) {
	logger.Info("GetLoggedInUser", "Checking if user is logged in")

	user, err := authentication.GetUserFromToken(req)
	var response api.Response

	if err != nil {
		logger.Error("GetLoggedInUser", err.Error())
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
