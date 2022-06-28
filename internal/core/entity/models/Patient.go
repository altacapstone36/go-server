package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Patient struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	ResidentRegistration string `json:"resident_registration"`
	FullName string `json:"full_name" sql:"unique"`
	Address string `json:"address"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthdate"`
	BloodType string `json:"blood_type"`

	MedicRecord *[]MedicRecord `json:"medic_record"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	var id int64

	tx.Model(&p).Count(&id)
	p.Code = fmt.Sprintf("RM%05d", id+1)

	return
}
