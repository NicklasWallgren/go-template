package checkers

import (
	"context"
	"github.com/NicklasWallgren/go-template/adapters/driven/health"

	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQHealthChecker struct {
	rabbitmqPublisher *rabbitmq.Publisher // nolint: unused
}

func NewRabbitMQHealthChecker() health.HealthChecker {
	return &RabbitMQHealthChecker{}
}

func (r RabbitMQHealthChecker) Check(ctx context.Context) health.Health {
	return health.NewHealth(health.Healthy, "rabbitmq") // TODO, implement
}
