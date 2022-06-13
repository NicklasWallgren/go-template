package converters

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/request"
	"github.com/NicklasWallgren/go-template/domain/common"
)

type APIRequestCreateConverter[T request.APIRequest, R common.EntityConstraint] interface {
	EntityOf(*T) R
}

type APIRequestUpdateConverter[T request.APIRequest, R common.EntityConstraint] interface {
	UpdatedEntityOf(ctx context.Context, request *T) (*R, error)
}
