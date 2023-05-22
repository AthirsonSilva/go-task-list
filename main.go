package main

import (
	"go-mongo/database"
	"go-mongo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Database.Connect()

	router := gin.Default()

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
