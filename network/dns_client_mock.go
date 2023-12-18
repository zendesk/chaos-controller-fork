// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package network

import (
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// DNSClientMock is an autogenerated mock type for the DNSClient type
type DNSClientMock struct {
	mock.Mock
}

type DNSClientMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DNSClientMock) EXPECT() *DNSClientMock_Expecter {
	return &DNSClientMock_Expecter{mock: &_m.Mock}
}

// Resolve provides a mock function with given fields: host
func (_m *DNSClientMock) Resolve(host string) ([]net.IP, error) {
	ret := _m.Called(host)

	if len(ret) == 0 {
		panic("no return value specified for Resolve")
	}

	var r0 []net.IP
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]net.IP, error)); ok {
		return rf(host)
	}
	if rf, ok := ret.Get(0).(func(string) []net.IP); ok {
		r0 = rf(host)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]net.IP)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(host)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DNSClientMock_Resolve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolve'
type DNSClientMock_Resolve_Call struct {
	*mock.Call
}

// Resolve is a helper method to define mock.On call
//   - host string
func (_e *DNSClientMock_Expecter) Resolve(host interface{}) *DNSClientMock_Resolve_Call {
	return &DNSClientMock_Resolve_Call{Call: _e.mock.On("Resolve", host)}
}

func (_c *DNSClientMock_Resolve_Call) Run(run func(host string)) *DNSClientMock_Resolve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DNSClientMock_Resolve_Call) Return(_a0 []net.IP, _a1 error) *DNSClientMock_Resolve_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DNSClientMock_Resolve_Call) RunAndReturn(run func(string) ([]net.IP, error)) *DNSClientMock_Resolve_Call {
	_c.Call.Return(run)
	return _c
}

// NewDNSClientMock creates a new instance of DNSClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDNSClientMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DNSClientMock {
	mock := &DNSClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
