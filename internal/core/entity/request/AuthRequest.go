package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Login struct {
	Email    string `json:"email" example:"alsyadahmad@holyhos.co.id"`
	Password string `json:"password" example:"password"`
}

func (lr Login) Validate() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required, is.EmailFormat),
		validation.Field(&lr.Password, validation.Required, validation.RuneLength(8, 0)),
	)
}
