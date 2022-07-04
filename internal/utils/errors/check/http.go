package check

import (
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors"
	"reflect"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type HTTPCheck struct {
	Code int
	Result interface{}
}

// check HTTP request error
// default response in code 200 for success
// and 417 for error
// it will return code from errors.RequestError
// if type is errors.RequestError
func HTTP(res interface{}, err error, msg string) (r HTTPCheck, ok bool) {
	r.Code = 200
	var error *errors.RequestError
	var errMsg any

	if err != nil {
		r.Code = 400
		if reflect.TypeOf(err).String() == "*errors.RequestError" {
			error = err.(*errors.RequestError)
			r.Code = error.Code()
		}

		if reflect.ValueOf(err).Kind().String() == "map" {
			errMsg = err
		} else {
			errMsg = err.Error()
		}

		r.Result = response.Error{
			Message: "Failed to " + msg,
			Error: errMsg,
		}

		return
	}

	ok = true

	if res != nil {
		r.Result = response.MessageData{
			Message: "Success to " + msg,
			Data: res,
		}
		return
	}
	r.Result = response.MessageOnly{
			Message: "Success to " + msg,
	}

	return
}

func ParamQuery(args ...interface{}) (r HTTPCheck, ok bool ) {
	for _, v := range args {
		err := validation.Validate(&v, validation.Date("2006-01-02"))
		if err != nil {
			r.Code = 417
			r.Result =  response.Error{
				Message: "Failed to Validate Parameter",
				Error: err.Error(),
			}

			return
		}
	}
	ok = true
	return
}
