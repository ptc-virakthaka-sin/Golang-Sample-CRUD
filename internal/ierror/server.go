package ierror

import "net/http"

type ServerError struct {
	Status  int
	Code    string
	Message string
}

func (e *ServerError) Error() string {
	return e.Message
}

func NewServerError(code, message string) *ServerError {
	return &ServerError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Code:    code,
	}
}
