package infrastructure

import (
	database "github.com/NicklasWallgren/go-template/infrastructure/database"
	"github.com/NicklasWallgren/go-template/infrastructure/env"
	"github.com/NicklasWallgren/go-template/infrastructure/health"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(env.NewEnv),
	fx.Provide(logger.NewLogger),
	fx.Provide(database.NewDatabase),
	fx.Provide(database.NewGooseMigrator),
	dbChecker,
	healthCheckerManager,
	healthCheckers,
)

var dbChecker = fx.Provide(fx.Annotated{
	Name:   "dbChecker",
	Target: health.NewDBHealthChecker,
})

type HealthCheckerParams struct {
	fx.In
	DBChecker health.HealthChecker `name:"dbChecker"`
}

var healthCheckers = fx.Provide(func(healthCheckers HealthCheckerParams) []health.HealthChecker {
	return []health.HealthChecker{healthCheckers.DBChecker}
})

var healthCheckerManager = fx.Provide(health.NewHealthCheckerManager)
