package health

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
)

type DBHealthChecker struct {
	database persistence.Database
	logger   logger.Logger
}

// TODO, move to persistence driver package?

func NewDBHealthChecker(database persistence.Database, logger logger.Logger) HealthChecker {
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
