package users

import (
	"context"
	repository "github.com/NicklasWallgren/go-template/adapters/driver/persistence/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/domain/validation"
	"gorm.io/gorm"
)

type UserValidator struct {
	userRepository repository.UserRepository
}

func NewUserValidator(userRepository repository.UserRepository) *UserValidator {
	return &UserValidator{userRepository: userRepository}
}

// To ensure that UserValidator implements the validation.EntityCreationValidator interface
var _ validation.EntityCreationValidator[entities.User] = (*UserValidator)(nil)

// To ensure that UserValidator implements the validation.EntityUpdateValidator interface
var _ validation.EntityUpdateValidator[entities.User] = (*UserValidator)(nil)

// To ensure that UserValidator implements the validation.EntityDeleteValidator interface
var _ validation.EntityDeleteValidator[entities.User] = (*UserValidator)(nil)

func (u UserValidator) WithTx(tx *gorm.DB) validation.EntityValidator[entities.User] {
	// WithTx that the transaction (*gorm.DB) is only available in the returned UserValidator
	// Otherwise we would pollute the main instance.
	u.userRepository = u.userRepository.WithTx(tx)

	return u
}

func (u UserValidator) ValidateToCreate(user *entities.User) error {
	validationMethods := []validation.ValidationFunc[entities.User]{
		u.validateName,
		u.validateAge,
		u.validateUniqueEmail,
	}

	return validation.Validate(user, validationMethods)
}

func (u UserValidator) ValidateToUpdate(origin *entities.User, updated *entities.User) error {
	validationSteps := []validation.ValidationFunc[entities.User]{
		validation.ValidateChangeStep(u.validateName, origin.Name, updated.Name),
		validation.ValidateChangeStep(u.validateAge, origin.Age, updated.Age),
	}

	return validation.Validate(updated, validationSteps)
}

func (u UserValidator) ValidateToDelete(user *entities.User) error {
	validationMethods := []validation.ValidationFunc[entities.User]{}

	return validation.Validate(user, validationMethods)
}

func (u UserValidator) validateName(user *entities.User) error {
	if len(user.Name) <= 0 {
		return &validation.ValidationFieldError{Field: "Name", Message: "Invalid name", Value: user.Name}
	}

	return nil
}

func (u UserValidator) validateAge(user *entities.User) error {
	if user.Age < 18 {
		return &validation.ValidationFieldError{Field: "Age", Message: "Invalid age", Value: user.Age}
	}

	return nil
}

func (u UserValidator) validateUniqueEmail(user *entities.User) error {
	user, err := u.userRepository.FindOneByEmailWithExclusiveLock(context.TODO(), user.Email) // TODO, pass correct context
	if user != nil && user.ID > 0 {
		return &validation.ValidationFieldError{Field: "Email", Message: "The email has already been reserved", Value: user.Email}
	}

	return err
}
