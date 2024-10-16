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

    // Define routes
    router.POST("/register", handlers.RegisterUser)
    router.POST("/login", handlers.LoginUser)
    router.GET("/projects", handlers.GetProjects)
    router.GET("/tasks", handlers.GetTasks)
    router.POST("/tasks", handlers.CreateTask)

    return router
}