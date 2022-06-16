package models

type Session struct {
	ID uint `gorm:"primaryKey"`
	Name string
	TimeStart string
	TimeEnd string
}
