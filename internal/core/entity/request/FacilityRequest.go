package request

import (
	vs "go-hospital-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Facility struct {
	ID   int    `json:"-"`
	Name string `json:"name" example:"General"`
}

func (s Facility) Validate() (err error) {
	err = v.ValidateStruct(&s,
		v.Field(&s.Name, v.Required, v.Length(4, 0), v.By(vs.Duplicate("medical_facilities", "name", s.ID))),
	)
	return
}
