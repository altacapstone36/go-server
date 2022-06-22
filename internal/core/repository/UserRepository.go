package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type UserRepository interface {
	FindAll() ([]response.User, error)
	FindByID(int) (response.User, error)
	Create(models.User) (error)
	Update(models.User) (error)
	Delete(int) (error)
}
