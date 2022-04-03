package database

import (
	"fmt"
	"time"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/avast/retry-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(config *config.AppConfig, logger logger.Logger) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	// TODO, lazy load database connection pool, or fail fast?
	db, err := connect(url, logger)
	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	return Database{DB: db}
}

func connect(url string, logger logger.Logger) (db *gorm.DB, err error) {
	// The database might not have been initialized yet, retry
	err = retry.Do(func() (err error) {
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})
		return err
	}, retry.Delay(200*time.Millisecond), retry.Attempts(10))

	return db, err
}
