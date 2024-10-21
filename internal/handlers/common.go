// /internal/handlers/common.go
package handlers

import (
	"time"

    "github.com/go-playground/validator/v10"
	"task_management_api/config"
)

var validate = validator.New()
var cfg = config.LoadConfig()

type BaseResponse struct {
    ID        uint      `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}