package health

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/health/checker"
	"go.uber.org/fx"
)

var healthCheckers = fx.Provide(
	fx.Annotate(checker.NewDBHealthChecker, fx.ResultTags(`group:"checkers"`)),
	fx.Annotate(checker.NewRabbitMQHealthChecker, fx.ResultTags(`group:"checkers"`)),
)

var healthCheckerManager = fx.Provide(
	fx.Annotate(NewHealthCheckerManager, fx.ParamTags(`group:"checkers"`)),
)

var Module = fx.Options(
	healthCheckers,
	healthCheckerManager,
)
