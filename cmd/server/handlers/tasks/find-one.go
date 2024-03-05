package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// FindOneTaskById @Summary Find one task by ID
//
//	@Tags		tasks
//	@Produce	json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		id				path		string	true	"Task ID"
//	@Param		Authorization	header		string	true	"Authorization"
//	@Router		/api/v1/tasks/{id} [get]
func FindOneTaskById(res http.ResponseWriter, req *http.Request) {
	id := api.PathVar(req, 1)
	var response api.Response

	if id == "" {
		logger.Error("FindOneTaskById", "ID is required")
		api.Error(res, req, "ID is required", errors.New("ID is required"), http.StatusBadRequest)
		return
	}

	task, err := repositories.FindTaskById(id)
	if err != nil {
		logger.Error("FindOneTaskById", err.Error())
		api.Error(res, req, "Error while finding task", err, http.StatusInternalServerError)
		return
	}

	taskResponse := task.ToResponse()
	response = api.Response{
		Message: "Task found",
		Data:    taskResponse,
	}
	api.JSON(res, response, http.StatusOK)
}
