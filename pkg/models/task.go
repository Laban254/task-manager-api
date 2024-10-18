package models

import (
    "gorm.io/gorm"
)

type Task struct {
    gorm.Model
    Title     string `json:"title" gorm:"not null"`
    Status    string `json:"status" gorm:"default:'todo'"`
    ProjectID uint   `json:"project_id"`
    Project   Project `json:"project" gorm:"foreignKey:ProjectID"`
}