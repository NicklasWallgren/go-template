package checker

import (
	"context"
)

type HealthChecker interface {
	Check(ctx context.Context) Health
}
