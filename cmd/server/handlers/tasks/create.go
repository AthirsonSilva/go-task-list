package handlers

import (
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// CreateTask @Summary Creates a task
//
//	@Tags		tasks
//	@Accept		application/json
//	@Produce	application/json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		Authorization	header		string				true	"Authorization"
//	@Param		task			body		models.TaskRequest	true	"Task request"
//	@Router		/api/v1/tasks [post]
func CreateTask(res http.ResponseWriter, req *http.Request) {
	var request models.TaskRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		api.Error(res, req, "Malformed request", err, http.StatusBadRequest)
		return
	}

	task := request.ToModel()
	task.Finished = false

	token, err := api.AuthToken(req)
	if err != nil {
		api.Error(res, req, "Error while creating task", err, http.StatusUnauthorized)
		return
	}

	claims, err := authentication.GetTokenInfo(token)
	if err != nil {
		api.Error(res, req, "Error while creating task", err, http.StatusUnauthorized)
		return
	}

	user, err := repositories.FindUserByEmail(claims.Username)
	if err != nil {
		api.Error(res, req, "Error while creating task", err, http.StatusInternalServerError)
		return
	}

	task.User = user
	task, err = repositories.CreateTask(task)
	if err != nil {
		api.Error(res, req, "Error while creating task", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Task created",
		Data:    task,
	}
	api.JSON(res, response, http.StatusCreated)
}
