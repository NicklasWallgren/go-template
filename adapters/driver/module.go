package driver

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/users"
	"github.com/NicklasWallgren/go-template/adapters/driver/rabbitmq"
	"go.uber.org/fx"
)

// PersistenceModule exports dependency
var PersistenceModule = fx.Options(
	fx.Provide(users.NewUserRepository),
)

var ConsumerModules = fx.Options(fx.Provide(
	func(consumer rabbitmq.Consumer) rabbitmq.Consumers {
		return []rabbitmq.Consumer{consumer}
	}), fx.Provide(rabbitmq.NewConsumerManager), fx.Provide(rabbitmq.NewRabbitMQConsumer))

var ConsumerRunnerModules = fx.Options(fx.Provide(
	func(simpleConsumerRunner rabbitmq.ConsumerRunner) rabbitmq.ConsumerRunners {
		return []rabbitmq.ConsumerRunner{simpleConsumerRunner}
	}), fx.Provide(rabbitmq.NewSimpleConsumerRunner))

var PublisherModules = fx.Options(fx.Provide(rabbitmq.NewRabbitMQPublisher))

// TODO, use fx.In and fx.Out

// Module exports dependency
var Module = fx.Options(
	PersistenceModule,
	ConsumerModules,
	ConsumerRunnerModules,
	PublisherModules,
)
