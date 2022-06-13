package integration

import (
	"testing"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/tests"
	"github.com/NicklasWallgren/go-template/tests/integration/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

type (
	TestRunner any
	InvokeFunc any
)

func WithPersistence() fx.Option {
	return tests.TestPersistenceModule
}

func WithPersistenceAndApplyMigration(opts ...fx.Option) fx.Option {
	return fx.Options(tests.TestPersistenceModule, fx.Options(opts...), fx.Invoke(MigrationUp))
}

func WithApplicationAndApplyMigration(opts ...fx.Option) fx.Option {
	return fx.Options(tests.ApplicationModule, fx.Options(opts...), fx.Invoke(MigrationUp))
}

func WithDatabaseName(t *testing.T, testFuncName string) fx.Option {
	t.Helper()

	return fx.Decorate(func(appConfig *config.AppConfig) *config.AppConfig {
		databaseName, err := utils.CreateDatabase(utils.ToDatabaseName(testFuncName), appConfig)

		utils.AssertNilOrFail(t, err)

		t.Cleanup(func() {
			// TODO, should we truncate instead? Add options to either truncate or drop, more performant to keep database
			utils.DropDatabase(databaseName, appConfig) // nolint:errcheck, gosec
		})

		// Makes the new database name available in the application context
		appConfig.Database.Name = databaseName

		return appConfig
	})
}

func Runner(tb testing.TB, test TestRunner, option fx.Option, initializers ...interface{}) *fx.App {
	tb.Helper()

	runner := tests.NewForTest(tb, option, fx.Invoke(initializers...), fx.Invoke(test))

	require.NoError(tb, runner.Err())

	return runner
}
