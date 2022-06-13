package bootstrap

import (
	"context"
	"log"

	"github.com/NicklasWallgren/go-template/adapters/driven/api"
	"github.com/NicklasWallgren/go-template/adapters/driver"
	"github.com/NicklasWallgren/go-template/cmd"
	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain"
	infra "github.com/NicklasWallgren/go-template/infrastructure"
	"github.com/NicklasWallgren/go-template/infrastructure/cli"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var CommonModules = fx.Options(
	infra.Module,
	domain.Module,
	api.Module,
	driver.Module,
)

// App root of application.
type App struct {
	*cli.RootCommand
	assets *config.Assets
}

func NewApp(assets *config.Assets) *App {
	app := &App{RootCommand: cli.NewRootCommand(assets), assets: assets}
	app.RootCommand.Add(cmd.NewHttpServerCommand(), app.boot)
	app.RootCommand.Add(cmd.NewMigrationCommand(), app.boot)
	app.RootCommand.Add(cmd.NewRabbitMQCommand(), app.boot)

	return app
}

func (a App) boot(runner cli.CommandRunner) {
	opts := fx.Options(
		fx.WithLogger(func(logger logger.Logger) fxevent.Logger {
			return logger.GetFxLogger()
		}),
		fx.Invoke(runner),
	)
	ctx := context.Background()
	app := fx.New(fx.Provide(func() *config.Assets { return a.assets }), config.Module, CommonModules, opts)

	if err := app.Start(ctx); err != nil {
		log.Println(err)
	}
	defer app.Stop(ctx) // nolint:errcheck
}
