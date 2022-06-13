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

func (v ValidationGoPlaygroundErrorHandler) Handle(err error) response.APIResponseEnvelop {
	validationErrors := validator.ValidationErrors{}
	errors.As(err, &validationErrors)

	fieldErrors := make([]response.APIError, len(validationErrors))
	for i, v := range validationErrors { // nolint: wsl
		message := fmt.Sprintf("Invalid value for field '%s'. Cause: '%s'. Value: '%s'", v.Field(), v.Tag(), v.Value())

		fieldErrors[i] = response.NewAPIWithFieldError(message, v.Field(), v.Value())
	}

	return response.New(http.StatusBadRequest, response.WithResponse(response.NewAPIErrorResponse(fieldErrors)))
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
