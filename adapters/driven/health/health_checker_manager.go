package health

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/health/checker"
)

type HealthResult struct {
	Status     checker.HealthStatus
	Components []checker.Health
}

type HealthCheckerManager struct {
	healthCheckers []checker.HealthChecker
}

func NewHealthCheckerManager(healthCheckers []checker.HealthChecker) *HealthCheckerManager {
	return &HealthCheckerManager{healthCheckers: healthCheckers}
}

func (h HealthCheckerManager) Check(ctx context.Context) HealthResult {
	result := HealthResult{Status: checker.Unknown, Components: make([]checker.Health, len(h.healthCheckers))}

	for index, checker := range h.healthCheckers {
		health := checker.Check(ctx)

		if health.Status > result.Status {
			result.Status = health.Status
		}

		result.Components[index] = health
	}

	return result
}
