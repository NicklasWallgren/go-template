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
	healthCheckers,
	healthCheckerManager,
)

var healthCheckers = fx.Provide(
	fx.Annotate(health.NewDBHealthChecker, fx.ResultTags(`group:"checkers"`)),
	fx.Annotate(health.NewRabbitMQHealthChecker, fx.ResultTags(`group:"checkers"`)),
)

var healthCheckerManager = fx.Provide(
	fx.Annotate(health.NewHealthCheckerManager, fx.ParamTags(`group:"checkers"`)),
)
