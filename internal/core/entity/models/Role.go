package models


type Role struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`
}
