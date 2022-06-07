package handlers

import (
	"errors"
	"github.com/mariomac/gostream/order"
	"github.com/mariomac/gostream/stream"
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
)

type ErrorResponseManager interface {
	Handle(err error) response.ApiResponseEnvelop
}

type errorResponseManager struct {
	errorTypeHandlers []ErrorTypeResponseHandler
}

func NewErrorResponseManager(errorTypeHandlers []ErrorTypeResponseHandler) ErrorResponseManager {
	// Sort the error handlers based on priority
	sortedTypeHandlers := stream.Sorted(stream.OfSlice(errorTypeHandlers), func(h1, h2 ErrorTypeResponseHandler) int {
		return order.Natural(h1.Priority(), h2.Priority())
	}).ToSlice()

	return &errorResponseManager{errorTypeHandlers: sortedTypeHandlers}
}

func (e errorResponseManager) Handle(err error) response.ApiResponseEnvelop {
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
