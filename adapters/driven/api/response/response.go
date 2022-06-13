package response

type ApiResponse interface{}

type ApiResponseEnvelop interface {
	Status() int
	Response() ApiResponse
}

type responseEnvelop struct {
	status   int
	response ApiResponse
}

func (r responseEnvelop) Status() int {
	return r.status
}

func (r responseEnvelop) Response() ApiResponse {
	return r.response
}

type ResponseOption func(envelop *responseEnvelop)

func WithResponse(response ApiResponse) ResponseOption {
	return func(responseEnvelop *responseEnvelop) {
		responseEnvelop.response = response
	}
}

func New(status int, options ...ResponseOption) ApiResponseEnvelop {
	responseEnvelop := &responseEnvelop{status: status}

	for _, option := range options {
		option(responseEnvelop)
	}

	return responseEnvelop
}

func NewWithResponse(status int, payload ApiResponse) ApiResponseEnvelop {
	return New(status, WithResponse(payload))
}

type PageableResponse[T any] struct {
	Content          []T
	Empty            bool
	Number           int
	NumberOfElements int
	TotalElements    int
	TotalPages       int
}

func NewPageableResponse[T any](content []T, empty bool, number int, numberOfElements int, totalElements int, totalPages int) *PageableResponse[T] {
	return &PageableResponse[T]{
		Content:          content,
		Empty:            empty,
		Number:           number,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		TotalPages:       totalPages,
	}
}
