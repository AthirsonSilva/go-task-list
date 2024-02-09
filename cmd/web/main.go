package main

import (
	"github.com/AthirsonSilva/music-streaming-api/database"
	"github.com/AthirsonSilva/music-streaming-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	database.Database.Connect()

	server := echo.New()

	albums := server.Group("/api/v1/albums")

	albums.GET("/", routes.FindAll)
	albums.GET("/:id", routes.FindOne)
	albums.POST("/", routes.Create)
	albums.PUT("/:id", routes.Update)
	albums.DELETE(":id", routes.Delete)

	server.Logger.Fatal(server.Start(":8080"))
}
