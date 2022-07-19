package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	sqldb *gorm.DB
}

func NewScheduleRepository(sqldb *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{sqldb: sqldb}
}

func (repo ScheduleRepository) Create(ds models.Schedule) (err error) {
	db := repo.sqldb.Create(&ds)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (repo ScheduleRepository) Update(ds models.Schedule) (err error) {
	db := repo.sqldb.Debug().Updates(&ds)

	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo ScheduleRepository) Delete(id int) (err error) {
	db := repo.sqldb.Delete(models.Schedule{}, id)

	err = check.DBRecord(db, check.DELETE)
	return
}
