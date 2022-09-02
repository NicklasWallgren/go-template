package handlers

import (
	"errors"
	errorTypes "github.com/NicklasWallgren/go-template/adapters/driver/api/errors"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
)

type APIErrorTypeHandler struct{}

func NewAPIErrorTypeHandler() ErrorTypeResponseHandler {
	return &APIErrorTypeHandler{}
}

func (a APIErrorTypeHandler) Handle(err error) *response2.APIResponseEnvelope {
	actualError := &errorTypes.APIError{}
	errors.As(err, &actualError)

	errorList := []response2.APIError{
		response2.NewAPIError(actualError.Message),
	}

	return response2.NewEnvelope(actualError.HTTPStatus, response2.WithResponse(response2.NewAPIErrorResponse(errorList)))
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
