package main

import (
	"github.com/AthirsonSilva/music-streaming-api/database"
	"github.com/AthirsonSilva/music-streaming-api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	database.Database.Connect()

	server := echo.New()

	albums := server.Group("/api/v1/albums")

	albums.GET("/", handlers.FindAll)
	albums.GET("/:id", handlers.FindOne)
	albums.POST("/", handlers.Create)
	albums.PUT("/:id", handlers.Update)
	albums.DELETE(":id", handlers.Delete)

	server.Logger.Fatal(server.Start(":8080"))
}
