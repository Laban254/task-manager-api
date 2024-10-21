// /pkg/middleware/auth.go
package middleware

import (
    "net/http"
    "strings"
    
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "task_management_api/config"
    "task_management_api/pkg/models"
    "task_management_api/pkg/database"
)

// AuthMiddleware checks if the user is authenticated and has the required role
func AuthMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        parts := strings.Split(tokenString, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }
        tokenString = parts[1]

        cfg := config.LoadConfig()

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorSignatureInvalid)
            }
            return []byte(cfg.JWTSecret), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        username, exists := claims["username"].(string)
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in token"})
            c.Abort()
            return
        }

        var user models.User
        if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        c.Set("userID", user.ID)
        role, exists := claims["role"].(string)
        if !exists || !strings.EqualFold(role, requiredRole) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: insufficient permissions"})
            c.Abort()
            return
        }

        c.Next()
    }
}