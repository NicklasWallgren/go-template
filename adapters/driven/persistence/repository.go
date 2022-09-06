package persistence

import (
	"context"
	"fmt"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/drivers"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/transaction"
	sqlTemplate "github.com/NicklasWallgren/sqlTemplate/pkg"

	"github.com/NicklasWallgren/go-template/config"
	"gorm.io/gorm"
)

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	RawSQL(ctx context.Context, sql string, values ...any) error
	Gorm() *gorm.DB
	WrapError(err error) error
}

type repository struct {
	Database
	Logger              logger.Logger
	config              *config.AppConfig
	queryTemplateEngine sqlTemplate.QueryTemplateEngine
}

func NewRepository(
	database Database,
	logger logger.Logger,
	config *config.AppConfig,
	queryTemplateEngine sqlTemplate.QueryTemplateEngine,
) Repository {
	return &repository{
		Database:            database,
		Logger:              logger,
		config:              config,
		queryTemplateEngine: queryTemplateEngine,
	}
}

func (r repository) WithTx(tx *gorm.DB) Repository {
	// The transaction (*gorm.DB) is only available in the returned Repository
	// Otherwise we would pollute the main instance.
	cloned := r
	cloned.DB = tx

	return cloned
}

func (r repository) RawSQL(ctx context.Context, sql string, values ...any) error {
	if err := r.DB.WithContext(ctx).Exec(sql, values...).Error; err != nil {
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

// FindUsingTemplate cannot be part of the repository until https://github.com/golang/go/issues/49085 is fixed.
func FindUsingTemplate[T any, C any](
	ctx context.Context,
	gorm *gorm.DB,
	queryTemplateEngine sqlTemplate.QueryTemplateEngine,
	namespace string,
	name string,
	criteriaAndPagination *models.CriteriaAndPagination[C],
) (*T, error) {
	tmpl, err := queryTemplateEngine.ParseWithValuesFromStruct(namespace, name, criteriaAndPagination)
	if err != nil {
		return nil, err
	}

	var result *T
	if len(tmpl.GetParams()) > 0 {
		err = gorm.WithContext(ctx).Raw(tmpl.GetQuery(), tmpl.GetParams()...).Scan(&result).Error
	} else {
		err = gorm.WithContext(ctx).Raw(tmpl.GetQuery()).Scan(&result).Error
	}

	return result, err
}

// FindUsingTemplatePageable cannot be part of the repository until https://github.com/golang/go/issues/49085 is fixed.
func FindUsingTemplatePageable[T any, C any](
	ctx context.Context,
	gorm *gorm.DB,
	queryTemplateEngine sqlTemplate.QueryTemplateEngine,
	namespace string,
	name string,
	criteriaAndPagination *models.CriteriaAndPagination[C],
) (*models.Page[*T], error) {
	result, err := FindUsingTemplate[[]*T, C](ctx, gorm, queryTemplateEngine, namespace, name, criteriaAndPagination)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return models.NewEmptyPage[*T](), nil
	}

	newPage, err := models.NewPageWith[*T](*result, criteriaAndPagination.Pagination, func() (int, error) {
		result, err := FindUsingTemplate[int, C](
			ctx, gorm, queryTemplateEngine, namespace, fmt.Sprintf("%s_count", name), criteriaAndPagination)
		if err != nil {
			return 0, err
		}

		return *result, nil
	})

	return newPage, err
}
