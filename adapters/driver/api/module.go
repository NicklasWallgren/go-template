package api

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	handlers2 "github.com/NicklasWallgren/go-template/adapters/driver/api/errors/handlers"
	health2 "github.com/NicklasWallgren/go-template/adapters/driver/api/health"
	middlewares2 "github.com/NicklasWallgren/go-template/adapters/driver/api/middlewares"
	routes2 "github.com/NicklasWallgren/go-template/adapters/driver/api/routes"
	routeHandlers "github.com/NicklasWallgren/go-template/adapters/driver/api/routes/handlers"
	users2 "github.com/NicklasWallgren/go-template/adapters/driver/api/users"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/fx"
)

var routers = fx.Options(
	fx.Provide(users2.NewUserRoutes),
	fx.Provide(health2.NewHealthRoutes),
	fx.Provide(routes2.NewSwaggerRoutes),
	fx.Provide(routes2.NewRoutes),
)

var controllers = fx.Options(
	fx.Provide(users2.NewUserController),
	fx.Provide(health2.NewHealthController),
)

var middlewares = fx.Options(
	fx.Provide(middlewares2.NewCorsMiddleware),
	fx.Provide(middlewares2.NewObservabilityMiddleware),
	fx.Provide(middlewares2.NewMiddlewares),
)

var apiConverters = fx.Options(
	fx.Provide(users2.NewUserAPIConverter),
	fx.Provide(health2.NewHealthAPIConverter),
)

var errorTypeHandlers = fx.Provide(
	fx.Annotate(handlers2.NewValidationFieldErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers2.NewValidationErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers2.NewValidationGoPlaygroundErrorHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers2.NewDomainErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
	fx.Annotate(handlers2.NewAPIErrorTypeHandler, fx.ResultTags(`group:"error_type_handlers"`)),
)

var errorResponseManager = fx.Provide(
	fx.Annotate(handlers2.NewErrorResponseManager, fx.ParamTags(`group:"error_type_handlers"`)),
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
