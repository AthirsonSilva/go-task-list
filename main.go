package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/albums")
	{
		v1.GET("/", findAll)
		v1.GET("/:id", findOne)
		v1.POST("/", create)
		v1.PUT("/:id", update)
		v1.DELETE("/:id", delete)
	}

	router.Run("localhost:8080")
}

func findAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func findOne(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func create(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func update(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var updateAlbum album

			if err := c.BindJSON(&updateAlbum); err != nil {
				return
			}

			albums[i] = updateAlbum

			c.IndentedJSON(http.StatusOK, updateAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func delete(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)

			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
