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

func (repo *userRepository) FindAll() (res []response.User, err error) {
	err = repo.sqldb.Model(&models.User{}).
		Select(`users.*, roles.name as role, medical_facilities.name as facility`).
		Joins("left join roles on users.role_id = roles.id").
		Joins("left join medical_facilities on medical_facilities.id = users.medical_facility_id").
		Scan(&res).Error
	return
}

func (repo *userRepository) FindByID(id int) (res response.User, err error) {
	err = repo.sqldb.Model(&models.User{}).
		Select(`users.*, roles.name as role, medical_facilities.name as facility`).
		Joins("left join roles on users.role_id = roles.id").
		Joins("left join medical_facilities on medical_facilities.id = users.medical_facility_id").
		Where("users.id = ?", id).
		Scan(&res).Error
	return
}

func (repo *userRepository) Create(us models.User) (err error) {
	err = repo.sqldb.Create(&us).Error
	return
}

func (repo *userRepository) Update(us models.User) (err error) {
	up := repo.sqldb.Updates(&us)
	if up.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (repo *userRepository) Delete(id int) (err error) {
	del := repo.sqldb.Delete(models.Patient{}, id)
	if del.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}
