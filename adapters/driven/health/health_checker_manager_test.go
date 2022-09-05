package health

import (
	"context"
	"github.com/NicklasWallgren/go-template/adapters/driven/health/checker"
	mocks "github.com/NicklasWallgren/go-template/tests/mocks/adapters/driven/health/checker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("givenHealthyCheckers_whenCheck_thenIsHealthy", func(t *testing.T) {
		ctx := context.TODO()

		mockedHealthyChecker := mocks.NewHealthChecker(t)
		mockedHealthyChecker.EXPECT().Check(ctx).Return(checker.NewHealth(checker.Healthy, "name of checker"))

		healthCheckerManager := NewHealthCheckerManager([]checker.HealthChecker{mockedHealthyChecker})
		healthResult := healthCheckerManager.Check(ctx)

		assert.Equal(t, checker.Healthy, healthResult.Status)
	})

	t.Run("givenHealthyAndUnhealthyCheckers_whenCheck_thenIsUnhealthy", func(t *testing.T) {
		ctx := context.TODO()

		mockedHealthyChecker := mocks.NewHealthChecker(t)
		mockedHealthyChecker.EXPECT().Check(ctx).Return(checker.NewHealth(checker.Healthy, "name of healthy checker"))

		mockedUnhealthyChecker := mocks.NewHealthChecker(t)
		mockedUnhealthyChecker.EXPECT().Check(ctx).Return(checker.NewHealth(checker.Unhealthy, "name of unhealthy checker"))

		healthCheckerManager := NewHealthCheckerManager([]checker.HealthChecker{mockedHealthyChecker, mockedUnhealthyChecker})
		healthResult := healthCheckerManager.Check(ctx)

		assert.Equal(t, checker.Unhealthy, healthResult.Status)
	})

	t.Run("givenNoCheckers_whenCheck_thenIsUnknown", func(t *testing.T) {
		ctx := context.TODO()

		healthCheckerManager := NewHealthCheckerManager([]checker.HealthChecker{})
		healthResult := healthCheckerManager.Check(ctx)

		assert.Equal(t, checker.Unknown, healthResult.Status)
	})

}
