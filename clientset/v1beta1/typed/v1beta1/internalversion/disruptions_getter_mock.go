// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

package internalversion

import mock "github.com/stretchr/testify/mock"

// DisruptionsGetterMock is an autogenerated mock type for the DisruptionsGetter type
type DisruptionsGetterMock struct {
	mock.Mock
}

type DisruptionsGetterMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DisruptionsGetterMock) EXPECT() *DisruptionsGetterMock_Expecter {
	return &DisruptionsGetterMock_Expecter{mock: &_m.Mock}
}

// Disruptions provides a mock function with given fields: namespace
func (_m *DisruptionsGetterMock) Disruptions(namespace string) DisruptionInterface {
	ret := _m.Called(namespace)

	if len(ret) == 0 {
		panic("no return value specified for Disruptions")
	}

	var r0 DisruptionInterface
	if rf, ok := ret.Get(0).(func(string) DisruptionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DisruptionInterface)
		}
	}

	return r0
}

// DisruptionsGetterMock_Disruptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disruptions'
type DisruptionsGetterMock_Disruptions_Call struct {
	*mock.Call
}

// Disruptions is a helper method to define mock.On call
//   - namespace string
func (_e *DisruptionsGetterMock_Expecter) Disruptions(namespace interface{}) *DisruptionsGetterMock_Disruptions_Call {
	return &DisruptionsGetterMock_Disruptions_Call{Call: _e.mock.On("Disruptions", namespace)}
}

func (_c *DisruptionsGetterMock_Disruptions_Call) Run(run func(namespace string)) *DisruptionsGetterMock_Disruptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DisruptionsGetterMock_Disruptions_Call) Return(_a0 DisruptionInterface) *DisruptionsGetterMock_Disruptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DisruptionsGetterMock_Disruptions_Call) RunAndReturn(run func(string) DisruptionInterface) *DisruptionsGetterMock_Disruptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewDisruptionsGetterMock creates a new instance of DisruptionsGetterMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisruptionsGetterMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisruptionsGetterMock {
	mock := &DisruptionsGetterMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
