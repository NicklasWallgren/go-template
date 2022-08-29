// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	response "github.com/NicklasWallgren/go-template/adapters/driven/api/response"
)

// RouteHandler is an autogenerated mock type for the RouteHandler type
type RouteHandler struct {
	mock.Mock
}

type RouteHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *RouteHandler) EXPECT() *RouteHandler_Expecter {
	return &RouteHandler_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: c
func (_m *RouteHandler) Execute(c *gin.Context) (response.APIResponseEnvelope, error) {
	ret := _m.Called(c)

	var r0 response.APIResponseEnvelope
	if rf, ok := ret.Get(0).(func(*gin.Context) response.APIResponseEnvelope); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.APIResponseEnvelope)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RouteHandler_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type RouteHandler_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//  - c *gin.Context
func (_e *RouteHandler_Expecter) Execute(c interface{}) *RouteHandler_Execute_Call {
	return &RouteHandler_Execute_Call{Call: _e.mock.On("Execute", c)}
}

func (_c *RouteHandler_Execute_Call) Run(run func(c *gin.Context)) *RouteHandler_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *RouteHandler_Execute_Call) Return(_a0 response.APIResponseEnvelope, _a1 error) *RouteHandler_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewRouteHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewRouteHandler creates a new instance of RouteHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRouteHandler(t mockConstructorTestingTNewRouteHandler) *RouteHandler {
	mock := &RouteHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
