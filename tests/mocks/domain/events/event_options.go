// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	events "github.com/NicklasWallgren/go-template/domain/event"
	mock "github.com/stretchr/testify/mock"
)

// EventOptions is an autogenerated mock type for the EventOptions type
type EventOptions struct {
	mock.Mock
}

type EventOptions_Expecter struct {
	mock *mock.Mock
}

func (_m *EventOptions) EXPECT() *EventOptions_Expecter {
	return &EventOptions_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: event
func (_m *EventOptions) Execute(event *events.Event) {
	_m.Called(event)
}

// EventOptions_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type EventOptions_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//  - event *event.Event
func (_e *EventOptions_Expecter) Execute(event interface{}) *EventOptions_Execute_Call {
	return &EventOptions_Execute_Call{Call: _e.mock.On("Execute", event)}
}

func (_c *EventOptions_Execute_Call) Run(run func(event *events.Event)) *EventOptions_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*events.Event))
	})
	return _c
}

func (_c *EventOptions_Execute_Call) Return() *EventOptions_Execute_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewEventOptions interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventOptions creates a new instance of EventOptions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventOptions(t mockConstructorTestingTNewEventOptions) *EventOptions {
	mock := &EventOptions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
