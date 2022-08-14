package errors

import (
	"fmt"
	"net/http"
)

type APIError struct {
	err        error
	HTTPStatus int
	Message    string
}

// APIErrorOption definition.
type APIErrorOption func(apiError *APIError)

func WithError(err error) APIErrorOption {
	return func(apiError *APIError) {
		apiError.err = err
	}
}

func WithMessage(message string) APIErrorOption {
	return func(apiError *APIError) {
		apiError.Message = message
	}
}

func WithStatusAndError(httpStatus int, err error) APIErrorOption {
	return func(apiError *APIError) {
		apiError.HTTPStatus = httpStatus
		apiError.err = err
	}
}

func WithStatusAndMessage(httpStatus int, message string) APIErrorOption {
	return func(apiError *APIError) {
		apiError.HTTPStatus = httpStatus
		apiError.Message = message
	}
}

func WithStatusAndMessageAndError(httpStatus int, message string, err error) APIErrorOption {
	return func(apiError *APIError) {
		apiError.HTTPStatus = httpStatus
		apiError.Message = message
		apiError.err = err
	}
}

func NewApiError(options ...APIErrorOption) *APIError {
	apiError := &APIError{
		HTTPStatus: http.StatusInternalServerError,
		Message:    "Unable to process the request, please try again",
	}

	for _, option := range options {
		option(apiError)
	}

	return apiError
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s %s", a.Message, a.err) // TODO
}

func (a APIError) Unwrap() error {
	if a.err != nil {
		return a.err
	}

	return a
}
