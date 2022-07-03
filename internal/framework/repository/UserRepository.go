package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"

	"gorm.io/gorm"
)


type userRepository struct {
	sqldb *gorm.DB
}

func NewUserRepository(sqldb *gorm.DB) *userRepository{
	return &userRepository{sqldb: sqldb}
}

func (repo *userRepository) FindAll() (res []response.User, err error) {
	db := repo.sqldb.Model(&models.User{}).
		Select("users.*", "Role.name as role", "MedicalFacility.name as facility").
		Joins("Role").Joins("MedicalFacility").
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *userRepository) FindByID(id int) (res response.User, err error) {
	db := repo.sqldb.Model(&models.User{}).
		Select(`users.*, Role.name as role, MedicalFacility.name as facility`).
		Joins("Role").Joins("MedicalFacility").
		Where("users.id = ?", id).
		Scan(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *userRepository) Create(us models.User) (err error) {
	err = repo.sqldb.Create(&us).Error
	return
}

func (repo *userRepository) Update(us models.User) (err error) {
	db := repo.sqldb.Updates(&us)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo *userRepository) Delete(id int) (err error) {
	db := repo.sqldb.Delete(&models.User{}, id)
	err = check.DBRecord(db, check.DELETE)
	return
}
