package services

import "net/http"

type HttpError struct {
	code    int
	message string
}

func (e HttpError) Error() string {
	return e.message
}

func BadRequest(message string) error {
	return HttpError{code: http.StatusBadRequest, message: message}
}
