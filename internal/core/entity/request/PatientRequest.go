package request

import (
	vs "go-hospital-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Patient struct {
	ID         int    `json:"-"`
	NationalID string `json:"national_id" example:"8729301745162748"`
	FullName   string `json:"full_name" example:"Faizur Ramadhan"`
	Address    string `json:"address" example:"Sumenep"`
	Gender     string `json:"gender" example:"Male"`
	BirthDate  string `json:"birthdate" example:"2001-04-14"`
	BloodType  string `json:"blood_type" example:"A"`
}

func (ur Patient) Validate() (err error) {
	err = v.ValidateStruct(&ur,
		v.Field(&ur.NationalID, v.Required, v.RuneLength(16, 16), is.Digit, v.By(vs.Duplicate("patients", "national_id", ur.ID))),
		v.Field(&ur.FullName, v.Required, v.Length(5, 0), v.By(vs.Duplicate("patients", "full_name", ur.ID))),
		v.Field(&ur.Address, v.Required, v.Length(4, 0)),
		v.Field(&ur.Gender, v.Required, v.In("Male", "Female").Error("please select between Male or Female")),
		v.Field(&ur.BirthDate, v.Required, v.Date("2006-01-02")),
		v.Field(&ur.BloodType, v.Required),
	)
	return
}
