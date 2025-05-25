package ierror

import (
	"learn-fiber/api/response"
	"net/http"
)

type ValidationError struct {
	Status  int
	Code    string
	Message string
	Errors  []response.ValidationError
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(errors []response.ValidationError) *ValidationError {
	return &ValidationError{
		Code:    ErrCodeValidationError,
		Message: ErrCodeValidationError,
		Status:  http.StatusBadRequest,
		Errors:  errors,
	}
}
