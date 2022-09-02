package rabbitmq

import "go.uber.org/fx"

var consumers = fx.Provide(
	fx.Annotate(NewRabbitMQConsumer, fx.ResultTags(`name:"rabbitmq_consumer"`)),
)

var consumerRunners = fx.Provide(
	fx.Annotate(NewSimpleConsumerRunner,
		fx.ParamTags(`name:"rabbitmq_consumer"`), fx.ResultTags(`group:"amqp_consumers"`)),
)

var consumerManager = fx.Provide(
	fx.Annotate(NewConsumerManager, fx.ParamTags(`group:"amqp_consumers"`)),
)

var publishers = fx.Provide(NewRabbitMQPublisher)

var Module = fx.Options(
	consumers,
	consumerRunners,
	consumerManager,
	publishers,
)
