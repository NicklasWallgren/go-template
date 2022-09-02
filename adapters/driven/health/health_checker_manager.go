package health

import "context"

type HealthResult struct {
	Status     HealthStatus
	Components []Health
}

type HealthCheckerManager struct {
	healthCheckers []HealthChecker
}

func NewHealthCheckerManager(healthCheckers []HealthChecker) *HealthCheckerManager {
	return &HealthCheckerManager{healthCheckers: healthCheckers}
}

func (h HealthCheckerManager) Check(ctx context.Context) HealthResult {
	result := HealthResult{Status: Unknown, Components: make([]Health, len(h.healthCheckers))}

	for index, checker := range h.healthCheckers {
		health := checker.Check(ctx)

		if health.Status > result.Status {
			result.Status = health.Status
		}

		result.Components[index] = health
	}

	return result
}
