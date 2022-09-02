package checkers

import (
	"context"
	"github.com/NicklasWallgren/go-template/adapters/driven/health"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
)

type DBHealthChecker struct {
	database persistence.Database
	logger   logger.Logger
}

func NewDBHealthChecker(database persistence.Database, logger logger.Logger) health.HealthChecker {
	return &DBHealthChecker{database: database, logger: logger}
}

func (d DBHealthChecker) Check(ctx context.Context) health.Health {
	result := health.NewHealth(health.Healthy, "db")

	if err := d.database.WithContext(ctx).Exec("SELECT 1").Error; err != nil {
		result.Status = health.Unhealthy

		d.logger.Fatalf("The database isn't in a healthy state. Cause: %v", err)
	}

	return result
}
