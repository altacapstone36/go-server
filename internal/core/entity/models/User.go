package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Code string `json:"code"`
	FullName string `json:"full_name"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Password string `json:"password"`
	Status int `json:"status" sql:"default:1"`
	RoleID int `json:"role_id"`
	Role Role `json:"roles"`
	MedicalFacilityID *uint `json:"facility_id"`
	MedicalFacility MedicalFacility `json:"facility"`
	Schedule []Schedule `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var email string
	var name string
	var c int64

	tx.Model(&u).Select("email").
		Where("email = ?", u.Email).
		Scan(&email);
	
	tx.Model(&u).Select("full_name").
		Where("full_name = ?", u.FullName).
		Scan(&name);
		
	tx.Table("medical_facilities").Where("id = ?", u.MedicalFacilityID).Count(&c)
	if c == 0 {
		errMsg := fmt.Sprintf("No Medical Facility with ID %d", u.MedicalFacilityID)
		err = errors.New(417, errMsg)
		return
	}
	
	if name != "" {
		msg := fmt.Sprintf("duplicate name for %s found, please use another name", u.FullName)
		err = errors.New(203, msg)
		return
	}

	if email != "" {
		msg := fmt.Sprintf("duplicate email for %s found, please use another email", u.Email)
		err = errors.New(203, msg)
		return
	}
	
	var id int64
	var code string

	tx.Model(&u).Where("role_id = ?", u.RoleID).Count(&id)
	tx.Model(&u.Role).Select("code").
		Where("id = ?", u.RoleID).Scan(&code)
	u.Code = fmt.Sprintf("%s%05d", code, id+1)

	return
}
