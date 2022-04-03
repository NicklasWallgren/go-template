package factories

import (
	"github.com/NicklasWallgren/go-template/domain/common"
	"github.com/NicklasWallgren/go-template/tests/fakers"
)

type EntityFactory[T common.EntityConstraint, U fakers.EntityFaker[T]] interface {
	Any() (*T, error)
	Many(numberOfEntities int) ([]T, error)
	Given(entity T) (*T, error) // TODO, pass faker?
}
