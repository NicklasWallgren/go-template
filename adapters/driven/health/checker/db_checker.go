package checker

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
)

type DBHealthChecker struct {
	repository persistence.Repository
	logger     logger.Logger
}

func NewDBHealthChecker(repository persistence.Repository, logger logger.Logger) HealthChecker {
	return &DBHealthChecker{repository: repository, logger: logger}
}

func (d DBHealthChecker) Check(ctx context.Context) Health {
	result := NewHealth(Healthy, "db")

	if err := d.repository.RawSql(ctx, "SELECT 1").Error; err != nil {
		result.Status = Unhealthy

		d.logger.Errorf("The database isn't in a healthy state. Cause: %v", err)
	}

	return result
}
