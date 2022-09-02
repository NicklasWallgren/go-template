package checker

import (
	"context"
)

const (
	Unknown HealthStatus = iota
	Healthy
	Unhealthy
)

type HealthChecker interface {
	Check(ctx context.Context) Health
}
