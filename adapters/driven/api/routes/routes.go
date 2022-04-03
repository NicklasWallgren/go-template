package routes

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users"
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(userRoutes users.UserRoutes, swaggerRoutes SwaggerRoutes, healthRoutes health.HealthRoutes) Routes {
	return Routes{
		userRoutes,
		swaggerRoutes,
		healthRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
