package api

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/health"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/routes"
	routeHandler "github.com/NicklasWallgren/go-template/adapters/driver/api/routes/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/users"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/fx"
)

var routers = fx.Provide(
	fx.Annotate(users.NewUserRoutes, fx.ResultTags(`group:"routers"`)),
	fx.Annotate(health.NewHealthRoutes, fx.ResultTags(`group:"routers"`)),
	fx.Annotate(routes.NewSwaggerRoutes, fx.ResultTags(`group:"routers"`)),
)

var route = fx.Provide(
	fx.Annotate(routes.NewRoutes, fx.ParamTags(`group:"routers"`)),
)

var controllers = fx.Options(
	fx.Provide(users.NewUserController),
	fx.Provide(health.NewHealthController),
)

var httpMiddlewares = fx.Provide(
	fx.Annotate(middlewares.NewCorsMiddleware, fx.ResultTags(`group:"http_middlewares"`)),
	fx.Annotate(middlewares.NewObservabilityMiddleware, fx.ResultTags(`group:"http_middlewares"`)),
)

var httpMiddleware = fx.Provide(
	fx.Annotate(middlewares.NewMiddlewares, fx.ParamTags(`group:"http_middlewares"`)),
)

var apiConverters = fx.Options(
	fx.Provide(users.NewUserAPIConverter),
	fx.Provide(users.NewUserOverviewAPIConverter),
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
	fx.Provide(routeHandler.NewRootRouteHandler),
	routers,
	route,
	controllers,
	httpMiddlewares,
	httpMiddleware,
	apiConverters,
	validators,
)
