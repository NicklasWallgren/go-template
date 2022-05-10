package health

import (
	"context"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQHealthChecker struct {
	rabbitmqPublisher *rabbitmq.Publisher
}

func NewRabbitMQHealthChecker() HealthChecker {
	return &RabbitMQHealthChecker{}
}

func (r RabbitMQHealthChecker) Check(ctx context.Context) Health {
	// TODO, implement

	return NewHealth(Healthy, "rabbitmq")
}
