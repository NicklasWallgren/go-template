package cmd

import (
	"context"
	"errors"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/routes"
	"log"
	"net/http"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type HTTPServerCommand struct{}

func NewHTTPServerCommand() Command {
	return &HTTPServerCommand{}
}

func (s *HTTPServerCommand) Use() string {
	return "start"
}

func (s *HTTPServerCommand) Short() string {
	return "serve application"
}

func (s *HTTPServerCommand) Setup(cmd *cobra.Command) {}

func (s *HTTPServerCommand) Run(cmd *cobra.Command) CommandRunner {
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

		srv := &http.Server{Addr: ":" + config.HTTPServer.Port, ReadHeaderTimeout: 10 * time.Second, Handler: router.Gin}

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
