package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/pubsub"
	"github.com/wagslane/go-rabbitmq"
)

// To ensure that RabbitMQPublisher implements the pubsub.AMQPPublisher interface
var _ pubsub.AMQPPublisher = (*RabbitMQPublisher)(nil)

type RabbitMQPublisher struct {
	rabbitmqPublisher *rabbitmq.Publisher
	initialized       bool
	config            *config.RabbitMQ
}

func NewRabbitMQPublisher(config *config.AppConfig) pubsub.AMQPPublisher {
	return &RabbitMQPublisher{config: config.RabbitMQ}
}

func (r *RabbitMQPublisher) Publish(data any, routingKey string) error {
	// Lazy initialize to improve bootup time of the application
	if err := r.lazyLoad(); err != nil {
		return err
	}

	dataByteArray, err := json.Marshal(data)
	if err != nil {
		return err // TODO, wrap in error
	}

	return r.rabbitmqPublisher.Publish(
		dataByteArray,
		[]string{routingKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		// rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("events"),
	)
}

func (r *RabbitMQPublisher) setup() (err error) {
	r.rabbitmqPublisher, err = rabbitmq.NewPublisher(r.config.ToDsn(), rabbitmq.Config{}, rabbitmq.WithPublisherOptionsLogging)
	if err != nil {
		return err
	}

	confirmations := r.rabbitmqPublisher.NotifyPublish()
	go func() {
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
