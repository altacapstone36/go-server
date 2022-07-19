package repository

import (
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

func (repo sessionRepository) FindAllSession() (session []response.Session, err error) {
	db := repo.sqldb.Model(models.Session{}).
		Scan(&session)

	err = check.DBRecord(db, check.FIND)
	return
}
