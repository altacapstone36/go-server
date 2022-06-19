package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"

	"gorm.io/gorm"
)

type outPatientRepository struct {
	sqldb *gorm.DB
}

func NewOutPatientRepository(sqldb *gorm.DB) *outPatientRepository {
	return &outPatientRepository{sqldb: sqldb}
}

func (repo outPatientRepository) NewMedicalRecord(ms models.MedicRecord) (err error) {
	err = repo.sqldb.Create(&ms).Error
	return
}

func (repo outPatientRepository) Proceed(mr models.MedicRecord) (err error) {
	err = repo.sqldb.Updates(&mr).Error
	return
}

func (repo outPatientRepository) ListForDoctor(id int) (res []response.OutPatientResponse, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						medical_staffs.full_name as doctor`).
		Joins("join patients on patients.id = medic_records.patient_id").
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join sessions on medical_sessions.session_id = sessions.id").
		Joins("join medical_staffs on medical_staffs.id = medical_sessions.medical_staff_id").
		Where("medical_sessions.medical_staff_id = ?", id).
		Where("medic_records.status = 0").
		Scan(&res).Error

	return
}

func (repo outPatientRepository) ListForNurse(facility string) (res []response.OutPatientResponse, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						medical_staffs.full_name as doctor`).
		Joins("join patients on patients.id = medic_records.patient_id").
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join sessions on medical_sessions.session_id = sessions.id").
		Joins("join medical_facilities on medical_sessions.medical_facility_id = medical_facilities.id").
		Joins("join medical_staffs on medical_staffs.id = medical_sessions.medical_staff_id").
		Where("medic_records.blood_tension = 0").
		Where("medic_records.body_temperature = 0").
		Where("medic_records.height = 0").
		Where("medic_records.weight = 0").
		Where("medic_records.status = 0").
		Where("medical_facilities.name = ?", facility).
		Scan(&res).Error

	return
}

func (repo outPatientRepository) Report() (res []response.OutPatientReportResponse, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						medical_staffs.full_name as doctor`).
		Joins("join patients on patients.id = medic_records.patient_id").
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join sessions on medical_sessions.session_id = sessions.id").
		Joins("join medical_staffs on medical_staffs.id = medical_sessions.medical_staff_id").
		Where("medic_records.status = 1").
		Scan(&res).Error

	return
}

func (repo outPatientRepository) FilterByDate(date_start, date_end string) (res []response.OutPatientResponse, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						medical_staffs.full_name as doctor`).
		Joins("left join patients on patients.id = medic_records.patient_id").
		Joins("left join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("left join sessions on medical_sessions.session_id = sessions.id").
		Joins("left join medical_staffs on medical_staffs.id = medical_sessions.medical_staff_id").
		Where("medical_sessions.date_check BETWEEN ? AND ?", date_start, date_end).
		Scan(&res).Error

	return
}
