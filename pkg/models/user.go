package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string    `json:"username" gorm:"unique;not null"`
    Password string    `json:"password" gorm:"not null"`
    Role     string    `json:"role" gorm:"default:'user'"`
    Projects []Project  `json:"projects" gorm:"foreignKey:UserID"`
}