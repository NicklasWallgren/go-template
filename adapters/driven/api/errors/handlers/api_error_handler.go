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

func (a ApiErrorTypeHandler) Handle(err error) response.APIResponseEnvelop {
	actualError, _ := (err).(*errorTypes.APIError) // nolint:errorlint

	errorList := []response.APIError{
		response.NewAPIError(actualError.Message),
	}

	return response.New(actualError.HTTPStatus, response.WithResponse(response.NewAPIErrorResponse(errorList)))
}

func (a ApiErrorTypeHandler) IsSupported(err error) bool {
	domainError := &errorTypes.APIError{}

	return errors.As(err, &domainError)
}

func (a ApiErrorTypeHandler) ErrorType() error {
	return &errorTypes.APIError{}
}

func (a ApiErrorTypeHandler) Priority() int {
	return 5
}
