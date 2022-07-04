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
	db := repo.sqldb.Model(models.Patient{}).
		Where("id = ?", id).
		Scan(&patient)

	err = check.DBRecord(db, check.FIND)
	if err != nil {
		return
	}

	err = repo.sqldb.Model(models.MedicalSession{}).
		Select("MedicRecord.*", "medical_sessions.*",
			"MC.blood_tension AS MedicCheck__blood_tension",
			"MC.height AS MedicCheck__height",
			"MC.weight AS MedicCheck__weight",
			"MC.body_temperature AS MedicCheck__body_temperature",
			"U.full_name AS MedicCheck__nurse",
			"User.full_name as doctor", "MedicalFacility.name as facility").
		Joins("MedicRecord").Joins("User").Joins("MedicalFacility").
		Joins("LEFT JOIN medic_checks MC ON MC.medic_record_id = MedicRecord.id").
		Joins("LEFT JOIN users U ON U.code = MC.user_code").
		Where("MedicRecord.patient_code = ? AND MedicRecord.status = 1", patient.Code).
		Find(&patient.MedicRecord).Error

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
	db := repo.sqldb.Debug().Updates(&patient)

	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo patientRepository) DeletePatientByID(id uint) (err error) {
	db := repo.sqldb.Delete(&models.Patient{}, id)

	err = check.DBRecord(db, check.DELETE)
	return
}
