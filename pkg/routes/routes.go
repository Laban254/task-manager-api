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

    // Auth routes
    router.POST("/auth/register", handlers.RegisterUser)
    router.POST("/auth/login", handlers.LoginUser)
    router.GET("/auth/google/login", handlers.GoogleLogin)
    router.GET("/auth/google/callback", handlers.GoogleCallback)

    // Protected routes for authenticated users
    protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.AuthMiddleware("user"))

    // Project and Task routes (available to authenticated users)
    protectedRoutes.GET("/projects", handlers.GetProjects)
    protectedRoutes.GET("/tasks", handlers.GetTasks)
    protectedRoutes.POST("/tasks", handlers.CreateTask)

    // Admin routes (restricted to admin users)
    adminRoutes := router.Group("/admin")
    adminRoutes.Use(middleware.AuthMiddleware("admin"))

    adminRoutes.POST("/users/register", handlers.AdminRegisterUser)

    return router
}
