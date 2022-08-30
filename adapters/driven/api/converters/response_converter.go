package converters

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/models"
	"github.com/NicklasWallgren/go-template/domain/common"
	"github.com/mariomac/gostream/stream"
)

type APIResponseConverter[T common.EntityConstraint, R response.APIResponse] interface {
	ResponseOf(T) R
}

type PageableResponseConverter[T common.EntityConstraint, R response.APIResponse] struct{}

func (p PageableResponseConverter[T, R]) ResponseOf(
	page *models.Page[T],
	converter APIResponseConverter[T, R],
) *response.PageableResponse[R] {
	contentSlice := stream.Map(stream.OfSlice(page.Content), converter.ResponseOf).ToSlice()

	return response.NewPageableResponse[R](
		contentSlice,
		page.IsEmpty(),
		page.TotalPages(),
		page.NumberOfElements(),
		page.TotalNumberOfElements,
		page.TotalPages())
}
