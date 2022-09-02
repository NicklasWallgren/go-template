package response

type PageableResponse[T any] struct {
	Content          []T
	Empty            bool
	Number           int
	NumberOfElements int
	TotalElements    int
	TotalPages       int
}

func NewPageableResponse[T any](
	content []T,
	empty bool,
	number int,
	numberOfElements int,
	totalElements int,
	totalPages int,
) *PageableResponse[T] {
	return &PageableResponse[T]{
		Content:          content,
		Empty:            empty,
		Number:           number,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		TotalPages:       totalPages,
	}
}
