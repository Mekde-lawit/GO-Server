package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Health check (public)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 group
	api := router.Group("/api/v1")
	{
		// Setup all route groups
		// SetupAuthRoutes(api)
		SetupPostRoutes(api)
		// SetupUserRoutes(api)
		// SetupAdminRoutes(api)
	}
}