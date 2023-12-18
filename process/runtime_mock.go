// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package process

import mock "github.com/stretchr/testify/mock"

// RuntimeMock is an autogenerated mock type for the Runtime type
type RuntimeMock struct {
	mock.Mock
}

type RuntimeMock_Expecter struct {
	mock *mock.Mock
}

func (_m *RuntimeMock) EXPECT() *RuntimeMock_Expecter {
	return &RuntimeMock_Expecter{mock: &_m.Mock}
}

// GOMAXPROCS provides a mock function with given fields: _a0
func (_m *RuntimeMock) GOMAXPROCS(_a0 int) int {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GOMAXPROCS")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// RuntimeMock_GOMAXPROCS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GOMAXPROCS'
type RuntimeMock_GOMAXPROCS_Call struct {
	*mock.Call
}

// GOMAXPROCS is a helper method to define mock.On call
//   - _a0 int
func (_e *RuntimeMock_Expecter) GOMAXPROCS(_a0 interface{}) *RuntimeMock_GOMAXPROCS_Call {
	return &RuntimeMock_GOMAXPROCS_Call{Call: _e.mock.On("GOMAXPROCS", _a0)}
}

func (_c *RuntimeMock_GOMAXPROCS_Call) Run(run func(_a0 int)) *RuntimeMock_GOMAXPROCS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *RuntimeMock_GOMAXPROCS_Call) Return(_a0 int) *RuntimeMock_GOMAXPROCS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuntimeMock_GOMAXPROCS_Call) RunAndReturn(run func(int) int) *RuntimeMock_GOMAXPROCS_Call {
	_c.Call.Return(run)
	return _c
}

// LockOSThread provides a mock function with given fields:
func (_m *RuntimeMock) LockOSThread() {
	_m.Called()
}

// RuntimeMock_LockOSThread_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LockOSThread'
type RuntimeMock_LockOSThread_Call struct {
	*mock.Call
}

// LockOSThread is a helper method to define mock.On call
func (_e *RuntimeMock_Expecter) LockOSThread() *RuntimeMock_LockOSThread_Call {
	return &RuntimeMock_LockOSThread_Call{Call: _e.mock.On("LockOSThread")}
}

func (_c *RuntimeMock_LockOSThread_Call) Run(run func()) *RuntimeMock_LockOSThread_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuntimeMock_LockOSThread_Call) Return() *RuntimeMock_LockOSThread_Call {
	_c.Call.Return()
	return _c
}

func (_c *RuntimeMock_LockOSThread_Call) RunAndReturn(run func()) *RuntimeMock_LockOSThread_Call {
	_c.Call.Return(run)
	return _c
}

// UnlockOSThread provides a mock function with given fields:
func (_m *RuntimeMock) UnlockOSThread() {
	_m.Called()
}

// RuntimeMock_UnlockOSThread_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnlockOSThread'
type RuntimeMock_UnlockOSThread_Call struct {
	*mock.Call
}

// UnlockOSThread is a helper method to define mock.On call
func (_e *RuntimeMock_Expecter) UnlockOSThread() *RuntimeMock_UnlockOSThread_Call {
	return &RuntimeMock_UnlockOSThread_Call{Call: _e.mock.On("UnlockOSThread")}
}

func (_c *RuntimeMock_UnlockOSThread_Call) Run(run func()) *RuntimeMock_UnlockOSThread_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuntimeMock_UnlockOSThread_Call) Return() *RuntimeMock_UnlockOSThread_Call {
	_c.Call.Return()
	return _c
}

func (_c *RuntimeMock_UnlockOSThread_Call) RunAndReturn(run func()) *RuntimeMock_UnlockOSThread_Call {
	_c.Call.Return(run)
	return _c
}

// NewRuntimeMock creates a new instance of RuntimeMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRuntimeMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RuntimeMock {
	mock := &RuntimeMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
