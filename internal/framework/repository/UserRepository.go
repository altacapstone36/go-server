package repository

import (
	"fmt"
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"
	"strings"

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

func (repo *userRepository) FindByRFS(role_id, facility_id, session_id int) (res []response.User, err error) {
	var where []string
	var join string

	if session_id != 0 {
		join = "JOIN schedules S ON S.user_id = users.id"
		w := fmt.Sprintf("S.session_id = %d", session_id)
		where = append(where, w)
	}

	if role_id != 0 {
		w := fmt.Sprintf("users.role_id = %d", role_id)
		where = append(where, w)
	}

	if facility_id != 0 {
		w := fmt.Sprintf("users.medical_facility_id = %d", facility_id)
		where = append(where, w)
	}

	db := repo.sqldb.Debug().Model(&models.User{}).
		Select("users.*", "Role.name as role", "MedicalFacility.name as facility").
		Joins("Role").Joins("MedicalFacility").Joins(join).
		Where(strings.Join(where, " AND ")).
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
	db := repo.sqldb.Create(&us)
	err = check.DBRecord(db, check.CREATE)
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
