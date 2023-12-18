// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package safemode

import (
	mock "github.com/stretchr/testify/mock"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
)

// SafemodeMock is an autogenerated mock type for the Safemode type
type SafemodeMock struct {
	mock.Mock
}

type SafemodeMock_Expecter struct {
	mock *mock.Mock
}

func (_m *SafemodeMock) EXPECT() *SafemodeMock_Expecter {
	return &SafemodeMock_Expecter{mock: &_m.Mock}
}

// Init provides a mock function with given fields: disruption, _a1
func (_m *SafemodeMock) Init(disruption v1beta1.Disruption, _a1 client.Client) {
	_m.Called(disruption, _a1)
}

// SafemodeMock_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type SafemodeMock_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - disruption v1beta1.Disruption
//   - _a1 client.Client
func (_e *SafemodeMock_Expecter) Init(disruption interface{}, _a1 interface{}) *SafemodeMock_Init_Call {
	return &SafemodeMock_Init_Call{Call: _e.mock.On("Init", disruption, _a1)}
}

func (_c *SafemodeMock_Init_Call) Run(run func(disruption v1beta1.Disruption, _a1 client.Client)) *SafemodeMock_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(v1beta1.Disruption), args[1].(client.Client))
	})
	return _c
}

func (_c *SafemodeMock_Init_Call) Return() *SafemodeMock_Init_Call {
	_c.Call.Return()
	return _c
}

func (_c *SafemodeMock_Init_Call) RunAndReturn(run func(v1beta1.Disruption, client.Client)) *SafemodeMock_Init_Call {
	_c.Call.Return(run)
	return _c
}

// NewSafemodeMock creates a new instance of SafemodeMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSafemodeMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *SafemodeMock {
	mock := &SafemodeMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
