package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Patient struct {
	ID uint `json:"id" gorm:"primaryKey"`
	ResidentRegistration string `json:"resident_registration"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	Gender string `json:"gender"`
	BirthDate time.Time `json:"birthdate"`
	BloodType string `json:"blood_type"`
}

func (p Patient) Validate() (err error) {
	err = validation.ValidateStruct(&p,
		validation.Field(&p.ResidentRegistration, validation.Required, validation.RuneLength(16, 16), is.Digit),
		validation.Field(&p.FullName, validation.Required),
		validation.Field(&p.Address, validation.Required),
		validation.Field(&p.Gender, validation.Required),
		validation.Field(&p.BirthDate, validation.Required),
	)
	return
}
