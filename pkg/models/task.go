// /pkg/models/task.go
package models

// Task represents a task in the task management system
type Task struct {
    ID          string `json:"id" binding:"required"`
    Title       string `json:"title" binding:"required,min=3,max=100"`
    Description string `json:"description" binding:"max=255"` 
    Status      string `json:"status" binding:"required,oneof=To-do In-progress Completed"`
}