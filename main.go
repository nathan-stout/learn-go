package main

import (
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup all routes
	routes.SetupRoutes(router)

	router.Run("localhost:8080")
}
