package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserRequest struct {
	Email string `json:"email" example:"milimnava@holyhos.co.id"`
	Password string `json:"password" example:"password"`
	FullName string `json:"full_name" example:"Milim Nava"`
	Gender string `json:"gender" example:"Female"`
	RoleID int `json:"role_id" example:"2"`
	MedicalFacilityID *uint `json:"facility_id" example:"3"`
}

func (ur UserRequest) Validate() (err error) {
	err = validation.ValidateStruct(&ur, 
		validation.Field(&ur.Email, validation.Required, is.EmailFormat),
		validation.Field(&ur.FullName, validation.Required, validation.Length(5, 0)),
		validation.Field(&ur.Gender, validation.Required, validation.In("Male", "Female")),
		validation.Field(&ur.RoleID, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&ur.MedicalFacilityID, validation.Required, validation.NilOrNotEmpty),
	)
	return
}
