package handlers

import "github.com/NicklasWallgren/go-template/adapters/driven/api/response"

type ErrorTypeResponseHandler interface {
	IsSupported(err error) bool
	Handle(err error) response.APIResponseEnvelop
	ErrorType() error
	Priority() int
}
