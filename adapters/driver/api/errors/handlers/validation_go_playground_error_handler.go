package handlers

import (
	"errors"
	"fmt"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationGoPlaygroundErrorHandler struct{}

func NewValidationGoPlaygroundErrorHandler() ErrorTypeResponseHandler {
	return &ValidationGoPlaygroundErrorHandler{}
}

func (v ValidationGoPlaygroundErrorHandler) Handle(err error) *response2.APIResponseEnvelope {
	validationErrors := validator.ValidationErrors{}
	errors.As(err, &validationErrors)

	fieldErrors := make([]response2.APIError, len(validationErrors))
	for i, v := range validationErrors { // nolint: wsl
		message := fmt.Sprintf("Invalid value for field '%s'. Cause: '%s'. Value: '%s'", v.Field(), v.Tag(), v.Value())

		fieldErrors[i] = response2.NewAPIErrorWithField(message, v.Field(), v.Value())
	}

	return response2.NewEnvelope(http.StatusBadRequest, response2.WithResponse(response2.NewAPIErrorResponse(fieldErrors)))
}

func (v ValidationGoPlaygroundErrorHandler) IsSupported(err error) bool {
	return errors.As(err, &validator.ValidationErrors{})
}

func (v ValidationGoPlaygroundErrorHandler) ErrorType() error {
	return &validator.ValidationErrors{}
}

func (v ValidationGoPlaygroundErrorHandler) Priority() int {
	return 3
}
