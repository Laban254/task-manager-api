// /cmd/main.go
package main

import (
    "task_management_api/pkg/routes"
	"task_management_api/pkg/database"
    "task_management_api/internal/handlers"
)

func main() {
	database.ConnectDatabase()
    handlers.SetupAdmin()
    router := routes.SetupRouter()
    if err := router.Run(":8080"); err != nil {
        panic(err)
    }
}