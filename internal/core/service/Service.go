package service

import "go-hospital-server/internal/core/repository"


type Service struct {
	Auth *AuthService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(r.Auth),
	}
}
