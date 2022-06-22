package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"

	"gorm.io/gorm"
)

type patientRepository struct {
	sqldb *gorm.DB
}

func NewPatientRepository(sqldb *gorm.DB) *patientRepository {
	return &patientRepository{sqldb: sqldb}
}

func (repo patientRepository) GetPatientByID(id uint) (patient response.PatientDetails, err error) {
	err = repo.sqldb.Model(models.Patient{}).
		Where("id = ?", id).
		Scan(&patient).
		Error

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_facilities.name as facility, medical_sessions.date_check,
						users.full_name as doctor`).
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join users on users.id = medical_sessions.user_id").
		Joins("join medical_facilities on medical_facilities.id = medical_sessions.medical_facility_id").
		Where("medic_records.patient_id = ?", id).
		Scan(&patient.MedicRecord).
		Error

	return
}

func (repo patientRepository) GetPatientByName(name string) (patient []response.Patient, err error) {
	name = "%" + name + "%"
	err = repo.sqldb.Model(models.Patient{}).
		Where("full_name = ?", name).
		Scan(&patient).
		Error

	return
}

func (repo patientRepository) GetAllPatient() (patient []response.Patient, err error) {
	err = repo.sqldb.Model(models.Patient{}).
		Scan(&patient).
		Error

	return
}

func (repo patientRepository) CreatePatient(patient models.Patient) (err error) {
	err = repo.sqldb.Create(&patient).Error

	return
}

func (repo patientRepository) UpdatePatient(patient models.Patient) (err error) {
	err = repo.sqldb.Updates(&patient).Error

	return
}

func (repo patientRepository) DeletePatientByID(id uint) (err error) {
	del := repo.sqldb.Delete(models.Patient{}, id)
	if del.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}
