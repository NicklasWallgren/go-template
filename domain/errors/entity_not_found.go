package errors

import (
	"fmt"
)

type EntityNotFoundError struct {
	ID uint
}

func NewEntityNotFoundError(id uint) *EntityNotFoundError {
	return &EntityNotFoundError{ID: id}
}

func (e EntityNotFoundError) Error() string {
	return fmt.Sprintf("entity not found with ID %d", e.ID)
}
