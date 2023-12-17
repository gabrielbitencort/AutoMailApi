package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run server
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
