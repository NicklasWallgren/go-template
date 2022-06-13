package response

type APIErrorResponse struct {
	Errors []APIError
}

func NewAPIErrorResponse(errors []APIError) APIResponse {
	return &APIErrorResponse{Errors: errors}
}

type APIError struct {
	Message string `json:"message" example:"invalid id"`
	Field   string `json:"field,omitempty" example:"id"`
	Value   any    `json:"value,omitempty" example:"1"`
}

func NewAPIError(message string) APIError {
	return APIError{Message: message}
}

func NewAPIWithFieldError(message string, field string, value any) APIError {
	return APIError{Message: message, Field: field, Value: value}
}
