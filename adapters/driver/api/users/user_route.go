package users

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/routes"
	routeHandlers "github.com/NicklasWallgren/go-template/adapters/driver/api/routes/handlers"
)

// UserRoutes struct.
type UserRoutes struct {
	logger         logger.Logger
	handler        common.RequestHandler
	userController UserController
	rootHandler    *routeHandlers.RootRouteHandler
}

// Setup user routes.
func (s UserRoutes) Setup() {
	api := s.handler.Gin.Group("/api/users")
	{
		api.GET("/", s.rootHandler.Handle(s.userController.FindAllUsers))
		api.GET("/:id", s.rootHandler.Handle(s.userController.FindOneUserByID))
		api.POST("/", s.rootHandler.Handle(s.userController.CreateUser))
		api.POST("/:id", s.rootHandler.Handle(s.userController.UpdateUser))
		api.DELETE("/:id", s.rootHandler.Handle(s.userController.DeleteUserByID))
	}
}

// NewUserRoutes creates new user routes.
func NewUserRoutes(
	logger logger.Logger,
	handler common.RequestHandler,
	userController UserController,
	rootHandler *routeHandlers.RootRouteHandler,
) routes.Route {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		rootHandler:    rootHandler,
	}
}
