package health_integration_test

import (
	"context"
	"net/http"
	"testing"

	healthDriven "github.com/NicklasWallgren/go-template/adapters/driven/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/health/checker"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/common"
	healthAPI "github.com/NicklasWallgren/go-template/adapters/driver/api/health"
	"github.com/NicklasWallgren/go-template/tests"
	"github.com/NicklasWallgren/go-template/tests/integration/utils"
	mocks "github.com/NicklasWallgren/go-template/tests/mocks/adapters/driven/health/checker"
	"github.com/gkampitakis/go-snaps/snaps"
	"go.uber.org/fx"

	. "github.com/NicklasWallgren/go-template/tests/integration"
)

// nolint: funlen, paralleltest
func Test(t *testing.T) {
	t.Run("GivenHealthyStatus_WhenHealth_ThenMatchSnapshot", func(t *testing.T) {
		t.Parallel()

		testFunc := func(requestHandler common.RequestHandler) {
			request := utils.NewHTTPRequest(t, "GET", "/api/health/", nil)
			healthResponse := healthAPI.HealthResponse{}
			utils.DoHTTPRequestWithResponse(
				t, requestHandler.Gin, request, &healthResponse, utils.ExpectHTTPStatus(http.StatusOK))

			snaps.MatchSnapshot(t, healthResponse)
		}

		healthChecker := fx.Provide(
			fx.Annotate(
				func() checker.HealthChecker { return healthChecker(t, checker.Healthy) },
				fx.ResultTags(`group:"checkers"`)),
		)

		// nolint:typecheck
		Runner(t, testFunc, fx.Options(dependencyTree(), healthChecker), InitializeMiddlewareAndRoutes)
	})

	t.Run("GivenUnhealthyStatus_WhenHealth_ThenMatchSnapshot", func(t *testing.T) {
		t.Parallel()

		testFunc := func(requestHandler common.RequestHandler) {
			request := utils.NewHTTPRequest(t, "GET", "/api/health/", nil)
			healthResponse := healthAPI.HealthResponse{}
			utils.DoHTTPRequestWithResponse(
				t, requestHandler.Gin, request, &healthResponse, utils.ExpectHTTPStatus(http.StatusInternalServerError))

			snaps.MatchSnapshot(t, healthResponse)
		}

		healthChecker := fx.Provide(
			fx.Annotate(
				func() checker.HealthChecker { return healthChecker(t, checker.Unhealthy) },
				fx.ResultTags(`group:"checkers"`)),
		)

		// nolint:typecheck
		Runner(t, testFunc, fx.Options(dependencyTree(), healthChecker), InitializeMiddlewareAndRoutes)
	})
}

func healthChecker(tb testing.TB, healthStatus checker.HealthStatus) checker.HealthChecker {
	tb.Helper()

	mockedHealthyChecker := mocks.NewHealthChecker(tb)
	mockedHealthyChecker.EXPECT().Check(context.TODO()).Return(checker.NewHealth(healthStatus, "name of checker"))

	return mockedHealthyChecker
}

func dependencyTree() fx.Option {
	return fx.Options(
		tests.DefaultModule,
		tests.DefaultHTTPModule,

		fx.Provide(fx.Annotate(healthAPI.NewHealthRoutes, fx.ResultTags(`group:"routers"`))),
		fx.Provide(healthAPI.NewHealthController),
		fx.Provide(healthAPI.NewHealthAPIConverter),
		fx.Provide(fx.Annotate(healthDriven.NewHealthCheckerManager, fx.ParamTags(`group:"checkers"`))),
	)
}
