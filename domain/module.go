package domain

import (
	"github.com/NicklasWallgren/go-template/domain/events"
	"github.com/NicklasWallgren/go-template/domain/users"
	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(users.NewUserService),
	fx.Provide(users.NewUserValidator),
	fx.Provide(events.NewAmqpEventPublisher),
	fx.Provide(events.NewEventPublisherManager),
)
