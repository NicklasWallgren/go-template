package handlers

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"github.com/gin-gonic/gin"
)

type RouteHandler func(c *gin.Context) (*response.APIResponseEnvelope, error)

type RootRouteHandler struct {
	ErrorHandler handlers.ErrorResponseManager
}

func NewRootRouteHandler(errorHandler handlers.ErrorResponseManager) *RootRouteHandler {
	return &RootRouteHandler{ErrorHandler: errorHandler}
}

func (r RootRouteHandler) Handle(handler RouteHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiResponseEnvelop, err := handler(ctx)
		if err != nil {
			handleErrorResponse(ctx, err, r.ErrorHandler)

			return
		}

		if apiResponseEnvelop.Response() != nil {
			ctx.JSON(apiResponseEnvelop.Status(), apiResponseEnvelop.Response())

			return
		}

		ctx.Status(apiResponseEnvelop.Status())
	}
}

func handleErrorResponse(c *gin.Context, err error, errorHandler handlers.ErrorResponseManager) {
	apiResponse := errorHandler.Handle(err)

	// TODO, should we use c.AbortWithStatus if 4xx,5xx?

	c.JSON(apiResponse.Status(), apiResponse.Response())
}
