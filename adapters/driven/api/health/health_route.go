package health

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	routeHandlers "github.com/NicklasWallgren/go-template/adapters/driven/api/routes/handlers"
	logger "github.com/NicklasWallgren/go-template/infrastructure/logger"
)

// HealthRoutes struct.
type HealthRoutes struct {
	logger           logger.Logger
	handler          common.RequestHandler
	healthController HealthController
	rootHandler      *routeHandlers.RootHandler
}

// Setup health routes.
func (s HealthRoutes) Setup() {
	api := s.handler.Gin.Group("/api/health")
	{
		api.GET("/", s.rootHandler.Handle(s.healthController.Health))
	}
}

// NewHealthRoutes creates new health routes.
func NewHealthRoutes(
	logger logger.Logger,
	handler common.RequestHandler,
	healthController HealthController,
	rootHandler *routeHandlers.RootHandler,
) HealthRoutes {
	return HealthRoutes{
		handler:          handler,
		logger:           logger,
		healthController: healthController,
		rootHandler:      rootHandler,
	}
}
