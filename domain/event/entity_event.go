package event

import (
	"time"
)

type EventAction int

const (
	Created = iota
	Updated
	Deleted
)

func (e EventAction) String() string {
	return [...]string{"Created", "Updated", "Deleted"}[e]
}

const EntityEventName Name = "entity"

type EntityEvent struct {
	Time   time.Time
	Entity any
	Action EventAction
}

func (c EntityEvent) Name() Name {
	return EntityEventName
}
