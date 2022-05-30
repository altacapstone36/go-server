package response

import m "go-hospital-server/internal/core/entity/models"

type (
	User struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
		Level m.Level `json:"level"`
		Status int `json:"status"`
	}
)
