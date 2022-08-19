package health

import (
	"context"

	"github.com/wagslane/go-rabbitmq"
)

// TODO, move to rabbitmq driver package?

type RabbitMQHealthChecker struct {
	rabbitmqPublisher *rabbitmq.Publisher // nolint: unused
}

func NewRabbitMQHealthChecker() HealthChecker {
	return &RabbitMQHealthChecker{}
}

func (r RabbitMQHealthChecker) Check(ctx context.Context) Health {
	return NewHealth(Healthy, "rabbitmq") // TODO, implement
}
