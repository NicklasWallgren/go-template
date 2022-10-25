package models

type CriteriaAndPagination[T any] struct {
	Criteria   T
	Pagination Pagination
}

func NewCriteriaAndPagination[T any](criteria T) CriteriaAndPagination[T] {
	return CriteriaAndPagination[T]{Criteria: criteria, Pagination: NewPaginationWithDefaults()}
}
