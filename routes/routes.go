package routes

import (
	"codebysachin.com/go-gin-try/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)
}
