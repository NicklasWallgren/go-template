package handlers

import (
	"errors"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	domainErrors "github.com/NicklasWallgren/go-template/domain/errors"
)

type DomainErrorTypeHandler struct{}

func NewDomainErrorTypeHandler() ErrorTypeResponseHandler {
	return &DomainErrorTypeHandler{}
}

func (d DomainErrorTypeHandler) Handle(err error) response.ApiResponseEnvelop {
	domainError := &domainErrors.DomainError{}
	errors.As(err, &domainError)

	errors := []response.ApiErrorConstraint{
		response.NewApiError(domainError.Message),
	}

	return response.NewApiResponseEnvelop(http.StatusBadRequest, response.WithPayload(response.NewApiErrorResponse(errors)))
}

func (d DomainErrorTypeHandler) IsSupported(err error) bool {
	domainError := &domainErrors.DomainError{}
	return errors.As(err, &domainError)
}

func (d DomainErrorTypeHandler) ErrorType() error {
	return &domainErrors.DomainError{}
}
