package users

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"
	repository "github.com/NicklasWallgren/go-template/adapters/driven/persistence/users"

	domainErrors "github.com/NicklasWallgren/go-template/domain/errors"
	"github.com/NicklasWallgren/go-template/domain/event"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/domain/validation"
	"gorm.io/gorm"
)

type UserService interface {
	WithTx(tx *gorm.DB) UserService
	FindOneUserByID(ctx context.Context, id uint) (user *entities.User, err error)
	FindOneUserByIDForUpdate(ctx context.Context, id uint) (*entities.User, error)
	FindAllUser(ctx context.Context, pagination *models.Pagination) (users *models.Page[*entities.User], err error)
	FindAllUserByCriteria(ctx context.Context, criteria *repository.FindAllCriteria, pagination *models.Pagination) (users *models.Page[*entities.User], err error)
	CreateUser(ctx context.Context, toBeCreated entities.User) (user *entities.User, err error)
	UpdateUser(ctx context.Context, updated *entities.User) (user *entities.User, err error)
	DeleteUserByID(ctx context.Context, id uint) error
}

// userService service layer.
type userService struct {
	userValidator   validation.EntityValidator[entities.User]
	logger          logger.Logger
	repository      repository.UserRepository
	eventDispatcher *event.Dispatcher
}

// NewUserService creates a new [UserService].
func NewUserService(
	userValidator *UserValidator,
	logger logger.Logger,
	repository repository.UserRepository,
	eventDispatcher *event.Dispatcher,
) UserService {
	return userService{
		userValidator:   userValidator,
		logger:          logger,
		repository:      repository,
		eventDispatcher: eventDispatcher,
	}
}

// WithTx delegates transaction to user repository.
func (s userService) WithTx(tx *gorm.DB) UserService {
	// Ensures that the transaction (*gorm.DB) is only available in the returned UserRepository
	// Otherwise we would pollute the main instance.
	cloned := s
	// Returns a copy of UserRepository with the TX applied
	cloned.repository = s.repository.WithTx(tx)

	return cloned
}

// FindOneUserByID retrieves a user by the provided ID.
func (s userService) FindOneUserByID(ctx context.Context, id uint) (user *entities.User, err error) {
	if user, err = s.repository.FindOneByID(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.NewEntityNotFoundError(id)
		}

		return nil, domainErrors.NewDomainError(fmt.Sprintf("could not retrieve the user with id %d", id), err)
	}

	return user, nil
}

// FindOneUserByIDForUpdate retrieves a user by the provided ID.
// Applies a database row lock if executed in a transaction.
func (s userService) FindOneUserByIDForUpdate(ctx context.Context, id uint) (user *entities.User, err error) {
	if user, err = s.repository.FindOneByIDForUpdate(ctx, id); err != nil {
		return nil, domainErrors.NewDomainError(fmt.Sprintf("could not retrieve the user id %d for update", id), err)
	}

	return user, nil
}

// FindAllUser retrieves a paginated list of users.
func (s userService) FindAllUser(
	ctx context.Context,
	pagination *models.Pagination,
) (users *models.Page[*entities.User], err error) {
	if users, err = s.repository.FindAll(ctx, pagination); err != nil {
		return nil, domainErrors.NewDomainError("unable to retrieve the available users", err)
	}

	return users, nil
}

// FindAllUserByCriteria retrieves a paginated list of users.
func (s userService) FindAllUserByCriteria(
	ctx context.Context,
	criteria *repository.FindAllCriteria,
	pagination *models.Pagination,
) (users *models.Page[*entities.User], err error) {
	if users, err = s.repository.FindAllByCriteria(ctx, criteria, pagination); err != nil {
		return nil, domainErrors.NewDomainError("unable to retrieve the available users", err)
	}

	return users, nil
}

// CreateUser creates a new user.
func (s userService) CreateUser(ctx context.Context, toBeCreated entities.User) (user *entities.User, err error) {
	err = s.repository.TransactWithDefaultRetry(ctx, func(tx *gorm.DB) error {
		txService := s.WithTx(tx).(userService) // nolint: forcetypeassert

		if err := txService.userValidator.WithTx(tx).ValidateToCreate(ctx, &toBeCreated); err != nil {
			return domainErrors.NewDomainError("unable to create user", err)
		}

		if user, err = txService.repository.Create(ctx, &toBeCreated); err != nil {
			return domainErrors.NewDomainError("unable to create user", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	s.eventDispatcher.Dispatch(
		ctx, event.EntityEvent{Time: time.Now().UTC(), Entity: toBeCreated, Action: event.Created})

	return user, nil
}

// UpdateUser updates the provided user.
func (s userService) UpdateUser(ctx context.Context, updated *entities.User) (persistedUser *entities.User, err error) {
	err = s.repository.TransactWithDefaultRetry(ctx, func(tx *gorm.DB) error {
		txService := s.WithTx(tx).(userService) // nolint: forcetypeassert

		// Apply optimistic locking before updating the user entity
		originUser, err := txService.FindOneUserByIDForUpdate(ctx, uint(updated.ID))
		if err != nil {
			return domainErrors.NewDomainError("unable to retrieve the user before updating", err)
		}

		if err = txService.userValidator.WithTx(tx).ValidateToUpdate(ctx, originUser, updated); err != nil {
			return domainErrors.NewDomainError("unable to update user", err)
		}

		if persistedUser, err = txService.repository.Save(ctx, updated); err != nil {
			return domainErrors.NewDomainError("unable to update user", err)
		}

		return nil
	})

	return persistedUser, err
}

// DeleteUserByID deletes a user by ID.
func (s userService) DeleteUserByID(ctx context.Context, id uint) error {
	err := s.repository.TransactWithDefaultRetry(ctx, func(tx *gorm.DB) error {
		txService := s.WithTx(tx).(userService) // nolint: forcetypeassert

		// Apply optimistic locking before updating the user entity
		user, err := txService.FindOneUserByIDForUpdate(ctx, id)
		if err != nil {
			return domainErrors.NewDomainError("unable to retrieve the user before deleting", err)
		}

		if err := s.userValidator.ValidateToDelete(ctx, user); err != nil {
			return domainErrors.NewDomainError("unable to delete user", err)
		}

		return txService.repository.DeleteByID(ctx, id)
	})

	return err
}
