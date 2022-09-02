// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	common "github.com/NicklasWallgren/go-template/domain/common"
	mock "github.com/stretchr/testify/mock"
)

// EntityConstraint is an autogenerated mock type for the EntityConstraint type
type EntityConstraint struct {
	mock.Mock
}

type EntityConstraint_Expecter struct {
	mock *mock.Mock
}

func (_m *EntityConstraint) EXPECT() *EntityConstraint_Expecter {
	return &EntityConstraint_Expecter{mock: &_m.Mock}
}

// Id provides a mock function with given fields:
func (_m *EntityConstraint) Id() common.PrimaryID {
	ret := _m.Called()

	var r0 common.PrimaryID
	if rf, ok := ret.Get(0).(func() common.PrimaryID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.PrimaryID)
	}

	return r0
}

// EntityConstraint_Id_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Id'
type EntityConstraint_Id_Call struct {
	*mock.Call
}

// Id is a helper method to define mock.On call
func (_e *EntityConstraint_Expecter) Id() *EntityConstraint_Id_Call {
	return &EntityConstraint_Id_Call{Call: _e.mock.On("Id")}
}

func (_c *EntityConstraint_Id_Call) Run(run func()) *EntityConstraint_Id_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityConstraint_Id_Call) Return(_a0 common.PrimaryID) *EntityConstraint_Id_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEntityConstraint interface {
	mock.TestingT
	Cleanup(func())
}

// NewEntityConstraint creates a new instance of EntityConstraint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEntityConstraint(t mockConstructorTestingTNewEntityConstraint) *EntityConstraint {
	mock := &EntityConstraint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
