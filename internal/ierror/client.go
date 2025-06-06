package ierror

type ClientError struct {
	Status  int
	Code    string
	Message string
}

func (e *ClientError) Error() string {
	return e.Message
}

func NewClientError(status int, code, message string) *ClientError {
	return &ClientError{
		Message: message,
		Status:  status,
		Code:    code,
	}
}
