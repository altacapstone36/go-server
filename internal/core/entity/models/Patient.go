package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
	var name string

	tx.Model(&p).Select("full_name").
		Where("full_name = ?", p.FullName).
		Scan(&name);

	if name != "" {
		msg := fmt.Sprintf("duplicate name for %s found, please use another name", p.FullName)
		err = errors.New(203, msg)
		return
	}
	
	var id int64

	tx.Model(&p).Count(&id)
	p.Code = fmt.Sprintf("RM%05d", id+1)

	return
}

func (p Patient) Validate() (err error) {
	err = validation.ValidateStruct(&p,
		validation.Field(&p.ResidentRegistration, validation.Required, validation.RuneLength(16, 16), is.Digit),
		validation.Field(&p.FullName, validation.Required),
		validation.Field(&p.Address, validation.Required),
		validation.Field(&p.Gender, validation.Required),
		validation.Field(&p.BirthDate, validation.Required, validation.Date("2006-01-02")),
	)
	return
}
