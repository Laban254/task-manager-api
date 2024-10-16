// /pkg/database/database.go
package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "task_management_api/config" 
)

var DB *gorm.DB

// ConnectDatabase establishes a connection to the PostgreSQL database
func ConnectDatabase() {
    // Load configuration values from the config package
    cfg := config.LoadConfig()

    // Construct the Data Source Name (DSN) using environment variables
    dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"

    // Establish the database connection
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connection established")

    // Perform database migrations
    MigrateDatabase() // Call to the migration function
}