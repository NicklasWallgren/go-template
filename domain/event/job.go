package event

import "context"

type job struct {
	ctx   context.Context
	event Event
}
