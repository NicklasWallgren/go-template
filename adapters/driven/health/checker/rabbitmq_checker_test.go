package checker

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRabbitMQChecker(t *testing.T) {

	t.Run("givenHealthyRabbitMQ_whenCheck_thenIsHealthy", func(t *testing.T) {
		rabbitMQHealthChecker := NewRabbitMQHealthChecker()
		health := rabbitMQHealthChecker.Check(context.TODO())

		assert.Equal(t, Healthy, health.Status)
	})

}
