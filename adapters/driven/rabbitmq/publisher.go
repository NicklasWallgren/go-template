package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/NicklasWallgren/go-template/adapters/driven/pubsub"
	"log"

	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/wagslane/go-rabbitmq"
)

// To ensure that RabbitMQPublisher implements the pubsub.AMQPPublisher interface.
var _ pubsub.AMQPPublisher = (*RabbitMQPublisher)(nil)

type RabbitMQPublisher struct {
	rabbitmqPublisher *rabbitmq.Publisher
	initialized       bool
	config            *config.RabbitMQ
	logger            logger.Logger
}

func NewRabbitMQPublisher(config *config.AppConfig, logger logger.Logger) pubsub.AMQPPublisher {
	return &RabbitMQPublisher{config: config.RabbitMQ, logger: logger}
}

func (r *RabbitMQPublisher) Publish(ctx context.Context, data any, routingKey string) error {
	// Lazy initialize to improve bootup time of the application
	if err := r.lazyLoad(); err != nil {
		return err
	}

	// TODO, handle error before emitting the span
	opts := []ddtrace.StartSpanOption{
		tracer.ServiceName("rabbitmq"),
		tracer.ResourceName("publish/" + routingKey),
		tracer.SpanType(ext.SpanTypeMessageProducer),
		tracer.Tag("amqp.command", "basic.publish"),
		tracer.Tag("amqp.exchange", "events"), // TODO
		tracer.Tag("amqp.routing_key", routingKey),
		tracer.Measured(),
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "amqp.publish", opts...) // nolint:ineffassign, staticcheck
	defer span.Finish()

	dataByteArray, err := json.Marshal(data)
	if err != nil {
		return err // TODO, wrap in error
	}

	r.logger.Info("Publishing to rabbitmq")

	return r.rabbitmqPublisher.Publish(
		dataByteArray,
		[]string{routingKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		// rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
	)
}

func (r *RabbitMQPublisher) setup() (err error) {
	r.rabbitmqPublisher, err = rabbitmq.NewPublisher(r.config.ToDsn(), rabbitmq.Config{},
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeKind("topic"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return err
	}

	confirmations := r.rabbitmqPublisher.NotifyPublish()
	go func() { // nolint: wsl
		for c := range confirmations {
			// TODO, handle confirmation from server. Retry on nack?
			// TODO, use correct logger
			log.Printf("message confirmed from server. tag: %v, ack: %v", c.DeliveryTag, c.Ack)
		}
	}()

	return nil
}

func (r *RabbitMQPublisher) lazyLoad() error {
	if !r.initialized {
		if err := r.setup(); err != nil {
			return err // TODO, wrap in error for more context
		}
	}

	return nil
}
