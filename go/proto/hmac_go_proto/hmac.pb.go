// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: third_party/tink/proto/hmac.proto

package hmac_go_proto

import (
	common_go_proto "github.com/google/tink/go/proto/common_go_proto"
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

type HmacParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash    common_go_proto.HashType `protobuf:"varint,1,opt,name=hash,proto3,enum=google.crypto.tink.HashType" json:"hash,omitempty"` // HashType is an enum.
	TagSize uint32                   `protobuf:"varint,2,opt,name=tag_size,json=tagSize,proto3" json:"tag_size,omitempty"`
}

func (x *HmacParams) Reset() {
	*x = HmacParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hmac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HmacParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HmacParams) ProtoMessage() {}

func (x *HmacParams) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hmac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HmacParams.ProtoReflect.Descriptor instead.
func (*HmacParams) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hmac_proto_rawDescGZIP(), []int{0}
}

func (x *HmacParams) GetHash() common_go_proto.HashType {
	if x != nil {
		return x.Hash
	}
	return common_go_proto.HashType(0)
}

func (x *HmacParams) GetTagSize() uint32 {
	if x != nil {
		return x.TagSize
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.HmacKey
type HmacKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint32      `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params   *HmacParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue []byte      `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (x *HmacKey) Reset() {
	*x = HmacKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hmac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HmacKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HmacKey) ProtoMessage() {}

func (x *HmacKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hmac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HmacKey.ProtoReflect.Descriptor instead.
func (*HmacKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hmac_proto_rawDescGZIP(), []int{1}
}

func (x *HmacKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *HmacKey) GetParams() *HmacParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *HmacKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

type HmacKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params  *HmacParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize uint32      `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	Version uint32      `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *HmacKeyFormat) Reset() {
	*x = HmacKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hmac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HmacKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HmacKeyFormat) ProtoMessage() {}

func (x *HmacKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hmac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HmacKeyFormat.ProtoReflect.Descriptor instead.
func (*HmacKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hmac_proto_rawDescGZIP(), []int{2}
}

func (x *HmacKeyFormat) GetParams() *HmacParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *HmacKeyFormat) GetKeySize() uint32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

func (x *HmacKeyFormat) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_third_party_tink_proto_hmac_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_hmac_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6d, 0x61, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x1a, 0x23, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0a,
	0x48, 0x6d, 0x61, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x61, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x74, 0x61, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x78, 0x0a, 0x07, 0x48, 0x6d, 0x61, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x2e, 0x48, 0x6d, 0x61, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x7c, 0x0a, 0x0d, 0x48, 0x6d, 0x61, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x6d, 0x61, 0x63, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x4c, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x6d, 0x61, 0x63, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_hmac_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_hmac_proto_rawDescData = file_third_party_tink_proto_hmac_proto_rawDesc
)

func file_third_party_tink_proto_hmac_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_hmac_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_hmac_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_hmac_proto_rawDescData)
	})
	return file_third_party_tink_proto_hmac_proto_rawDescData
}

var file_third_party_tink_proto_hmac_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_third_party_tink_proto_hmac_proto_goTypes = []interface{}{
	(*HmacParams)(nil),            // 0: google.crypto.tink.HmacParams
	(*HmacKey)(nil),               // 1: google.crypto.tink.HmacKey
	(*HmacKeyFormat)(nil),         // 2: google.crypto.tink.HmacKeyFormat
	(common_go_proto.HashType)(0), // 3: google.crypto.tink.HashType
}
var file_third_party_tink_proto_hmac_proto_depIdxs = []int32{
	3, // 0: google.crypto.tink.HmacParams.hash:type_name -> google.crypto.tink.HashType
	0, // 1: google.crypto.tink.HmacKey.params:type_name -> google.crypto.tink.HmacParams
	0, // 2: google.crypto.tink.HmacKeyFormat.params:type_name -> google.crypto.tink.HmacParams
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_hmac_proto_init() }
func file_third_party_tink_proto_hmac_proto_init() {
	if File_third_party_tink_proto_hmac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_hmac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HmacParams); i {
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
		file_third_party_tink_proto_hmac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HmacKey); i {
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
		file_third_party_tink_proto_hmac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HmacKeyFormat); i {
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
			RawDescriptor: file_third_party_tink_proto_hmac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_hmac_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_hmac_proto_depIdxs,
		MessageInfos:      file_third_party_tink_proto_hmac_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_hmac_proto = out.File
	file_third_party_tink_proto_hmac_proto_rawDesc = nil
	file_third_party_tink_proto_hmac_proto_goTypes = nil
	file_third_party_tink_proto_hmac_proto_depIdxs = nil
}
