package events

import (
	"github.com/NicklasWallgren/go-template/infrastructure/pubsub"
)

type EventPublisher interface {
	Publish(event *Event) error
}

type AmqpEventPublisher struct {
	publisher pubsub.AMQPPublisher
}

func NewAmqpEventPublisher(publisher pubsub.AMQPPublisher) EventPublisher {
	return &AmqpEventPublisher{publisher: publisher}
}

func (r AmqpEventPublisher) Publish(event *Event) error {
	convertedData, err := event.Converter(event)
	if err != nil {
		return err // TODO, wrap in error for more context
	}

	// TODO, wrap error for more context?

	return r.publisher.Publish(convertedData, event.RoutingKey)
}
