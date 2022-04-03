package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	err        error
	HttpStatus int
	Message    string
}

// ApiErrorOption definition.
type ApiErrorOption func(apiError *ApiError)

func WithError(err error) ApiErrorOption {
	return func(apiError *ApiError) {
		apiError.err = err
	}
}

func WithMessage(message string) ApiErrorOption {
	return func(apiError *ApiError) {
		apiError.Message = message
	}
}

func WithStatusAndError(httpStatus int, err error) ApiErrorOption {
	return func(apiError *ApiError) {
		apiError.HttpStatus = httpStatus
		apiError.err = err
	}
}

func WithStatusAndMessage(httpStatus int, message string) ApiErrorOption {
	return func(apiError *ApiError) {
		apiError.HttpStatus = httpStatus
		apiError.Message = message
	}
}

func WithStatusAndMessageAndError(httpStatus int, message string, err error) ApiErrorOption {
	return func(apiError *ApiError) {
		apiError.HttpStatus = httpStatus
		apiError.Message = message
		apiError.err = err
	}
}

func NewApiError(message string, httpStatus int) *ApiError {
	return &ApiError{Message: message, HttpStatus: httpStatus}
}

func NewApiErrorWith(options ...ApiErrorOption) *ApiError {
	apiError := &ApiError{HttpStatus: http.StatusInternalServerError, Message: "Unable to process the request, please try again"}

	for _, option := range options {
		option(apiError)
	}

	return apiError
}

func (a ApiError) Error() string {
	return fmt.Sprintf("%s %s", a.Message, a.err) // TODO
}

func (a ApiError) Unwrap() error {
	if a.err != nil {
		return a.err
	}

	return a
}
