package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/go-playground/validator/v10"
)

type ValidationGoPlaygroundErrorHandler struct{}

func NewValidationGoPlaygroundErrorHandler() ErrorTypeResponseHandler {
	return &ValidationGoPlaygroundErrorHandler{}
}

func (v ValidationGoPlaygroundErrorHandler) Handle(err error) response.ApiResponseEnvelop {
	validationErrors := validator.ValidationErrors{}
	errors.As(err, &validationErrors)

	fieldErrors := make([]response.ApiErrorConstraint, len(validationErrors))
	for i, v := range validationErrors {
		message := fmt.Sprintf("Invalid value for field '%s'. Cause: '%s'. Value: '%s'", v.Field(), v.Tag(), v.Value())

		fieldErrors[i] = response.NewApiFieldError(message, v.Field(), v.Value())
	}

	return response.NewApiResponseEnvelop(http.StatusBadRequest, response.WithPayload(response.NewApiErrorResponse(fieldErrors)))
}

func (v ValidationGoPlaygroundErrorHandler) IsSupported(err error) bool {
	return errors.As(err, &validator.ValidationErrors{})
}

func (v ValidationGoPlaygroundErrorHandler) ErrorType() error {
	return &validator.ValidationErrors{}
}
