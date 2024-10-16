// /pkg/database/migrations.go
package database

import (
    "log"
    "task_management_api/pkg/models" 
)

// MigrateDatabase performs the database migrations for all models
func MigrateDatabase() {
    // Automatically create/update the users table based on the User model
    if err := DB.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    log.Println("Database migration completed")
}