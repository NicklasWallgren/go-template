package response

type ApiErrorConstraint interface{}

type ApiErrorResponse struct {
	Errors []ApiErrorConstraint
}

func NewApiErrorResponse(errors []ApiErrorConstraint) *ApiErrorResponse {
	return &ApiErrorResponse{Errors: errors}
}

type ApiError struct {
	Message string `json:"message"`
}

func NewApiError(message string) *ApiError {
	return &ApiError{Message: message}
}

type ApiFieldError struct {
	ApiError
	Field string
	Value any
}

func NewApiFieldError(message string, field string, value any) *ApiFieldError {
	return &ApiFieldError{ApiError: ApiError{Message: message}, Field: field, Value: value}
}
