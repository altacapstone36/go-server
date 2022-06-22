package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"

	"gorm.io/gorm"
)


type facilityRepository struct {
	sqldb *gorm.DB
}

func NewFacilityRepository(sqldb *gorm.DB) *facilityRepository{
	return &facilityRepository{sqldb: sqldb}
}

func (repo *facilityRepository) FindAll() (res []response.Facility, err error) {
	err = repo.sqldb.Model(&models.MedicalFacility{}).Scan(&res).Error
	return
}

func (repo *facilityRepository) FindByID(id int) (res response.FacilityDetails, err error) {
	err = repo.sqldb.Model(&models.MedicalFacility{}).
		Where("id = ?", id).
		Scan(&res).Error

	err = repo.sqldb.Model(&models.User{}).
		Select("users.*, roles.name as role").
		Joins("join roles on roles.id = users.role_id").
		Where("medical_facility_id = ?", id).
		Scan(&res.Staff).Error
	return
}

func (repo *facilityRepository) Create(mf models.MedicalFacility) (err error) {
	err = repo.sqldb.Create(&mf).Error
	return
}

func (u *facilityRepository) Update(us models.MedicalFacility) (err error) {
	err = u.sqldb.Updates(&us).Error
	return
}

func (u *facilityRepository) Delete(id int) (err error) {
	err = u.sqldb.Delete(models.MedicalFacility{}, id).Error
	return
}
