package handlers

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "task_management_api/pkg/database"
    "task_management_api/pkg/models"
)

type TaskResponse struct {
    BaseResponse
    Title       string `json:"title"`
    Status      string `json:"status"`
    ProjectID   uint   `json:"project_id"`
    Description string `json:"description"`
}

func BuildTaskResponse(task models.Task) TaskResponse {
    return TaskResponse{
        BaseResponse: BaseResponse{
            ID:        task.ID,
            CreatedAt: task.CreatedAt,
            UpdatedAt: task.UpdatedAt,
        },
        Title:       task.Title,
        Status:      task.Status,
        ProjectID:   task.ProjectID,
        Description: task.Description,
    }
}

func validateTask(task *models.Task) error {
    if task.Title == "" {
        return errors.New("title is required")
    }
    if task.Status == "" {
        return errors.New("status is required")
    }
    if task.ProjectID == 0 {
        return errors.New("project ID is required")
    }
    return nil
}

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
        return
    }

    if validationErr := validateTask(&task); validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        return
    }

    var project models.Project
    if err := database.DB.First(&project, task.ProjectID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        return
    }

    if err := database.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }

    c.JSON(http.StatusCreated, BuildTaskResponse(task))
}

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    projectID := c.Param("projectID")

    pid, err := strconv.ParseUint(projectID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
        return
    }

    if err := database.DB.Where("project_id = ?", pid).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
        return
    }

    response := make([]TaskResponse, len(tasks))
    for i, task := range tasks {
        response[i] = BuildTaskResponse(task)
    }

    c.JSON(http.StatusOK, response)
}

func GetTask(c *gin.Context) {
    var task models.Task
    taskID := c.Param("id")

    tid, err := strconv.ParseUint(taskID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    if err := database.DB.First(&task, tid).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, BuildTaskResponse(task))
}

func UpdateTask(c *gin.Context) {
    var existingTask models.Task
    taskID := c.Param("id")

    tid, err := strconv.ParseUint(taskID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    if err := database.DB.First(&existingTask, tid).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
        return
    }

    if validationErr := validateTask(&updatedTask); validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        return
    }

    existingTask.Title = updatedTask.Title
    existingTask.Status = updatedTask.Status
    existingTask.Description = updatedTask.Description

    if err := database.DB.Save(&existingTask).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }

    c.JSON(http.StatusOK, BuildTaskResponse(existingTask))
}

func DeleteTask(c *gin.Context) {
    taskID := c.Param("id")

    tid, err := strconv.ParseUint(taskID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    var task models.Task
    if err := database.DB.First(&task, tid).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
        return
    }

    if err := database.DB.Delete(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
