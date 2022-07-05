package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Login struct {
	Email    string `json:"email" example:"alsyadahmad@holyhos.co.id"`
	Password string `json:"password" example:"password"`
}

type FindEmail struct {
	Email string `json:"email"`
}

type ChangePassword struct {
	Code string `json:"-"`
	Password string `json:"password"`
}

func (lr Login) Validate() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required, is.EmailFormat),
		validation.Field(&lr.Password, validation.Required, validation.RuneLength(8, 0)),
	)
}

func (fe FindEmail) Validate() error {
	return validation.ValidateStruct(&fe,
		validation.Field(&fe.Email, validation.Required, is.EmailFormat),
	)
}

func (cp ChangePassword) Validate() error {
	return validation.ValidateStruct(&cp,
		validation.Field(&cp.Password, validation.Required, validation.RuneLength(8, 0)),
	)
}
