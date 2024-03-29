// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	errors "github.com/NicklasWallgren/go-template/adapters/driven/persistence/errors"
	mock "github.com/stretchr/testify/mock"
)

// DBErrorOption is an autogenerated mock type for the DBErrorOption type
type DBErrorOption struct {
	mock.Mock
}

type DBErrorOption_Expecter struct {
	mock *mock.Mock
}

func (_m *DBErrorOption) EXPECT() *DBErrorOption_Expecter {
	return &DBErrorOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: dbError
func (_m *DBErrorOption) Execute(dbError *errors.DBError) {
	_m.Called(dbError)
}

// DBErrorOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type DBErrorOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//  - dbError *errors.DBError
func (_e *DBErrorOption_Expecter) Execute(dbError interface{}) *DBErrorOption_Execute_Call {
	return &DBErrorOption_Execute_Call{Call: _e.mock.On("Execute", dbError)}
}

func (_c *DBErrorOption_Execute_Call) Run(run func(dbError *errors.DBError)) *DBErrorOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*errors.DBError))
	})
	return _c
}

func (_c *DBErrorOption_Execute_Call) Return() *DBErrorOption_Execute_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewDBErrorOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewDBErrorOption creates a new instance of DBErrorOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBErrorOption(t mockConstructorTestingTNewDBErrorOption) *DBErrorOption {
	mock := &DBErrorOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
