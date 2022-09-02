package cmd

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/rabbitmq"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/spf13/cobra"
)

type RabbitMQCommand struct{}

func NewRabbitMQCommand() *RabbitMQCommand {
	return &RabbitMQCommand{}
}

func (s *RabbitMQCommand) Use() string {
	return "start-consumers"
}

func (s *RabbitMQCommand) Short() string {
	return "starts pubsub consumers"
}

func (s *RabbitMQCommand) Setup(cmd *cobra.Command) {}

func (s *RabbitMQCommand) Run(cmd *cobra.Command) CommandRunner {
	return func(
		logger logger.Logger,
		consumerManager *rabbitmq.ConsumerManager, // TODO, generic pub sub consumer, not rabbitmq specific?
	) {
		logger.Info("Starting consumers")

		consumerManager.Run(cmd.Context())
	}
}
