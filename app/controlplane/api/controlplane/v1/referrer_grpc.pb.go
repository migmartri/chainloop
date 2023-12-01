//
// Copyright 2023 The Chainloop Authors.
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: controlplane/v1/referrer.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ReferrerService_DiscoverPrivate_FullMethodName      = "/controlplane.v1.ReferrerService/DiscoverPrivate"
	ReferrerService_DiscoverPublicShared_FullMethodName = "/controlplane.v1.ReferrerService/DiscoverPublicShared"
)

// ReferrerServiceClient is the client API for ReferrerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReferrerServiceClient interface {
	// DiscoverPrivate returns the referrer item for a given digest in the organizations of the logged-in user
	DiscoverPrivate(ctx context.Context, in *ReferrerServiceDiscoverPrivateRequest, opts ...grpc.CallOption) (*ReferrerServiceDiscoverPrivateResponse, error)
	// DiscoverPublicShared returns the referrer item for a given digest in the public shared index
	DiscoverPublicShared(ctx context.Context, in *DiscoverPublicSharedRequest, opts ...grpc.CallOption) (*DiscoverPublicSharedResponse, error)
}

type referrerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReferrerServiceClient(cc grpc.ClientConnInterface) ReferrerServiceClient {
	return &referrerServiceClient{cc}
}

func (c *referrerServiceClient) DiscoverPrivate(ctx context.Context, in *ReferrerServiceDiscoverPrivateRequest, opts ...grpc.CallOption) (*ReferrerServiceDiscoverPrivateResponse, error) {
	out := new(ReferrerServiceDiscoverPrivateResponse)
	err := c.cc.Invoke(ctx, ReferrerService_DiscoverPrivate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referrerServiceClient) DiscoverPublicShared(ctx context.Context, in *DiscoverPublicSharedRequest, opts ...grpc.CallOption) (*DiscoverPublicSharedResponse, error) {
	out := new(DiscoverPublicSharedResponse)
	err := c.cc.Invoke(ctx, ReferrerService_DiscoverPublicShared_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReferrerServiceServer is the server API for ReferrerService service.
// All implementations must embed UnimplementedReferrerServiceServer
// for forward compatibility
type ReferrerServiceServer interface {
	// DiscoverPrivate returns the referrer item for a given digest in the organizations of the logged-in user
	DiscoverPrivate(context.Context, *ReferrerServiceDiscoverPrivateRequest) (*ReferrerServiceDiscoverPrivateResponse, error)
	// DiscoverPublicShared returns the referrer item for a given digest in the public shared index
	DiscoverPublicShared(context.Context, *DiscoverPublicSharedRequest) (*DiscoverPublicSharedResponse, error)
	mustEmbedUnimplementedReferrerServiceServer()
}

// UnimplementedReferrerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReferrerServiceServer struct {
}

func (UnimplementedReferrerServiceServer) DiscoverPrivate(context.Context, *ReferrerServiceDiscoverPrivateRequest) (*ReferrerServiceDiscoverPrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverPrivate not implemented")
}
func (UnimplementedReferrerServiceServer) DiscoverPublicShared(context.Context, *DiscoverPublicSharedRequest) (*DiscoverPublicSharedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverPublicShared not implemented")
}
func (UnimplementedReferrerServiceServer) mustEmbedUnimplementedReferrerServiceServer() {}

// UnsafeReferrerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReferrerServiceServer will
// result in compilation errors.
type UnsafeReferrerServiceServer interface {
	mustEmbedUnimplementedReferrerServiceServer()
}

func RegisterReferrerServiceServer(s grpc.ServiceRegistrar, srv ReferrerServiceServer) {
	s.RegisterService(&ReferrerService_ServiceDesc, srv)
}

func _ReferrerService_DiscoverPrivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferrerServiceDiscoverPrivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferrerServiceServer).DiscoverPrivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReferrerService_DiscoverPrivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferrerServiceServer).DiscoverPrivate(ctx, req.(*ReferrerServiceDiscoverPrivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferrerService_DiscoverPublicShared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverPublicSharedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferrerServiceServer).DiscoverPublicShared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReferrerService_DiscoverPublicShared_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferrerServiceServer).DiscoverPublicShared(ctx, req.(*DiscoverPublicSharedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReferrerService_ServiceDesc is the grpc.ServiceDesc for ReferrerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReferrerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.v1.ReferrerService",
	HandlerType: (*ReferrerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoverPrivate",
			Handler:    _ReferrerService_DiscoverPrivate_Handler,
		},
		{
			MethodName: "DiscoverPublicShared",
			Handler:    _ReferrerService_DiscoverPublicShared_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/v1/referrer.proto",
}