package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/wagslane/go-rabbitmq"
)

const DefaultExchange = "events"

type Consumer interface {
	Consume(ctx context.Context, runner ConsumerRunner)
}

type RabbitMQConsumer struct {
	config *config.AppConfig
	logger logger.Logger
}

func NewRabbitMQConsumer(config *config.AppConfig, logger logger.Logger) Consumer {
	return &RabbitMQConsumer{config: config, logger: logger}
}

func (c *RabbitMQConsumer) Consume(ctx context.Context, runner ConsumerRunner) {
	url := fmt.Sprintf("amqp://%s:%s@%s", c.config.RabbitMQ.Name, c.config.RabbitMQ.Password, c.config.RabbitMQ.Host)

	// TODO, pass compatible logger
	consumer, err := rabbitmq.NewConsumer(url, rabbitmq.Config{}, rabbitmq.WithConsumerOptionsLogging)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close() // nolint:errcheck

	runnerOptions := runner.Options()

	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			runner.Do(ctx) // TODO Unmarshall and pass d.Body
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		runnerOptions.queue,
		runnerOptions.routingKeys,
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsBindingExchangeName(runnerOptions.exchangeName),
		rabbitmq.WithConsumeOptionsBindingExchangeKind(runnerOptions.exchangeKind),
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
		rabbitmq.WithConsumeOptionsConsumerName(runnerOptions.name),
	)

	if err != nil {
		log.Println(err)
	}

	<-ctx.Done()

	fmt.Println("The consumer is stopping") // nolint:forbidigo
}
