package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type AuthRepository interface {
	Login(string) (response.User, error)
	Register(models.User) error
	FindEmail(string) error
	ChangePassword(string) (response.User, error)
	SaveToken(models.Token) error
	UpdateToken(models.Token, models.Token) error
	RevokeToken(string) error
}
