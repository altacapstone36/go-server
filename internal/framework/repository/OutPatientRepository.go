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

func (repo outPatientRepository) ProceedDoctor(mr models.MedicRecord) (err error) {
	up := repo.sqldb.Updates(&mr)
	if up.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (repo outPatientRepository) ProceedNurse(mr models.MedicCheck) (err error) {
	up := repo.sqldb.Updates(&mr)
	if up.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (repo outPatientRepository) DoctorFindAll(id int) (res []response.OutPatient, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						users.full_name as doctor`).
		Joins("join patients on patients.id = medic_records.patient_id").
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join sessions on medical_sessions.session_id = sessions.id").
		Joins("join users on users.id = medical_sessions.user_id").
		Where("medical_sessions.user_id = ?", id).
		Where("medic_records.status = 0").
		Scan(&res).Error

	return
}

func (repo outPatientRepository) NurseFindAll(id int) (res []response.OutPatient, err error) {

	err = repo.sqldb.Debug().Model(models.MedicCheck{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						users.full_name as doctor`).
		Joins("join medic_records on medic_records.id = medic_checks.medic_record_id").
		Joins("join patients on patients.id = medic_records.patient_id").
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join sessions on medical_sessions.session_id = sessions.id").
		Joins("join users on users.id = medic_checks.user_id").
		Where("medic_checks.user_id = ?", id).
		Where("medic_checks.status = 0").
		Scan(&res).Error

	return
}

func (repo outPatientRepository) Report() (res []response.OutPatientReport, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`patients.resident_registration, medic_records.*, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						users.full_name as doctor, medical_facilities.name as facility`).
		Joins("join patients on patients.id = medic_records.patient_id").
		Joins("join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("join sessions on medical_sessions.session_id = sessions.id").
		Joins("join medical_facilities on medical_facilities.id = medical_sessions.medical_facility_id").
		Joins("join users on users.id = medical_sessions.user_id").
		Where("medic_records.status = 1").
		Scan(&res).Error

	return
}

func (repo outPatientRepository) ReportLog(id int, role string) (res []response.OutPatientReportLog, err error) {
	var where, join string
	if role == "doctor" {
		join = "join medical_sessions on medical_sessions.medic_record_id = medic_records.id"
		where = "medic_records.status = 1 AND medical_sessions.user_id = ?"
	} else if role == "nurse" {
		where = "medic_checks.status = 1 AND medic_checks.user_id = ?"
		join = "join medic_checks on medic_checks.medic_record_id = medic_records.id"
	}

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select("serial_number").
		Joins(join).
		Where(where, id).
		Scan(&res).Error

	return
}

func (repo outPatientRepository) FindByDate(date_start, date_end string) (res []response.OutPatient, err error) {

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

func (repo outPatientRepository) FindByID(id int) (res []response.OutPatientSimple, err error) {

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, patients.full_name, patients.code`).
		Joins("left join patients on patients.id = medic_records.patient_id").
		Where("medic_records.id = ?", id).
		Scan(&res).Error

	return
}


func (repo outPatientRepository) AssignNurse(mc models.MedicCheck) (err error) {
	err = repo.sqldb.Create(&mc).Error
	return
}
