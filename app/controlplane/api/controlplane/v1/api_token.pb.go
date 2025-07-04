//
// Copyright 2024-2025 The Chainloop Authors.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: controlplane/v1/api_token.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type APITokenServiceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description *string              `protobuf:"bytes,1,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ExpiresIn   *durationpb.Duration `protobuf:"bytes,2,opt,name=expires_in,json=expiresIn,proto3,oneof" json:"expires_in,omitempty"`
}

func (x *APITokenServiceCreateRequest) Reset() {
	*x = APITokenServiceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceCreateRequest) ProtoMessage() {}

func (x *APITokenServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*APITokenServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{0}
}

func (x *APITokenServiceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *APITokenServiceCreateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *APITokenServiceCreateRequest) GetExpiresIn() *durationpb.Duration {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

type APITokenServiceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *APITokenServiceCreateResponse_APITokenFull `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *APITokenServiceCreateResponse) Reset() {
	*x = APITokenServiceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceCreateResponse) ProtoMessage() {}

func (x *APITokenServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*APITokenServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{1}
}

func (x *APITokenServiceCreateResponse) GetResult() *APITokenServiceCreateResponse_APITokenFull {
	if x != nil {
		return x.Result
	}
	return nil
}

type APITokenServiceRevokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *APITokenServiceRevokeRequest) Reset() {
	*x = APITokenServiceRevokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceRevokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceRevokeRequest) ProtoMessage() {}

func (x *APITokenServiceRevokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceRevokeRequest.ProtoReflect.Descriptor instead.
func (*APITokenServiceRevokeRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{2}
}

func (x *APITokenServiceRevokeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type APITokenServiceRevokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *APITokenServiceRevokeResponse) Reset() {
	*x = APITokenServiceRevokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceRevokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceRevokeResponse) ProtoMessage() {}

func (x *APITokenServiceRevokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceRevokeResponse.ProtoReflect.Descriptor instead.
func (*APITokenServiceRevokeResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{3}
}

type APITokenServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeRevoked bool `protobuf:"varint,1,opt,name=include_revoked,json=includeRevoked,proto3" json:"include_revoked,omitempty"`
}

func (x *APITokenServiceListRequest) Reset() {
	*x = APITokenServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceListRequest) ProtoMessage() {}

func (x *APITokenServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceListRequest.ProtoReflect.Descriptor instead.
func (*APITokenServiceListRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{4}
}

func (x *APITokenServiceListRequest) GetIncludeRevoked() bool {
	if x != nil {
		return x.IncludeRevoked
	}
	return false
}

type APITokenServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*APITokenItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *APITokenServiceListResponse) Reset() {
	*x = APITokenServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceListResponse) ProtoMessage() {}

func (x *APITokenServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceListResponse.ProtoReflect.Descriptor instead.
func (*APITokenServiceListResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{5}
}

func (x *APITokenServiceListResponse) GetResult() []*APITokenItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type APITokenServiceCreateResponse_APITokenFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *APITokenItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Jwt  string        `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *APITokenServiceCreateResponse_APITokenFull) Reset() {
	*x = APITokenServiceCreateResponse_APITokenFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_api_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APITokenServiceCreateResponse_APITokenFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APITokenServiceCreateResponse_APITokenFull) ProtoMessage() {}

func (x *APITokenServiceCreateResponse_APITokenFull) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_api_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APITokenServiceCreateResponse_APITokenFull.ProtoReflect.Descriptor instead.
func (*APITokenServiceCreateResponse_APITokenFull) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_api_token_proto_rawDescGZIP(), []int{1, 0}
}

func (x *APITokenServiceCreateResponse_APITokenFull) GetItem() *APITokenItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *APITokenServiceCreateResponse_APITokenFull) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

var File_controlplane_v1_api_token_proto protoreflect.FileDescriptor

var file_controlplane_v1_api_token_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x1c, 0x41, 0x50, 0x49,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x1d,
	0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x50,
	0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x1a, 0x53, 0x0a, 0x0c, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x75,
	0x6c, 0x6c, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x1c, 0x41, 0x50, 0x49, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x97, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x82, 0x01, 0xba, 0x48, 0x7f, 0xba, 0x01, 0x7c,
	0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x64, 0x6e, 0x73, 0x2d, 0x31, 0x31, 0x32, 0x33, 0x12,
	0x3a, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x6f, 0x6e,
	0x6c, 0x79, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x73, 0x2e, 0x1a, 0x2f, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x24, 0x27, 0x29, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x1a, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x22, 0x54, 0x0a, 0x1b, 0x41, 0x50,
	0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xc6, 0x02, 0x0a, 0x0f, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x67, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controlplane_v1_api_token_proto_rawDescOnce sync.Once
	file_controlplane_v1_api_token_proto_rawDescData = file_controlplane_v1_api_token_proto_rawDesc
)

func file_controlplane_v1_api_token_proto_rawDescGZIP() []byte {
	file_controlplane_v1_api_token_proto_rawDescOnce.Do(func() {
		file_controlplane_v1_api_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_controlplane_v1_api_token_proto_rawDescData)
	})
	return file_controlplane_v1_api_token_proto_rawDescData
}

var file_controlplane_v1_api_token_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_controlplane_v1_api_token_proto_goTypes = []interface{}{
	(*APITokenServiceCreateRequest)(nil),               // 0: controlplane.v1.APITokenServiceCreateRequest
	(*APITokenServiceCreateResponse)(nil),              // 1: controlplane.v1.APITokenServiceCreateResponse
	(*APITokenServiceRevokeRequest)(nil),               // 2: controlplane.v1.APITokenServiceRevokeRequest
	(*APITokenServiceRevokeResponse)(nil),              // 3: controlplane.v1.APITokenServiceRevokeResponse
	(*APITokenServiceListRequest)(nil),                 // 4: controlplane.v1.APITokenServiceListRequest
	(*APITokenServiceListResponse)(nil),                // 5: controlplane.v1.APITokenServiceListResponse
	(*APITokenServiceCreateResponse_APITokenFull)(nil), // 6: controlplane.v1.APITokenServiceCreateResponse.APITokenFull
	(*durationpb.Duration)(nil),                        // 7: google.protobuf.Duration
	(*APITokenItem)(nil),                               // 8: controlplane.v1.APITokenItem
}
var file_controlplane_v1_api_token_proto_depIdxs = []int32{
	7, // 0: controlplane.v1.APITokenServiceCreateRequest.expires_in:type_name -> google.protobuf.Duration
	6, // 1: controlplane.v1.APITokenServiceCreateResponse.result:type_name -> controlplane.v1.APITokenServiceCreateResponse.APITokenFull
	8, // 2: controlplane.v1.APITokenServiceListResponse.result:type_name -> controlplane.v1.APITokenItem
	8, // 3: controlplane.v1.APITokenServiceCreateResponse.APITokenFull.item:type_name -> controlplane.v1.APITokenItem
	0, // 4: controlplane.v1.APITokenService.Create:input_type -> controlplane.v1.APITokenServiceCreateRequest
	4, // 5: controlplane.v1.APITokenService.List:input_type -> controlplane.v1.APITokenServiceListRequest
	2, // 6: controlplane.v1.APITokenService.Revoke:input_type -> controlplane.v1.APITokenServiceRevokeRequest
	1, // 7: controlplane.v1.APITokenService.Create:output_type -> controlplane.v1.APITokenServiceCreateResponse
	5, // 8: controlplane.v1.APITokenService.List:output_type -> controlplane.v1.APITokenServiceListResponse
	3, // 9: controlplane.v1.APITokenService.Revoke:output_type -> controlplane.v1.APITokenServiceRevokeResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_controlplane_v1_api_token_proto_init() }
func file_controlplane_v1_api_token_proto_init() {
	if File_controlplane_v1_api_token_proto != nil {
		return
	}
	file_controlplane_v1_response_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controlplane_v1_api_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_api_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_api_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceRevokeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_api_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceRevokeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_api_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_api_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controlplane_v1_api_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APITokenServiceCreateResponse_APITokenFull); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_controlplane_v1_api_token_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controlplane_v1_api_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controlplane_v1_api_token_proto_goTypes,
		DependencyIndexes: file_controlplane_v1_api_token_proto_depIdxs,
		MessageInfos:      file_controlplane_v1_api_token_proto_msgTypes,
	}.Build()
	File_controlplane_v1_api_token_proto = out.File
	file_controlplane_v1_api_token_proto_rawDesc = nil
	file_controlplane_v1_api_token_proto_goTypes = nil
	file_controlplane_v1_api_token_proto_depIdxs = nil
}
