// /cmd/main.go
package main

import (
    "task_management_api/pkg/routes"
	"task_management_api/pkg/database"
)

func main() {
	database.ConnectDatabase()
    router := routes.SetupRouter()
    if err := router.Run(":8080"); err != nil {
        panic(err)
    }
}