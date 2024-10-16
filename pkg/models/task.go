// /pkg/models/task.go
package models

// Task represents a task in the task management system
type Task struct {
    ID          string `json:"id" binding:"required"`
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
    Status      string `json:"status" binding:"required"` // e.g., "To-do", "In-progress", "Completed"
}
