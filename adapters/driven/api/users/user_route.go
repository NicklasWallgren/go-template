package users

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	routeHandlers "github.com/NicklasWallgren/go-template/adapters/driven/api/routes/handlers"
	logger "github.com/NicklasWallgren/go-template/infrastructure/logger"
)

// UserRoutes struct.
type UserRoutes struct {
	logger         logger.Logger
	handler        common.RequestHandler
	userController UserController
	rootHandler    *routeHandlers.RootHandler
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
	rootHandler *routeHandlers.RootHandler,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		rootHandler:    rootHandler,
	}
}
