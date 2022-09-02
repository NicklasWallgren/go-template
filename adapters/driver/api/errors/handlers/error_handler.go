package handlers

import (
	"fmt"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"net/http"

	"github.com/google/uuid"

	"github.com/mariomac/gostream/order"
	"github.com/mariomac/gostream/stream"
)

type ErrorResponseManager interface {
	Handle(err error) *response2.APIResponseEnvelope
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

func (e errorResponseManager) Handle(err error) *response2.APIResponseEnvelope {
	for _, handler := range e.errorTypeHandlers {
		if !handler.IsSupported(err) {
			continue
		}

		return handler.Handle(err)
	}

	return response2.NewEnvelope(http.StatusInternalServerError,
		response2.WithResponse(fmt.Sprintf("Error occurred with id %s, please try again", uuid.New())))
}
