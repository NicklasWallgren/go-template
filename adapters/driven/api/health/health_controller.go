package health

import (
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/NicklasWallgren/go-template/infrastructure/health"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthCheckerManager *health.HealthCheckerManager
	apiConverter         *HealthApiConverter
}

func NewHealthController(healthCheckerManager *health.HealthCheckerManager, apiConverter *HealthApiConverter) HealthController {
	return HealthController{healthCheckerManager: healthCheckerManager, apiConverter: apiConverter}
}

func (h HealthController) Health(ctx *gin.Context) (response.ApiResponseEnvelop, error) {
	healthResult := h.healthCheckerManager.Check(ctx.Request.Context())

	return response.NewApiResponseEnvelop(HealthToHttpStatus(healthResult.Status), response.WithPayload(h.apiConverter.ResponseOf(healthResult))), nil
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
