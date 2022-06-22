package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Patient struct {
	Code string `json:"code" example:"RM0001"`
	ResidentRegistration string `json:"resident_registration" example:"8729301745162748"`
	FullName string `json:"full_name" example:"Faizur Ramadhan"`
	Address string `json:"address" example:"Sumenep"`
	Gender string `json:"gender" example:"Male"`
	BirthDate string `json:"birthdate" example:"2001-04-14"`
	BloodType string `json:"blood_type" example:"A"`
}

func (ur Patient) Validate() (err error) {
	err = validation.ValidateStruct(&ur, 
		validation.Field(&ur.ResidentRegistration, validation.Required, is.Digit),
		validation.Field(&ur.FullName, validation.Required, validation.Length(5, 0)),
		validation.Field(&ur.Address, validation.Required, validation.Length(4, 0)),
		validation.Field(&ur.Gender, validation.Required, validation.In("Male", "Female")),
		validation.Field(&ur.BirthDate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&ur.BloodType, validation.Required),
	)
	return
}

