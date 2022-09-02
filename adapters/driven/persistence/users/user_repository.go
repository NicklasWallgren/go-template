package users

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
	models2 "github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	WithTx(tx *gorm.DB) UserRepository
	TransactWithDefaultRetry(ctx context.Context, operation func(tx *gorm.DB) error) error
	FindOneByID(ctx context.Context, id uint) (user *entities.User, err error)
	FindOneByIDForUpdate(ctx context.Context, id uint) (*entities.User, error)
	FindOneByEmailWithExclusiveLock(ctx context.Context, email string) (*entities.User, error)
	FindAll(ctx context.Context, pagination *models2.Pagination) (page *models2.Page[*entities.User], err error)
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Save(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteByID(ctx context.Context, id uint) error
	Count(ctx context.Context) (int64, error)
}

// userRepository database structure.
type userRepository struct {
	persistence.Repository[entities.User]
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db persistence.Database, logger logger.Logger, config *config.AppConfig) UserRepository {
	return &userRepository{persistence.NewRepository[entities.User](db, entities.User{}, logger, config)}
}

// WithTx delegates transaction to user repository.
func (r userRepository) WithTx(tx *gorm.DB) UserRepository {
	// Ensures that the transaction (*gorm.DB) is only available in the returned UserRepository
	// Otherwise we would pollute the main instance.
	cloned := r
	// Returns a copy of Repository with the TX applied
	cloned.Repository = r.Repository.WithTx(tx)

	return cloned
}

// nolint:wsl
func (r userRepository) TransactWithDefaultRetry(ctx context.Context, operation func(tx *gorm.DB) error) error {
	// TODO, pass repository instead of gorm.DB?

	return r.Repository.TransactWithDefaultRetry(func(tx *gorm.DB) error {
		return operation(tx.WithContext(ctx))
	})
}

func (r userRepository) FindOneByID(ctx context.Context, id uint) (*entities.User, error) {
	return r.Repository.FindOneByID(ctx, id)
}

func (r userRepository) FindOneByIDForUpdate(ctx context.Context, id uint) (*entities.User, error) {
	return r.Repository.FindOneByIDForUpdate(ctx, id)
}

func (r userRepository) FindOneByEmailWithExclusiveLock(ctx context.Context, email string) (*entities.User, error) {
	var user *entities.User
	if err := r.Gorm().
		WithContext(ctx).
		Where("email = ?", email).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&user).Error; err != nil {
		return nil, r.WrapError(err)
	}

	return user, nil
}

func (r userRepository) FindAll(
	ctx context.Context, pagination *models2.Pagination,
) (page *models2.Page[*entities.User], err error) {
	return r.Repository.FindAll(ctx, pagination)
}

func (r userRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	return r.Repository.Create(ctx, user)
}

func (r userRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	return r.Repository.Save(ctx, user)
}

func (r userRepository) DeleteByID(ctx context.Context, id uint) error {
	return r.Repository.DeleteByID(ctx, id)
}

func (r userRepository) Count(ctx context.Context) (int64, error) {
	return r.Repository.Count(ctx)
}
