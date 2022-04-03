package response

type ApiResponseEnvelop interface {
	Status() int
	Payload() ApiResponse
}

type ResponseEnvelop struct {
	status  int
	payload ApiResponse
}

func (r ResponseEnvelop) Status() int {
	return r.status
}

func (r ResponseEnvelop) Payload() ApiResponse {
	return r.payload
}

type ResponseEnvelopOption func(envelop *ResponseEnvelop)

func WithPayload(payload ApiResponse) ResponseEnvelopOption {
	return func(responseEnvelop *ResponseEnvelop) {
		responseEnvelop.payload = payload
	}
}

func NewApiResponseEnvelop(status int, options ...ResponseEnvelopOption) ApiResponseEnvelop {
	responseEnvelop := &ResponseEnvelop{status: status}

	for _, option := range options {
		option(responseEnvelop)
	}

	return responseEnvelop
}

func NewApiResponseWithPayload(status int, payload ApiResponse) ApiResponseEnvelop {
	return NewApiResponseEnvelop(status, WithPayload(payload))
}

type ApiResponse interface{}

type ApiPaginatedResponse interface {
	*ApiResponse
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
