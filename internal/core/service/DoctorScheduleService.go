package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
)

type DoctorScheduleService struct {
	repo repository.DoctorScheduleRepository
}

func NewDoctorScheduleService(repo repository.DoctorScheduleRepository) *DoctorScheduleService {
	return &DoctorScheduleService{repo: repo}
}

func (srv DoctorScheduleService) FindAll() (schedule []response.ScheduleList, err error) {
	schedule, err = srv.repo.FindAll()
	return
}

func (srv DoctorScheduleService) Create(req request.DoctorSchedule) (err error) {
	m, _ := utils.TypeConverter[models.DoctorSchedule](req)
	err = srv.repo.Create(m)
	return
}

func (srv DoctorScheduleService) Update(req request.DoctorSchedule) (err error) {
	m, _ := utils.TypeConverter[models.DoctorSchedule](req)
	m.SessionID = uint(req.SessionID)
	err = srv.repo.Update(m)
	return
}

func (srv DoctorScheduleService) Delete(id int) (err error) {
	err = srv.repo.Delete(id)
	return
}
