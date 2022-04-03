package events

type EventAction int

const (
	CREATED EventAction = iota
	UPDATED
	DELETED
)

type EventPayloadConverter func(event *Event) (any, error)

func defaultPayloadConverter(event *Event) (any, error) {
	return event.Payload, nil
}

type Event struct {
	Action     EventAction
	Payload    any
	RoutingKey string
	Converter  EventPayloadConverter
}

type EventOptions func(event *Event)

func NewEvent(action EventAction, payload any, opts ...EventOptions) *Event {
	event := &Event{Action: action, Payload: payload, RoutingKey: "routing_key", Converter: defaultPayloadConverter}

	for _, opt := range opts {
		opt(event)
	}

	return event
}

func WithRouting(routingKey string) EventOptions {
	return func(event *Event) {
		event.RoutingKey = routingKey
	}
}

func WithConverter(converter EventPayloadConverter) EventOptions {
	return func(event *Event) {
		event.Converter = converter
	}
}
