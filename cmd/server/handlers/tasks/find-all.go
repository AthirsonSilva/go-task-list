package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// FindAllTasks @Summary Find all tasks
//
//	@Tags		tasks
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
//	@Router		/api/v1/tasks [get]
func FindAllTasks(res http.ResponseWriter, req *http.Request) {
	var response api.Response

	paginationInfo, errorData := api.GetPaginationInfo(req)
	if errorData.Message != "" {
		logger.Error("FindAllTasks", errorData.Message)
		api.Error(res, req, errorData.Message, errors.New(errorData.Message), http.StatusBadRequest)
		return
	}

	user, err := authentication.GetUserFromToken(req)
	if err != nil {
		logger.Error("FindAllTasks", err.Error())
		api.Error(res, req, "You need to be logged in to get tasks", err, http.StatusInternalServerError)
		return
	}

	tasks, err := repositories.FindAllTasks(paginationInfo, user.ID)
	if err != nil {
		logger.Error("FindAllTasks", err.Error())
		api.Error(res, req, "Error while finding tasks", err, http.StatusInternalServerError)
		return
	}

	var tasksResponse []models.TaskResponse
	for _, task := range tasks {
		response := task.ToResponse()
		tasksResponse = append(tasksResponse, response)
	}

	response = api.Response{
		Message: "Tasks found",
		Data:    tasksResponse,
	}
	api.JSON(res, response, http.StatusOK)
}
