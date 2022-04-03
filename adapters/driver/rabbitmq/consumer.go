package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/wagslane/go-rabbitmq"
)

const DefaultExchange = "events"

type ConsumerRunnerOptions struct {
	name         string
	exchangeName string
	exchangeKind string
	queue        string
	routingKeys  []string
}

type ConsumerRunner interface {
	Run(ctx context.Context)
	Do(ctx context.Context)
	Options() ConsumerRunnerOptions
}

type Consumer interface {
	Consume(ctx context.Context, runner ConsumerRunner)
}

type (
	Consumers       []Consumer
	ConsumerRunners []ConsumerRunner
)

type ConsumerManager struct {
	ConsumerRunners ConsumerRunners
}

func NewConsumerManager(consumerRunners ConsumerRunners) *ConsumerManager {
	return &ConsumerManager{ConsumerRunners: consumerRunners}
}

func (c *ConsumerManager) AddConsumerRunner(consumerRunner ConsumerRunner) {
	c.ConsumerRunners = append(c.ConsumerRunners, consumerRunner)
}

func (c ConsumerManager) Run(ctx context.Context) {
	for _, runner := range c.ConsumerRunners {
		go runner.Run(ctx)
	}

	// Wait until all go routines has finished
	<-ctx.Done()

	fmt.Println("stopping consumer")
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

	defer consumer.Close()

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
		log.Fatal(err)
	}

	<-ctx.Done()

	fmt.Println("The consumer is stopping")
}

type SimpleConsumerRunner struct {
	consumer Consumer
}

func NewSimpleConsumerRunner(consumer Consumer) ConsumerRunner {
	return &SimpleConsumerRunner{consumer: consumer}
}

func (s SimpleConsumerRunner) Run(ctx context.Context) {
	s.consumer.Consume(ctx, s)
}

func (s SimpleConsumerRunner) Do(ctx context.Context) {
	fmt.Println("Inside simple consumer runner")

	// TODO, unmarshall body
	// TODO, return rabbitmq.action?
}

func (s SimpleConsumerRunner) Options() ConsumerRunnerOptions {
	return ConsumerRunnerOptions{
		name:         "simple-consumer-runner",
		exchangeName: "events",
		exchangeKind: "topic",
		queue:        "my_queue",
		routingKeys:  []string{"routing_key"},
	}
}
