package routes

import (
	routes "github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	docs "github.com/NicklasWallgren/go-template/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerRoutes struct
type SwaggerRoutes struct {
	handler routes.RequestHandler
}

// Setup user routes
func (s SwaggerRoutes) Setup() {
	docs.SwaggerInfo.BasePath = ""
	s.handler.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// NewSwaggerRoutes creates new user controller
func NewSwaggerRoutes(handler routes.RequestHandler) SwaggerRoutes {
	return SwaggerRoutes{
		handler: handler,
	}
}
