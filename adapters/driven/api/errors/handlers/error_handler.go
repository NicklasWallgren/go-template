package handlers

import (
	"errors"
	"github.com/mariomac/gostream/order"
	"github.com/mariomac/gostream/stream"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
)

type ErrorTypeResponseHandler interface {
	IsSupported(error error) bool
	Handle(error error) response.ApiResponseEnvelop
	ErrorType() error
	Priority() int
}

type ErrorResponseManager struct {
	errorTypeHandlers []ErrorTypeResponseHandler
}

func NewErrorResponseManager(errorTypeHandlers []ErrorTypeResponseHandler) *ErrorResponseManager {
	// Sort the error handlers based on priority
	sortedTypeHandlers := stream.Sorted(stream.OfSlice(errorTypeHandlers), func(h1, h2 ErrorTypeResponseHandler) int {
		return order.Natural(h1.Priority(), h2.Priority())
	}).ToSlice()

	return &ErrorResponseManager{errorTypeHandlers: sortedTypeHandlers}
}

func (e ErrorResponseManager) Handle(err error) response.ApiResponseEnvelop {
	for _, handler := range e.errorTypeHandlers {
		if !handler.IsSupported(err) {
			continue
		}

		errorType := handler.ErrorType()
		if !errors.As(err, &errorType) {
			// TODO, generic errors, UUID for error tracing, log
			return response.NewApiResponseEnvelop(http.StatusInternalServerError, response.WithPayload("INSERT MESSAGE AND UUID"))
		}

		return handler.Handle(errorType)
	}

	// TODO, generic errors, UUID for error tracing, log
	return response.NewApiResponseEnvelop(http.StatusInternalServerError, response.WithPayload("INSERT MESSAGE AND UUID"))
}
