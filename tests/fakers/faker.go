package fakers

import "github.com/NicklasWallgren/go-template/domain/common"

type EntityFaker[T common.EntityConstraint] interface {
	Any() T
	Create() T
}
