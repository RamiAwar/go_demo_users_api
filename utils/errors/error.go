package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func BadRequestError(err error) *RestError {
	return &RestError{
		Message: "Invalid Request",
		Status: http.StatusBadRequest,
		Error: err.Error(),
	}
}