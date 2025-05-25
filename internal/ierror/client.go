package ierror

import "net/http"

/*
Bad Request 400
- Business logic error
*/
type ClientError struct {
	Status  int
	Code    string
	Message string
}

func (e *ClientError) Error() string {
	return e.Message
}

func NewClientError(httpStatus int, code, message string) *ClientError {
	status := http.StatusBadRequest

	if httpStatus >= 400 && httpStatus < 500 {
		status = httpStatus
	}

	return &ClientError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
