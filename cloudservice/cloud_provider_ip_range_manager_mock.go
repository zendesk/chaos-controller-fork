// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package cloudservice

import (
	types "github.com/DataDog/chaos-controller/cloudservice/types"
	mock "github.com/stretchr/testify/mock"
)

// CloudProviderIPRangeManagerMock is an autogenerated mock type for the CloudProviderIPRangeManager type
type CloudProviderIPRangeManagerMock struct {
	mock.Mock
}

type CloudProviderIPRangeManagerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *CloudProviderIPRangeManagerMock) EXPECT() *CloudProviderIPRangeManagerMock_Expecter {
	return &CloudProviderIPRangeManagerMock_Expecter{mock: &_m.Mock}
}

// ConvertToGenericIPRanges provides a mock function with given fields: ipRangeData
func (_m *CloudProviderIPRangeManagerMock) ConvertToGenericIPRanges(ipRangeData []byte) (*types.CloudProviderIPRangeInfo, error) {
	ret := _m.Called(ipRangeData)

	var r0 *types.CloudProviderIPRangeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (*types.CloudProviderIPRangeInfo, error)); ok {
		return rf(ipRangeData)
	}
	if rf, ok := ret.Get(0).(func([]byte) *types.CloudProviderIPRangeInfo); ok {
		r0 = rf(ipRangeData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.CloudProviderIPRangeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(ipRangeData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConvertToGenericIPRanges'
type CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call struct {
	*mock.Call
}

// ConvertToGenericIPRanges is a helper method to define mock.On call
//   - ipRangeData []byte
func (_e *CloudProviderIPRangeManagerMock_Expecter) ConvertToGenericIPRanges(ipRangeData interface{}) *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call {
	return &CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call{Call: _e.mock.On("ConvertToGenericIPRanges", ipRangeData)}
}

func (_c *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call) Run(run func(ipRangeData []byte)) *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call) Return(_a0 *types.CloudProviderIPRangeInfo, _a1 error) *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call) RunAndReturn(run func([]byte) (*types.CloudProviderIPRangeInfo, error)) *CloudProviderIPRangeManagerMock_ConvertToGenericIPRanges_Call {
	_c.Call.Return(run)
	return _c
}

// IsNewVersion provides a mock function with given fields: ipRangeData, version
func (_m *CloudProviderIPRangeManagerMock) IsNewVersion(ipRangeData []byte, version string) (bool, error) {
	ret := _m.Called(ipRangeData, version)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string) (bool, error)); ok {
		return rf(ipRangeData, version)
	}
	if rf, ok := ret.Get(0).(func([]byte, string) bool); ok {
		r0 = rf(ipRangeData, version)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(ipRangeData, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudProviderIPRangeManagerMock_IsNewVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNewVersion'
type CloudProviderIPRangeManagerMock_IsNewVersion_Call struct {
	*mock.Call
}

// IsNewVersion is a helper method to define mock.On call
//   - ipRangeData []byte
//   - version string
func (_e *CloudProviderIPRangeManagerMock_Expecter) IsNewVersion(ipRangeData interface{}, version interface{}) *CloudProviderIPRangeManagerMock_IsNewVersion_Call {
	return &CloudProviderIPRangeManagerMock_IsNewVersion_Call{Call: _e.mock.On("IsNewVersion", ipRangeData, version)}
}

func (_c *CloudProviderIPRangeManagerMock_IsNewVersion_Call) Run(run func(ipRangeData []byte, version string)) *CloudProviderIPRangeManagerMock_IsNewVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(string))
	})
	return _c
}

func (_c *CloudProviderIPRangeManagerMock_IsNewVersion_Call) Return(_a0 bool, _a1 error) *CloudProviderIPRangeManagerMock_IsNewVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloudProviderIPRangeManagerMock_IsNewVersion_Call) RunAndReturn(run func([]byte, string) (bool, error)) *CloudProviderIPRangeManagerMock_IsNewVersion_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCloudProviderIPRangeManagerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewCloudProviderIPRangeManagerMock creates a new instance of CloudProviderIPRangeManagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCloudProviderIPRangeManagerMock(t mockConstructorTestingTNewCloudProviderIPRangeManagerMock) *CloudProviderIPRangeManagerMock {
	mock := &CloudProviderIPRangeManagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}