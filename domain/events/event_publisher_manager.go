package events

import "context"

type EventPublisherManager interface {
	Publish(ctx context.Context, event *Event) error
}

type eventPublisherManager struct {
	publisher EventPublisher // TODO, support multiple?
}

func NewEventPublisherManager(publisher EventPublisher) EventPublisherManager {
	return &eventPublisherManager{publisher: publisher}
}

func (e eventPublisherManager) Publish(ctx context.Context, event *Event) error {
	// TODO, publish using go routine to not block current routine?
	return e.publisher.Publish(ctx, event)
}
