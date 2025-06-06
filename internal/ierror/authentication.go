package ierror

import (
	"net/http"
)

type AuthenticationError struct {
	Status  int
	Code    string
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(code, message string) *AuthenticationError {
	return &AuthenticationError{
		Status:  http.StatusUnauthorized,
		Message: message,
		Code:    code,
	}
}
