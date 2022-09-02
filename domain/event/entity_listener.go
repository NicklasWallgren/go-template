package event

import (
	"context"
	"log"


	"github.com/NicklasWallgren/go-template/adapters/driven/pubsub"

	"github.com/NicklasWallgren/go-template/domain/users/entities"
	userEvent "github.com/NicklasWallgren/go-template/domain/users/event"
)

type EntityConverter[T any] func(entity T) any

// To ensure that EntityEventListener implements the event.Listener interface.
var _ Listener = (*EntityEventListener)(nil)

type EntityEventListener struct {
	publisher pubsub.AMQPPublisher
}

func NewEntityEventListener(publisher pubsub.AMQPPublisher) *EntityEventListener {
	return &EntityEventListener{
		publisher: publisher,
	}
}

func (u EntityEventListener) Listen(ctx context.Context, event Event) {
	actualEntity, ok := event.(EntityEvent)
	if !ok {
		log.Printf("registered an invalid entity event: %T\n", event)
	}

	u.handleEvent(ctx, actualEntity)
}

func (u EntityEventListener) handleEvent(ctx context.Context, event EntityEvent) {
	switch entity := event.Entity.(type) { // nolint: gocritic
	case entities.User:
		u.publisher.Publish(ctx, userEvent.ResponseOf(entity), "routing_key") // nolint: errcheck
	}
}
