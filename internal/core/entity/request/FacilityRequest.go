package request

import validation "github.com/go-ozzo/ozzo-validation/v4"


type Facility struct {
	Name string `json:"name"`
}

func (s Facility) Validate() (err error) {
	err = validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required, validation.Length(4, 0)),
	)
	return
}
