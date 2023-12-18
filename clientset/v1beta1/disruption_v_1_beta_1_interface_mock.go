// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package v1beta1

import mock "github.com/stretchr/testify/mock"

// DisruptionV1Beta1InterfaceMock is an autogenerated mock type for the DisruptionV1Beta1Interface type
type DisruptionV1Beta1InterfaceMock struct {
	mock.Mock
}

type DisruptionV1Beta1InterfaceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DisruptionV1Beta1InterfaceMock) EXPECT() *DisruptionV1Beta1InterfaceMock_Expecter {
	return &DisruptionV1Beta1InterfaceMock_Expecter{mock: &_m.Mock}
}

// Disruptions provides a mock function with given fields: namespace
func (_m *DisruptionV1Beta1InterfaceMock) Disruptions(namespace string) DisruptionInterface {
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

// DisruptionV1Beta1InterfaceMock_Disruptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disruptions'
type DisruptionV1Beta1InterfaceMock_Disruptions_Call struct {
	*mock.Call
}

// Disruptions is a helper method to define mock.On call
//   - namespace string
func (_e *DisruptionV1Beta1InterfaceMock_Expecter) Disruptions(namespace interface{}) *DisruptionV1Beta1InterfaceMock_Disruptions_Call {
	return &DisruptionV1Beta1InterfaceMock_Disruptions_Call{Call: _e.mock.On("Disruptions", namespace)}
}

func (_c *DisruptionV1Beta1InterfaceMock_Disruptions_Call) Run(run func(namespace string)) *DisruptionV1Beta1InterfaceMock_Disruptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DisruptionV1Beta1InterfaceMock_Disruptions_Call) Return(_a0 DisruptionInterface) *DisruptionV1Beta1InterfaceMock_Disruptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DisruptionV1Beta1InterfaceMock_Disruptions_Call) RunAndReturn(run func(string) DisruptionInterface) *DisruptionV1Beta1InterfaceMock_Disruptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewDisruptionV1Beta1InterfaceMock creates a new instance of DisruptionV1Beta1InterfaceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisruptionV1Beta1InterfaceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisruptionV1Beta1InterfaceMock {
	mock := &DisruptionV1Beta1InterfaceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
