package persistence

import (
	"errors"
	"time"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	gormTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	mysqlGorm "gorm.io/driver/mysql"
	postgresGorm "gorm.io/driver/postgres"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/avast/retry-go"
	"gorm.io/gorm"
)

// Database modal.
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance.
func NewDatabase(config *config.AppConfig, logger logger.Logger) (Database, error) {
	db, err := connect(config.Database, logger)
	if err != nil {
		return Database{}, err
	}

	return Database{DB: db}, nil
}

func connect(config *config.DatabaseConfig, logger logger.Logger) (db *gorm.DB, err error) {
	// The database might not have been initialized yet, retry
	err = retry.Do(func() (err error) {
		dialector, err := gormDialector(config)
		if err != nil {
			return err
		}

		db, err = gormTrace.Open(dialector, &gorm.Config{Logger: logger.GetGormLogger()}, gormTrace.WithServiceName("mysql"))

		return err
	}, retry.Delay(200*time.Millisecond), retry.Attempts(10)) // nolint:gomnd

	return db, err
}

func gormDialector(config *config.DatabaseConfig) (gorm.Dialector, error) {
	url := DSN(config)

	switch config.Driver {
	case "mysql":
		return mysqlGorm.New(mysqlGorm.Config{DSN: url}), nil
	case "postgres":
		return postgresGorm.New(postgresGorm.Config{DSN: url}), nil
	}

	return nil, errors.New("unsupported db driver " + config.Driver) // nolint:goerr113
}
