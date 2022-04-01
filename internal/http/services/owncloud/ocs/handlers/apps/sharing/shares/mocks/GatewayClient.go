// Copyright 2018-2022 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	collaborationv1beta1 "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"

	gatewayv1beta1 "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"

	groupv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/group/v1beta1"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"

	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
)

// GatewayClient is an autogenerated mock type for the GatewayClient type
type GatewayClient struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) Authenticate(ctx context.Context, in *gatewayv1beta1.AuthenticateRequest, opts ...grpc.CallOption) (*gatewayv1beta1.AuthenticateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gatewayv1beta1.AuthenticateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gatewayv1beta1.AuthenticateRequest, ...grpc.CallOption) *gatewayv1beta1.AuthenticateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gatewayv1beta1.AuthenticateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gatewayv1beta1.AuthenticateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateShare provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) CreateShare(ctx context.Context, in *collaborationv1beta1.CreateShareRequest, opts ...grpc.CallOption) (*collaborationv1beta1.CreateShareResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.CreateShareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.CreateShareRequest, ...grpc.CallOption) *collaborationv1beta1.CreateShareResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.CreateShareResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.CreateShareRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroup provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetGroup(ctx context.Context, in *groupv1beta1.GetGroupRequest, opts ...grpc.CallOption) (*groupv1beta1.GetGroupResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *groupv1beta1.GetGroupResponse
	if rf, ok := ret.Get(0).(func(context.Context, *groupv1beta1.GetGroupRequest, ...grpc.CallOption) *groupv1beta1.GetGroupResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*groupv1beta1.GetGroupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *groupv1beta1.GetGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupByClaim provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetGroupByClaim(ctx context.Context, in *groupv1beta1.GetGroupByClaimRequest, opts ...grpc.CallOption) (*groupv1beta1.GetGroupByClaimResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *groupv1beta1.GetGroupByClaimResponse
	if rf, ok := ret.Get(0).(func(context.Context, *groupv1beta1.GetGroupByClaimRequest, ...grpc.CallOption) *groupv1beta1.GetGroupByClaimResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*groupv1beta1.GetGroupByClaimResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *groupv1beta1.GetGroupByClaimRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPath provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetPath(ctx context.Context, in *providerv1beta1.GetPathRequest, opts ...grpc.CallOption) (*providerv1beta1.GetPathResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *providerv1beta1.GetPathResponse
	if rf, ok := ret.Get(0).(func(context.Context, *providerv1beta1.GetPathRequest, ...grpc.CallOption) *providerv1beta1.GetPathResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*providerv1beta1.GetPathResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *providerv1beta1.GetPathRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReceivedShare provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetReceivedShare(ctx context.Context, in *collaborationv1beta1.GetReceivedShareRequest, opts ...grpc.CallOption) (*collaborationv1beta1.GetReceivedShareResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.GetReceivedShareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.GetReceivedShareRequest, ...grpc.CallOption) *collaborationv1beta1.GetReceivedShareResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.GetReceivedShareResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.GetReceivedShareRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShare provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetShare(ctx context.Context, in *collaborationv1beta1.GetShareRequest, opts ...grpc.CallOption) (*collaborationv1beta1.GetShareResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.GetShareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.GetShareRequest, ...grpc.CallOption) *collaborationv1beta1.GetShareResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.GetShareResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.GetShareRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetUser(ctx context.Context, in *userv1beta1.GetUserRequest, opts ...grpc.CallOption) (*userv1beta1.GetUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *userv1beta1.GetUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userv1beta1.GetUserRequest, ...grpc.CallOption) *userv1beta1.GetUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.GetUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userv1beta1.GetUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByClaim provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) GetUserByClaim(ctx context.Context, in *userv1beta1.GetUserByClaimRequest, opts ...grpc.CallOption) (*userv1beta1.GetUserByClaimResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *userv1beta1.GetUserByClaimResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userv1beta1.GetUserByClaimRequest, ...grpc.CallOption) *userv1beta1.GetUserByClaimResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.GetUserByClaimResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userv1beta1.GetUserByClaimRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListContainer provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) ListContainer(ctx context.Context, in *providerv1beta1.ListContainerRequest, opts ...grpc.CallOption) (*providerv1beta1.ListContainerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *providerv1beta1.ListContainerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *providerv1beta1.ListContainerRequest, ...grpc.CallOption) *providerv1beta1.ListContainerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*providerv1beta1.ListContainerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *providerv1beta1.ListContainerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReceivedShares provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) ListReceivedShares(ctx context.Context, in *collaborationv1beta1.ListReceivedSharesRequest, opts ...grpc.CallOption) (*collaborationv1beta1.ListReceivedSharesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.ListReceivedSharesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ListReceivedSharesRequest, ...grpc.CallOption) *collaborationv1beta1.ListReceivedSharesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.ListReceivedSharesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.ListReceivedSharesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListShares provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) ListShares(ctx context.Context, in *collaborationv1beta1.ListSharesRequest, opts ...grpc.CallOption) (*collaborationv1beta1.ListSharesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.ListSharesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.ListSharesRequest, ...grpc.CallOption) *collaborationv1beta1.ListSharesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.ListSharesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.ListSharesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveShare provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) RemoveShare(ctx context.Context, in *collaborationv1beta1.RemoveShareRequest, opts ...grpc.CallOption) (*collaborationv1beta1.RemoveShareResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.RemoveShareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.RemoveShareRequest, ...grpc.CallOption) *collaborationv1beta1.RemoveShareResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.RemoveShareResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.RemoveShareRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) Stat(ctx context.Context, in *providerv1beta1.StatRequest, opts ...grpc.CallOption) (*providerv1beta1.StatResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *providerv1beta1.StatResponse
	if rf, ok := ret.Get(0).(func(context.Context, *providerv1beta1.StatRequest, ...grpc.CallOption) *providerv1beta1.StatResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*providerv1beta1.StatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *providerv1beta1.StatRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReceivedShare provides a mock function with given fields: ctx, in, opts
func (_m *GatewayClient) UpdateReceivedShare(ctx context.Context, in *collaborationv1beta1.UpdateReceivedShareRequest, opts ...grpc.CallOption) (*collaborationv1beta1.UpdateReceivedShareResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *collaborationv1beta1.UpdateReceivedShareResponse
	if rf, ok := ret.Get(0).(func(context.Context, *collaborationv1beta1.UpdateReceivedShareRequest, ...grpc.CallOption) *collaborationv1beta1.UpdateReceivedShareResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collaborationv1beta1.UpdateReceivedShareResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *collaborationv1beta1.UpdateReceivedShareRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}