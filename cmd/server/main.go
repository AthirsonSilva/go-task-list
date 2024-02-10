package main

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/internal/database"
	"github.com/AthirsonSilva/music-streaming-api/internal/routes"
)

const port = ":8080"

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
