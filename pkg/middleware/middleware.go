// /pkg/middleware/middleware.go
package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

// Logger is a middleware for logging requests
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next() // Call the next handler
        log.Printf("%s %s took %v", c.Request.Method, c.Request.URL, time.Since(start))
    }
}

// Recovery from panics
func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from panic: %v", err)
                c.JSON(500, gin.H{"error": "Internal Server Error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}