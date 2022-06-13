package handlers

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/gin-gonic/gin"
)

type RouteHandler func(c *gin.Context) (response.ApiResponseEnvelop, error)

type RootHandler struct {
	ErrorHandler handlers.ErrorResponseManager
}

func NewRootHandler(errorHandler handlers.ErrorResponseManager) *RootHandler {
	return &RootHandler{ErrorHandler: errorHandler}
}

func (r RootHandler) Handle(handler RouteHandler) gin.HandlerFunc {
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
