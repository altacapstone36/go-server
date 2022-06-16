package repository

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

type patientRepository struct {
	sqldb *gorm.DB
}

func NewPatientRepository(sqldb *gorm.DB) *patientRepository {
	return &patientRepository{sqldb: sqldb}
}

func (repo patientRepository) GetPatientByID(id uint) (patient models.Patient, err error) {
	err = repo.sqldb.Preload("MedicRecord").First(&patient, id).Error

	return
}

func (repo patientRepository) GetPatientByName(name string) (patient []models.Patient, err error) {
	name = "%" + name + "%"
	err = repo.sqldb.Find(&patient, "full_name LIKE ?", name).Error

	return
}

func (repo patientRepository) GetAllPatient() (patient []models.Patient, err error) {
	err = repo.sqldb.Find(&patient).Error

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
