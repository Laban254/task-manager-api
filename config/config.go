// /pkg/config/config.go
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config holds the configuration values for the application
type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
    JWTSecret  string
    GOOGLE_CLIENT_ID string
    GOOGLE_CLIENT_SECRET string
    GOOGLE_REDIRECT_URL string
    ADMIN_USERNAME string
    ADMIN_PASSWORD string

}

func LoadConfig() *Config {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBPort:     os.Getenv("DB_PORT"),
        JWTSecret:  os.Getenv("JWT_SECRET"),
        GOOGLE_CLIENT_ID:  os.Getenv("GOOGLE_CLIENT_ID"),
        GOOGLE_CLIENT_SECRET:  os.Getenv("GOOGLE_CLIENT_SECRET"),
        GOOGLE_REDIRECT_URL:  os.Getenv("GOOGLE_REDIRECT_URL"), 
        ADMIN_USERNAME:  os.Getenv("ADMIN_USERNAME"), 
        ADMIN_PASSWORD:   os.Getenv("ADMIN_PASSWORD"), 
    }
}