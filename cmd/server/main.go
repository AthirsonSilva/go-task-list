package main

import (
	"context"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"log"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/database"
	_ "github.com/AthirsonSilva/music-streaming-api/cmd/server/docs"
	handlers "github.com/AthirsonSilva/music-streaming-api/cmd/server/handlers/users"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/routes"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const (
	PORT = ":8080"
)

// @title			Task list Swagger API
// @version		1.0
// @author			Athirson Silva
// @description	Swagger API for Golang Project Music Streaming API
// @termsOfService	http://swagger.io/terms/
// @license.name	MIT
// @license.url	https://github.com/AthrsonsSilva/music-streaming-api
// @BasePath		/
func main() {
	database.Database.Connect()
	defer database.Database.Client.Disconnect(context.TODO())

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("main", "Error loading .env file => "+err.Error())
	} else {
		logger.Info("main", "Loaded .env file")
	}

	defer close(handlers.EmailChannel)
	go handlers.ListenForEmail()

	server := &http.Server{
		Addr:    PORT,
		Handler: routes.Routes(),
	}

	logger.Info("main", "Server running on port http://localhost"+PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
