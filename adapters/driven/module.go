package driven

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/env"
	"github.com/NicklasWallgren/go-template/adapters/driven/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/health/checkers"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/migration"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/users"
	"github.com/NicklasWallgren/go-template/adapters/driven/rabbitmq"
	"go.uber.org/fx"
)

var PersistenceModule = fx.Options(
	fx.Provide(users.NewUserRepository),
)

var consumers = fx.Provide(
	fx.Annotate(rabbitmq.NewRabbitMQConsumer, fx.ResultTags(`name:"rabbitmq_consumer"`)),
)

var consumerRunners = fx.Provide(
	fx.Annotate(rabbitmq.NewSimpleConsumerRunner,
		fx.ParamTags(`name:"rabbitmq_consumer"`), fx.ResultTags(`group:"amqp_consumers"`)),
)

var consumerManager = fx.Provide(
	fx.Annotate(rabbitmq.NewConsumerManager, fx.ParamTags(`group:"amqp_consumers"`)),
)

var publishers = fx.Provide(rabbitmq.NewRabbitMQPublisher)

var healthCheckers = fx.Provide(
	fx.Annotate(checkers.NewDBHealthChecker, fx.ResultTags(`group:"checkers"`)),
	fx.Annotate(checkers.NewRabbitMQHealthChecker, fx.ResultTags(`group:"checkers"`)),
)

var healthCheckerManager = fx.Provide(
	fx.Annotate(health.NewHealthCheckerManager, fx.ParamTags(`group:"checkers"`)),
)

var Module = fx.Options(
	PersistenceModule,
	fx.Provide(env.NewEnv),
	fx.Provide(logger.NewLogger),
	fx.Provide(persistence.NewDatabase),
	fx.Provide(migration.NewGooseMigrator),
	healthCheckers,
	healthCheckerManager,
	consumers,
	consumerRunners,
	consumerManager,
	publishers,
)
