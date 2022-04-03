package handlers

import (
	"errors"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/NicklasWallgren/go-template/domain/validation"
)

type ValidationFieldErrorTypeHandler struct{}

// To ensure that ValidationFieldErrorTypeHandler implements the ErrorTypeResponseHandler interface
var _ ErrorTypeResponseHandler = (*ValidationFieldErrorTypeHandler)(nil)

func NewValidationFieldErrorTypeHandler() *ValidationFieldErrorTypeHandler {
	return &ValidationFieldErrorTypeHandler{}
}

func (a ValidationFieldErrorTypeHandler) Handle(err error) response.ApiResponseEnvelop {
	validationFieldError := &validation.ValidationFieldError{}
	errors.As(err, &validationFieldError)

	errors := []response.ApiErrorConstraint{
		response.NewApiFieldError(validationFieldError.Message, validationFieldError.Field, validationFieldError.Value),
	}

	return response.NewApiResponseEnvelop(http.StatusBadRequest, response.WithPayload(response.NewApiErrorResponse(errors)))
}

func (a ValidationFieldErrorTypeHandler) IsSupported(err error) bool {
	validationFieldError := &validation.ValidationFieldError{}
	return errors.As(err, &validationFieldError)
}

func (a ValidationFieldErrorTypeHandler) ErrorType() error {
	return &validation.ValidationFieldError{}
}
