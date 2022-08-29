package domain

import (
	"github.com/NicklasWallgren/go-template/domain/event"
	"github.com/NicklasWallgren/go-template/domain/users"
	"go.uber.org/fx"
)

var eventDispatcher = fx.Provide(func(entityEventListener *event.EntityEventListener) (*event.Dispatcher, error) {
	dispatcher := event.NewDispatcher()

	if err := dispatcher.Register(entityEventListener, event.EntityEventName); err != nil {
		return nil, err
	}

	return dispatcher, nil
})

var Module = fx.Options(
	fx.Provide(event.NewEntityEventListener),
	eventDispatcher,
	fx.Provide(users.NewUserService),
	fx.Provide(users.NewUserValidator),
)
