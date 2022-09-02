// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	mock "github.com/stretchr/testify/mock"
)

// ErrorResponseManager is an autogenerated mock type for the ErrorResponseManager type
type ErrorResponseManager struct {
	mock.Mock
}

type ErrorResponseManager_Expecter struct {
	mock *mock.Mock
}

func (_m *ErrorResponseManager) EXPECT() *ErrorResponseManager_Expecter {
	return &ErrorResponseManager_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: err
func (_m *ErrorResponseManager) Handle(err error) *response.APIResponseEnvelope {
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

// ErrorResponseManager_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type ErrorResponseManager_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - err error
func (_e *ErrorResponseManager_Expecter) Handle(err interface{}) *ErrorResponseManager_Handle_Call {
	return &ErrorResponseManager_Handle_Call{Call: _e.mock.On("Handle", err)}
}

func (_c *ErrorResponseManager_Handle_Call) Run(run func(err error)) *ErrorResponseManager_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *ErrorResponseManager_Handle_Call) Return(_a0 *response.APIResponseEnvelope) *ErrorResponseManager_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewErrorResponseManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewErrorResponseManager creates a new instance of ErrorResponseManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewErrorResponseManager(t mockConstructorTestingTNewErrorResponseManager) *ErrorResponseManager {
	mock := &ErrorResponseManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
