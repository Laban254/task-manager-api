// /pkg/routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "task_management_api/internal/handlers"
    "task_management_api/pkg/middleware"
)

// SetupRouter initializes the API routes
func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Middleware
    router.Use(middleware.Logger())    
    router.Use(middleware.Recovery())  

    // routes
    router.POST("/auth/register", handlers.RegisterUser)
    router.POST("/auth/login", handlers.LoginUser)
    router.GET("/projects", handlers.GetProjects)
    router.GET("/tasks", handlers.GetTasks)
    router.POST("/tasks", handlers.CreateTask)
    router.GET("/auth/google/login", handlers.GoogleLogin)
    router.GET("/auth/google/callback", handlers.GoogleCallback)

    return router
}