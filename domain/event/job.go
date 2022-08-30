package event

import "context"

type job struct {
	ctx   context.Context // nolint: containedctx
	event Event
}
