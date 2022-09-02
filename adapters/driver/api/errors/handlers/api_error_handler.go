package handlers

import (
	"errors"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"

	errorTypes "github.com/NicklasWallgren/go-template/adapters/driver/api/errors"
)

type APIErrorTypeHandler struct{}

func NewAPIErrorTypeHandler() ErrorTypeResponseHandler {
	return &APIErrorTypeHandler{}
}

func (a APIErrorTypeHandler) Handle(err error) *response.APIResponseEnvelope {
	actualError := &errorTypes.APIError{}
	errors.As(err, &actualError)

	errorList := []response.APIError{
		response.NewAPIError(actualError.Message),
	}

	return response.NewEnvelope(actualError.HTTPStatus, response.WithResponse(response.NewAPIErrorResponse(errorList)))
}

func (a APIErrorTypeHandler) IsSupported(err error) bool {
	apiError := &errorTypes.APIError{}

	return errors.As(err, &apiError)
}

func (a APIErrorTypeHandler) ErrorType() error {
	return &errorTypes.APIError{}
}

func (a APIErrorTypeHandler) Priority() int {
	return 5
}
