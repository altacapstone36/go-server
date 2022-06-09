package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils/errors"
)

type PatientService struct {
	repo repository.PatientRepository
}

func NewPatientService (repo repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (srv PatientService) GetAllPatient() (patient []models.Patient, err error) {
	patient, err = srv.repo.GetAllPatient()
	err = errors.CheckError(patient, err)
	return
}

func (srv PatientService) GetPatientByID(id uint) (patient models.Patient, err error) {
	patient, err = srv.repo.GetPatientByID(id)
	err = errors.CheckError(patient, err)
	return
}

func (srv PatientService) GetPatientByName(name string) (patient []models.Patient, err error) {
	patient, err = srv.repo.GetPatientByName(name)
	err = errors.CheckError(patient, err)
	return
}

func (srv PatientService) CreatePatient(patient models.Patient) (err error) {
	if (models.Patient{}) == patient {
		err = errors.New(404, "error no data for create provided")
		return
	}

	err = srv.repo.CreatePatient(patient)
	err = errors.CheckError(nil, err)
	return
}

func (srv PatientService) UpdatePatient(id uint, patient models.Patient) (err error) {
	if (models.Patient{}) == patient {
		err = errors.New(404, "error no data for update provided")
		return
	}

	patient.ID = id
	err = srv.repo.UpdatePatient(patient)
	err = errors.CheckError(patient, err)
	return
}

func (srv PatientService) DeletePatient(id uint) (err error) {
	err = srv.repo.DeletePatientByID(id)
	err = errors.CheckError(nil, err)
	return
}
