package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"

	"gorm.io/gorm"
)


type userRepository struct {
	sqldb *gorm.DB
}

func NewUserRepository(sqldb *gorm.DB) *userRepository{
	return &userRepository{sqldb: sqldb}
}

func (u *userRepository) FindAll() (res []response.User, err error) {
	err = u.sqldb.Model(&models.User{}).
		Select(`users.*, roles.name as role, medical_facilities.name as facility`).
		Joins("left join roles on users.role_id = roles.id").
		Joins("left join medical_facilities on medical_facilities.id = users.medical_facility_id").
		Scan(&res).Error
	return
}

func (u *userRepository) FindByID(id int) (res response.User, err error) {
	err = u.sqldb.Model(&models.User{}).
		Select(`users.*, roles.name as role, medical_facilities.name as facility`).
		Joins("left join roles on users.role_id = roles.id").
		Joins("left join medical_facilities on medical_facilities.id = users.medical_facility_id").
		Where("users.id = ?", id).
		Scan(&res).Error
	return
}

func (u *userRepository) Create(us models.User) (err error) {
	err = u.sqldb.Create(&us).Error
	return
}

func (u *userRepository) Update(us models.User) (err error) {
	err = u.sqldb.Updates(&us).Error
	return
}

func (u *userRepository) Delete(id int) (err error) {
	err = u.sqldb.Delete(models.User{}, id).Error
	return
}
