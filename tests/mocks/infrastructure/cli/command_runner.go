// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CommandRunner is an autogenerated mock type for the CommandRunner type
type CommandRunner struct {
	mock.Mock
}

type CommandRunner_Expecter struct {
	mock *mock.Mock
}

func (_m *CommandRunner) EXPECT() *CommandRunner_Expecter {
	return &CommandRunner_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewCommandRunner interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommandRunner creates a new instance of CommandRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommandRunner(t mockConstructorTestingTNewCommandRunner) *CommandRunner {
	mock := &CommandRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
