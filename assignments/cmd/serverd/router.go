package main

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and defines routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			// Check database connection
			err := db.Ping()
			if err != nil {
				c.JSON(500, gin.H{
					"status":  "error",
					"message": "Database connection error",
				})
				return
			}

			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Service is healthy",
			})
		})

		// Add more API routes here
	}

	return router
}
