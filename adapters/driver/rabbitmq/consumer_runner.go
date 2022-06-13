package rabbitmq

import (
	"context"
	"fmt"
)

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
	fmt.Println("Inside simple consumer runner") // nolint:forbidigo

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
