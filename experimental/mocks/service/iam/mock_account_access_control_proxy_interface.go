// Code generated by mockery v2.53.2. DO NOT EDIT.

package iam

import (
	context "context"

	iam "github.com/databricks/databricks-sdk-go/service/iam"
	mock "github.com/stretchr/testify/mock"
)

// MockAccountAccessControlProxyInterface is an autogenerated mock type for the AccountAccessControlProxyInterface type
type MockAccountAccessControlProxyInterface struct {
	mock.Mock
}

type MockAccountAccessControlProxyInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountAccessControlProxyInterface) EXPECT() *MockAccountAccessControlProxyInterface_Expecter {
	return &MockAccountAccessControlProxyInterface_Expecter{mock: &_m.Mock}
}

// GetAssignableRolesForResource provides a mock function with given fields: ctx, request
func (_m *MockAccountAccessControlProxyInterface) GetAssignableRolesForResource(ctx context.Context, request iam.GetAssignableRolesForResourceRequest) (*iam.GetAssignableRolesForResourceResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetAssignableRolesForResource")
	}

	var r0 *iam.GetAssignableRolesForResourceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetAssignableRolesForResourceRequest) (*iam.GetAssignableRolesForResourceResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetAssignableRolesForResourceRequest) *iam.GetAssignableRolesForResourceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.GetAssignableRolesForResourceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetAssignableRolesForResourceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssignableRolesForResource'
type MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call struct {
	*mock.Call
}

// GetAssignableRolesForResource is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetAssignableRolesForResourceRequest
func (_e *MockAccountAccessControlProxyInterface_Expecter) GetAssignableRolesForResource(ctx interface{}, request interface{}) *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call {
	return &MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call{Call: _e.mock.On("GetAssignableRolesForResource", ctx, request)}
}

func (_c *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call) Run(run func(ctx context.Context, request iam.GetAssignableRolesForResourceRequest)) *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetAssignableRolesForResourceRequest))
	})
	return _c
}

func (_c *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call) Return(_a0 *iam.GetAssignableRolesForResourceResponse, _a1 error) *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call) RunAndReturn(run func(context.Context, iam.GetAssignableRolesForResourceRequest) (*iam.GetAssignableRolesForResourceResponse, error)) *MockAccountAccessControlProxyInterface_GetAssignableRolesForResource_Call {
	_c.Call.Return(run)
	return _c
}

// GetRuleSet provides a mock function with given fields: ctx, request
func (_m *MockAccountAccessControlProxyInterface) GetRuleSet(ctx context.Context, request iam.GetRuleSetRequest) (*iam.RuleSetResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetRuleSet")
	}

	var r0 *iam.RuleSetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetRuleSetRequest) (*iam.RuleSetResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetRuleSetRequest) *iam.RuleSetResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.RuleSetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetRuleSetRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountAccessControlProxyInterface_GetRuleSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRuleSet'
type MockAccountAccessControlProxyInterface_GetRuleSet_Call struct {
	*mock.Call
}

// GetRuleSet is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetRuleSetRequest
func (_e *MockAccountAccessControlProxyInterface_Expecter) GetRuleSet(ctx interface{}, request interface{}) *MockAccountAccessControlProxyInterface_GetRuleSet_Call {
	return &MockAccountAccessControlProxyInterface_GetRuleSet_Call{Call: _e.mock.On("GetRuleSet", ctx, request)}
}

func (_c *MockAccountAccessControlProxyInterface_GetRuleSet_Call) Run(run func(ctx context.Context, request iam.GetRuleSetRequest)) *MockAccountAccessControlProxyInterface_GetRuleSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetRuleSetRequest))
	})
	return _c
}

func (_c *MockAccountAccessControlProxyInterface_GetRuleSet_Call) Return(_a0 *iam.RuleSetResponse, _a1 error) *MockAccountAccessControlProxyInterface_GetRuleSet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountAccessControlProxyInterface_GetRuleSet_Call) RunAndReturn(run func(context.Context, iam.GetRuleSetRequest) (*iam.RuleSetResponse, error)) *MockAccountAccessControlProxyInterface_GetRuleSet_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRuleSet provides a mock function with given fields: ctx, request
func (_m *MockAccountAccessControlProxyInterface) UpdateRuleSet(ctx context.Context, request iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRuleSet")
	}

	var r0 *iam.RuleSetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.UpdateRuleSetRequest) *iam.RuleSetResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.RuleSetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.UpdateRuleSetRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountAccessControlProxyInterface_UpdateRuleSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRuleSet'
type MockAccountAccessControlProxyInterface_UpdateRuleSet_Call struct {
	*mock.Call
}

// UpdateRuleSet is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.UpdateRuleSetRequest
func (_e *MockAccountAccessControlProxyInterface_Expecter) UpdateRuleSet(ctx interface{}, request interface{}) *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call {
	return &MockAccountAccessControlProxyInterface_UpdateRuleSet_Call{Call: _e.mock.On("UpdateRuleSet", ctx, request)}
}

func (_c *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call) Run(run func(ctx context.Context, request iam.UpdateRuleSetRequest)) *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.UpdateRuleSetRequest))
	})
	return _c
}

func (_c *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call) Return(_a0 *iam.RuleSetResponse, _a1 error) *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call) RunAndReturn(run func(context.Context, iam.UpdateRuleSetRequest) (*iam.RuleSetResponse, error)) *MockAccountAccessControlProxyInterface_UpdateRuleSet_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountAccessControlProxyInterface creates a new instance of MockAccountAccessControlProxyInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountAccessControlProxyInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountAccessControlProxyInterface {
	mock := &MockAccountAccessControlProxyInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
