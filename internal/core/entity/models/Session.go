package models

type Session struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	TimeStart string
	TimeEnd   string

	Schedule Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
