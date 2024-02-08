package main

import (
	"go-mongo/database"
	"go-mongo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Database.Connect()

	router := echo.New()

	v1 := router.Group("/api/v1/albums")
	{
		v1.GET("/", routes.FindAll)
		v1.GET("/:id", routes.FindOne)
		v1.POST("/", routes.Create)
		v1.PUT("/:id", routes.Update)
		v1.DELETE("/:id", routes.Delete)
	}

	router.Run("localhost:8000")
}
