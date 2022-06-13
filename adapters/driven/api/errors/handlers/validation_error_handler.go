package handlers

import (
	"errors"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/NicklasWallgren/go-template/domain/validation"
)

type ValidationErrorTypeHandler struct{}

func NewValidationErrorTypeHandler() ErrorTypeResponseHandler {
	return &ValidationErrorTypeHandler{}
}

func (a ValidationErrorTypeHandler) Handle(err error) response.APIResponseEnvelop {
	validationError := &validation.ValidationError{}
	errors.As(err, &validationError)

	errors := []response.APIError{response.NewAPIError(validationError.Message)}

	return response.New(http.StatusBadRequest, response.WithResponse(response.NewAPIErrorResponse(errors)))
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
