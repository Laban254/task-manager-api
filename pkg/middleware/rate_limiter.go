// pkg/middleware/rate_limiter.go
package middleware

import (
    "net/http"
    "time"
	"fmt"

    "github.com/gin-gonic/gin"
    "github.com/ulule/limiter/v3"
    "github.com/ulule/limiter/v3/drivers/store/memory"
)

func NewRateLimiter() gin.HandlerFunc {
    rate, err := limiter.NewRateFromFormatted("10-M")
    if err != nil {
        panic(err)
    }

    store := memory.NewStore()

    rateLimiter := limiter.New(store, rate)

    return gin.HandlerFunc(func(c *gin.Context) {
        // Limit the request
        ctx, err := rateLimiter.Get(c.Request.Context(), c.ClientIP())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limit error"})
            c.Abort()
            return
        }

        // Check if rate limit is exceeded
        if ctx.Reached {
            resetTime := time.Unix(ctx.Reset, 0).Format(time.RFC1123)
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error":     "Too many requests",
                "reset":     resetTime,
                "limit":     ctx.Limit,              
                "remaining": ctx.Remaining,          
            })
            c.Abort()
            return
        }

        c.Header("X-Rate-Limit-Limit", fmt.Sprintf("%d", ctx.Limit))  
        c.Header("X-Rate-Limit-Remaining", fmt.Sprintf("%d", ctx.Remaining))
        c.Header("X-Rate-Limit-Reset", time.Unix(ctx.Reset, 0).Format(time.RFC1123))

        c.Next()
    })
}
