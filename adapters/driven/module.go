package driven

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/migration"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/users"
	"github.com/NicklasWallgren/go-template/adapters/driven/rabbitmq"
	"go.uber.org/fx"
)

var PersistenceRepositories = fx.Options(
	fx.Provide(persistence.NewRepository),
	fx.Provide(users.NewUserRepository),
)

var Module = fx.Options(
	logger.Module,
	fx.Provide(persistence.NewDatabase),
	fx.Provide(migration.NewGooseMigrator),
	PersistenceRepositories,
	health.Module,
	rabbitmq.Module,
)
