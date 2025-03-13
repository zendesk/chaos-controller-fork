// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

package internalversion

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"
)

// ChaosInterfaceMock is an autogenerated mock type for the ChaosInterface type
type ChaosInterfaceMock struct {
	mock.Mock
}

type ChaosInterfaceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ChaosInterfaceMock) EXPECT() *ChaosInterfaceMock_Expecter {
	return &ChaosInterfaceMock_Expecter{mock: &_m.Mock}
}

// DisruptionCrons provides a mock function with given fields: namespace
func (_m *ChaosInterfaceMock) DisruptionCrons(namespace string) DisruptionCronInterface {
	ret := _m.Called(namespace)

	if len(ret) == 0 {
		panic("no return value specified for DisruptionCrons")
	}

	var r0 DisruptionCronInterface
	if rf, ok := ret.Get(0).(func(string) DisruptionCronInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DisruptionCronInterface)
		}
	}

	return r0
}

// ChaosInterfaceMock_DisruptionCrons_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisruptionCrons'
type ChaosInterfaceMock_DisruptionCrons_Call struct {
	*mock.Call
}

// DisruptionCrons is a helper method to define mock.On call
//   - namespace string
func (_e *ChaosInterfaceMock_Expecter) DisruptionCrons(namespace interface{}) *ChaosInterfaceMock_DisruptionCrons_Call {
	return &ChaosInterfaceMock_DisruptionCrons_Call{Call: _e.mock.On("DisruptionCrons", namespace)}
}

func (_c *ChaosInterfaceMock_DisruptionCrons_Call) Run(run func(namespace string)) *ChaosInterfaceMock_DisruptionCrons_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ChaosInterfaceMock_DisruptionCrons_Call) Return(_a0 DisruptionCronInterface) *ChaosInterfaceMock_DisruptionCrons_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChaosInterfaceMock_DisruptionCrons_Call) RunAndReturn(run func(string) DisruptionCronInterface) *ChaosInterfaceMock_DisruptionCrons_Call {
	_c.Call.Return(run)
	return _c
}

// Disruptions provides a mock function with given fields: namespace
func (_m *ChaosInterfaceMock) Disruptions(namespace string) DisruptionInterface {
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

// ChaosInterfaceMock_Disruptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disruptions'
type ChaosInterfaceMock_Disruptions_Call struct {
	*mock.Call
}

// Disruptions is a helper method to define mock.On call
//   - namespace string
func (_e *ChaosInterfaceMock_Expecter) Disruptions(namespace interface{}) *ChaosInterfaceMock_Disruptions_Call {
	return &ChaosInterfaceMock_Disruptions_Call{Call: _e.mock.On("Disruptions", namespace)}
}

func (_c *ChaosInterfaceMock_Disruptions_Call) Run(run func(namespace string)) *ChaosInterfaceMock_Disruptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ChaosInterfaceMock_Disruptions_Call) Return(_a0 DisruptionInterface) *ChaosInterfaceMock_Disruptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChaosInterfaceMock_Disruptions_Call) RunAndReturn(run func(string) DisruptionInterface) *ChaosInterfaceMock_Disruptions_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with no fields
func (_m *ChaosInterfaceMock) RESTClient() rest.Interface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RESTClient")
	}

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// ChaosInterfaceMock_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type ChaosInterfaceMock_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *ChaosInterfaceMock_Expecter) RESTClient() *ChaosInterfaceMock_RESTClient_Call {
	return &ChaosInterfaceMock_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *ChaosInterfaceMock_RESTClient_Call) Run(run func()) *ChaosInterfaceMock_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChaosInterfaceMock_RESTClient_Call) Return(_a0 rest.Interface) *ChaosInterfaceMock_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChaosInterfaceMock_RESTClient_Call) RunAndReturn(run func() rest.Interface) *ChaosInterfaceMock_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

// NewChaosInterfaceMock creates a new instance of ChaosInterfaceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChaosInterfaceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChaosInterfaceMock {
	mock := &ChaosInterfaceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
