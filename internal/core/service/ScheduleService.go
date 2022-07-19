package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
)

type ScheduleService struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (srv ScheduleService) Create(req request.Schedule) (err error) {
	m, _ := utils.TypeConverter[models.Schedule](req)
	err = srv.repo.Create(m)
	return
}

func (srv ScheduleService) Update(req request.Schedule) (err error) {
	m, _ := utils.TypeConverter[models.Schedule](req)
	m.SessionID = uint(req.SessionID)
	err = srv.repo.Update(m)
	return
}

func (srv ScheduleService) Delete(id int) (err error) {
	err = srv.repo.Delete(id)
	return
}
