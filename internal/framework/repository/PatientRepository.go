package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"

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

	db := repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_facilities.name as facility, medical_sessions.date_check,
						users.full_name as doctor`).
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join users on users.id = medical_sessions.user_id").
		Joins("join medical_facilities on medical_facilities.id = medical_sessions.medical_facility_id").
		Where("medic_records.patient_id = ?", id).
		Scan(&patient.MedicRecord)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo patientRepository) GetPatientByName(name string) (patient []response.Patient, err error) {
	name = "%" + name + "%"
	db := repo.sqldb.Model(models.Patient{}).
		Where("full_name LIKE ?", name).
		Scan(&patient)
	
	err = check.DBRecord(db, check.FIND)
	return
}

func (repo patientRepository) GetAllPatient() (patient []response.Patient, err error) {
	db := repo.sqldb.Model(models.Patient{}).
		Scan(&patient)
	
	err = check.DBRecord(db, check.FIND)
	return
}

func (repo patientRepository) CreatePatient(patient models.Patient) (err error) {
	err = repo.sqldb.Create(&patient).Error

	return
}

func (repo patientRepository) UpdatePatient(patient models.Patient) (err error) {
	db := repo.sqldb.Updates(&patient)
	
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo patientRepository) DeletePatientByID(id uint) (err error) {
	db := repo.sqldb.Delete(models.Patient{}, id)
	
	err = check.DBRecord(db, check.DELETE)
	return
}
