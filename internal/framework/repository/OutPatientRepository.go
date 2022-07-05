package repository

import (
	m "go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type outPatientRepository struct {
	sqldb *gorm.DB
}

func NewOutPatientRepository(sqldb *gorm.DB) *outPatientRepository {
	return &outPatientRepository{sqldb: sqldb}
}

func (repo outPatientRepository) NewMedicalRecord(ms m.MedicRecord) (err error) {
	db := repo.sqldb.Create(&ms)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (repo outPatientRepository) ProceedDoctor(mr m.MedicRecord) (err error) {
	db := repo.sqldb.Updates(&mr)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo outPatientRepository) ProceedNurse(mr m.MedicCheck) (err error) {
	db := repo.sqldb.Updates(&mr)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo outPatientRepository) DoctorFindAll(code string) (res []response.OutPatient, err error) {

	db := repo.sqldb.Model(m.MedicalSession{}).
		Select("MedicRecord.*", "medical_sessions.*", "User.full_name as doctor", "Session.time_start as session",
					 "MedicalFacility.name as facility", "P.full_name", "P.code as code").
		Joins("MedicRecord").Joins("User").Joins("MedicalFacility").Joins("Session").
		Joins("JOIN patients P ON MedicRecord.patient_code = P.code").
		Where("medical_sessions.user_code = ? AND MedicRecord.status = 0", code).
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo outPatientRepository) NurseFindAll(code string) (res []response.OutPatient, err error) {

	db := repo.sqldb.Model(m.MedicCheck{}).
		Select("MedicRecord.*", "User.full_name as doctor").
		Joins("MedicRecord").Joins("User").
		Where("medic_checks.user_code = ? AND medic_checks.status = 0", code).
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}


func (repo outPatientRepository) Report() (res []response.OutPatientDetails, err error) {

	db := repo.sqldb.Model(m.MedicalSession{}).
		Select("MedicRecord.*", "medical_sessions.*",
						"P.full_name as patient_name", "P.national_id",
						"MC.blood_tension AS MedicCheck__blood_tension",
						"MC.height AS MedicCheck__height",
						"MC.weight AS MedicCheck__weight",
						"MC.body_temperature AS MedicCheck__body_temperature",
						"U.full_name AS MedicCheck__nurse",
						"User.full_name as doctor", "MedicalFacility.name as facility").
		Joins("MedicRecord").Joins("User").Joins("MedicalFacility").
		Joins("LEFT JOIN medic_checks MC ON MC.medic_record_id = MedicRecord.id").
		Joins("LEFT JOIN patients P ON P.code = MedicRecord.patient_code").
		Joins("LEFT JOIN users U ON U.code = MC.user_code").
		Where("MedicRecord.status = 1").
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo outPatientRepository) ReportLog(code, role string) (res []response.OutPatientReportLog, err error) {
	var where, join string
	if role == "doctor" {
		join = "MedicalSession"
		where = "medic_records.status = 1 AND MedicalSession.user_code = ?"
	} else if role == "nurse" {
		where = "MedicCheck.status = 1 AND MedicCheck.user_code = ?"
		join = "MedicCheck"
	}

	db := repo.sqldb.Model(m.MedicRecord{}).
		Select("serial_number").
		Joins(join).
		Where(where, code).
		Scan(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo outPatientRepository) FindByDate(date_start, date_end string) (res []response.OutPatient, err error) {

	db := repo.sqldb.Model(m.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						users.full_name as doctor`).
		Joins("left join patients on patients.id = medic_records.patient_id").
		Joins("left join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("left join sessions on medical_sessions.session_id = sessions.id").
		Joins("left join users on users.id = medical_sessions.user_id").
		Where("medical_sessions.date_check BETWEEN ? AND ?", date_start, date_end).
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo outPatientRepository) FindByID(id int) (res response.OutPatientDetails, err error) {

	db := repo.sqldb.Model(m.MedicalSession{}).
		Select("MedicRecord.*", "medical_sessions.*",
						"P.full_name as patient_name", "P.national_id",
						"MC.blood_tension AS MedicCheck__blood_tension",
						"MC.height AS MedicCheck__height",
						"MC.weight AS MedicCheck__weight",
						"MC.body_temperature AS MedicCheck__body_temperature",
						"U.full_name AS MedicCheck__nurse",
						"User.full_name as doctor", "MedicalFacility.name as facility").
		Joins("MedicRecord").Joins("User").Joins("MedicalFacility").
		Joins("LEFT JOIN medic_checks MC ON MC.medic_record_id = MedicRecord.id").
		Joins("LEFT JOIN patients P ON P.code = MedicRecord.patient_code").
		Joins("LEFT JOIN users U ON U.code = MC.user_code").
		Where("MedicRecord.id = ?", id).
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}


func (repo outPatientRepository) AssignNurse(mc m.MedicCheck) (err error) {
	err = repo.sqldb.Create(&mc).Error
	return
}
