package handlers

import (
	"errors"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"net/http"

	"github.com/NicklasWallgren/go-template/domain/validation"
)

type ValidationErrorTypeHandler struct{}

func NewValidationErrorTypeHandler() ErrorTypeResponseHandler {
	return &ValidationErrorTypeHandler{}
}

func (a ValidationErrorTypeHandler) Handle(err error) *response2.APIResponseEnvelope {
	validationError := &validation.ValidationError{}
	errors.As(err, &validationError)

	errorList := []response2.APIError{response2.NewAPIError(validationError.Message)}

	return response2.NewEnvelope(http.StatusBadRequest, response2.WithResponse(response2.NewAPIErrorResponse(errorList)))
}

func (a ValidationErrorTypeHandler) IsSupported(err error) bool {
	validationError := &validation.ValidationError{}

	return errors.As(err, &validationError)
}

func (a ValidationErrorTypeHandler) ErrorType() error {
	return &validation.ValidationError{}
}

func (a ValidationErrorTypeHandler) Priority() int {
	return 2
}
