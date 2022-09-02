package middlewares

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/env"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

// ObservabilityMiddleware middleware for cors.
type ObservabilityMiddleware struct {
	handler common.RequestHandler
	logger  logger.Logger
	env     env.Env
}

// NewObservabilityMiddleware creates new cors middleware.
func NewObservabilityMiddleware(
	handler common.RequestHandler, logger logger.Logger, env env.Env,
) ObservabilityMiddleware {
	return ObservabilityMiddleware{
		handler: handler,
		logger:  logger,
		env:     env, // TODO, use appConfig instead
	}
}

// Setup sets up observability middleware.
func (m ObservabilityMiddleware) Setup() {
	m.logger.Info("Setting up observability middleware")

	// Use the tracer middleware with your desired service name.
	m.handler.Gin.Use(gintrace.Middleware("go-template"))
}
