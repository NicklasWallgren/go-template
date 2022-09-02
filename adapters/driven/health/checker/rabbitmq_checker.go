package checker

import (
	"context"

	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQHealthChecker struct {
	rabbitmqPublisher *rabbitmq.Publisher // nolint: unused
}

func NewRabbitMQHealthChecker() HealthChecker {
	return &RabbitMQHealthChecker{}
}

func (r RabbitMQHealthChecker) Check(ctx context.Context) Health {
	return NewHealth(Healthy, "rabbitmq") // TODO, implement
}
