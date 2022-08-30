package api

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/health"
	routeMiddlewares "github.com/NicklasWallgren/go-template/adapters/driven/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/routes"
	routeHandlers "github.com/NicklasWallgren/go-template/adapters/driven/api/routes/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/fx"
)

var routers = fx.Options(
	fx.Provide(users.NewUserRoutes),
	fx.Provide(health.NewHealthRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewRoutes),
)

var controllers = fx.Options(
	fx.Provide(users.NewUserController),
	fx.Provide(health.NewHealthController),
)

var middlewares = fx.Options(
	fx.Provide(routeMiddlewares.NewCorsMiddleware),
	fx.Provide(routeMiddlewares.NewObservabilityMiddleware),
	fx.Provide(routeMiddlewares.NewMiddlewares),
)

var apiConverters = fx.Options(
	fx.Provide(users.NewUserAPIConverter),
	fx.Provide(health.NewHealthAPIConverter),
)

var errorTypeHandlers = fx.Provide(
	fx.Annotate(handlers.NewValidationFieldErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewValidationErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewValidationGoPlaygroundErrorHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewDomainErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers.NewAPIErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
)

var errorResponseManager = fx.Provide(
	fx.Annotate(handlers.NewErrorResponseManager, fx.ParamTags(`group:"error_type_handlers"`)),
)

// A copy of the validator so we can reuse it even if gin.DisableBindValidation() has been called.
var validator binding.StructValidator = binding.Validator

var validators = fx.Options(
	fx.Provide(func() binding.StructValidator { return validator }),
)

// Module exports dependency.
var Module = fx.Options(
	errorTypeHandlers,
	errorResponseManager,
	fx.Provide(common.NewRequestHandler),
	fx.Provide(routeHandlers.NewRootHandler),
	routers,
	controllers,
	middlewares,
	apiConverters,
	validators,
)
