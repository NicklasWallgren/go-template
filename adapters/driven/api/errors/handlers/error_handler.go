package handlers

import (
	"net/http"

	"github.com/mariomac/gostream/order"
	"github.com/mariomac/gostream/stream"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
)

type ErrorResponseManager interface {
	Handle(err error) *response.APIResponseEnvelope
}

type errorResponseManager struct {
	errorTypeHandlers []ErrorTypeResponseHandler
}

func NewErrorResponseManager(errorTypeHandlers []ErrorTypeResponseHandler) ErrorResponseManager {
	// Sort the error handlers by priority
	sortedTypeHandlers := stream.Sorted(stream.OfSlice(errorTypeHandlers), func(h1, h2 ErrorTypeResponseHandler) int {
		return order.Natural(h1.Priority(), h2.Priority())
	}).ToSlice()

	return &errorResponseManager{errorTypeHandlers: sortedTypeHandlers}
}

func (e errorResponseManager) Handle(err error) *response.APIResponseEnvelope {
	for _, handler := range e.errorTypeHandlers {
		if !handler.IsSupported(err) {
			continue
		}

		// errorType := handler.ErrorType()
		// if !errors.As(err, &errorType) { // nolint: govet
		//	// TODO, generic errors, UUID for error tracing, log
		//	return response.NewEnvelope(http.StatusInternalServerError, response.WithResponse("INSERT MESSAGE AND UUID"))
		//}

		return handler.Handle(err)
	}

	// TODO, generic errors, UUID for error tracing, log
	return response.NewEnvelope(http.StatusInternalServerError, response.WithResponse("INSERT MESSAGE AND UUID"))
}
