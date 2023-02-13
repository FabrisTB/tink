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

// Definitions for Elliptic Curve Digital Signature Algorithm (ECDSA).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: third_party/tink/proto/ecdsa.proto

package ecdsa_go_proto

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

type EcdsaSignatureEncoding int32

const (
	EcdsaSignatureEncoding_UNKNOWN_ENCODING EcdsaSignatureEncoding = 0
	// The signature's format is r || s, where r and s are zero-padded and have
	// the same size in bytes as the order of the curve. For example, for NIST
	// P-256 curve, r and s are zero-padded to 32 bytes.
	EcdsaSignatureEncoding_IEEE_P1363 EcdsaSignatureEncoding = 1
	// The signature is encoded using ASN.1
	// (https://tools.ietf.org/html/rfc5480#appendix-A):
	//
	//	ECDSA-Sig-Value :: = SEQUENCE {
	//	 r INTEGER,
	//	 s INTEGER
	//	}
	EcdsaSignatureEncoding_DER EcdsaSignatureEncoding = 2
)

// Enum value maps for EcdsaSignatureEncoding.
var (
	EcdsaSignatureEncoding_name = map[int32]string{
		0: "UNKNOWN_ENCODING",
		1: "IEEE_P1363",
		2: "DER",
	}
	EcdsaSignatureEncoding_value = map[string]int32{
		"UNKNOWN_ENCODING": 0,
		"IEEE_P1363":       1,
		"DER":              2,
	}
)

func (x EcdsaSignatureEncoding) Enum() *EcdsaSignatureEncoding {
	p := new(EcdsaSignatureEncoding)
	*p = x
	return p
}

func (x EcdsaSignatureEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EcdsaSignatureEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_ecdsa_proto_enumTypes[0].Descriptor()
}

func (EcdsaSignatureEncoding) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_ecdsa_proto_enumTypes[0]
}

func (x EcdsaSignatureEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EcdsaSignatureEncoding.Descriptor instead.
func (EcdsaSignatureEncoding) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ecdsa_proto_rawDescGZIP(), []int{0}
}

// Protos for Ecdsa.
type EcdsaParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	HashType common_go_proto.HashType `protobuf:"varint,1,opt,name=hash_type,json=hashType,proto3,enum=google.crypto.tink.HashType" json:"hash_type,omitempty"`
	// Required.
	Curve common_go_proto.EllipticCurveType `protobuf:"varint,2,opt,name=curve,proto3,enum=google.crypto.tink.EllipticCurveType" json:"curve,omitempty"`
	// Required.
	Encoding EcdsaSignatureEncoding `protobuf:"varint,3,opt,name=encoding,proto3,enum=google.crypto.tink.EcdsaSignatureEncoding" json:"encoding,omitempty"`
}

func (x *EcdsaParams) Reset() {
	*x = EcdsaParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdsaParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdsaParams) ProtoMessage() {}

func (x *EcdsaParams) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdsaParams.ProtoReflect.Descriptor instead.
func (*EcdsaParams) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ecdsa_proto_rawDescGZIP(), []int{0}
}

func (x *EcdsaParams) GetHashType() common_go_proto.HashType {
	if x != nil {
		return x.HashType
	}
	return common_go_proto.HashType(0)
}

func (x *EcdsaParams) GetCurve() common_go_proto.EllipticCurveType {
	if x != nil {
		return x.Curve
	}
	return common_go_proto.EllipticCurveType(0)
}

func (x *EcdsaParams) GetEncoding() EcdsaSignatureEncoding {
	if x != nil {
		return x.Encoding
	}
	return EcdsaSignatureEncoding_UNKNOWN_ENCODING
}

// key_type: type.googleapis.com/google.crypto.tink.EcdsaPublicKey
type EcdsaPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	Params *EcdsaParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// Affine coordinates of the public key in bigendian representation. The
	// public key is a point (x, y) on the curve defined by params.curve. For
	// ECDH, it is crucial to verify whether the public key point (x, y) is on the
	// private's key curve. For ECDSA, such verification is a defense in depth.
	// Required.
	X []byte `protobuf:"bytes,3,opt,name=x,proto3" json:"x,omitempty"`
	// Required.
	Y []byte `protobuf:"bytes,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *EcdsaPublicKey) Reset() {
	*x = EcdsaPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdsaPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdsaPublicKey) ProtoMessage() {}

func (x *EcdsaPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdsaPublicKey.ProtoReflect.Descriptor instead.
func (*EcdsaPublicKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ecdsa_proto_rawDescGZIP(), []int{1}
}

func (x *EcdsaPublicKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EcdsaPublicKey) GetParams() *EcdsaParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *EcdsaPublicKey) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *EcdsaPublicKey) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.EcdsaPrivateKey
type EcdsaPrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Required.
	PublicKey *EcdsaPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Unsigned big integer in bigendian representation.
	// Required.
	KeyValue []byte `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (x *EcdsaPrivateKey) Reset() {
	*x = EcdsaPrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdsaPrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdsaPrivateKey) ProtoMessage() {}

func (x *EcdsaPrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdsaPrivateKey.ProtoReflect.Descriptor instead.
func (*EcdsaPrivateKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ecdsa_proto_rawDescGZIP(), []int{2}
}

func (x *EcdsaPrivateKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EcdsaPrivateKey) GetPublicKey() *EcdsaPublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *EcdsaPrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

type EcdsaKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Params  *EcdsaParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Version uint32       `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *EcdsaKeyFormat) Reset() {
	*x = EcdsaKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdsaKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdsaKeyFormat) ProtoMessage() {}

func (x *EcdsaKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ecdsa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdsaKeyFormat.ProtoReflect.Descriptor instead.
func (*EcdsaKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ecdsa_proto_rawDescGZIP(), []int{3}
}

func (x *EcdsaKeyFormat) GetParams() *EcdsaParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *EcdsaKeyFormat) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_third_party_tink_proto_ecdsa_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_ecdsa_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x1a, 0x23, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01,
	0x0a, 0x0b, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x39, 0x0a,
	0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x68, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x6c, 0x6c,
	0x69, 0x70, 0x74, 0x69, 0x63, 0x43, 0x75, 0x72, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05,
	0x63, 0x75, 0x72, 0x76, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x63, 0x64,
	0x73, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x7f, 0x0a,
	0x0e, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45,
	0x63, 0x64, 0x73, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x79, 0x22, 0x8b,
	0x01, 0x0a, 0x0f, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x63, 0x0a, 0x0e,
	0x45, 0x63, 0x64, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x37,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2a, 0x47, 0x0a, 0x16, 0x45, 0x63, 0x64, 0x73, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x10, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x45, 0x45, 0x45, 0x5f, 0x50, 0x31, 0x33, 0x36, 0x33, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x52, 0x10, 0x02, 0x42, 0x4d, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x64, 0x73, 0x61,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_third_party_tink_proto_ecdsa_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_ecdsa_proto_rawDescData = file_third_party_tink_proto_ecdsa_proto_rawDesc
)

func file_third_party_tink_proto_ecdsa_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_ecdsa_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_ecdsa_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_ecdsa_proto_rawDescData)
	})
	return file_third_party_tink_proto_ecdsa_proto_rawDescData
}

var file_third_party_tink_proto_ecdsa_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_third_party_tink_proto_ecdsa_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_third_party_tink_proto_ecdsa_proto_goTypes = []interface{}{
	(EcdsaSignatureEncoding)(0),            // 0: google.crypto.tink.EcdsaSignatureEncoding
	(*EcdsaParams)(nil),                    // 1: google.crypto.tink.EcdsaParams
	(*EcdsaPublicKey)(nil),                 // 2: google.crypto.tink.EcdsaPublicKey
	(*EcdsaPrivateKey)(nil),                // 3: google.crypto.tink.EcdsaPrivateKey
	(*EcdsaKeyFormat)(nil),                 // 4: google.crypto.tink.EcdsaKeyFormat
	(common_go_proto.HashType)(0),          // 5: google.crypto.tink.HashType
	(common_go_proto.EllipticCurveType)(0), // 6: google.crypto.tink.EllipticCurveType
}
var file_third_party_tink_proto_ecdsa_proto_depIdxs = []int32{
	5, // 0: google.crypto.tink.EcdsaParams.hash_type:type_name -> google.crypto.tink.HashType
	6, // 1: google.crypto.tink.EcdsaParams.curve:type_name -> google.crypto.tink.EllipticCurveType
	0, // 2: google.crypto.tink.EcdsaParams.encoding:type_name -> google.crypto.tink.EcdsaSignatureEncoding
	1, // 3: google.crypto.tink.EcdsaPublicKey.params:type_name -> google.crypto.tink.EcdsaParams
	2, // 4: google.crypto.tink.EcdsaPrivateKey.public_key:type_name -> google.crypto.tink.EcdsaPublicKey
	1, // 5: google.crypto.tink.EcdsaKeyFormat.params:type_name -> google.crypto.tink.EcdsaParams
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_ecdsa_proto_init() }
func file_third_party_tink_proto_ecdsa_proto_init() {
	if File_third_party_tink_proto_ecdsa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_ecdsa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdsaParams); i {
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
		file_third_party_tink_proto_ecdsa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdsaPublicKey); i {
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
		file_third_party_tink_proto_ecdsa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdsaPrivateKey); i {
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
		file_third_party_tink_proto_ecdsa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdsaKeyFormat); i {
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
			RawDescriptor: file_third_party_tink_proto_ecdsa_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_ecdsa_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_ecdsa_proto_depIdxs,
		EnumInfos:         file_third_party_tink_proto_ecdsa_proto_enumTypes,
		MessageInfos:      file_third_party_tink_proto_ecdsa_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_ecdsa_proto = out.File
	file_third_party_tink_proto_ecdsa_proto_rawDesc = nil
	file_third_party_tink_proto_ecdsa_proto_goTypes = nil
	file_third_party_tink_proto_ecdsa_proto_depIdxs = nil
}
