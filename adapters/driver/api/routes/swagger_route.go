package routes

import (
	routes "github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	"github.com/NicklasWallgren/go-template/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerRoutes struct.
type SwaggerRoutes struct {
	handler routes.RequestHandler
}

// Setup user routes.
func (s SwaggerRoutes) Setup() {
	docs.SwaggerInfo.BasePath = "/api"
	s.handler.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // nolint: wsl
}

// NewSwaggerRoutes creates new user controller.
func NewSwaggerRoutes(handler routes.RequestHandler) SwaggerRoutes {
	return SwaggerRoutes{
		handler: handler,
	}
}
