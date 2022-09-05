package checker

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRabbitMQChecker(t *testing.T) {
	t.Parallel()

	t.Run("givenHealthyRabbitMQ_whenCheck_thenIsHealthy", func(t *testing.T) {
		t.Parallel()

		rabbitMQHealthChecker := NewRabbitMQHealthChecker()
		health := rabbitMQHealthChecker.Check(context.TODO())

		assert.Equal(t, Healthy, health.Status)
	})
}
