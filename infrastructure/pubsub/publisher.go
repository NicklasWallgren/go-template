package pubsub

import "context"

type AMQPPublisher interface {
	Publish(ctx context.Context, data any, routingKey string) error
}
