package error

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	Code    int
	Message string
}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.Code, e.Message)
}

// NewBadRequest creates a new ClientError with status code 400
func NewBadRequest(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

// NewNotFound creates a new ClientError with status code 404
func NewNotFound(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

// NewForbidden creates a new ClientError with status code 403
func NewForbidden(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

// NewUnauthorized creates a new ClientError with status code 401
func NewUnauthorized(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

// NewInternalServerError creates a new ClientError with status code 500
func NewInternalServerError(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
