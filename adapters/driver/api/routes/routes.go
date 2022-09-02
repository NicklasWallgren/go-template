package routes

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/health"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/users"
)

// Routes contains multiple routes.
type Routes []Route

// Route interface.
type Route interface {
	Setup()
}

// NewRoutes sets up routes.
// nolint: interfacer
func NewRoutes(userRoutes users.UserRoutes, swaggerRoutes SwaggerRoutes, healthRoutes health.HealthRoutes) Routes {
	return Routes{
		userRoutes,
		swaggerRoutes,
		healthRoutes,
	}
}

// Setup all the route.
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
