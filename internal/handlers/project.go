// /internal/handlers/project.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "task_management_api/pkg/database"
    "task_management_api/pkg/models"
)

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

    c.JSON(http.StatusOK, project)
}

func GetProjects(c *gin.Context) {
    userID, _ := c.Get("userID")
    var projects []models.Project
    if err := database.DB.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects"})
        return
    }

    c.JSON(http.StatusOK, projects)
}

func UpdateProject(c *gin.Context) {
    var project models.Project
    if err := c.ShouldBindJSON(&project); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Save(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
        return
    }

    c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
    var project models.Project
    if err := database.DB.Where("id = ?", c.Param("id")).Delete(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}