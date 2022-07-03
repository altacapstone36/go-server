package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Login struct {
	Email    string `json:"email" example:"alsyadahmad@holyhos.co.id"`
	Password string `json:"password" example:"password"`
}

type RegisterRequest struct {
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Gender            string `json:"gender"`
	RoleID            int    `json:"role_id"`
	MedicalFacilityID *uint  `json:"facility_id"`
}

type FindEmail struct {
	Email string `json:"email"`
}

type ChangePassword struct {
	Password string `json:"password"`
}

func (lr Login) Validate() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required, is.EmailFormat),
		validation.Field(&lr.Password, validation.Required, validation.RuneLength(8, 0)),
	)
}

func (reg RegisterRequest) Validate() (err error) {
	err = validation.ValidateStruct(&reg,
		validation.Field(&reg.Email, validation.Required, is.EmailFormat),
		validation.Field(&reg.FullName, validation.Required, validation.Length(5, 0)),
		validation.Field(&reg.Gender, validation.Required, validation.In("Male", "Female")),
		validation.Field(&reg.RoleID, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&reg.MedicalFacilityID, validation.Required, validation.NilOrNotEmpty),
	)
	return
}
