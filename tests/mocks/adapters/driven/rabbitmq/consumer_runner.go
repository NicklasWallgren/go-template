// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	rabbitmq "github.com/NicklasWallgren/go-template/adapters/driven/rabbitmq"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerRunner is an autogenerated mock type for the ConsumerRunner type
type ConsumerRunner struct {
	mock.Mock
}

type ConsumerRunner_Expecter struct {
	mock *mock.Mock
}

func (_m *ConsumerRunner) EXPECT() *ConsumerRunner_Expecter {
	return &ConsumerRunner_Expecter{mock: &_m.Mock}
}

// Do provides a mock function with given fields: ctx
func (_m *ConsumerRunner) Do(ctx context.Context) {
	_m.Called(ctx)
}

// ConsumerRunner_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type ConsumerRunner_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//  - ctx context.Context
func (_e *ConsumerRunner_Expecter) Do(ctx interface{}) *ConsumerRunner_Do_Call {
	return &ConsumerRunner_Do_Call{Call: _e.mock.On("Do", ctx)}
}

func (_c *ConsumerRunner_Do_Call) Run(run func(ctx context.Context)) *ConsumerRunner_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ConsumerRunner_Do_Call) Return() *ConsumerRunner_Do_Call {
	_c.Call.Return()
	return _c
}

// Options provides a mock function with given fields:
func (_m *ConsumerRunner) Options() rabbitmq.ConsumerRunnerOptions {
	ret := _m.Called()

	var r0 rabbitmq.ConsumerRunnerOptions
	if rf, ok := ret.Get(0).(func() rabbitmq.ConsumerRunnerOptions); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rabbitmq.ConsumerRunnerOptions)
	}

	return r0
}

// ConsumerRunner_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type ConsumerRunner_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
func (_e *ConsumerRunner_Expecter) Options() *ConsumerRunner_Options_Call {
	return &ConsumerRunner_Options_Call{Call: _e.mock.On("Options")}
}

func (_c *ConsumerRunner_Options_Call) Run(run func()) *ConsumerRunner_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConsumerRunner_Options_Call) Return(_a0 rabbitmq.ConsumerRunnerOptions) *ConsumerRunner_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *ConsumerRunner) Run(ctx context.Context) {
	_m.Called(ctx)
}

// ConsumerRunner_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type ConsumerRunner_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//  - ctx context.Context
func (_e *ConsumerRunner_Expecter) Run(ctx interface{}) *ConsumerRunner_Run_Call {
	return &ConsumerRunner_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *ConsumerRunner_Run_Call) Run(run func(ctx context.Context)) *ConsumerRunner_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ConsumerRunner_Run_Call) Return() *ConsumerRunner_Run_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewConsumerRunner interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerRunner creates a new instance of ConsumerRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerRunner(t mockConstructorTestingTNewConsumerRunner) *ConsumerRunner {
	mock := &ConsumerRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
