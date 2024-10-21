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
    router.Use(middleware.NewRateLimiter())

    // Auth routes (not protected)
    setupAuthRoutes(router)

    // Protected routes for authenticated users
    protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.AuthMiddleware("user"))
    setupProtectedRoutes(protectedRoutes)

    // Admin routes (restricted to admin users)
    adminRoutes := router.Group("/admin")
    adminRoutes.Use(middleware.AuthMiddleware("admin"))
    setupAdminRoutes(adminRoutes)

    return router
}

func setupAuthRoutes(router *gin.Engine) {
    router.POST("/auth/register", handlers.RegisterUser)
    router.POST("/auth/login", handlers.LoginUser)
    router.GET("/auth/google/login", handlers.GoogleLogin)
    router.GET("/auth/google/callback", handlers.GoogleCallback)
}

func setupProtectedRoutes(router *gin.RouterGroup) {
    projectRoutes := router.Group("/projects")
    {
        projectRoutes.POST("/", handlers.CreateProject)
        projectRoutes.GET("/", handlers.GetProjects)
        projectRoutes.PUT("/:id", handlers.UpdateProject)
        projectRoutes.DELETE("/:id", handlers.DeleteProject)
    }

    taskRoutes := router.Group("/tasks")
    {
        taskRoutes.POST("/", handlers.CreateTask)
        taskRoutes.GET("/:projectID", handlers.GetTasks)
        taskRoutes.PUT("/:id", handlers.UpdateTask)
        taskRoutes.DELETE("/:id", handlers.DeleteTask)
    }
}

func setupAdminRoutes(router *gin.RouterGroup) {
    router.POST("/users/register", handlers.AdminRegisterUser)
}
