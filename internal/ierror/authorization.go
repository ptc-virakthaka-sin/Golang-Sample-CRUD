package ierror

import (
	"net/http"
)

type AuthorizationError struct {
	Status  int
	Code    string
	Message string
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

func NewAuthorizationError(code, message string) *AuthorizationError {
	return &AuthorizationError{
		Status:  http.StatusForbidden,
		Message: message,
		Code:    code,
	}
}
