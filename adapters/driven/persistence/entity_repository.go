package persistence

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/drivers"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/transaction"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EntityRepository[T common.EntityConstraint] interface {
	WithTx(tx *gorm.DB) EntityRepository[T]
	TransactWithDefaultRetry(operation func(tx *gorm.DB) error) error
	FindOneByID(ctx context.Context, id uint) (entity *T, err error)
	FindOneByIDForUpdate(ctx context.Context, id uint) (entity *T, err error)
	FindAll(ctx context.Context, pagination *models.Pagination) (page *models.Page[*T], err error)
	FindAllByCriteria(
		ctx context.Context, criteriaAndPagination any, pagination *models.Pagination) (page *models.Page[*T], err error)
	Create(ctx context.Context, entity *T) (*T, error)
	Save(ctx context.Context, entity *T) (*T, error)
	DeleteByID(ctx context.Context, id uint) error
	Count(ctx context.Context) (int64, error)
	Gorm() *gorm.DB
	WrapError(err error) error
}

type entityRepository[T common.EntityConstraint] struct {
	Database
	Logger logger.Logger
	config *config.AppConfig
	entity T
}

func NewEntityRepository[T common.EntityConstraint](
	database Database,
	entity T,
	logger logger.Logger,
	config *config.AppConfig,
) EntityRepository[T] {
	return &entityRepository[T]{Database: database, entity: entity, Logger: logger, config: config}
}

func (r entityRepository[T]) WithTx(tx *gorm.DB) EntityRepository[T] {
	// The transaction (*gorm.DB) is only available in the returned EntityRepository
	// Otherwise we would pollute the main instance.
	cloned := r
	cloned.DB = tx

	return cloned
}

func (r entityRepository[T]) TransactWithDefaultRetry(operation func(tx *gorm.DB) error) error { // nolint: wsl
	return transaction.TransactWithDefaultRetry(r.DB, func(tx *gorm.DB) error {
		return operation(tx)
	})
}

func (r entityRepository[T]) FindOneByID(ctx context.Context, id uint) (entity *T, err error) {
	dbSession := r.DB.WithContext(ctx).Limit(1).Find(&entity, id)

	if err := dbSession.Error; err != nil {
		return nil, r.WrapError(err)
	}

	if dbSession.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return entity, nil
}

func (r entityRepository[T]) FindOneByIDForUpdate(ctx context.Context, id uint) (entity *T, err error) {
	dbSession := r.DB.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&entity, id)

	if err := dbSession.Error; err != nil {
		return nil, r.WrapError(err)
	}

	if dbSession.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return entity, nil
}

func (r entityRepository[T]) FindAll(
	ctx context.Context,
	pagination *models.Pagination,
) (page *models.Page[*T], err error) {
	tx := r.DB.WithContext(ctx).Offset(pagination.Offset()).Limit(pagination.Limit).Order(pagination.Order())

	content := &[]*T{}
	if tx.Find(content).Error != nil {
		return page, r.WrapError(err)
	}

	newPage, err := models.NewPageWith[*T](*content, pagination, func() (int, error) { return r.totalCountSupplier(ctx) })
	if err != nil {
		return page, r.WrapError(err)
	}

	return newPage, nil
}

func (r entityRepository[T]) FindAllByCriteria(
	ctx context.Context,
	criteria any,
	pagination *models.Pagination,
) (page *models.Page[*T], err error) {
	tx := r.DB.WithContext(ctx).Where(criteria).Offset(pagination.Offset()).Limit(pagination.Limit).
		Order(pagination.Order())

	content := &[]*T{}
	if tx.Find(content).Error != nil {
		return page, r.WrapError(err)
	}

	newPage, err := models.NewPageWith[*T](*content, pagination, func() (int, error) { return r.totalCountSupplier(ctx) })
	if err != nil {
		return page, r.WrapError(err)
	}

	return newPage, nil
}

func (r entityRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	if err := r.DB.WithContext(ctx).Create(entity).Error; err != nil {
		return nil, r.WrapError(err)
	}

	return entity, nil
}

func (r entityRepository[T]) Save(ctx context.Context, entity *T) (*T, error) {
	if err := r.DB.WithContext(ctx).Save(entity).Error; err != nil {
		return nil, r.WrapError(err)
	}

	return entity, nil
}

func (r entityRepository[T]) DeleteByID(ctx context.Context, id uint) error {
	var entity *T
	if err := r.DB.WithContext(ctx).Delete(&entity, id).Error; err != nil {
		return err
	}

	return nil
}

func (r entityRepository[T]) Count(ctx context.Context) (int64, error) {
	var totalCount int64
	if err := r.DB.WithContext(ctx).Model(r.entity).Count(&totalCount).Error; err != nil {
		return 0, r.WrapError(err)
	}

	return totalCount, nil
}

func (r entityRepository[T]) Gorm() *gorm.DB {
	return r.DB
}

func (r entityRepository[T]) WrapError(err error) error {
	nillableDriver := drivers.GetDriverOrNil(r.config.Database.Driver)
	if nillableDriver == nil {
		return err
	}

	return nillableDriver.ConvertError(err)
}

func (r *entityRepository[T]) totalCountSupplier(ctx context.Context) (int, error) {
	result, err := r.Count(ctx)

	return int(result), err
}
