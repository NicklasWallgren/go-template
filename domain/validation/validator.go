package validation

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/NicklasWallgren/go-template/domain/common"
)

type EntityValidator[T common.EntityConstraint] interface {
	WithTx(tx *gorm.DB) EntityValidator[T]
	EntityCreationValidator[T]
	EntityUpdateValidator[T]
	EntityDeleteValidator[T]
}

type EntityCreationValidator[T common.EntityConstraint] interface {
	ValidateToCreate(entity *T) error
}

type EntityUpdateValidator[T common.EntityConstraint] interface {
	ValidateToUpdate(origin *T, updated *T) error
}

type EntityDeleteValidator[T common.EntityConstraint] interface {
	ValidateToDelete(entity *T) error
}

type ValidationError struct {
	Message string
}

func (v ValidationError) Unwrap() error {
	return nil
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error %s", v.Message)
}

type ValidationFieldError struct {
	Field   string
	Message string
	Value   any
}

func (v ValidationFieldError) Unwrap() error {
	return nil
}

func (v ValidationFieldError) Error() string {
	return fmt.Sprintf("Field validation error %s %s %s", v.Message, v.Field, v.Value)
}

type ValidationFunc[T any] func(subject *T) error

func HasValueChanged[T comparable](value T, value2 T) bool {
	return value != value2
}

func ValidationStep[T any](toBeValidated *T, v ValidationFunc[T]) error {
	return v(toBeValidated)
}

func ValidateChangeStep[T any, V comparable](validationMethod ValidationFunc[T], value1 V, value2 V) func(subject *T) error {
	return func(subject *T) error {
		if !HasValueChanged(value1, value2) {
			return nil
		}

		return validationMethod(subject)
	}
}

func Validate[T any](subject *T, validationMethods []ValidationFunc[T]) error {
	for _, method := range validationMethods {
		if err := ValidationStep[T](subject, method); err != nil {
			return err
		}
	}

	return nil
}
