// /internal/handlers/task_handler.go
package handlers

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "task_management_api/pkg/models"
    "task_management_api/pkg/database"
)

// Secret key for JWT signing
var jwtSecret = []byte("your_secret_key") // Change this in production

// In-memory store for tasks (you might want to use a database in production)
var tasks []models.Task

// User Registration
func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Store user in the database
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// User Login
func LoginUser(c *gin.Context) {
    var loginData models.User
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    // Retrieve user from the database
    if err := database.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Check password
    if user.Password != loginData.Password { // Note: Use hashed passwords in production
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT token
    token := generateToken(user.Username)
    c.JSON(http.StatusOK, gin.H{"token": token})
}

// Generate JWT token
func generateToken(username string) string {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString(jwtSecret)
    return tokenString
}

// GetProjects handles fetching all projects
func GetProjects(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"projects": []string{}}) // Placeholder for projects
}

// GetTasks handles fetching all tasks
func GetTasks(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"tasks": tasks}) // Return the tasks
}

// CreateTask handles creating a new task
func CreateTask(c *gin.Context) {
    var newTask models.Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    tasks = append(tasks, newTask) // Add the new task to the tasks slice
    c.JSON(http.StatusCreated, newTask)
}