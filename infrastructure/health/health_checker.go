package health

import "context"

type HealthStatus int

func (h HealthStatus) String() string {
	return [...]string{"Unknown", "Healthy", "Unhealthy"}[h]
}

const (
	Unknown HealthStatus = iota
	Healthy
	Unhealthy
)

type HealthChecker interface {
	Check(ctx context.Context) Health
}

type Health struct {
	Status  HealthStatus
	Name    string
	Details map[string]any
}

func NewHealth(status HealthStatus, name string) Health {
	return Health{Status: status, Name: name, Details: make(map[string]any)}
}

func NewHealthWithDetails(status HealthStatus, name string, details map[string]any) *Health {
	return &Health{Status: status, Name: name, Details: details}
}

func (h *Health) withDetail(key string, value any) { // nolint: unused
	h.Details[key] = value
}

func (h *Health) withDetails(details map[string]any) { // nolint: unused
	h.Details = details
}

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
