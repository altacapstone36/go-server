package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
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

	if err = tx.Model(&u).Where("email = ?", u.Email).Error;
	err == nil {
		msg := fmt.Sprintf("duplicate email for %s found, please use another email", u.Email)
		err = errors.New(203, msg)
		return
	}

	return
}
