// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Pageable is an autogenerated mock type for the Pageable type
type Pageable struct {
	mock.Mock
}

type Pageable_Expecter struct {
	mock *mock.Mock
}

func (_m *Pageable) EXPECT() *Pageable_Expecter {
	return &Pageable_Expecter{mock: &_m.Mock}
}

// Offset provides a mock function with given fields:
func (_m *Pageable) Offset() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Pageable_Offset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Offset'
type Pageable_Offset_Call struct {
	*mock.Call
}

// Offset is a helper method to define mock.On call
func (_e *Pageable_Expecter) Offset() *Pageable_Offset_Call {
	return &Pageable_Offset_Call{Call: _e.mock.On("Offset")}
}

func (_c *Pageable_Offset_Call) Run(run func()) *Pageable_Offset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Pageable_Offset_Call) Return(_a0 int) *Pageable_Offset_Call {
	_c.Call.Return(_a0)
	return _c
}

// PageNumber provides a mock function with given fields:
func (_m *Pageable) PageNumber() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Pageable_PageNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PageNumber'
type Pageable_PageNumber_Call struct {
	*mock.Call
}

// PageNumber is a helper method to define mock.On call
func (_e *Pageable_Expecter) PageNumber() *Pageable_PageNumber_Call {
	return &Pageable_PageNumber_Call{Call: _e.mock.On("PageNumber")}
}

func (_c *Pageable_PageNumber_Call) Run(run func()) *Pageable_PageNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Pageable_PageNumber_Call) Return(_a0 int) *Pageable_PageNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

// PageSize provides a mock function with given fields:
func (_m *Pageable) PageSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Pageable_PageSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PageSize'
type Pageable_PageSize_Call struct {
	*mock.Call
}

// PageSize is a helper method to define mock.On call
func (_e *Pageable_Expecter) PageSize() *Pageable_PageSize_Call {
	return &Pageable_PageSize_Call{Call: _e.mock.On("PageSize")}
}

func (_c *Pageable_PageSize_Call) Run(run func()) *Pageable_PageSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Pageable_PageSize_Call) Return(_a0 int) *Pageable_PageSize_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewPageable interface {
	mock.TestingT
	Cleanup(func())
}

// NewPageable creates a new instance of Pageable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPageable(t mockConstructorTestingTNewPageable) *Pageable {
	mock := &Pageable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
