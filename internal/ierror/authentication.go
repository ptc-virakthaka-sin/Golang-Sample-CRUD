package ierror

import (
	"net/http"
)

/*
Unauthorized 401
- No authorization header or missing token
- Invalid token or expired
*/
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
		Code:    code,
		Message: message,
	}
}
