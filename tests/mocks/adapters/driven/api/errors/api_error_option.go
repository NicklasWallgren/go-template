// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	"github.com/NicklasWallgren/go-template/adapters/driver/api/errors"
	mock "github.com/stretchr/testify/mock"
)

// APIErrorOption is an autogenerated mock type for the APIErrorOption type
type APIErrorOption struct {
	mock.Mock
}

type APIErrorOption_Expecter struct {
	mock *mock.Mock
}

func (_m *APIErrorOption) EXPECT() *APIErrorOption_Expecter {
	return &APIErrorOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: apiError
func (_m *APIErrorOption) Execute(apiError *errors.APIError) {
	_m.Called(apiError)
}

// APIErrorOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type APIErrorOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - apiError *errors.APIError
func (_e *APIErrorOption_Expecter) Execute(apiError interface{}) *APIErrorOption_Execute_Call {
	return &APIErrorOption_Execute_Call{Call: _e.mock.On("Execute", apiError)}
}

func (_c *APIErrorOption_Execute_Call) Run(run func(apiError *errors.APIError)) *APIErrorOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*errors.APIError))
	})
	return _c
}

func (_c *APIErrorOption_Execute_Call) Return() *APIErrorOption_Execute_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewAPIErrorOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewAPIErrorOption creates a new instance of APIErrorOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAPIErrorOption(t mockConstructorTestingTNewAPIErrorOption) *APIErrorOption {
	mock := &APIErrorOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
