package service

import (
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
)

type SessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (srv SessionService) FindAllSession() (session []response.Session, err error) {
	session, err = srv.repo.FindAllSession()
	return
}
