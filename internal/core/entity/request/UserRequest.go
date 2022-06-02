package request

import "go-hospital-server/internal/core/entity/models"


type User struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Level models.Level `json:"level"`
	Status int `json:"status"`
}

