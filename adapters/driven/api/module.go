package api

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/routes"
	routeHandlers "github.com/NicklasWallgren/go-template/adapters/driven/api/routes/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/fx"
)

var routerModule = fx.Options(
	fx.Provide(users.NewUserRoutes),
	fx.Provide(health.NewHealthRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewRoutes),
)

var controllersModule = fx.Options(
	fx.Provide(users.NewUserController),
	fx.Provide(health.NewHealthController),
)

var middlewareModule = fx.Options(
	fx.Provide(middlewares.NewCorsMiddleware),
	fx.Provide(middlewares.NewObservabilityMiddleware),
	fx.Provide(middlewares.NewMiddlewares),
)

var converterModule = fx.Options(
	fx.Provide(users.NewUserAPIConverter),
	fx.Provide(health.NewHealthAPIConverter),
)

var errorTypeHandlers = fx.Provide(
	fx.Annotate(handlers.NewValidationFieldErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewValidationErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewValidationGoPlaygroundErrorHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewDomainErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewApiErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
)

var errorResponseModule = fx.Provide(
	fx.Annotate(handlers.NewErrorResponseManager, fx.ParamTags(`group:"error_type_handlers"`)),
)

// A copy of the validator so we can reuse it even if gin.DisableBindValidation() has been called.
var validator binding.StructValidator = binding.Validator

var validatorModule = fx.Options(
	fx.Provide(func() binding.StructValidator { return validator }),
)

// Module exports dependency.
var Module = fx.Options(
	errorTypeHandlers,
	errorResponseModule,
	fx.Provide(common.NewRequestHandler),
	fx.Provide(routeHandlers.NewRootHandler),
	routerModule,
	controllersModule,
	middlewareModule,
	converterModule,
	validatorModule,
)
