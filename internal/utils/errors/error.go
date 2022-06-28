package errors

var ErrNoChange = New(204, "Nothing Changed")
var ErrNoRecord = New(404, "No Record Found")
var ErrInternalServer = New(500, "Internal Server Error")

func New(code int, text string) error {
	return &RequestError{
		code: code,
		message: text,
	}
}

type RequestError struct {
	code int
	message string
}

func (e *RequestError) Error() string {
	return e.message
}

func (e *RequestError) Code() int {
	return e.code
}
