package bootstrap

import (
	"context"
	"log"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driver/cmd"

	"github.com/NicklasWallgren/go-template/adapters/driven"
	"github.com/NicklasWallgren/go-template/adapters/driver/api"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var CommonModules = fx.Options(
	domain.Module,
	api.Module,
	driven.Module,
)

// App root of application.
type App struct {
	*cmd.RootCommand
	assets *config.Assets
}

func NewApp(assets *config.Assets) *App {
	app := &App{RootCommand: cmd.NewRootCommand(assets), assets: assets}
	app.RootCommand.Add(cmd.NewHTTPServerCommand(), app.boot)
	app.RootCommand.Add(cmd.NewMigrationCommand(), app.boot)
	app.RootCommand.Add(cmd.NewRabbitMQCommand(), app.boot)

	return app
}

func (a App) boot(runner cmd.CommandRunner) {
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
