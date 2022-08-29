package response

type APIResponse interface{}

func WithResponse(response APIResponse) ResponseEnvelopeOption {
	return func(responseEnvelop *APIResponseEnvelope) {
		responseEnvelop.response = response
	}
}

func NewWithResponse(status int, payload APIResponse) *APIResponseEnvelope {
	return NewEnvelope(status, WithResponse(payload))
}
