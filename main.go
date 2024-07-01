package main

import (
	"github.com/gin-gonic/gin"

	"codebysachin.com/go-gin-try/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)
	router.Run("localhost:8080")
}
