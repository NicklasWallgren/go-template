package pubsub

type AMQPPublisher interface {
	Publish(data any, routingKey string) error
}
