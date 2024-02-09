package main

import (
	"github.com/AthirsonSilva/music-streaming-api/database"
	"github.com/AthirsonSilva/music-streaming-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	database.Database.Connect()

	router := echo.New()

	router.GET("/api/v1/albums", routes.FindAll)
	router.GET("/api/v1/albums/:id", routes.FindOne)
	router.POST("/api/v1/albums", routes.Create)
	router.PUT("/api/v1/albums/:id", routes.Update)
	router.DELETE("/api/v1/albums/:id", routes.Delete)

	router.Logger.Fatal(router.Start(":8080"))
}
