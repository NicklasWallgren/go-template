package handlers

import (
	"errors"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"

	domainErrors "github.com/NicklasWallgren/go-template/domain/errors"
)

type DomainErrorTypeHandler struct{}

func NewDomainErrorTypeHandler() ErrorTypeResponseHandler {
	return &DomainErrorTypeHandler{}
}

func (d DomainErrorTypeHandler) Handle(err error) *response.APIResponseEnvelope {
	if entityNotFoundError := AsError[domainErrors.EntityNotFoundError](err); entityNotFoundError != nil {
		errorList := []response.APIError{
			response.NewAPIErrorWithField(entityNotFoundError.Error(), "Id", entityNotFoundError.ID),
		}

		return response.NewEnvelope(http.StatusNotFound, response.WithResponse(response.NewAPIErrorResponse(errorList)))
	}

	domainError := &domainErrors.DomainError{}
	errors.As(err, &domainError)

	errorList := []response.APIError{response.NewAPIError(domainError.Message)}

	return response.NewEnvelope(http.StatusBadRequest, response.WithResponse(response.NewAPIErrorResponse(errorList)))
}

func (d DomainErrorTypeHandler) IsSupported(err error) bool {
	if AsError[domainErrors.EntityNotFoundError](err) != nil {
		return true
	}

	domainError := &domainErrors.DomainError{}

	return errors.As(err, &domainError)
}

func (d DomainErrorTypeHandler) Priority() int {
	return 4
}
