package middlewares

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/env"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	cors "github.com/rs/cors/wrapper/gin"
)

// CorsMiddleware middleware for cors.
type CorsMiddleware struct {
	handler common.RequestHandler
	logger  logger.Logger
	env     env.Env
}

// NewCorsMiddleware creates new cors middleware.
func NewCorsMiddleware(handler common.RequestHandler, logger logger.Logger, env env.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware.
func (m CorsMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	debug := m.env.Environment == "development"
	m.handler.Gin.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		Debug:            debug,
	}))
}
