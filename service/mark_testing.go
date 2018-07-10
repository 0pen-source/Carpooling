package service

import (
	"github.com/gin-gonic/gin"
)

// MarkTesting returns a middleware that marks the request as testing
func MarkTesting() gin.HandlerFunc {
	return func(c *gin.Context) {
		testing := c.Request.Header.Get("Origin") == "test"

		c.Set("testing", testing)

		c.Next()
	}
}
