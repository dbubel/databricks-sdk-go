// Code generated by mockery v2.53.2. DO NOT EDIT.

package compute

import (
	context "context"

	compute "github.com/databricks/databricks-sdk-go/service/compute"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockPolicyComplianceForClustersInterface is an autogenerated mock type for the PolicyComplianceForClustersInterface type
type MockPolicyComplianceForClustersInterface struct {
	mock.Mock
}

type MockPolicyComplianceForClustersInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPolicyComplianceForClustersInterface) EXPECT() *MockPolicyComplianceForClustersInterface_Expecter {
	return &MockPolicyComplianceForClustersInterface_Expecter{mock: &_m.Mock}
}

// EnforceCompliance provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForClustersInterface) EnforceCompliance(ctx context.Context, request compute.EnforceClusterComplianceRequest) (*compute.EnforceClusterComplianceResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for EnforceCompliance")
	}

	var r0 *compute.EnforceClusterComplianceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.EnforceClusterComplianceRequest) (*compute.EnforceClusterComplianceResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, compute.EnforceClusterComplianceRequest) *compute.EnforceClusterComplianceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.EnforceClusterComplianceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, compute.EnforceClusterComplianceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForClustersInterface_EnforceCompliance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnforceCompliance'
type MockPolicyComplianceForClustersInterface_EnforceCompliance_Call struct {
	*mock.Call
}

// EnforceCompliance is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.EnforceClusterComplianceRequest
func (_e *MockPolicyComplianceForClustersInterface_Expecter) EnforceCompliance(ctx interface{}, request interface{}) *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call {
	return &MockPolicyComplianceForClustersInterface_EnforceCompliance_Call{Call: _e.mock.On("EnforceCompliance", ctx, request)}
}

func (_c *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call) Run(run func(ctx context.Context, request compute.EnforceClusterComplianceRequest)) *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.EnforceClusterComplianceRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call) Return(_a0 *compute.EnforceClusterComplianceResponse, _a1 error) *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call) RunAndReturn(run func(context.Context, compute.EnforceClusterComplianceRequest) (*compute.EnforceClusterComplianceResponse, error)) *MockPolicyComplianceForClustersInterface_EnforceCompliance_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompliance provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForClustersInterface) GetCompliance(ctx context.Context, request compute.GetClusterComplianceRequest) (*compute.GetClusterComplianceResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetCompliance")
	}

	var r0 *compute.GetClusterComplianceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.GetClusterComplianceRequest) (*compute.GetClusterComplianceResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, compute.GetClusterComplianceRequest) *compute.GetClusterComplianceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.GetClusterComplianceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, compute.GetClusterComplianceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForClustersInterface_GetCompliance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompliance'
type MockPolicyComplianceForClustersInterface_GetCompliance_Call struct {
	*mock.Call
}

// GetCompliance is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.GetClusterComplianceRequest
func (_e *MockPolicyComplianceForClustersInterface_Expecter) GetCompliance(ctx interface{}, request interface{}) *MockPolicyComplianceForClustersInterface_GetCompliance_Call {
	return &MockPolicyComplianceForClustersInterface_GetCompliance_Call{Call: _e.mock.On("GetCompliance", ctx, request)}
}

func (_c *MockPolicyComplianceForClustersInterface_GetCompliance_Call) Run(run func(ctx context.Context, request compute.GetClusterComplianceRequest)) *MockPolicyComplianceForClustersInterface_GetCompliance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.GetClusterComplianceRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_GetCompliance_Call) Return(_a0 *compute.GetClusterComplianceResponse, _a1 error) *MockPolicyComplianceForClustersInterface_GetCompliance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_GetCompliance_Call) RunAndReturn(run func(context.Context, compute.GetClusterComplianceRequest) (*compute.GetClusterComplianceResponse, error)) *MockPolicyComplianceForClustersInterface_GetCompliance_Call {
	_c.Call.Return(run)
	return _c
}

// GetComplianceByClusterId provides a mock function with given fields: ctx, clusterId
func (_m *MockPolicyComplianceForClustersInterface) GetComplianceByClusterId(ctx context.Context, clusterId string) (*compute.GetClusterComplianceResponse, error) {
	ret := _m.Called(ctx, clusterId)

	if len(ret) == 0 {
		panic("no return value specified for GetComplianceByClusterId")
	}

	var r0 *compute.GetClusterComplianceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*compute.GetClusterComplianceResponse, error)); ok {
		return rf(ctx, clusterId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *compute.GetClusterComplianceResponse); ok {
		r0 = rf(ctx, clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.GetClusterComplianceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComplianceByClusterId'
type MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call struct {
	*mock.Call
}

// GetComplianceByClusterId is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterId string
func (_e *MockPolicyComplianceForClustersInterface_Expecter) GetComplianceByClusterId(ctx interface{}, clusterId interface{}) *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call {
	return &MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call{Call: _e.mock.On("GetComplianceByClusterId", ctx, clusterId)}
}

func (_c *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call) Run(run func(ctx context.Context, clusterId string)) *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call) Return(_a0 *compute.GetClusterComplianceResponse, _a1 error) *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call) RunAndReturn(run func(context.Context, string) (*compute.GetClusterComplianceResponse, error)) *MockPolicyComplianceForClustersInterface_GetComplianceByClusterId_Call {
	_c.Call.Return(run)
	return _c
}

// ListCompliance provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForClustersInterface) ListCompliance(ctx context.Context, request compute.ListClusterCompliancesRequest) listing.Iterator[compute.ClusterCompliance] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListCompliance")
	}

	var r0 listing.Iterator[compute.ClusterCompliance]
	if rf, ok := ret.Get(0).(func(context.Context, compute.ListClusterCompliancesRequest) listing.Iterator[compute.ClusterCompliance]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[compute.ClusterCompliance])
		}
	}

	return r0
}

// MockPolicyComplianceForClustersInterface_ListCompliance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCompliance'
type MockPolicyComplianceForClustersInterface_ListCompliance_Call struct {
	*mock.Call
}

// ListCompliance is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.ListClusterCompliancesRequest
func (_e *MockPolicyComplianceForClustersInterface_Expecter) ListCompliance(ctx interface{}, request interface{}) *MockPolicyComplianceForClustersInterface_ListCompliance_Call {
	return &MockPolicyComplianceForClustersInterface_ListCompliance_Call{Call: _e.mock.On("ListCompliance", ctx, request)}
}

func (_c *MockPolicyComplianceForClustersInterface_ListCompliance_Call) Run(run func(ctx context.Context, request compute.ListClusterCompliancesRequest)) *MockPolicyComplianceForClustersInterface_ListCompliance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.ListClusterCompliancesRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_ListCompliance_Call) Return(_a0 listing.Iterator[compute.ClusterCompliance]) *MockPolicyComplianceForClustersInterface_ListCompliance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_ListCompliance_Call) RunAndReturn(run func(context.Context, compute.ListClusterCompliancesRequest) listing.Iterator[compute.ClusterCompliance]) *MockPolicyComplianceForClustersInterface_ListCompliance_Call {
	_c.Call.Return(run)
	return _c
}

// ListComplianceAll provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForClustersInterface) ListComplianceAll(ctx context.Context, request compute.ListClusterCompliancesRequest) ([]compute.ClusterCompliance, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListComplianceAll")
	}

	var r0 []compute.ClusterCompliance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.ListClusterCompliancesRequest) ([]compute.ClusterCompliance, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, compute.ListClusterCompliancesRequest) []compute.ClusterCompliance); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]compute.ClusterCompliance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, compute.ListClusterCompliancesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForClustersInterface_ListComplianceAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListComplianceAll'
type MockPolicyComplianceForClustersInterface_ListComplianceAll_Call struct {
	*mock.Call
}

// ListComplianceAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.ListClusterCompliancesRequest
func (_e *MockPolicyComplianceForClustersInterface_Expecter) ListComplianceAll(ctx interface{}, request interface{}) *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call {
	return &MockPolicyComplianceForClustersInterface_ListComplianceAll_Call{Call: _e.mock.On("ListComplianceAll", ctx, request)}
}

func (_c *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call) Run(run func(ctx context.Context, request compute.ListClusterCompliancesRequest)) *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.ListClusterCompliancesRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call) Return(_a0 []compute.ClusterCompliance, _a1 error) *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call) RunAndReturn(run func(context.Context, compute.ListClusterCompliancesRequest) ([]compute.ClusterCompliance, error)) *MockPolicyComplianceForClustersInterface_ListComplianceAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPolicyComplianceForClustersInterface creates a new instance of MockPolicyComplianceForClustersInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPolicyComplianceForClustersInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPolicyComplianceForClustersInterface {
	mock := &MockPolicyComplianceForClustersInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
