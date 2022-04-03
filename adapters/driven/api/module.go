package api

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/errors/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/middlewares"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/routes"
	handlers2 "github.com/NicklasWallgren/go-template/adapters/driven/api/routes/handlers"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/fx"
)

// Module exports dependency to container
var routerModule = fx.Options(
	fx.Provide(users.NewUserRoutes),
	fx.Provide(health.NewHealthRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewRoutes),
)

// Module Controller exported
var controllersModule = fx.Options(
	fx.Provide(users.NewUserController),
	fx.Provide(health.NewHealthController),
)

// Module Middleware exported
var middlewareModule = fx.Options(
	fx.Provide(middlewares.NewCorsMiddleware),
	fx.Provide(middlewares.NewMiddlewares),
)

// Module Converter exported
var converterModule = fx.Options(
	fx.Provide(users.NewUserApiConverter),
	fx.Provide(health.NewHealthApiConverter),
)

// TODO
var errorTypeHandlers = []handlers.ErrorTypeResponseHandler{
	handlers.NewValidationFieldErrorTypeHandler(),
	handlers.NewValidationErrorTypeHandler(),
	handlers.NewValidationGoPlaygroundErrorHandler(),
	handlers.NewDomainErrorTypeHandler(),
	handlers.NewApiErrorTypeHandler(),
}

// Module Error Response exported
var errorResponseModule = fx.Options(
	fx.Provide(func() *handlers.ErrorResponseHandler { return handlers.NewErrorResponseHandler(errorTypeHandlers) }),
)

// A copy of the validator so we can reuse it even if gin.DisableBindValidation() has been called
var validator binding.StructValidator = binding.Validator

// Module Error Response exported
var validatorModule = fx.Options(
	fx.Provide(func() binding.StructValidator { return validator }),
)

// Module exports dependency
var Module = fx.Options(
	errorResponseModule,
	fx.Provide(common.NewRequestHandler),
	fx.Provide(handlers2.NewRootHandler),
	routerModule,
	controllersModule,
	middlewareModule,
	converterModule,
	validatorModule,
)
