package checker

type HealthStatus int

func (h HealthStatus) String() string {
	return [...]string{"Unknown", "Healthy", "Unhealthy"}[h]
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
