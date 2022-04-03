package users

import (
	"context"
	"fmt"
	"github.com/NicklasWallgren/go-template/domain/validation"

	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/models"
	repository "github.com/NicklasWallgren/go-template/adapters/driver/persistence/users"
	domainErrors "github.com/NicklasWallgren/go-template/domain/errors"
	"github.com/NicklasWallgren/go-template/domain/events"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"gorm.io/gorm"
)

type UserService interface {
	WithTx(tx *gorm.DB) UserService
	FindOneUserById(ctx context.Context, id uint) (user *entities.User, err error)
	FindOneUserByIdForUpdate(ctx context.Context, id uint) (*entities.User, error)
	FindAllUser(ctx context.Context, pagination *models.Pagination) (users *models.Page[entities.User], err error)
	CreateUser(ctx context.Context, toBeCreated entities.User) (user *entities.User, err error)
	UpdateUser(ctx context.Context, updated *entities.User) (user *entities.User, err error)
	DeleteUserById(ctx context.Context, id uint) error
}

// userService service layer
type userService struct {
	userValidator         validation.EntityValidator[entities.User]
	logger                logger.Logger
	repository            repository.UserRepository
	eventPublisherManager events.EventPublisherManager
	ctx                   context.Context
}

// NewUserService creates a new userService
func NewUserService(userValidator *UserValidator, logger logger.Logger, repository repository.UserRepository, eventPublisherManager events.EventPublisherManager) UserService {
	return userService{
		userValidator:         userValidator,
		logger:                logger,
		repository:            repository,
		eventPublisherManager: eventPublisherManager,
		ctx:                   context.TODO(),
	}
}

// WithTx delegates transaction to user repository
func (s userService) WithTx(tx *gorm.DB) UserService {
	// Ensures that the transaction (*gorm.DB) is only available in the returned UserRepository
	// Otherwise we would pollute the main instance.
	cloned := s
	// Returns a copy of UserRepository with the TX applied
	cloned.repository = s.repository.WithTx(tx)

	return cloned
}

// FindOneUserById gets one user
func (s userService) FindOneUserById(ctx context.Context, id uint) (user *entities.User, err error) {
	if user, err = s.repository.FindOneById(ctx, id); err != nil {
		return nil, domainErrors.NewDomainError(fmt.Sprintf("could not retrieve the user id %d", id), err)
	}

	return user, nil
}

// FindOneUserByIdForUpdate gets one user
func (s userService) FindOneUserByIdForUpdate(ctx context.Context, id uint) (user *entities.User, err error) {
	if user, err = s.repository.FindOneByIdForUpdate(ctx, id); err != nil {
		return nil, domainErrors.NewDomainError(fmt.Sprintf("could not retrieve the user id %d for update", id), err)
	}

	return user, nil
}

// FindAllUser get all the user
func (s userService) FindAllUser(ctx context.Context, pagination *models.Pagination) (users *models.Page[entities.User], err error) {
	// TODO, support filter by predicate/criteria
	if users, err = s.repository.FindAll(ctx, pagination); err != nil {
		return nil, domainErrors.NewDomainError("unable to retrieve the available users", err)
	}

	return users, nil
}

// CreateUser call to create the user
func (s userService) CreateUser(ctx context.Context, toBeCreated entities.User) (user *entities.User, err error) {
	err = s.repository.TransactWithDefaultRetry(ctx, func(tx *gorm.DB) error {
		txService := s.WithTx(tx).(userService)

		if err := txService.userValidator.WithTx(tx).ValidateToCreate(&toBeCreated); err != nil {
			return domainErrors.NewDomainError("could not create user", err)
		}

		if user, err = txService.repository.Create(ctx, &toBeCreated); err != nil {
			return domainErrors.NewDomainError("unable to create user", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	Publish(s.eventPublisherManager, events.CREATED, user)

	return user, nil
}

// UpdateUser updates the user
func (s userService) UpdateUser(ctx context.Context, updated *entities.User) (persistedUser *entities.User, err error) {
	err = s.repository.TransactWithDefaultRetry(ctx, func(tx *gorm.DB) error {
		txService := s.WithTx(tx).(userService)

		// Apply optimistic locking before updating the user entity
		originUser, err := txService.FindOneUserByIdForUpdate(ctx, uint(updated.ID))
		if err != nil {
			return domainErrors.NewDomainError("unable to retrieve the user before updating", err)
		}

		if err = txService.userValidator.WithTx(tx).ValidateToUpdate(originUser, updated); err != nil {
			return domainErrors.NewDomainError("unable to update user", err)
		}

		if persistedUser, err = txService.repository.Save(ctx, updated); err != nil {
			return domainErrors.NewDomainError("unable to update user", err)
		}

		return nil
	})

	Publish(s.eventPublisherManager, events.UPDATED, persistedUser)

	return persistedUser, err
}

// DeleteUserById deletes the user by id
func (s userService) DeleteUserById(ctx context.Context, id uint) error {
	err := s.repository.TransactWithDefaultRetry(ctx, func(tx *gorm.DB) error {
		txService := s.WithTx(tx).(userService)

		// Apply optimistic locking before updating the user entity
		user, err := txService.FindOneUserByIdForUpdate(ctx, id)
		if err != nil {
			return domainErrors.NewDomainError("unable to retrieve the user before deleting", err)
		}

		if err := s.userValidator.ValidateToDelete(user); err != nil {
			return domainErrors.NewDomainError("unable to delete user", err)
		}

		return txService.repository.DeleteById(ctx, id)
	})

	//Publish(s.eventPublisherManager, events.DELETED, persistedUser)

	return err
}

func Publish(eventPublisherManager events.EventPublisherManager, action events.EventAction, user *entities.User) {
	// TODO, handle error, retry, dead-letter queue?
	eventPublisherManager.Publish(events.NewEvent(action, user,
		events.WithRouting("routing_key"),
		events.WithConverter(ResponseOf),
	))
}
