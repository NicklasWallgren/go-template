package driven

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	persistence "github.com/NicklasWallgren/go-template/adapters/driven/persistence"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/migration"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/users"
	"github.com/NicklasWallgren/go-template/adapters/driven/rabbitmq"
	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	sqlTemplate "github.com/NicklasWallgren/sqlTemplate/pkg"
	"go.uber.org/fx"
)

var QueryTemplateEngine = fx.Provide(func(config *config.AppConfig) (sqlTemplate.QueryTemplateEngine, error) {
	sqlT := sqlTemplate.NewQueryTemplateEngine()

	if err := sqlT.Register("users", config.Assets.TemplateSQL, ".tsql"); err != nil {
		return nil, err
	}

	return sqlT, nil
})

var userEntityRepository = fx.Annotate(func(
	db persistence.Database,
	logger logger.Logger,
	config *config.AppConfig,
) persistence.EntityRepository[entities.User] {
	return persistence.NewEntityRepository[entities.User](db, entities.User{}, logger, config)
}, fx.ResultTags(`name:"userEntityRepository"`))

var PersistenceRepositories = fx.Options(
	fx.Provide(persistence.NewRepository),
	fx.Provide(userEntityRepository),
	fx.Provide(fx.Annotate(users.NewUserRepository, fx.ParamTags(`name:"userEntityRepository"`))),
)

var Module = fx.Options(
	logger.Module,
	fx.Provide(persistence.NewDatabase),
	fx.Provide(migration.NewGooseMigrator),
	PersistenceRepositories,
	health.Module,
	rabbitmq.Module,
	QueryTemplateEngine,
)
