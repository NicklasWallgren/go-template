package event

import "context"

type Listener interface {
	Listen(ctx context.Context, event Event)
}
