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

func NewServerError(httpStatus int, code, message string) *ServerError {
	status := http.StatusInternalServerError

	if httpStatus >= 500 && httpStatus < 600 {
		status = httpStatus
	}

	return &ServerError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
