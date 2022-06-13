package response

type APIResponse interface{}

type APIResponseEnvelop interface {
	Status() int
	Response() APIResponse
}

type responseEnvelop struct {
	status   int
	response APIResponse
}

func (r responseEnvelop) Status() int {
	return r.status
}

func (r responseEnvelop) Response() APIResponse {
	return r.response
}

type ResponseOption func(envelop *responseEnvelop)

func WithResponse(response APIResponse) ResponseOption {
	return func(responseEnvelop *responseEnvelop) {
		responseEnvelop.response = response
	}
}

func New(status int, options ...ResponseOption) APIResponseEnvelop {
	responseEnvelop := &responseEnvelop{status: status}

	for _, option := range options {
		option(responseEnvelop)
	}

	return responseEnvelop
}

func NewWithResponse(status int, payload APIResponse) APIResponseEnvelop {
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

func NewPageableResponse[T any](
	content []T, empty bool, number int, numberOfElements int, totalElements int, totalPages int,
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
