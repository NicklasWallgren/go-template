package handlers

import (
	"errors"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"net/http"

	"github.com/NicklasWallgren/go-template/domain/validation"
)

type ValidationFieldErrorTypeHandler struct{}

func NewValidationFieldErrorTypeHandler() ErrorTypeResponseHandler {
	return &ValidationFieldErrorTypeHandler{}
}

func (a ValidationFieldErrorTypeHandler) Handle(err error) *response2.APIResponseEnvelope {
	validationFieldError := &validation.ValidationFieldError{}
	errors.As(err, &validationFieldError)

	errorList := []response2.APIError{
		response2.NewAPIErrorWithField(validationFieldError.Message, validationFieldError.Field, validationFieldError.Value),
	}

	return response2.NewEnvelope(http.StatusBadRequest, response2.WithResponse(response2.NewAPIErrorResponse(errorList)))
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
