package health

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/health"
	"github.com/mariomac/gostream/stream"
)

type HealthAPIConverter struct{}

func NewHealthAPIConverter() *HealthAPIConverter {
	return &HealthAPIConverter{}
}

func (h HealthAPIConverter) ResponseOf(result health.HealthResult) HealthResponse {
	componentSlice := stream.Map(stream.OfSlice(result.Components), HealthToResponseHealth).ToSlice()

	return HealthResponse{result.Status.String(), componentSlice}
}

func HealthToResponseHealth(health health.Health) Health {
	return Health{Status: health.Status.String(), Name: health.Name, Details: health.Details}
}
