package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Invalid Request",
	}
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Resource not found",
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
}
