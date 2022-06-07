package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	Email string `json:"email"`
	Password string `json:"password"`
	Status int `json:"status"`
	LevelID int `json:"level_id"`
	Level Level `json:"level"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (*User) TableName() string {
	return "users"
}
