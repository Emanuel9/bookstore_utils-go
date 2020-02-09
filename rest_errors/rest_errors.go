package rest_errors

import (
	"errors"
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status    int    `json:"status"`
	Error   string `json:"error"`
	Causes []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestError {
	result := &RestError{
		Message: message,
		Status:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}

func NewRestError(message string, status int, error string, causes []interface{}) *RestError {
	return &RestError{
		Message: message,
		Status:  status,
		Error:   error,
		Causes:  causes,
	}
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: "unable to retrieve use information from given access_token",
		Status: http.StatusUnauthorized,
		Error: "unauthorized",
	}
}
