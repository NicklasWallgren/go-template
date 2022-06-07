package handlers

import "github.com/NicklasWallgren/go-template/adapters/driven/api/response"

type ErrorTypeResponseHandler interface {
	IsSupported(error error) bool
	Handle(error error) response.ApiResponseEnvelop
	ErrorType() error
	Priority() int
}
