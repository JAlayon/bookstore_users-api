package errors

import "net/http"

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func InternalServerError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
