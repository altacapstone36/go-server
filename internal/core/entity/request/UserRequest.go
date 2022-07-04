package request

import (
	vs "go-hospital-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserRequest struct {
	ID                int    `json:"-"`
	Email             string `json:"email" example:"milimnava@holyhos.co.id"`
	Password          string `json:"password" example:"password"`
	FullName          string `json:"full_name" example:"Milim Nava"`
	Gender            string `json:"gender" example:"Female"`
	RoleID            int    `json:"role_id" example:"2"`
	MedicalFacilityID int    `json:"facility_id" example:"3"`
}

func (ur UserRequest) Validate() (err error) {
	err = v.ValidateStruct(&ur,
		v.Field(&ur.Email, v.Required, is.EmailFormat, v.By(vs.Duplicate("users", "email", ur.ID))),
		v.Field(&ur.FullName, v.Required, v.Length(5, 0), v.By(vs.Duplicate("users", "full_name", ur.ID))),
		v.Field(&ur.Gender, v.Required, v.In("Male", "Female").Error("please select between Male or Female")),
		v.Field(&ur.RoleID, v.Required, v.NilOrNotEmpty, v.By(vs.RoleLimit(ur.MedicalFacilityID))),
		v.Field(&ur.MedicalFacilityID, v.Required, v.NilOrNotEmpty, v.By(vs.AvailableFacility)),
	)
	return
}
