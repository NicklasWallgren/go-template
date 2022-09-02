package health

import (
	"net/http"
	response2 "github.com/NicklasWallgren/go-template/adapters/driver/api/response"

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

func (h HealthController) Health(ctx *gin.Context) (*response2.APIResponseEnvelope, error) {
	healthResult := h.healthCheckerManager.Check(ctx.Request.Context())

	return response2.NewEnvelope(
		HealthToHTTPStatus(healthResult.Status), response2.WithResponse(h.apiConverter.ResponseOf(healthResult))), nil
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
