package converters

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	"github.com/mariomac/gostream/stream"
)

type APIResponseConverter[T any, R response.APIResponse] interface {
	ResponseOf(T) R
}

type PageableResponseConverter[T any, R response.APIResponse] struct{}

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
