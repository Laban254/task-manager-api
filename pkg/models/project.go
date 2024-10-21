package models

import (
    "gorm.io/gorm"
)

type Project struct {
    gorm.Model
    Name        string  `json:"name" gorm:"not null"`
    Description string  `json:"description" gorm:"not null"` 
    UserID     uint    `json:"user_id"`
    User       User    `json:"user" gorm:"foreignKey:UserID"`
    Tasks      []Task  `json:"tasks" gorm:"foreignKey:ProjectID"`
}