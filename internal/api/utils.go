package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, data Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func BadRequest(w http.ResponseWriter, err error) {
	Error(w, err, http.StatusBadRequest)
}

func ReadBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&v)
}

func PathVar(r *http.Request, order int) string {
	// id := strings.TrimPrefix(r.URL.Path, "/api/v1/"+varName+"/")
	// return id
	path := strings.Split(r.URL.Path, "/")
	lastIndex := len(path) - order
	pathVar := path[lastIndex]
	return pathVar
}

func Param(r *http.Request, param string) string {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	queryParams := u.Query()
	log.Printf("Query params: %v", queryParams)

	if len(queryParams[param]) == 0 {
		log.Printf("Param %s not found", param)
		return ""
	}

	return queryParams["id"][0]
}
