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

func (srv SessionService) FindAll() (session []response.Session, err error) {
	session, err = srv.repo.FindAll()
	return
}

func (srv SessionService) FindByDateID(id int, date string) (schedule response.SessionDetails, err error) {
	schedule, err = srv.repo.FindByDateID(id, date)
	return
}

