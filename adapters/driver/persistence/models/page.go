package models

import "math"

// TODO next Pageable, map.
type Page[T any] struct {
	Content               []T
	Pageable              Pageable
	TotalNumberOfElements int
}

func NewPage[T any](content []T, pageable Pageable, total int) *Page[T] {
	return &Page[T]{
		Content:               content,
		Pageable:              pageable,
		TotalNumberOfElements: total,
	}
}

func (p Page[T]) NumberOfElements() int {
	return len(p.Content)
}

func (p Page[T]) TotalPages() int {
	if p.Pageable.PageSize() == 0 {
		return 1
	}

	return int(math.Ceil(float64(p.TotalNumberOfElements / p.Pageable.PageSize())))
}

func (p Page[T]) IsEmpty() bool {
	return len(p.Content) == 0
}

type TotalSupplier func() (int, error)

func NewPageWith[T any](content []T, pageable Pageable, totalSupplier TotalSupplier) (*Page[T], error) {
	if pageable.Offset() == 0 {
		// No more Content to be retrieved
		if pageable.PageSize() > len(content) {
			return NewPage[T](content, pageable, len(content)), nil
		}

		// More Content is available
		return newPageWithAdditionalContent(content, pageable, totalSupplier)
	}

	if pageable.PageSize() > len(content) {
		// No more Content to be retrieved
		return NewPage[T](content, pageable, pageable.Offset()+len(content)), nil
	}

	// More Content is available
	return newPageWithAdditionalContent(content, pageable, totalSupplier)
}

func newPageWithAdditionalContent[T any](
	content []T, pageable Pageable, totalSupplier TotalSupplier,
) (*Page[T], error) {
	total, err := totalSupplier()
	if err != nil {
		return nil, err
	}

	return NewPage[T](content, pageable, total), nil
}
