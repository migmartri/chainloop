// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	anypb "google.golang.org/protobuf/types/known/anypb"

	integrations "github.com/chainloop-dev/chainloop/app/controlplane/integrations"

	mock "github.com/stretchr/testify/mock"
)

// FanOut is an autogenerated mock type for the FanOut type
type FanOut struct {
	mock.Mock
}

// Describe provides a mock function with given fields:
func (_m *FanOut) Describe() *integrations.IntegrationInfo {
	ret := _m.Called()

	var r0 *integrations.IntegrationInfo
	if rf, ok := ret.Get(0).(func() *integrations.IntegrationInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*integrations.IntegrationInfo)
		}
	}

	return r0
}

// Execute provides a mock function with given fields: ctx, opts
func (_m *FanOut) Execute(ctx context.Context, opts *integrations.ExecuteReq) error {
	ret := _m.Called(ctx, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *integrations.ExecuteReq) error); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreAttach provides a mock function with given fields: ctx, c
func (_m *FanOut) PreAttach(ctx context.Context, c *integrations.BundledConfig) (*integrations.PreAttachment, error) {
	ret := _m.Called(ctx, c)

	var r0 *integrations.PreAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *integrations.BundledConfig) (*integrations.PreAttachment, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *integrations.BundledConfig) *integrations.PreAttachment); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*integrations.PreAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *integrations.BundledConfig) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PreRegister provides a mock function with given fields: ctx, req
func (_m *FanOut) PreRegister(ctx context.Context, req *anypb.Any) (*integrations.PreRegistration, error) {
	ret := _m.Called(ctx, req)

	var r0 *integrations.PreRegistration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *anypb.Any) (*integrations.PreRegistration, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *anypb.Any) *integrations.PreRegistration); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*integrations.PreRegistration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *anypb.Any) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// String provides a mock function with given fields:
func (_m *FanOut) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewFanOut interface {
	mock.TestingT
	Cleanup(func())
}

// NewFanOut creates a new instance of FanOut. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFanOut(t mockConstructorTestingTNewFanOut) *FanOut {
	mock := &FanOut{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}