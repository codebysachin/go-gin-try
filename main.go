package main

import (
	"codebysachin.com/go-gin-try/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRouter(router)
	router.Run("localhost:8080")
}
