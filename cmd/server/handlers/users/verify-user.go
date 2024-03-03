package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// VerifyUser @Summary Verifies a user
//
//	@Tags		users
//	@Produce	json
//	@Success	200		{object}	api.Response
//	@Failure	500		{object}	api.Exception
//	@Failure	400		{object}	api.Exception
//	@Failure	429		{object}	api.Exception
//	@Param		token	query		string	false	"token"
//	@Router		/api/v1/users/verify [get]
func VerifyUser(res http.ResponseWriter, req *http.Request) {
	token := api.Param(req, "token")
	if token == "" {
		api.Error(res, req, "No token provided", errors.New("no token provided"), http.StatusBadRequest)
		return
	}

	user, err := repositories.FindUserById(token)
	if err != nil {
		api.Error(res, req, "Invalid token", err, http.StatusBadRequest)
		return
	}

	user.Enabled = true
	user, err = repositories.UpdateUserByID(user)
	if err != nil {
		api.Error(res, req, "Error while verifying user", err, http.StatusInternalServerError)
		return
	}

	api.JSON(res, api.Response{
		Message: "User verified successfully, you can now login",
		Data:    user,
	}, http.StatusOK)
}
