package main

import (
	"context"
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
		log.Printf("Error loading .env file => %s", err)
	} else {
		log.Println("Loaded .env file")
	}

	defer close(handlers.EmailChannel)
	go handlers.ListenForEmail()

	server := &http.Server{
		Addr:    PORT,
		Handler: routes.Routes(),
	}

	log.Printf("Server running on port %s", PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
