package handlers

import (
	"errors"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"

	"github.com/NicklasWallgren/go-template/domain/validation"
)

type ValidationFieldErrorTypeHandler struct{}

func NewValidationFieldErrorTypeHandler() ErrorTypeResponseHandler {
	return &ValidationFieldErrorTypeHandler{}
}

func (a ValidationFieldErrorTypeHandler) Handle(err error) *response.APIResponseEnvelope {
	validationFieldError := &validation.ValidationFieldError{}
	errors.As(err, &validationFieldError)

	errorList := []response.APIError{
		response.NewAPIErrorWithField(validationFieldError.Message, validationFieldError.Field, validationFieldError.Value),
	}

	return response.NewEnvelope(http.StatusBadRequest, response.WithResponse(response.NewAPIErrorResponse(errorList)))
}

func (a ValidationFieldErrorTypeHandler) IsSupported(err error) bool {
	validationFieldError := &validation.ValidationFieldError{}

	return errors.As(err, &validationFieldError)
}

func (a ValidationFieldErrorTypeHandler) ErrorType() error {
	return &validation.ValidationFieldError{}
}

func (a ValidationFieldErrorTypeHandler) Priority() int {
	return 1
}
