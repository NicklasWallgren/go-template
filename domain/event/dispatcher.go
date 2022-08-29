package event

import (
	"context"
	"fmt"
)

type Dispatcher struct {
	jobs   chan job
	events map[Name]Listener
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		jobs:   make(chan job),
		events: make(map[Name]Listener),
	}

	go d.consume()

	return d
}

func (d *Dispatcher) Register(listener Listener, names ...Name) error {
	for _, name := range names {
		if _, ok := d.events[name]; ok {
			return fmt.Errorf("the '%s' event is already registered", name) // nolint:goerr113
		}

		d.events[name] = listener
	}

	return nil
}

func (d *Dispatcher) Dispatch(ctx context.Context, event Event) {
	if _, ok := d.events[event.Name()]; !ok {
		panic(fmt.Sprintf("the '%s' event is not registered", event.Name()))
	}

	d.jobs <- job{ctx: ctx, event: event}
}

func (d *Dispatcher) consume() {
	for job := range d.jobs {
		d.events[job.event.Name()].Listen(job.ctx, job.event)
	}
}
