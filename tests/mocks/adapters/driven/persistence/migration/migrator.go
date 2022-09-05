// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Migrator is an autogenerated mock type for the Migrator type
type Migrator struct {
	mock.Mock
}

type Migrator_Expecter struct {
	mock *mock.Mock
}

func (_m *Migrator) EXPECT() *Migrator_Expecter {
	return &Migrator_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: filename
func (_m *Migrator) Create(filename string) error {
	ret := _m.Called(filename)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Migrator_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - filename string
func (_e *Migrator_Expecter) Create(filename interface{}) *Migrator_Create_Call {
	return &Migrator_Create_Call{Call: _e.mock.On("Create", filename)}
}

func (_c *Migrator_Create_Call) Run(run func(filename string)) *Migrator_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Migrator_Create_Call) Return(_a0 error) *Migrator_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Down provides a mock function with given fields:
func (_m *Migrator) Down() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Down_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Down'
type Migrator_Down_Call struct {
	*mock.Call
}

// Down is a helper method to define mock.On call
func (_e *Migrator_Expecter) Down() *Migrator_Down_Call {
	return &Migrator_Down_Call{Call: _e.mock.On("Down")}
}

func (_c *Migrator_Down_Call) Run(run func()) *Migrator_Down_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Down_Call) Return(_a0 error) *Migrator_Down_Call {
	_c.Call.Return(_a0)
	return _c
}

// Fix provides a mock function with given fields:
func (_m *Migrator) Fix() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Fix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fix'
type Migrator_Fix_Call struct {
	*mock.Call
}

// Fix is a helper method to define mock.On call
func (_e *Migrator_Expecter) Fix() *Migrator_Fix_Call {
	return &Migrator_Fix_Call{Call: _e.mock.On("Fix")}
}

func (_c *Migrator_Fix_Call) Run(run func()) *Migrator_Fix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Fix_Call) Return(_a0 error) *Migrator_Fix_Call {
	_c.Call.Return(_a0)
	return _c
}

// Up provides a mock function with given fields:
func (_m *Migrator) Up() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Up_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Up'
type Migrator_Up_Call struct {
	*mock.Call
}

// Up is a helper method to define mock.On call
func (_e *Migrator_Expecter) Up() *Migrator_Up_Call {
	return &Migrator_Up_Call{Call: _e.mock.On("Up")}
}

func (_c *Migrator_Up_Call) Run(run func()) *Migrator_Up_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Up_Call) Return(_a0 error) *Migrator_Up_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMigrator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMigrator creates a new instance of Migrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMigrator(t mockConstructorTestingTNewMigrator) *Migrator {
	mock := &Migrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}