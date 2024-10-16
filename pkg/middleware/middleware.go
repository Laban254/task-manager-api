// /pkg/middleware/middleware.go
package middleware

import (
    "github.com/gin-gonic/gin"
)

// Logger is a middleware for logging requests
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Log the request details (this can be expanded)
        c.Next()
    }
}

// Recovery is a middleware for recovering from panics
func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // Handle the panic
                c.JSON(500, gin.H{"error": "Internal Server Error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}