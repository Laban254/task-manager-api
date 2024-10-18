// /pkg/models/user.go
package models

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique;not null" binding:"required,min=3,max=20"` 
    Password string `json:"password,omitempty" gorm:""`
}