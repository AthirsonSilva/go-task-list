package handlers

import (
	"errors"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// DeleteTaskById @Summary Deletes a task
//
//	@Tags		tasks
//	@Produce	json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Failure	404				{object}	api.Exception
//	@Param		id				path		string	true	"Task ID"
//	@Param		Authorization	header		string	true	"Authorization"
//	@Router		/api/v1/tasks/{id} [delete]
func DeleteTaskById(res http.ResponseWriter, req *http.Request) {
	id := api.PathVar(req, 1)
	var response api.Response

	if id == "" {
		api.Error(res, req, "ID is required", errors.New("ID is required"), http.StatusBadRequest)
		return
	}

	err := repositories.DeleteTaskById(id)
	if err != nil {
		api.Error(res, req, "Error while deleting task", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Task deleted",
		Data:    "Provided ID: " + id,
	}
	api.JSON(res, response, http.StatusOK)
}
