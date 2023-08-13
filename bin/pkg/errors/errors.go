package errors

import "net/http"

type ErrorString struct {
	code    int
	message string
}

func (e ErrorString) Code() int {
	return e.code
}

func (e ErrorString) Error() string {
	return e.message
}

func (e ErrorString) Message() string {
	return e.message
}

func BadRequest(msg string) error {
	return &ErrorString{
		code:    http.StatusBadRequest,
		message: msg,
	}
}

func Conflict(msg string) error {
	return &ErrorString{
		code:    http.StatusConflict,
		message: msg,
	}
}

func InternalServerError(msg string) error {
	return &ErrorString{
		code:    http.StatusInternalServerError,
		message: msg,
	}
}
