package models

import (
	"fmt"
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

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		var id int64
		var code string

		tx.Model(&u).Where("role_id = ?", u.RoleID).Count(&id)
		tx.Model(&u.Role).Select("code").
			Where("id = ?", u.RoleID).Scan(&code)
		u.Code = fmt.Sprintf("%s%05d", code, id+1)
	}
	return
}
