package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors"
	"go-hospital-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type doctorScheduleRepository struct {
	sqldb *gorm.DB
}

func NewDoctorScheduleRepository(sqldb *gorm.DB) *doctorScheduleRepository {
	return &doctorScheduleRepository{sqldb: sqldb}
}

func (repo *doctorScheduleRepository) FindAll(res []response.ScheduleList, err error) {
	db := repo.sqldb.Model(&models.DoctorSchedule{}).Scan(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *doctorScheduleRepository) Create(ds models.DoctorSchedule) (err error) {
	db := repo.sqldb.Create(&ds)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (repo *doctorScheduleRepository) Update(ds models.DoctorSchedule) (err error) {
	db := repo.sqldb.Create(&ds)
	if db.RowsAffected == 0 {
		err = errors.ErrNoChange
	}
	return
}

func (repo *doctorScheduleRepository) Delete(id int) (err error) {
	db := repo.sqldb.Delete(models.DoctorSchedule{}, id)
	if db.RowsAffected == 0 {
		err = errors.ErrNoRecord
	}
	return
}
