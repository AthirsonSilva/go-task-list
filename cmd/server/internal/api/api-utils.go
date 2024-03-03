package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Exception struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
}

type Pagination struct {
	PageNumber    int
	PageSize      int
	SortDirection int
	SortField     string
	SearchName    string
}

func JSON(w http.ResponseWriter, data Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		Error(w, nil, err.Error(), err, http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, r *http.Request, errorMessage string, originalError error, status int) {
	log.Printf("Error while processing request: %s", originalError.Error())

	apiError := Exception{
		Message:   errorMessage,
		Path:      r.URL.Path,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(apiError)
	if err != nil {
		Error(w, nil, err.Error(), err, http.StatusInternalServerError)
	}
}

func ReadBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(&v)
}

func PathVar(r *http.Request, order int) string {
	path := strings.Split(r.URL.Path, "/")
	lastIndex := len(path) - order
	pathVar := path[lastIndex]
	return pathVar
}

func Param(r *http.Request, param string) string {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(err)
		return ""
	}

	queryParam := u.Query().Get(param)
	if queryParam == "" {
		err := fmt.Sprintf("param %s not found", param)
		log.Println(err)
		return ""
	}

	log.Printf("Query param: %s => %s", param, queryParam)
	return queryParam
}

func AuthToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("authorization header not found")
	}

	auth = strings.TrimPrefix(auth, "Bearer ")

	if auth == "" {
		return "", errors.New("bearer token not found")
	}

	return auth, nil
}

func GetPaginationInfo(req *http.Request) (Pagination, Response) {
	var pagination Pagination
	var pageNumber, pageSize, sortDirection int
	var sortField, searchName string
	var err error

	pageNumberStr := Param(req, "page")
	if pageNumberStr == "" {
		pageNumberStr = "1"
	}

	pageNumber, err = strconv.Atoi(pageNumberStr)
	if err != nil {
		errorResponse := Response{
			Message: err.Error(),
			Data:    nil,
		}
		return pagination, errorResponse
	}

	pageSizeStr := Param(req, "size")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		errorResponse := Response{
			Message: err.Error(),
			Data:    nil,
		}
		return pagination, errorResponse
	}

	sortField = Param(req, "field")
	if sortField == "" {
		sortField = "created_at"
	}

	sortDirectionStr := Param(req, "direction")
	if sortDirectionStr == "" {
		sortDirectionStr = "-1"
	}

	sortDirection, err = strconv.Atoi(sortDirectionStr)
	if err != nil {
		errorResponse := Response{
			Message: err.Error(),
			Data:    nil,
		}
		return pagination, errorResponse
	}

	searchName = Param(req, "searchName")

	pagination = Pagination{
		PageNumber:    pageNumber,
		PageSize:      pageSize,
		SortField:     sortField,
		SortDirection: sortDirection,
		SearchName:    searchName,
	}
	return pagination, Response{}
}

func FileUpload(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(32 << 20) // 32MB max
	if err != nil {
		return "", err
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// This is path which we want to store the file
	f, err := os.OpenFile("/tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return handler.Filename, nil
}
