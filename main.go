package main

import (
	"context"
	"embed"
	"os"
	"os/signal"

	"github.com/NicklasWallgren/go-template/bootstrap"
	"github.com/NicklasWallgren/go-template/config"
)

//go:embed resources/database/migrations/*/*.sql
var embedMigrations embed.FS

// logo is generated via https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=Go%20Template
//
//go:embed resources/assets/logo.ascii
var logo string

// @title GO template API
// @description An template for implementing a hexagonal application.
// @version 1.0
// @host localhost:8080
// @BasePath /api/.
func main() {
	assets := config.NewAssets(logo, embedMigrations)
	app := bootstrap.NewApp(assets)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	app.ExecuteContext(ctx) // nolint:errcheck
}
