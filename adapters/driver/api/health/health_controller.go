package health

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"net/http"

	"github.com/NicklasWallgren/go-template/infrastructure/health"
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

func HealthToHTTPStatus(status health.HealthStatus) int {
	switch status {
	case health.Unhealthy:
	case health.Unknown:
		return http.StatusInternalServerError
	case health.Healthy:
		return http.StatusOK
	}

	return http.StatusInternalServerError
}
