// /internal/handlers/admin_setup.go
package handlers

import (
    "log"
    "task_management_api/pkg/database"
    "task_management_api/pkg/models"
)

// SetupAdmin checks if the admin user exists and creates one if it does not
func SetupAdmin() {
    var existingUser models.User
    err := database.DB.Where("username = ?", cfg.ADMIN_USERNAME).First(&existingUser).Error
    if err == nil {
        log.Println("Admin user already exists")
        return
    }

    adminUser := models.User{
        Username: cfg.ADMIN_USERNAME,
        Password: cfg.ADMIN_PASSWORD,
        Role:     "admin",
    }

    hashedPassword, err := hashPassword(adminUser.Password)
    if err != nil {
        log.Fatalf("Failed to hash admin password: %v", err)
    }
    adminUser.Password = hashedPassword

    if err := database.DB.Create(&adminUser).Error; err != nil {
        log.Fatalf("Failed to create admin user: %v", err)
    }

    log.Println("Admin user created successfully")
}
