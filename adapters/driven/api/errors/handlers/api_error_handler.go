package handlers

import (
	"errors"

	errorTypes "github.com/NicklasWallgren/go-template/adapters/driven/api/errors"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
)

type ApiErrorTypeHandler struct{}

func NewApiErrorTypeHandler() ErrorTypeResponseHandler {
	return &ApiErrorTypeHandler{}
}

func (a ApiErrorTypeHandler) Handle(err error) *response.APIResponseEnvelope {
	actualError := &errorTypes.APIError{}
	errors.As(err, &actualError)

	errorList := []response.APIError{
		response.NewAPIError(actualError.Message),
	}

	return response.NewEnvelope(actualError.HTTPStatus, response.WithResponse(response.NewAPIErrorResponse(errorList)))
}

func (a ApiErrorTypeHandler) IsSupported(err error) bool {
	apiError := &errorTypes.APIError{}

	return errors.As(err, &apiError)
}

func (a ApiErrorTypeHandler) ErrorType() error {
	return &errorTypes.APIError{}
}

func (a ApiErrorTypeHandler) Priority() int {
	return 5
}
