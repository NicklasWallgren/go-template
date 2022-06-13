package driver

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/users"
	"github.com/NicklasWallgren/go-template/adapters/driver/rabbitmq"
	"go.uber.org/fx"
)

var PersistenceModule = fx.Options(
	fx.Provide(users.NewUserRepository),
)

var consumers = fx.Provide(
	fx.Annotate(rabbitmq.NewRabbitMQConsumer, fx.ResultTags(`name:"rabbitmq_consumer"`)),
)

var consumerRunners = fx.Provide(
	fx.Annotate(rabbitmq.NewSimpleConsumerRunner,
		fx.ParamTags(`name:"rabbitmq_consumer"`), fx.ResultTags(`group:"amqp_consumers"`)),
)

var consumerManager = fx.Provide(
	fx.Annotate(rabbitmq.NewConsumerManager, fx.ParamTags(`group:"amqp_consumers"`)),
)

var publishers = fx.Provide(rabbitmq.NewRabbitMQPublisher)

var Module = fx.Options(
	PersistenceModule,
	consumers,
	consumerRunners,
	consumerManager,
	publishers,
)
