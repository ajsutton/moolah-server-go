package web

import "net/http"

type RequestError struct {
	code    int
	message string
}

func (e RequestError) Error() string {
	return e.message
}

func BadRequest(message string) error {
	return RequestError{code: http.StatusBadRequest, message: message}
}
