package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"net/http"
	"os"
)

// ExportToCsv @Summary Export all user's tasks to csv
//
//	@Tags		tasks
//	@Accept		application/json
//	@Produce	application/json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		Authorization	header		string	true	"Authorization"
//	@Router		/api/v1/tasks/export-csv [get]
func ExportToCsv(res http.ResponseWriter, req *http.Request) {
	f, e := os.Create("./tasks.csv")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()

	pagination := api.Pagination{
		PageNumber:    1,
		PageSize:      99999,
		SortDirection: 1,
		SortField:     "title",
		SearchName:    "",
	}

	user, err := authentication.GetUserFromToken(req)
	if err != nil {
		api.Error(res, req, "You need to be logged in to export tasks", err, http.StatusInternalServerError)
		return
	}

	tasks, err := repositories.FindAllTasks(pagination, user.ID)
	if err != nil {
		api.Error(res, req, "Error while getting tasks from database", err, http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		api.Error(res, req, "No tasks found", err, http.StatusNotFound)
		return
	}

	writer := csv.NewWriter(f)
	var data = make([][]string, len(tasks))

	for i, task := range tasks {
		data[i] = []string{task.Title, task.Description, task.EndDate.String()}
	}

	e = writer.WriteAll(data)
	if e != nil {
		api.Error(res, req, "Error while exporting tasks", e, http.StatusInternalServerError)
		return
	}

	fileBytes, err := os.ReadFile("./tasks.csv")
	if err != nil {
		api.Error(res, req, "Error while getting tasks from database", err, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", "attachment; filename=export.csv")
	res.Write(fileBytes)
	res.WriteHeader(http.StatusOK)

	writer.Flush()
	err = os.Remove("./tasks.csv")
	if err != nil {
		return
	}
}
