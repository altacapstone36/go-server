package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Code string `json:"code" gorm:"primary_key"`
	Email string `json:"email"`
	Password string `json:"password"`
	Status int `json:"status"`
	RoleID int `json:"role_id"`
	Role Role `json:"roles"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var email string

	tx.Model(&u).Select("email").
		Where("email = ?", u.Email).
		Scan(&email);

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
