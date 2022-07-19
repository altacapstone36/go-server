package repository

import (
	"fmt"
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type sessionRepository struct {
	sqldb *gorm.DB
}

func NewSessionRepository(sqldb *gorm.DB) *sessionRepository {
	return &sessionRepository{sqldb: sqldb}
}

func (repo sessionRepository) FindAll() (session []response.Session, err error) {
	db := repo.sqldb.Model(models.Session{}).
		Scan(&session)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo sessionRepository) FindByDateID(id int, date string) (schedule response.SessionDetails, err error) {
	var where string

	if date != "" {
		where = fmt.Sprintf("schedules.date = '%s'", date)
	}

	db := repo.sqldb.Model(models.Session{}).
		Joins("Schedule").
		Preload("Schedule").
		Where("sessions.id = ?", id).
		Scan(&schedule)

	err = check.DBRecord(db, check.FIND)
	if err != nil {
		return
	}

	err = repo.sqldb.Model(models.Schedule{}).
		Select("U.id", "U.code", "U.full_name", "schedules.date").
		Joins("JOIN users U ON U.code = schedules.user_code").
		Where(where).
		Where("schedules.session_id = ?", id).Scan(&schedule.Staff).Error
	return
}

