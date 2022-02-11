// Copyright 2020 Unistack LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: api/micro.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Defines the Micro configuration for a method. It contains a
// [MicroMethod][micro.api.MicroMethod]
type MicroMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timeout for handler
	Timeout int32 `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// ServerOptions contains options for server
	ServerOptions []string `protobuf:"bytes,2,rep,name=server_options,json=serverOptions,proto3" json:"server_options,omitempty"`
	// ClientOptions contains options for client
	ClientOptions []string `protobuf:"bytes,3,rep,name=client_options,json=clientOptions,proto3" json:"client_options,omitempty"`
	// Type rpc/pub/sub
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *MicroMethod) Reset() {
	*x = MicroMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroMethod) ProtoMessage() {}

func (x *MicroMethod) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroMethod.ProtoReflect.Descriptor instead.
func (*MicroMethod) Descriptor() ([]byte, []int) {
	return file_api_micro_proto_rawDescGZIP(), []int{0}
}

func (x *MicroMethod) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *MicroMethod) GetServerOptions() []string {
	if x != nil {
		return x.ServerOptions
	}
	return nil
}

func (x *MicroMethod) GetClientOptions() []string {
	if x != nil {
		return x.ClientOptions
	}
	return nil
}

func (x *MicroMethod) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Defines the Micro configuration for a method. It contains a
// [MicroService][micro.api.MicroService]
type MicroService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ServerWrappers contains wrappers for server
	ServerWrappers []string `protobuf:"bytes,1,rep,name=server_wrappers,json=serverWrappers,proto3" json:"server_wrappers,omitempty"`
	// ClientWrappers contains wrappers for client
	ClientWrappers []string `protobuf:"bytes,2,rep,name=client_wrappers,json=clientWrappers,proto3" json:"client_wrappers,omitempty"`
}

func (x *MicroService) Reset() {
	*x = MicroService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MicroService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroService) ProtoMessage() {}

func (x *MicroService) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroService.ProtoReflect.Descriptor instead.
func (*MicroService) Descriptor() ([]byte, []int) {
	return file_api_micro_proto_rawDescGZIP(), []int{1}
}

func (x *MicroService) GetServerWrappers() []string {
	if x != nil {
		return x.ServerWrappers
	}
	return nil
}

func (x *MicroService) GetClientWrappers() []string {
	if x != nil {
		return x.ClientWrappers
	}
	return nil
}

var File_api_micro_proto protoreflect.FileDescriptor

var file_api_micro_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x89, 0x01, 0x0a,
	0x0b, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x60, 0x0a, 0x0c, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x42, 0x4b, 0x0a, 0x09, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0a, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x04, 0x4d, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_micro_proto_rawDescOnce sync.Once
	file_api_micro_proto_rawDescData = file_api_micro_proto_rawDesc
)

func file_api_micro_proto_rawDescGZIP() []byte {
	file_api_micro_proto_rawDescOnce.Do(func() {
		file_api_micro_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_micro_proto_rawDescData)
	})
	return file_api_micro_proto_rawDescData
}

var file_api_micro_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_micro_proto_goTypes = []interface{}{
	(*MicroMethod)(nil),  // 0: micro.api.MicroMethod
	(*MicroService)(nil), // 1: micro.api.MicroService
}
var file_api_micro_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_micro_proto_init() }
func file_api_micro_proto_init() {
	if File_api_micro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_micro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroMethod); i {
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
		file_api_micro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MicroService); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_micro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_micro_proto_goTypes,
		DependencyIndexes: file_api_micro_proto_depIdxs,
		MessageInfos:      file_api_micro_proto_msgTypes,
	}.Build()
	File_api_micro_proto = out.File
	file_api_micro_proto_rawDesc = nil
	file_api_micro_proto_goTypes = nil
	file_api_micro_proto_depIdxs = nil
}