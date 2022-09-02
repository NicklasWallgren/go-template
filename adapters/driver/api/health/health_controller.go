package health

import (
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/health"
	"github.com/NicklasWallgren/go-template/adapters/driven/health/checker"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthCheckerManager *health.HealthCheckerManager
	apiConverter         *HealthAPIConverter
}

func NewHealthController(
	healthCheckerManager *health.HealthCheckerManager, apiConverter *HealthAPIConverter,
) HealthController {
	return HealthController{
		healthCheckerManager: healthCheckerManager, apiConverter: apiConverter,
	}
}

func (h HealthController) Health(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	healthResult := h.healthCheckerManager.Check(ctx.Request.Context())

	return response.NewEnvelope(
		HealthToHTTPStatus(healthResult.Status), response.WithResponse(h.apiConverter.ResponseOf(healthResult))), nil
}

func HealthToHTTPStatus(status checker.HealthStatus) int {
	switch status {
	case checker.Unhealthy:
	case checker.Unknown:
		return http.StatusInternalServerError
	case checker.Healthy:
		return http.StatusOK
	}

	return http.StatusInternalServerError
}
