package health

import (
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
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

func (h HealthController) Health(ctx *gin.Context) (response.APIResponseEnvelop, error) {
	healthResult := h.healthCheckerManager.Check(ctx.Request.Context())

	return response.New(
		HealthToHttpStatus(healthResult.Status), response.WithResponse(h.apiConverter.ResponseOf(healthResult))), nil
}

func HealthToHttpStatus(status health.HealthStatus) int {
	switch status {
	case health.Unhealthy:
	case health.Unknown:
		return http.StatusInternalServerError
	case health.Healthy:
		return http.StatusOK
	}

	return http.StatusInternalServerError
}
