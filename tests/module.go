package tests

import (
	"io/fs"
	"os"
	"testing"

	"github.com/NicklasWallgren/go-template/tests/integration/utils"

	"github.com/NicklasWallgren/go-template/adapters/driven/api"
	"github.com/NicklasWallgren/go-template/adapters/driver"
	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain"
	infra "github.com/NicklasWallgren/go-template/infrastructure"
	"github.com/NicklasWallgren/go-template/infrastructure/database"
	"github.com/NicklasWallgren/go-template/infrastructure/env"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/NicklasWallgren/go-template/tests/factories"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func MigrationFs() fs.FS {
	return os.DirFS(utils.TestDirectoryRoot + "/../")
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

// TestPersistenceModule is the test persistence module containing the dependency graph for the persistence module
var TestPersistenceModule = fx.Options(
	fx.Provide(func() env.Env { return env.NewEnvWithPath(utils.TestDirectoryRoot + "/.env") }),
	fx.Provide(func(env env.Env) *config.AppConfig {
		return config.NewAppConfig(&config.Assets{EmbedMigrations: MigrationFs()}, env)
	}),
	fx.Provide(database.NewDatabase), // retrieve from infrastructure module?
	fx.Provide(logger.NewLogger),
	driver.PersistenceModule,
	fx.Provide(factories.NewUserFactory),
	fx.Provide(database.NewGooseMigrator),
)

// ApplicationModule is the application module containing the dependency graph for the application
var ApplicationModule = fx.Options(
	fx.Decorate(func() env.Env { return env.NewEnvWithPath(utils.TestDirectoryRoot + "/.env") }),
	fx.Provide(func(env env.Env) *config.AppConfig {
		return config.NewAppConfig(&config.Assets{EmbedMigrations: MigrationFs()}, env)
	}),
	infra.Module,
	domain.Module,
	api.Module,
	driver.Module,
	fx.Provide(factories.NewUserFactory),
)
