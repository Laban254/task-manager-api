package handlers

import (
    "net/http"
    "time"
    "errors"

    "github.com/gin-gonic/gin"
    "task_management_api/pkg/database"
    "task_management_api/pkg/models"
    "gorm.io/gorm"
)

type BaseResponse struct {
    ID        uint      `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type ProjectResponse struct {
    BaseResponse
    Name   string `json:"name"`
    UserID uint   `json:"user_id"`
    Description string `json:"description"`
}

func BuildProjectResponse(project models.Project) ProjectResponse {
    return ProjectResponse{
        BaseResponse: BaseResponse{
            ID:        project.ID,
            CreatedAt: project.CreatedAt,
            UpdatedAt: project.UpdatedAt,
        },
        Name:   project.Name,
        UserID: project.UserID,
        Description: project.Description,
    }
}

func CreateProject(c *gin.Context) {
    var project models.Project
    if err := c.ShouldBindJSON(&project); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    project.UserID = userID.(uint)

    if err := database.DB.Create(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
        return
    }

    response := BuildProjectResponse(project)
    c.JSON(http.StatusOK, response)
}

func GetProjects(c *gin.Context) {
    userID, _ := c.Get("userID")
    var projects []models.Project

    if err := database.DB.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects"})
        return
    }

    response := make([]ProjectResponse, len(projects))
    for i, project := range projects {
        response[i] = BuildProjectResponse(project)
    }

    c.JSON(http.StatusOK, response)
}

func UpdateProject(c *gin.Context) {
    var updatedProject models.Project

    if err := c.ShouldBindJSON(&updatedProject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    userIDUint, ok := userID.(uint) 
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
        return
    }

    var existingProject models.Project
    if err := database.DB.First(&existingProject, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        return
    }

    if existingProject.UserID != userIDUint {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this project"})
        return
    }

    if updatedProject.Name != "" {
        existingProject.Name = updatedProject.Name
    }
    if updatedProject.Description != "" {
        existingProject.Description = updatedProject.Description
    }

    if err := database.DB.Save(&existingProject).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
        return
    }

    response := BuildProjectResponse(existingProject)
    c.JSON(http.StatusOK, response)
}

func DeleteProject(c *gin.Context) {
    var project models.Project

    if err := database.DB.First(&project, c.Param("id")).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve project"})
        return
    }

    if err := database.DB.Delete(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
