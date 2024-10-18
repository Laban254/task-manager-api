// /internal/handlers/task.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "task_management_api/pkg/database"
    "task_management_api/pkg/models"
)

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    projectID := c.Param("projectID")

    if err := database.DB.Where("project_id = ?", projectID).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    var task models.Task
    if err := database.DB.Where("id = ?", c.Param("id")).Delete(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
