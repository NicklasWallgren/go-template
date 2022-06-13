package users_test

import (
	"context"
	"testing"

	"github.com/NicklasWallgren/go-template/tests/mocks"

	"github.com/NicklasWallgren/go-template/domain/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/domain/validation"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

// nolint: funlen
func Test(t *testing.T) {
	t.Run("GivenInvalidName_whenValidateToCreate_thenExpectedValidationFieldError", func(t *testing.T) {
		t.Parallel()

		user := entities.NewUser("", gofakeit.Email(), uint8(gofakeit.Number(18, 150)), gofakeit.Date())

		validator := users.NewUserValidator(mocks.NewUserRepository(t))
		err := validator.ValidateToCreate(context.TODO(), &user)

		assertValidationFieldError(t, err, "Name", "", "Invalid name")
	})

	t.Run("GivenInvalidAge_whenValidateToCreate_thenExpectedValidationFieldError", func(t *testing.T) {
		t.Parallel()

		user := entities.NewUser(gofakeit.Name(), gofakeit.Email(), uint8(gofakeit.Number(0, 17)), gofakeit.Date())

		validator := users.NewUserValidator(mocks.NewUserRepository(t))
		err := validator.ValidateToCreate(context.TODO(), &user)

		assertValidationFieldError(t, err, "Age", user.Age, "Invalid age")
	})

	t.Run("GivenInvalidName_whenValidateToUpdate_thenExpectedValidationFieldError", func(t *testing.T) {
		t.Parallel()

		user := entities.NewUser(gofakeit.Name(), gofakeit.Email(), uint8(gofakeit.Number(18, 150)), gofakeit.Date())
		updatedUser := user
		updatedUser.Name = ""

		validator := users.NewUserValidator(mocks.NewUserRepository(t))
		err := validator.ValidateToUpdate(context.TODO(), &user, &updatedUser)

		assertValidationFieldError(t, err, "Name", "", "Invalid name")
	})

	t.Run("GivenInvalidAge_whenValidateToUpdate_thenExpectedValidationFieldError", func(t *testing.T) {
		t.Parallel()

		user := entities.NewUser(gofakeit.Name(), gofakeit.Email(), uint8(gofakeit.Number(18, 150)), gofakeit.Date())
		updatedUser := user
		updatedUser.Age = 17

		validator := users.NewUserValidator(mocks.NewUserRepository(t))
		err := validator.ValidateToUpdate(context.TODO(), &user, &updatedUser)

		assertValidationFieldError(t, err, "Age", updatedUser.Age, "Invalid age")
	})

	t.Run("GivenNotUniqueEmail_whenValidateToCreate_thenExpectedValidationFieldError", func(t *testing.T) {
		t.Parallel()

		user := entities.NewUserWithId(1, gofakeit.Name(), gofakeit.Email(), uint8(gofakeit.Number(18, 150)), gofakeit.Date())

		userRepository := mocks.NewUserRepository(t)
		userRepository.On("FindOneByEmailWithExclusiveLock", user.Email).Return(&user, nil)

		validator := users.NewUserValidator(userRepository)
		err := validator.ValidateToCreate(context.TODO(), &user)

		assertValidationFieldError(t, err, "Email", user.Email, "The email has already been reserved")
	})
}

func assertValidationFieldError(t *testing.T, err error, field string, value any, message string) {
	t.Helper()

	assert.NotNil(t, err)

	validationFieldError := &validation.ValidationFieldError{}
	assert.ErrorAsf(t, err, &validationFieldError, "")

	assert.Equal(t, field, validationFieldError.Field)
	assert.Equal(t, value, validationFieldError.Value)
	assert.Equal(t, message, validationFieldError.Message)
}
