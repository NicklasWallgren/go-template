package checker

import (
	"context"
	"fmt"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	persistenceMock "github.com/NicklasWallgren/go-template/tests/mocks/adapters/driven/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test(t *testing.T) {

	t.Run("givenHealthyDatabase_whenCheck_thenIsHealthy", func(t *testing.T) {
		ctx := context.TODO()

		repository := persistenceMock.NewRepository(t)
		repository.EXPECT().RawSql(ctx, mock.AnythingOfType("string")).Return(nil)

		dbHealthChecker := NewDBHealthChecker(repository, logger.NullLogger{})
		health := dbHealthChecker.Check(ctx)

		assert.Equal(t, Healthy, health.Status)
	})

	t.Run("givenUnhealthyDatabase_whenCheck_thenIsUnhealthy", func(t *testing.T) {
		err := fmt.Errorf("database error")
		ctx := context.TODO()

		repository := persistenceMock.NewRepository(t)
		repository.EXPECT().RawSql(ctx, mock.AnythingOfType("string")).Return(err)

		dbHealthChecker := NewDBHealthChecker(repository, logger.NullLogger{})
		health := dbHealthChecker.Check(ctx)

		assert.Equal(t, Unhealthy, health.Status)
	})
}
