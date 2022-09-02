package response

type APIResponseEnvelope struct {
	status   int
	response APIResponse
}

func (r APIResponseEnvelope) Status() int {
	return r.status
}

func (r APIResponseEnvelope) Response() APIResponse {
	return r.response
}

type ResponseEnvelopeOption func(envelop *APIResponseEnvelope)

func NewEnvelope(status int, options ...ResponseEnvelopeOption) *APIResponseEnvelope {
	responseEnvelop := &APIResponseEnvelope{status: status}

	for _, option := range options {
		option(responseEnvelop)
	}

	return responseEnvelop
}
