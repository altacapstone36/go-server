package repository

import (
	"go-hospital-server/internal/core/entity/response"
)

type SessionRepository interface {
	FindAllSession() ([]response.Session, error)
}
