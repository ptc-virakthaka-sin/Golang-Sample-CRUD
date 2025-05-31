package ierror

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
		Message: message,
		Status:  500,
		Code:    code,
	}
}
