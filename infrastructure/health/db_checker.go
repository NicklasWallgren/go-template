package health

import (
	"context"

	"github.com/NicklasWallgren/go-template/infrastructure/database"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
)

type DBHealthChecker struct {
	database database.Database
	logger   logger.Logger
}

func NewDBHealthChecker(database database.Database, logger logger.Logger) HealthChecker {
	return &DBHealthChecker{database: database, logger: logger}
}

func (d DBHealthChecker) Check(ctx context.Context) Health {
	health := NewHealth(Healthy, "db")

	if err := d.database.WithContext(ctx).Exec("SELECT 1").Error; err != nil {
		health.Status = Unhealthy

		d.logger.Fatalf("The database isn't in a healthy state. Cause: %v", err)
	}

	return health
}
