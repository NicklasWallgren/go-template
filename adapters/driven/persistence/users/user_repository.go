package users

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence"
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"

	"github.com/NicklasWallgren/go-template/domain/users/entities"
	userModels "github.com/NicklasWallgren/go-template/domain/users/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	sqlTemplate "github.com/NicklasWallgren/sqlTemplate/pkg"
)

type UserRepository interface {
	WithTx(tx *gorm.DB) UserRepository
	TransactWithDefaultRetry(ctx context.Context, operation func(tx *gorm.DB) error) error
	FindOneByID(ctx context.Context, id uint) (user *entities.User, err error)
	FindOneByIDForUpdate(ctx context.Context, id uint) (*entities.User, error)
	FindOneByEmailWithExclusiveLock(ctx context.Context, email string) (*entities.User, error)
	FindAll(ctx context.Context, pagination *models.Pagination) (page *models.Page[*entities.User], err error)
	FindAllByCriteria(
		ctx context.Context,
		criteriaAndPagination *models.CriteriaAndPagination[FindAllCriteria],
	) (page *models.Page[*entities.User], err error)
	Overview(
		ctx context.Context,
		criteriaAndPagination *models.CriteriaAndPagination[OverviewCriteria],
	) (page *models.Page[*userModels.SenderOverview], err error)
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Save(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteByID(ctx context.Context, id uint) error
	Count(ctx context.Context) (int64, error)
}

// userRepository database structure.
type userRepository struct {
	persistence.EntityRepository[entities.User]
	queryTemplateEngine sqlTemplate.QueryTemplateEngine
}

// NewUserRepository creates a new user repository.
func NewUserRepository(
	userEntityRepository persistence.EntityRepository[entities.User],
	queryTemplateEngine sqlTemplate.QueryTemplateEngine,
) UserRepository {
	return &userRepository{
		userEntityRepository,
		queryTemplateEngine,
	}
}

// WithTx delegates transaction to user repository.
func (r userRepository) WithTx(tx *gorm.DB) UserRepository {
	// Ensures that the transaction (*gorm.DB) is only available in the returned UserRepository
	// Otherwise we would pollute the main instance.
	cloned := r
	// Returns a copy of EntityRepository with the TX applied
	cloned.EntityRepository = r.EntityRepository.WithTx(tx)

	return cloned
}

// nolint:wsl
func (r userRepository) TransactWithDefaultRetry(ctx context.Context, operation func(tx *gorm.DB) error) error {
	// TODO, pass repository instead of gorm.DB?

	return r.EntityRepository.TransactWithDefaultRetry(func(tx *gorm.DB) error {
		return operation(tx.WithContext(ctx))
	})
}

func (r userRepository) FindOneByID(ctx context.Context, id uint) (*entities.User, error) {
	return r.EntityRepository.FindOneByID(ctx, id)
}

func (r userRepository) FindOneByIDForUpdate(ctx context.Context, id uint) (*entities.User, error) {
	return r.EntityRepository.FindOneByIDForUpdate(ctx, id)
}

func (r userRepository) FindOneByEmailWithExclusiveLock(ctx context.Context, email string) (*entities.User, error) {
	var user *entities.User
	if err := r.EntityRepository.Gorm().
		WithContext(ctx).
		Where("email = ?", email).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&user).Error; err != nil {
		return nil, r.EntityRepository.WrapError(err)
	}

	return user, nil
}

func (r userRepository) FindAll(
	ctx context.Context, pagination *models.Pagination,
) (page *models.Page[*entities.User], err error) {
	return r.EntityRepository.FindAll(ctx, pagination)
}

func (r userRepository) FindAllByCriteria(
	ctx context.Context,
	criteriaAndPagination *models.CriteriaAndPagination[FindAllCriteria],
) (page *models.Page[*entities.User], err error) {
	return r.EntityRepository.FindAllByCriteria(ctx, criteriaAndPagination.Criteria, &criteriaAndPagination.Pagination)
}

func (r userRepository) Overview(
	ctx context.Context,
	criteriaAndPagination *models.CriteriaAndPagination[OverviewCriteria],
) (page *models.Page[*userModels.SenderOverview], err error) {
	return persistence.FindUsingTemplatePageable[userModels.SenderOverview, OverviewCriteria](
		ctx, r.EntityRepository.Gorm(), r.queryTemplateEngine, "users", "overview", criteriaAndPagination)
}

func (r userRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	return r.EntityRepository.Create(ctx, user)
}

func (r userRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	return r.EntityRepository.Save(ctx, user)
}

func (r userRepository) DeleteByID(ctx context.Context, id uint) error {
	return r.EntityRepository.DeleteByID(ctx, id)
}

func (r userRepository) Count(ctx context.Context) (int64, error) {
	return r.EntityRepository.Count(ctx)
}
