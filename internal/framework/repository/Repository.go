package repository

import (
	"go-hospital-server/internal/core/repository"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *repository.Repository {
	return &repository.Repository{
		Auth: NewAuthRepository(db),
	}
}
