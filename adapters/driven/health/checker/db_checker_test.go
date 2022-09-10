package checker

import (
	"context"
	"fmt"
	"testing"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	persistenceMock "github.com/NicklasWallgren/go-template/tests/mocks/adapters/driven/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDBChecker(t *testing.T) {
	t.Parallel()

	t.Run("givenHealthyDatabase_whenCheck_thenIsHealthy", func(t *testing.T) {
		t.Parallel()

		ctx := context.TODO()

		repository := persistenceMock.NewRepository(t)
		repository.EXPECT().RawSQL(ctx, mock.AnythingOfType("string")).Return(nil)

		dbHealthChecker := NewDBHealthChecker(repository, logger.NopLogger{})
		health := dbHealthChecker.Check(ctx)

		assert.Equal(t, Healthy, health.Status)
	})

	t.Run("givenUnhealthyDatabase_whenCheck_thenIsUnhealthy", func(t *testing.T) {
		t.Parallel()

		err := fmt.Errorf("database error") // nolint:goerr113
		ctx := context.TODO()

		repository := persistenceMock.NewRepository(t)
		repository.EXPECT().RawSQL(ctx, mock.AnythingOfType("string")).Return(err)

		dbHealthChecker := NewDBHealthChecker(repository, logger.NopLogger{})
		health := dbHealthChecker.Check(ctx)

		assert.Equal(t, Unhealthy, health.Status)
	})
}
