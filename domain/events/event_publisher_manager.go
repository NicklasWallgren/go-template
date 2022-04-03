package events

type EventPublisherManager interface {
	Publish(event *Event) error
}

type eventPublisherManager struct {
	publisher EventPublisher // TODO, support multiple?
}

func NewEventPublisherManager(publisher EventPublisher) EventPublisherManager {
	return &eventPublisherManager{publisher: publisher}
}

func (e eventPublisherManager) Publish(event *Event) error {
	// TODO, publish using go routine to not block current routine?
	return e.publisher.Publish(event)
}
