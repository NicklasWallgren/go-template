// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	events "github.com/NicklasWallgren/go-template/domain/events"
	mock "github.com/stretchr/testify/mock"
)

// EventPublisherManager is an autogenerated mock type for the EventPublisherManager type
type EventPublisherManager struct {
	mock.Mock
}

type EventPublisherManager_Expecter struct {
	mock *mock.Mock
}

func (_m *EventPublisherManager) EXPECT() *EventPublisherManager_Expecter {
	return &EventPublisherManager_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: ctx, event
func (_m *EventPublisherManager) Publish(ctx context.Context, event *events.Event) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *events.Event) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventPublisherManager_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type EventPublisherManager_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//  - ctx context.Context
//  - event *events.Event
func (_e *EventPublisherManager_Expecter) Publish(ctx interface{}, event interface{}) *EventPublisherManager_Publish_Call {
	return &EventPublisherManager_Publish_Call{Call: _e.mock.On("Publish", ctx, event)}
}

func (_c *EventPublisherManager_Publish_Call) Run(run func(ctx context.Context, event *events.Event)) *EventPublisherManager_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*events.Event))
	})
	return _c
}

func (_c *EventPublisherManager_Publish_Call) Return(_a0 error) *EventPublisherManager_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEventPublisherManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventPublisherManager creates a new instance of EventPublisherManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventPublisherManager(t mockConstructorTestingTNewEventPublisherManager) *EventPublisherManager {
	mock := &EventPublisherManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
