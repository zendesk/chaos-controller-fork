// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024 Datadog, Inc.
package ebpf

import mock "github.com/stretchr/testify/mock"

// ConfigInformerMock is an autogenerated mock type for the ConfigInformer type
type ConfigInformerMock struct {
	mock.Mock
}

type ConfigInformerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ConfigInformerMock) EXPECT() *ConfigInformerMock_Expecter {
	return &ConfigInformerMock_Expecter{mock: &_m.Mock}
}

// GetKernelFeatures provides a mock function with given fields:
func (_m *ConfigInformerMock) GetKernelFeatures() (Features, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetKernelFeatures")
	}

	var r0 Features
	var r1 error
	if rf, ok := ret.Get(0).(func() (Features, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() Features); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Features)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigInformerMock_GetKernelFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKernelFeatures'
type ConfigInformerMock_GetKernelFeatures_Call struct {
	*mock.Call
}

// GetKernelFeatures is a helper method to define mock.On call
func (_e *ConfigInformerMock_Expecter) GetKernelFeatures() *ConfigInformerMock_GetKernelFeatures_Call {
	return &ConfigInformerMock_GetKernelFeatures_Call{Call: _e.mock.On("GetKernelFeatures")}
}

func (_c *ConfigInformerMock_GetKernelFeatures_Call) Run(run func()) *ConfigInformerMock_GetKernelFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigInformerMock_GetKernelFeatures_Call) Return(_a0 Features, _a1 error) *ConfigInformerMock_GetKernelFeatures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConfigInformerMock_GetKernelFeatures_Call) RunAndReturn(run func() (Features, error)) *ConfigInformerMock_GetKernelFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// GetMapTypes provides a mock function with given fields:
func (_m *ConfigInformerMock) GetMapTypes() MapTypes {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMapTypes")
	}

	var r0 MapTypes
	if rf, ok := ret.Get(0).(func() MapTypes); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(MapTypes)
	}

	return r0
}

// ConfigInformerMock_GetMapTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMapTypes'
type ConfigInformerMock_GetMapTypes_Call struct {
	*mock.Call
}

// GetMapTypes is a helper method to define mock.On call
func (_e *ConfigInformerMock_Expecter) GetMapTypes() *ConfigInformerMock_GetMapTypes_Call {
	return &ConfigInformerMock_GetMapTypes_Call{Call: _e.mock.On("GetMapTypes")}
}

func (_c *ConfigInformerMock_GetMapTypes_Call) Run(run func()) *ConfigInformerMock_GetMapTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigInformerMock_GetMapTypes_Call) Return(_a0 MapTypes) *ConfigInformerMock_GetMapTypes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigInformerMock_GetMapTypes_Call) RunAndReturn(run func() MapTypes) *ConfigInformerMock_GetMapTypes_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequiredSystemConfig provides a mock function with given fields:
func (_m *ConfigInformerMock) GetRequiredSystemConfig() KernelParams {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRequiredSystemConfig")
	}

	var r0 KernelParams
	if rf, ok := ret.Get(0).(func() KernelParams); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(KernelParams)
		}
	}

	return r0
}

// ConfigInformerMock_GetRequiredSystemConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequiredSystemConfig'
type ConfigInformerMock_GetRequiredSystemConfig_Call struct {
	*mock.Call
}

// GetRequiredSystemConfig is a helper method to define mock.On call
func (_e *ConfigInformerMock_Expecter) GetRequiredSystemConfig() *ConfigInformerMock_GetRequiredSystemConfig_Call {
	return &ConfigInformerMock_GetRequiredSystemConfig_Call{Call: _e.mock.On("GetRequiredSystemConfig")}
}

func (_c *ConfigInformerMock_GetRequiredSystemConfig_Call) Run(run func()) *ConfigInformerMock_GetRequiredSystemConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigInformerMock_GetRequiredSystemConfig_Call) Return(_a0 KernelParams) *ConfigInformerMock_GetRequiredSystemConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigInformerMock_GetRequiredSystemConfig_Call) RunAndReturn(run func() KernelParams) *ConfigInformerMock_GetRequiredSystemConfig_Call {
	_c.Call.Return(run)
	return _c
}

// IsKernelConfigAvailable provides a mock function with given fields:
func (_m *ConfigInformerMock) IsKernelConfigAvailable() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsKernelConfigAvailable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ConfigInformerMock_IsKernelConfigAvailable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsKernelConfigAvailable'
type ConfigInformerMock_IsKernelConfigAvailable_Call struct {
	*mock.Call
}

// IsKernelConfigAvailable is a helper method to define mock.On call
func (_e *ConfigInformerMock_Expecter) IsKernelConfigAvailable() *ConfigInformerMock_IsKernelConfigAvailable_Call {
	return &ConfigInformerMock_IsKernelConfigAvailable_Call{Call: _e.mock.On("IsKernelConfigAvailable")}
}

func (_c *ConfigInformerMock_IsKernelConfigAvailable_Call) Run(run func()) *ConfigInformerMock_IsKernelConfigAvailable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigInformerMock_IsKernelConfigAvailable_Call) Return(_a0 bool) *ConfigInformerMock_IsKernelConfigAvailable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigInformerMock_IsKernelConfigAvailable_Call) RunAndReturn(run func() bool) *ConfigInformerMock_IsKernelConfigAvailable_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateRequiredSystemConfig provides a mock function with given fields:
func (_m *ConfigInformerMock) ValidateRequiredSystemConfig() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ValidateRequiredSystemConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfigInformerMock_ValidateRequiredSystemConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateRequiredSystemConfig'
type ConfigInformerMock_ValidateRequiredSystemConfig_Call struct {
	*mock.Call
}

// ValidateRequiredSystemConfig is a helper method to define mock.On call
func (_e *ConfigInformerMock_Expecter) ValidateRequiredSystemConfig() *ConfigInformerMock_ValidateRequiredSystemConfig_Call {
	return &ConfigInformerMock_ValidateRequiredSystemConfig_Call{Call: _e.mock.On("ValidateRequiredSystemConfig")}
}

func (_c *ConfigInformerMock_ValidateRequiredSystemConfig_Call) Run(run func()) *ConfigInformerMock_ValidateRequiredSystemConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigInformerMock_ValidateRequiredSystemConfig_Call) Return(_a0 error) *ConfigInformerMock_ValidateRequiredSystemConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigInformerMock_ValidateRequiredSystemConfig_Call) RunAndReturn(run func() error) *ConfigInformerMock_ValidateRequiredSystemConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfigInformerMock creates a new instance of ConfigInformerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigInformerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigInformerMock {
	mock := &ConfigInformerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}