package errors

import (
	"fmt"
)

type EntityNotFoundError struct {
	Id uint
}

func NewEntityNotFoundError(id uint) *EntityNotFoundError {
	return &EntityNotFoundError{Id: id}
}

func (e EntityNotFoundError) Error() string {
	return fmt.Sprintf("entity not found with Id %d", e.Id)
}
