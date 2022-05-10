package handlers

import (
	"errors"
	errorTypes "github.com/NicklasWallgren/go-template/adapters/driven/api/errors"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
)

type ApiErrorTypeHandler struct{}

// To ensure that ApiErrorTypeHandler implements the ErrorTypeResponseHandler interface
var _ ErrorTypeResponseHandler = (*ApiErrorTypeHandler)(nil)

func NewApiErrorTypeHandler() *ApiErrorTypeHandler {
	return &ApiErrorTypeHandler{}
}

func (a ApiErrorTypeHandler) Handle(err error) response.ApiResponseEnvelop {
	actualError, _ := (err).(*errorTypes.ApiError)

	errors := []response.ApiErrorConstraint{
		response.NewApiError(actualError.Message),
	}

	return response.NewApiResponseEnvelop(actualError.HttpStatus, response.WithPayload(response.NewApiErrorResponse(errors)))
}

func (a ApiErrorTypeHandler) IsSupported(err error) bool {
	domainError := &errorTypes.ApiError{}
	return errors.As(err, &domainError)
}

func (a ApiErrorTypeHandler) ErrorType() error {
	return &errorTypes.ApiError{}
}
