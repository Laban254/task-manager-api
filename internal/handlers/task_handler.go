// /internal/handlers/task_handler.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "task_management_api/pkg/models"
)

var tasks []models.Task

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

    // Validate task
    if err := validate.Struct(newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tasks = append(tasks, newTask)
    c.JSON(http.StatusCreated, newTask)
}