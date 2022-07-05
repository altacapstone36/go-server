package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors"
	"go-hospital-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type facilityRepository struct {
	sqldb *gorm.DB
}

func NewFacilityRepository(sqldb *gorm.DB) *facilityRepository {
	return &facilityRepository{sqldb: sqldb}
}

func (repo *facilityRepository) FindAll() (res []response.Facility, err error) {
	db := repo.sqldb.Model(&models.MedicalFacility{}).
		Scan(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *facilityRepository) FindByID(id int) (res response.FacilityDetails, err error) {
	var mf models.MedicalFacility
	db := repo.sqldb.First(&mf, id).Scan(&res)

	err = check.DBRecord(db, check.FIND)
	if err != nil {
		return
	}

	err = repo.sqldb.Model(&models.User{}).
		Select("users.*, roles.name as role").
		Joins("join roles on roles.id = users.role_id").
		Where("medical_facility_id = ?", id).
		Scan(&res.Staff).Error
	return
}

func (repo *facilityRepository) Create(mf models.MedicalFacility) (err error) {
	db := repo.sqldb.Create(&mf)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (u *facilityRepository) Update(us models.MedicalFacility) (err error) {
	db := u.sqldb.Updates(&us)
	if db.RowsAffected == 0 {
		err = errors.ErrNoChange
	}
	return
}

func (u *facilityRepository) Delete(id int) (err error) {
	db := u.sqldb.Delete(models.MedicalFacility{}, id)
	if db.RowsAffected == 0 {
		err = errors.ErrNoRecord
	}
	return
}
