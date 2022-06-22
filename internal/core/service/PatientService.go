package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors/check"
)

type PatientService struct {
	repo repository.PatientRepository
}

func NewPatientService (repo repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (srv PatientService) GetAllPatient() (patient []response.Patient, err error) {
	patient, err = srv.repo.GetAllPatient()
	err = check.Record(patient, err)
	return
}

func (srv PatientService) GetPatientByID(id uint) (patient response.PatientDetails, err error) {
	patient, err = srv.repo.GetPatientByID(id)
	err = check.Record(patient, err)
	return
}

func (srv PatientService) GetPatientByName(name string) (patient []response.Patient, err error) {
	patient, err = srv.repo.GetPatientByName(name)
	err = check.Record(patient, err)
	return
}

func (srv PatientService) CreatePatient(patient request.Patient) (err error) {
	p, _ := utils.TypeConverter[models.Patient](patient)
	err = srv.repo.CreatePatient(p)
	err = check.Record(patient, err)
	return
}

func (srv PatientService) UpdatePatient(id uint, patient request.Patient) (err error) {
	p, _ := utils.TypeConverter[models.Patient](patient)
	p.ID = id
	err = srv.repo.UpdatePatient(p)
	err = check.Record(patient, err)
	return
}

func (srv PatientService) DeletePatient(id uint) (err error) {
	err = srv.repo.DeletePatientByID(id)
	err = check.Record(nil, err)
	return
}
