// /internal/handlers/common.go
package handlers

import (
    "github.com/go-playground/validator/v10"
	"task_management_api/config"
)

var validate = validator.New()
var cfg = config.LoadConfig()