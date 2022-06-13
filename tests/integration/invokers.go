package integration

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/routes"
	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/database"
	"github.com/NicklasWallgren/go-template/tests/integration/utils"
	"github.com/gin-gonic/gin"
)

// MigrationUp is a InvokeFunc that should be passed to fx.Invoke
func MigrationUp(migrator database.Migrator) error {
	return migrator.Up() // nolint:wrapcheck
}

// TruncateDatabase is a InvokeFunc that should be passed to fx.Invoke
func TruncateDatabase(db database.Database, config *config.AppConfig) error {
	return utils.TruncateDatabase(db, config) // nolint:wrapcheck
}

// InitializeMiddlewareAndRoutes is a InvokeFunc that should be passed to fx.Invoke
func InitializeMiddlewareAndRoutes(middleware middlewares.Middlewares, route routes.Routes) error {
	middleware.Setup()
	route.Setup()

	// Disables the binding.StructValidator, to be able to chain bind methods and field validation
	// Workaround for see https://github.com/gin-gonic/gin/issues/2535
	gin.DisableBindValidation()

	return nil
}
