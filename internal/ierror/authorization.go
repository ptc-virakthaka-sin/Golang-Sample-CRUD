package ierror

import (
	"net/http"
)

/*
Unauthorized 403
- Do not have permission to perform this action
- Invalid role
*/
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
		Code:    code,
		Message: message,
	}
}
