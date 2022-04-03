package health

import (
	"github.com/NicklasWallgren/go-template/infrastructure/health"
	"github.com/mariomac/gostream/stream"
)

type HealthApiConverter struct{}

func NewHealthApiConverter() *HealthApiConverter {
	return &HealthApiConverter{}
}

func (h HealthApiConverter) ResponseOf(result health.HealthResult) HealthResponse {
	componentSlice := stream.Map(stream.OfSlice(result.Components), HealthToResponseHealth).ToSlice()

	return HealthResponse{result.Status.String(), componentSlice}
}

func HealthToResponseHealth(health health.Health) Health {
	return Health{Status: health.Status.String(), Name: health.Name, Details: health.Details}
}
