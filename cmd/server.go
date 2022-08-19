package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/routes"
	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/cli"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type HttpServerCommand struct{}

func NewHttpServerCommand() cli.Command {
	return &HttpServerCommand{}
}

func (s *HttpServerCommand) Use() string {
	return "start"
}

func (s *HttpServerCommand) Short() string {
	return "serve application"
}

func (s *HttpServerCommand) Setup(cmd *cobra.Command) {}

func (s *HttpServerCommand) Run(cmd *cobra.Command) cli.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		router common.RequestHandler,
		route routes.Routes,
		logger logger.Logger,
		config *config.AppConfig,
	) {
		// TODO, pass via env
		tracer.Start(
			tracer.WithEnv("gotemplate-dev"),
			tracer.WithService("go-template"),
			tracer.WithServiceVersion("0.10.0"),
		)
		defer tracer.Stop()

		logger.Info(config.Assets.Logo)
		middleware.Setup()
		route.Setup()

		// Disables the binding.StructValidator, use the one defined in FX context instead
		gin.DisableBindValidation()

		srv := &http.Server{Addr: ":" + config.HttpServer.Port, ReadHeaderTimeout: 10 * time.Second, Handler: router.Gin}

		// Initializing the server in a goroutine so that it won't block
		// See validatorModule in adapters/driven/api/module.go
		go func() {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Listen for the interrupt signal.
		<-cmd.Context().Done()

		logger.Info("Shutting down the HTTP server")

		// TODO, ensure that this actually works

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // nolint: gomnd
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}
}
