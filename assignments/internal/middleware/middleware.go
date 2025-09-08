package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestValidator middleware checks incoming requests for required fields
func RequestValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for required headers
		if userAgent := c.GetHeader("User-Agent"); userAgent == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "User-Agent header is required",
			})
			return
		}

		// Continue to next middleware/handler if checks pass
		c.Next()
	}
}

// ApiKeyAuth checks for valid API key in the request
func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")

		// No API key provided
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key is required",
			})
			return
		}

		// Validate API key (this would typically check against a database)
		if !isValidApiKey(apiKey) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
			})
			return
		}

		// Add key info to context for handlers to use
		c.Set("apiKeyValid", true)

		c.Next()
	}
}

// Helper function to validate API key
func isValidApiKey(key string) bool {
	// In a real application, you would check the key against your database
	// This is just a simple example
	validKeys := []string{"valid-api-key-1", "valid-api-key-2"}

	for _, validKey := range validKeys {
		if key == validKey {
			return true
		}
	}

	return false
}

// RateLimiter middleware limits request frequency
func RateLimiter() gin.HandlerFunc {
	// In a real application, you would use a proper rate limiting library
	// This is a simplified example
	return func(c *gin.Context) {
		// Get client IP
		clientIP := c.ClientIP()

		// Check if rate limit exceeded (implementation omitted)
		if isRateLimitExceeded(clientIP) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Try again later.",
			})
			return
		}

		c.Next()
	}
}

// Helper function for rate limiting
func isRateLimitExceeded(clientIP string) bool {
	// In a real app, you would implement proper rate limiting
	// using Redis or another store to track requests
	return false
}
