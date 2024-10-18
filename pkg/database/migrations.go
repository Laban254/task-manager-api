// /pkg/database/migrations.go
package database

import (
    "log"
    "task_management_api/pkg/models" 
)

// MigrateDatabase performs the database migrations for all models
func MigrateDatabase() {
    if err := DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    log.Println("Database migration completed")
}
