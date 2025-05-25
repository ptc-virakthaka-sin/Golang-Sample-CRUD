package response

type ErrorResponse struct {
	Error ErrorObject `json:"error"`
}

type ErrorObject struct {
	Code             string            `json:"code"`
	Message          string            `json:"message"`
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
	TraceID          string            `json:"traceID"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
