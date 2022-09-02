// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	mock "github.com/stretchr/testify/mock"
)

// ErrorTypeResponseHandler is an autogenerated mock type for the ErrorTypeResponseHandler type
type ErrorTypeResponseHandler struct {
	mock.Mock
}

type ErrorTypeResponseHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *ErrorTypeResponseHandler) EXPECT() *ErrorTypeResponseHandler_Expecter {
	return &ErrorTypeResponseHandler_Expecter{mock: &_m.Mock}
}

// ErrorType provides a mock function with given fields:
func (_m *ErrorTypeResponseHandler) ErrorType() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ErrorTypeResponseHandler_ErrorType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ErrorType'
type ErrorTypeResponseHandler_ErrorType_Call struct {
	*mock.Call
}

// ErrorType is a helper method to define mock.On call
func (_e *ErrorTypeResponseHandler_Expecter) ErrorType() *ErrorTypeResponseHandler_ErrorType_Call {
	return &ErrorTypeResponseHandler_ErrorType_Call{Call: _e.mock.On("ErrorType")}
}

func (_c *ErrorTypeResponseHandler_ErrorType_Call) Run(run func()) *ErrorTypeResponseHandler_ErrorType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ErrorTypeResponseHandler_ErrorType_Call) Return(_a0 error) *ErrorTypeResponseHandler_ErrorType_Call {
	_c.Call.Return(_a0)
	return _c
}

// Handle provides a mock function with given fields: err
func (_m *ErrorTypeResponseHandler) Handle(err error) *response.APIResponseEnvelope {
	ret := _m.Called(err)

	var r0 *response.APIResponseEnvelope
	if rf, ok := ret.Get(0).(func(error) *response.APIResponseEnvelope); ok {
		r0 = rf(err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.APIResponseEnvelope)
		}
	}

	return r0
}

// ErrorTypeResponseHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type ErrorTypeResponseHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - err error
func (_e *ErrorTypeResponseHandler_Expecter) Handle(err interface{}) *ErrorTypeResponseHandler_Handle_Call {
	return &ErrorTypeResponseHandler_Handle_Call{Call: _e.mock.On("Handle", err)}
}

func (_c *ErrorTypeResponseHandler_Handle_Call) Run(run func(err error)) *ErrorTypeResponseHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *ErrorTypeResponseHandler_Handle_Call) Return(_a0 *response.APIResponseEnvelope) *ErrorTypeResponseHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

// IsSupported provides a mock function with given fields: err
func (_m *ErrorTypeResponseHandler) IsSupported(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ErrorTypeResponseHandler_IsSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSupported'
type ErrorTypeResponseHandler_IsSupported_Call struct {
	*mock.Call
}

// IsSupported is a helper method to define mock.On call
//   - err error
func (_e *ErrorTypeResponseHandler_Expecter) IsSupported(err interface{}) *ErrorTypeResponseHandler_IsSupported_Call {
	return &ErrorTypeResponseHandler_IsSupported_Call{Call: _e.mock.On("IsSupported", err)}
}

func (_c *ErrorTypeResponseHandler_IsSupported_Call) Run(run func(err error)) *ErrorTypeResponseHandler_IsSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *ErrorTypeResponseHandler_IsSupported_Call) Return(_a0 bool) *ErrorTypeResponseHandler_IsSupported_Call {
	_c.Call.Return(_a0)
	return _c
}

// Priority provides a mock function with given fields:
func (_m *ErrorTypeResponseHandler) Priority() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ErrorTypeResponseHandler_Priority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Priority'
type ErrorTypeResponseHandler_Priority_Call struct {
	*mock.Call
}

// Priority is a helper method to define mock.On call
func (_e *ErrorTypeResponseHandler_Expecter) Priority() *ErrorTypeResponseHandler_Priority_Call {
	return &ErrorTypeResponseHandler_Priority_Call{Call: _e.mock.On("Priority")}
}

func (_c *ErrorTypeResponseHandler_Priority_Call) Run(run func()) *ErrorTypeResponseHandler_Priority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ErrorTypeResponseHandler_Priority_Call) Return(_a0 int) *ErrorTypeResponseHandler_Priority_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewErrorTypeResponseHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewErrorTypeResponseHandler creates a new instance of ErrorTypeResponseHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewErrorTypeResponseHandler(t mockConstructorTestingTNewErrorTypeResponseHandler) *ErrorTypeResponseHandler {
	mock := &ErrorTypeResponseHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
