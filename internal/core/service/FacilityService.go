package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
)

type FacilityService struct {
	repo repository.FacilityRepository
}

func NewFacilityService(repo repository.FacilityRepository) *FacilityService {
	return &FacilityService{repo: repo}
}

func (srv FacilityService) Create(req request.Facility) (err error) {
	m, _ := utils.TypeConverter[models.MedicalFacility](req)
	err = srv.repo.Create(m)
	return
}

func (srv FacilityService) Update(id int, req request.Facility) (err error) {
	m, _ := utils.TypeConverter[models.MedicalFacility](req)
	m.ID = uint(id)
	err = srv.repo.Update(m)
	return
}

func (srv FacilityService) Delete(id int) (err error) {
	err = srv.repo.Delete(id)
	return
}

func (srv FacilityService) FindAll() (user []response.Facility, err error) {
	user, err = srv.repo.FindAll()
	return
}

func (srv FacilityService) FindByID(id int) (user response.FacilityDetails, err error) {
	user, err = srv.repo.FindByID(id)
	return
}
