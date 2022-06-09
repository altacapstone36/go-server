package models

import "time"

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

func (*User) TableName() string {
	return "users"
}
