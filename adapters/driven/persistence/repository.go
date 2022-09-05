package persistence

import (
	"context"
	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/drivers"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/transaction"

	"github.com/NicklasWallgren/go-template/config"
	"gorm.io/gorm"
)

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	RawSql(ctx context.Context, sql string, values ...any) error
	Gorm() *gorm.DB
	WrapError(err error) error
}

type repository struct {
	Database
	Logger logger.Logger
	config *config.AppConfig
}

func NewRepository(
	database Database,
	logger logger.Logger,
	config *config.AppConfig,
) Repository {
	return &repository{Database: database, Logger: logger, config: config}
}

func (r repository) WithTx(tx *gorm.DB) Repository {
	// The transaction (*gorm.DB) is only available in the returned Repository
	// Otherwise we would pollute the main instance.
	cloned := r
	cloned.DB = tx

	return cloned
}

func (r repository) RawSql(ctx context.Context, sql string, values ...any) error {
	if err := r.DB.WithContext(ctx).Exec(sql, values).Error; err != nil {
		return r.WrapError(err)
	}

	return nil
}

func (r repository) TransactWithDefaultRetry(operation func(tx *gorm.DB) error) error { // nolint: wsl
	return transaction.TransactWithDefaultRetry(r.DB, func(tx *gorm.DB) error {
		return operation(tx)
	})
}

func (r repository) Gorm() *gorm.DB {
	return r.DB
}

func (r repository) WrapError(err error) error {
	nillableDriver := drivers.GetDriverOrNil(r.config.Database.Driver)
	if nillableDriver == nil {
		return err
	}

	return nillableDriver.ConvertError(err)
}
