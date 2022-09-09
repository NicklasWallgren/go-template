package tests

import (
	"io/fs"
	"os"
	"testing"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/routes"
	routeHandler "github.com/NicklasWallgren/go-template/adapters/driver/api/routes/handlers"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/migration"

	"github.com/NicklasWallgren/go-template/adapters/driven"
	"github.com/NicklasWallgren/go-template/adapters/driven/env"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
	"github.com/NicklasWallgren/go-template/adapters/driver/api"

	"github.com/NicklasWallgren/go-template/tests/integration/utils"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain"
	"github.com/NicklasWallgren/go-template/tests/factories"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func MigrationFs() fs.FS {
	return os.DirFS(utils.TestDirectoryRoot + "/../")
}

func TemplateSQLFs() fs.FS {
	return os.DirFS(utils.TestDirectoryRoot + "/../resources/database/sql")
}

func NewForTest(tb testing.TB, opts ...fx.Option) *fx.App {
	tb.Helper()

	testOpts := []fx.Option{
		fx.Logger(fxtest.NewTestPrinter(tb)),
		fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(tb) }),
	}
	opts = append(testOpts, opts...)

	return fx.New(opts...)
}

// DefaultModule contains the dependency graph for the default application components.
var DefaultModule = fx.Options(
	fx.Provide(func() env.Env { return env.NewEnvWithPath(utils.TestDirectoryRoot + "/.env") }),
	fx.Provide(func(env env.Env) *config.AppConfig {
		return config.NewAppConfig(&config.AssetsConfig{EmbedMigrations: MigrationFs(), TemplateSQL: TemplateSQLFs()}, env)
	}),
)

// TestPersistenceModule contains the dependency graph for the persistence components.
var TestPersistenceModule = fx.Options(
	DefaultModule,

	logger.Module,
	fx.Provide(persistence.NewDatabase), // retrieve from infrastructure module?
	driven.PersistenceRepositories,
	driven.QueryTemplateEngine,
	fx.Provide(factories.NewUserFactory),
	fx.Provide(migration.NewGooseMigrator),
)

// ApplicationModule contains the dependency graph for the application components.
var ApplicationModule = fx.Options(
	DefaultModule,

	domain.Module,
	api.Module,
	driven.Module,
	fx.Provide(factories.NewUserFactory),
)

// DefaultHTTPModule contains the dependency graph for the default http api components.
var DefaultHTTPModule = fx.Options(
	logger.Module,
	fx.Provide(fx.Annotate(routes.NewRoutes, fx.ParamTags(`group:"routers"`))),
	fx.Provide(fx.Annotate(middlewares.NewMiddlewares, fx.ParamTags(`group:"http_middlewares"`))),
	fx.Provide(fx.Annotate(handlers.NewErrorResponseManager, fx.ParamTags(`group:"error_type_handlers"`))),
	fx.Provide(routeHandler.NewRootRouteHandler),
	fx.Provide(common.NewRequestHandler),
)
