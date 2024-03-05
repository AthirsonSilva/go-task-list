package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// UpdateTaskById @Summary Find all tasks
//
//	@Tags		tasks
//	@Accept		application/json
//	@Produce	application/json
//	@Param		task			body		models.TaskRequest	true	"Task request"
//	@Param		id				path		string				true	"Task ID"
//	@Param		Authorization	header		string				true	"Authorization"
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Router		/api/v1/tasks/{id} [put]
func UpdateTaskById(res http.ResponseWriter, req *http.Request) {
	var request models.TaskRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		logger.Error("UpdateTaskById", err.Error())
		api.Error(res, req, "Malformed request", err, http.StatusBadRequest)
		return
	}

	id := api.PathVar(req, 1)
	if id == "" {
		logger.Error("UpdateTaskById", "ID is required")
		api.Error(res, req, "ID is required", errors.New("ID is required"), http.StatusBadRequest)
		return
	}

	task := request.ToModel()
	task, err := repositories.UpdateTaskById(id, task)
	if err != nil {
		logger.Error("UpdateTaskById", err.Error())
		api.Error(res, req, "Error while updating task", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Task updated",
		Data:    task,
	}
	api.JSON(res, response, http.StatusOK)
}
