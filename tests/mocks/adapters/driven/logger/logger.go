// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	fxevent "go.uber.org/fx/fxevent"
	gormlogger "gorm.io/gorm/logger"

	mock "github.com/stretchr/testify/mock"

	zap "go.uber.org/zap"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

type Logger_Expecter struct {
	mock *mock.Mock
}

func (_m *Logger) EXPECT() *Logger_Expecter {
	return &Logger_Expecter{mock: &_m.Mock}
}

// Debug provides a mock function with given fields: args
func (_m *Logger) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type Logger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//  - args ...interface{}
func (_e *Logger_Expecter) Debug(args ...interface{}) *Logger_Debug_Call {
	return &Logger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{}, args...)...)}
}

func (_c *Logger_Debug_Call) Run(run func(args ...interface{})) *Logger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Logger_Debug_Call) Return() *Logger_Debug_Call {
	_c.Call.Return()
	return _c
}

// Debugf provides a mock function with given fields: template, args
func (_m *Logger) Debugf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Debugf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugf'
type Logger_Debugf_Call struct {
	*mock.Call
}

// Debugf is a helper method to define mock.On call
//  - template string
//  - args ...interface{}
func (_e *Logger_Expecter) Debugf(template interface{}, args ...interface{}) *Logger_Debugf_Call {
	return &Logger_Debugf_Call{Call: _e.mock.On("Debugf",
		append([]interface{}{template}, args...)...)}
}

func (_c *Logger_Debugf_Call) Run(run func(template string, args ...interface{})) *Logger_Debugf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Debugf_Call) Return() *Logger_Debugf_Call {
	_c.Call.Return()
	return _c
}

// Error provides a mock function with given fields: args
func (_m *Logger) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Logger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//  - args ...interface{}
func (_e *Logger_Expecter) Error(args ...interface{}) *Logger_Error_Call {
	return &Logger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{}, args...)...)}
}

func (_c *Logger_Error_Call) Run(run func(args ...interface{})) *Logger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Logger_Error_Call) Return() *Logger_Error_Call {
	_c.Call.Return()
	return _c
}

// Errorf provides a mock function with given fields: template, args
func (_m *Logger) Errorf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type Logger_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//  - template string
//  - args ...interface{}
func (_e *Logger_Expecter) Errorf(template interface{}, args ...interface{}) *Logger_Errorf_Call {
	return &Logger_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{template}, args...)...)}
}

func (_c *Logger_Errorf_Call) Run(run func(template string, args ...interface{})) *Logger_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Errorf_Call) Return() *Logger_Errorf_Call {
	_c.Call.Return()
	return _c
}

// GetFxLogger provides a mock function with given fields:
func (_m *Logger) GetFxLogger() fxevent.Logger {
	ret := _m.Called()

	var r0 fxevent.Logger
	if rf, ok := ret.Get(0).(func() fxevent.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fxevent.Logger)
		}
	}

	return r0
}

// Logger_GetFxLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFxLogger'
type Logger_GetFxLogger_Call struct {
	*mock.Call
}

// GetFxLogger is a helper method to define mock.On call
func (_e *Logger_Expecter) GetFxLogger() *Logger_GetFxLogger_Call {
	return &Logger_GetFxLogger_Call{Call: _e.mock.On("GetFxLogger")}
}

func (_c *Logger_GetFxLogger_Call) Run(run func()) *Logger_GetFxLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Logger_GetFxLogger_Call) Return(_a0 fxevent.Logger) *Logger_GetFxLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetGormLogger provides a mock function with given fields:
func (_m *Logger) GetGormLogger() gormlogger.Interface {
	ret := _m.Called()

	var r0 gormlogger.Interface
	if rf, ok := ret.Get(0).(func() gormlogger.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gormlogger.Interface)
		}
	}

	return r0
}

// Logger_GetGormLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGormLogger'
type Logger_GetGormLogger_Call struct {
	*mock.Call
}

// GetGormLogger is a helper method to define mock.On call
func (_e *Logger_Expecter) GetGormLogger() *Logger_GetGormLogger_Call {
	return &Logger_GetGormLogger_Call{Call: _e.mock.On("GetGormLogger")}
}

func (_c *Logger_GetGormLogger_Call) Run(run func()) *Logger_GetGormLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Logger_GetGormLogger_Call) Return(_a0 gormlogger.Interface) *Logger_GetGormLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetZapLogger provides a mock function with given fields:
func (_m *Logger) GetZapLogger() *zap.Logger {
	ret := _m.Called()

	var r0 *zap.Logger
	if rf, ok := ret.Get(0).(func() *zap.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zap.Logger)
		}
	}

	return r0
}

// Logger_GetZapLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetZapLogger'
type Logger_GetZapLogger_Call struct {
	*mock.Call
}

// GetZapLogger is a helper method to define mock.On call
func (_e *Logger_Expecter) GetZapLogger() *Logger_GetZapLogger_Call {
	return &Logger_GetZapLogger_Call{Call: _e.mock.On("GetZapLogger")}
}

func (_c *Logger_GetZapLogger_Call) Run(run func()) *Logger_GetZapLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Logger_GetZapLogger_Call) Return(_a0 *zap.Logger) *Logger_GetZapLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

// Info provides a mock function with given fields: args
func (_m *Logger) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type Logger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//  - args ...interface{}
func (_e *Logger_Expecter) Info(args ...interface{}) *Logger_Info_Call {
	return &Logger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{}, args...)...)}
}

func (_c *Logger_Info_Call) Run(run func(args ...interface{})) *Logger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Logger_Info_Call) Return() *Logger_Info_Call {
	_c.Call.Return()
	return _c
}

// Infof provides a mock function with given fields: template, args
func (_m *Logger) Infof(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type Logger_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//  - template string
//  - args ...interface{}
func (_e *Logger_Expecter) Infof(template interface{}, args ...interface{}) *Logger_Infof_Call {
	return &Logger_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{template}, args...)...)}
}

func (_c *Logger_Infof_Call) Run(run func(template string, args ...interface{})) *Logger_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Infof_Call) Return() *Logger_Infof_Call {
	_c.Call.Return()
	return _c
}

// Warn provides a mock function with given fields: args
func (_m *Logger) Warn(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type Logger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//  - args ...interface{}
func (_e *Logger_Expecter) Warn(args ...interface{}) *Logger_Warn_Call {
	return &Logger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{}, args...)...)}
}

func (_c *Logger_Warn_Call) Run(run func(args ...interface{})) *Logger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Logger_Warn_Call) Return() *Logger_Warn_Call {
	_c.Call.Return()
	return _c
}

// Warnf provides a mock function with given fields: template, args
func (_m *Logger) Warnf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logger_Warnf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnf'
type Logger_Warnf_Call struct {
	*mock.Call
}

// Warnf is a helper method to define mock.On call
//  - template string
//  - args ...interface{}
func (_e *Logger_Expecter) Warnf(template interface{}, args ...interface{}) *Logger_Warnf_Call {
	return &Logger_Warnf_Call{Call: _e.mock.On("Warnf",
		append([]interface{}{template}, args...)...)}
}

func (_c *Logger_Warnf_Call) Run(run func(template string, args ...interface{})) *Logger_Warnf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Warnf_Call) Return() *Logger_Warnf_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
