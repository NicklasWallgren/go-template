package response

type ApiErrorResponse struct {
	Errors []ApiError
}

func NewApiErrorResponse(errors []ApiError) ApiResponse {
	return &ApiErrorResponse{Errors: errors}
}

type ApiError struct {
	Message string `json:"message" example:"invalid id"`
	Field   string `json:"field,omitempty" example:"id"`
	Value   any    `json:"value,omitempty" example:"1"`
}

func NewApiError(message string) ApiError {
	return ApiError{Message: message}
}

func NewApiFieldError(message string, field string, value any) ApiError {
	return ApiError{Message: message, Field: field, Value: value}
}
