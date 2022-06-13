package rabbitmq

import (
	"context"
	"fmt"
)

type ConsumerManager struct {
	ConsumerRunners []ConsumerRunner
}

func NewConsumerManager(consumerRunners []ConsumerRunner) *ConsumerManager {
	return &ConsumerManager{ConsumerRunners: consumerRunners}
}

func (c ConsumerManager) Run(ctx context.Context) {
	for _, runner := range c.ConsumerRunners {
		go runner.Run(ctx)
	}

	// Wait until all go routines has finished
	<-ctx.Done()

	fmt.Println("stopping consumer") // nolint:forbidigo
}
