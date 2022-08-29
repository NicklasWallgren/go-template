package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

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

		return handler.Handle(err)
	}

	return response.NewEnvelope(http.StatusInternalServerError,
		response.WithResponse(fmt.Sprintf("Error occurred with id %s, please try again", uuid.New())))
}
