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

func (d DomainErrorTypeHandler) Handle(err error) *response.APIResponseEnvelope {
	if entityNotFoundError := AsEntityNotFoundError(err); entityNotFoundError != nil {
		errorList := []response.APIError{response.NewAPIErrorWithValue(entityNotFoundError.Error(), entityNotFoundError.Id)}

		return response.NewEnvelope(http.StatusNotFound, response.WithResponse(response.NewAPIErrorResponse(errorList)))
	}

	domainError := &domainErrors.DomainError{}
	errors.As(err, &domainError)

	errorList := []response.APIError{response.NewAPIError(domainError.Message)}

	return response.NewEnvelope(http.StatusBadRequest, response.WithResponse(response.NewAPIErrorResponse(errorList)))
}

func (d DomainErrorTypeHandler) IsSupported(err error) bool {
	if AsEntityNotFoundError(err) != nil {
		return true
	}

	domainError := &domainErrors.DomainError{}
	return errors.As(err, &domainError)
}

func (d DomainErrorTypeHandler) ErrorType() error {
	return &domainErrors.DomainError{}
}

func (d DomainErrorTypeHandler) Priority() int {
	return 4
}

func AsEntityNotFoundError(err error) *domainErrors.EntityNotFoundError {
	var entityNotFoundError *domainErrors.EntityNotFoundError
	if errors.As(err, &entityNotFoundError) {
		return entityNotFoundError
	}

	return nil
}
