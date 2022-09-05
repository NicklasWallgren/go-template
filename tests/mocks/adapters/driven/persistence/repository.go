// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	persistence "github.com/NicklasWallgren/go-template/adapters/driven/persistence"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Gorm provides a mock function with given fields:
func (_m *Repository) Gorm() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Repository_Gorm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Gorm'
type Repository_Gorm_Call struct {
	*mock.Call
}

// Gorm is a helper method to define mock.On call
func (_e *Repository_Expecter) Gorm() *Repository_Gorm_Call {
	return &Repository_Gorm_Call{Call: _e.mock.On("Gorm")}
}

func (_c *Repository_Gorm_Call) Run(run func()) *Repository_Gorm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_Gorm_Call) Return(_a0 *gorm.DB) *Repository_Gorm_Call {
	_c.Call.Return(_a0)
	return _c
}

// RawSql provides a mock function with given fields: ctx, sql, values
func (_m *Repository) RawSql(ctx context.Context, sql string, values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, sql, values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_RawSql_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RawSql'
type Repository_RawSql_Call struct {
	*mock.Call
}

// RawSql is a helper method to define mock.On call
//  - ctx context.Context
//  - sql string
//  - values ...interface{}
func (_e *Repository_Expecter) RawSql(ctx interface{}, sql interface{}, values ...interface{}) *Repository_RawSql_Call {
	return &Repository_RawSql_Call{Call: _e.mock.On("RawSql",
		append([]interface{}{ctx, sql}, values...)...)}
}

func (_c *Repository_RawSql_Call) Run(run func(ctx context.Context, sql string, values ...interface{})) *Repository_RawSql_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Repository_RawSql_Call) Return(_a0 error) *Repository_RawSql_Call {
	_c.Call.Return(_a0)
	return _c
}

// WithTx provides a mock function with given fields: tx
func (_m *Repository) WithTx(tx *gorm.DB) persistence.Repository {
	ret := _m.Called(tx)

	var r0 persistence.Repository
	if rf, ok := ret.Get(0).(func(*gorm.DB) persistence.Repository); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(persistence.Repository)
		}
	}

	return r0
}

// Repository_WithTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTx'
type Repository_WithTx_Call struct {
	*mock.Call
}

// WithTx is a helper method to define mock.On call
//  - tx *gorm.DB
func (_e *Repository_Expecter) WithTx(tx interface{}) *Repository_WithTx_Call {
	return &Repository_WithTx_Call{Call: _e.mock.On("WithTx", tx)}
}

func (_c *Repository_WithTx_Call) Run(run func(tx *gorm.DB)) *Repository_WithTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB))
	})
	return _c
}

func (_c *Repository_WithTx_Call) Return(_a0 persistence.Repository) *Repository_WithTx_Call {
	_c.Call.Return(_a0)
	return _c
}

// WrapError provides a mock function with given fields: err
func (_m *Repository) WrapError(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_WrapError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WrapError'
type Repository_WrapError_Call struct {
	*mock.Call
}

// WrapError is a helper method to define mock.On call
//  - err error
func (_e *Repository_Expecter) WrapError(err interface{}) *Repository_WrapError_Call {
	return &Repository_WrapError_Call{Call: _e.mock.On("WrapError", err)}
}

func (_c *Repository_WrapError_Call) Run(run func(err error)) *Repository_WrapError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *Repository_WrapError_Call) Return(_a0 error) *Repository_WrapError_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
