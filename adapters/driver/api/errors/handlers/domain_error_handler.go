package handlers

import (
	"errors"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"net/http"

	domainErrors "github.com/NicklasWallgren/go-template/domain/errors"
)

type DomainErrorTypeHandler struct{}

func NewDomainErrorTypeHandler() ErrorTypeResponseHandler {
	return &DomainErrorTypeHandler{}
}

func (d DomainErrorTypeHandler) Handle(err error) *response2.APIResponseEnvelope {
	if entityNotFoundError := AsEntityNotFoundError(err); entityNotFoundError != nil {
		errorList := []response2.APIError{response2.NewAPIErrorWithValue(entityNotFoundError.Error(), entityNotFoundError.ID)}

		return response2.NewEnvelope(http.StatusNotFound, response2.WithResponse(response2.NewAPIErrorResponse(errorList)))
	}

	domainError := &domainErrors.DomainError{}
	errors.As(err, &domainError)

	errorList := []response2.APIError{response2.NewAPIError(domainError.Message)}

	return response2.NewEnvelope(http.StatusBadRequest, response2.WithResponse(response2.NewAPIErrorResponse(errorList)))
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
