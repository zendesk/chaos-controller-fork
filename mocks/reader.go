// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package mocks

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"
)

// ReaderMock is an autogenerated mock type for the Reader type
type ReaderMock struct {
	mock.Mock
}

type ReaderMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ReaderMock) EXPECT() *ReaderMock_Expecter {
	return &ReaderMock_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, key, obj, opts
func (_m *ReaderMock) Get(ctx context.Context, key types.NamespacedName, obj client.Object, opts ...client.GetOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, obj)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.NamespacedName, client.Object, ...client.GetOption) error); ok {
		r0 = rf(ctx, key, obj, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReaderMock_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ReaderMock_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key types.NamespacedName
//   - obj client.Object
//   - opts ...client.GetOption
func (_e *ReaderMock_Expecter) Get(ctx interface{}, key interface{}, obj interface{}, opts ...interface{}) *ReaderMock_Get_Call {
	return &ReaderMock_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, key, obj}, opts...)...)}
}

func (_c *ReaderMock_Get_Call) Run(run func(ctx context.Context, key types.NamespacedName, obj client.Object, opts ...client.GetOption)) *ReaderMock_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]client.GetOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(client.GetOption)
			}
		}
		run(args[0].(context.Context), args[1].(types.NamespacedName), args[2].(client.Object), variadicArgs...)
	})
	return _c
}

func (_c *ReaderMock_Get_Call) Return(_a0 error) *ReaderMock_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReaderMock_Get_Call) RunAndReturn(run func(context.Context, types.NamespacedName, client.Object, ...client.GetOption) error) *ReaderMock_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, list, opts
func (_m *ReaderMock) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, list)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ObjectList, ...client.ListOption) error); ok {
		r0 = rf(ctx, list, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReaderMock_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ReaderMock_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - list client.ObjectList
//   - opts ...client.ListOption
func (_e *ReaderMock_Expecter) List(ctx interface{}, list interface{}, opts ...interface{}) *ReaderMock_List_Call {
	return &ReaderMock_List_Call{Call: _e.mock.On("List",
		append([]interface{}{ctx, list}, opts...)...)}
}

func (_c *ReaderMock_List_Call) Run(run func(ctx context.Context, list client.ObjectList, opts ...client.ListOption)) *ReaderMock_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]client.ListOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(client.ListOption)
			}
		}
		run(args[0].(context.Context), args[1].(client.ObjectList), variadicArgs...)
	})
	return _c
}

func (_c *ReaderMock_List_Call) Return(_a0 error) *ReaderMock_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReaderMock_List_Call) RunAndReturn(run func(context.Context, client.ObjectList, ...client.ListOption) error) *ReaderMock_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewReaderMock creates a new instance of ReaderMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReaderMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReaderMock {
	mock := &ReaderMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
