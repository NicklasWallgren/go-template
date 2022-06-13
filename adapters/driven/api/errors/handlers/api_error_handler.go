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

func (a ApiErrorTypeHandler) Handle(err error) response.ApiResponseEnvelop {
	actualError, _ := (err).(*errorTypes.ApiError) // nolint:errorlint

	errors := []response.ApiError{
		response.NewApiError(actualError.Message),
	}

	return response.New(actualError.HttpStatus, response.WithResponse(response.NewApiErrorResponse(errors)))
}

func (a ApiErrorTypeHandler) IsSupported(err error) bool {
	domainError := &errorTypes.ApiError{}

	return errors.As(err, &domainError)
}

func (a ApiErrorTypeHandler) ErrorType() error {
	return &errorTypes.ApiError{}
}

func (a ApiErrorTypeHandler) Priority() int {
	return 5
}
