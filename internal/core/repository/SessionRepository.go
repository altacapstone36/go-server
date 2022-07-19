package repository

import (
	"go-hospital-server/internal/core/entity/response"
)

type SessionRepository interface {
	FindAll() ([]response.Session, error)
	FindByDateID(int, string) (response.SessionDetails, error)
}
