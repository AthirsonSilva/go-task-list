package main

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/database"
	_ "github.com/AthirsonSilva/music-streaming-api/cmd/server/docs"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/routes"
)

const port = ":8080"

// @title Music Streaming Swagger API
// @version 1.0
// @author Athirson Silva
// @description Swagger API for Golang Project Music Streaming API
// @termsOfService http://swagger.io/terms/

// @license.name MIT
// @license.url https://github.com/AthrsonsSilva/music-streaming-api

// @BasePath /
func main() {
	database.Database.Connect()

	server := &http.Server{
		Addr:    port,
		Handler: routes.Routes(),
	}

	log.Printf("Server running on port %s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
