package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"strings"
	"time"

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

func (repo outPatientRepository) ListAvailable(id int,
	level string) (res []response.OutPatientResponse, err error) {

	var whereAdd []string
	if level == "doctor" {
		whereAdd = []string{"medic_records.diagnose = ''", "medic_records.prescription = ''"}
	} else if level == "nurse" {
		whereAdd = []string{"medic_records.blood_tension = ''", "medic_records.body_templrature = ''",
			"medic_records.height = ''", "medic_records.weight = ''"}
	}

	err = repo.sqldb.Model(models.MedicRecord{}).
		Select(`medic_records.*, medical_sessions.queue, medical_sessions.date_check,
						patients.full_name, patients.code, sessions.time_start,
						medical_staffs.full_name as doctor`).
		Joins("left join patients on patients.id = medic_records.patient_id").
		Joins("left join medical_sessions on medical_sessions.medic_record_id = medic_records.id").
		Joins("left join sessions on medical_sessions.session_id = sessions.id").
		Joins("left join medical_staffs on medical_staffs.id = medical_sessions.medical_staff_id").
		Where("medical_sessions.medical_staff_id = ?", id).
		Where(strings.Join(whereAdd, " AND ")).
		Scan(&res).Error

	return
}

func (repo outPatientRepository) FilterByDate(date_start, date_end time.Time) (res []response.OutPatientResponse, err error) {
	return
}
