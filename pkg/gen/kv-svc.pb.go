// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: kv-svc.proto

package kv_store_grpc

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KVMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVMessage) Reset() {
	*x = KVMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVMessage) ProtoMessage() {}

func (x *KVMessage) ProtoReflect() protoreflect.Message {
	mi := &file_kv_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVMessage.ProtoReflect.Descriptor instead.
func (*KVMessage) Descriptor() ([]byte, []int) {
	return file_kv_svc_proto_rawDescGZIP(), []int{0}
}

func (x *KVMessage) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetKVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKVRequest) Reset() {
	*x = GetKVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKVRequest) ProtoMessage() {}

func (x *GetKVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKVRequest.ProtoReflect.Descriptor instead.
func (*GetKVRequest) Descriptor() ([]byte, []int) {
	return file_kv_svc_proto_rawDescGZIP(), []int{1}
}

func (x *GetKVRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteKVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteKVRequest) Reset() {
	*x = DeleteKVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKVRequest) ProtoMessage() {}

func (x *DeleteKVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKVRequest.ProtoReflect.Descriptor instead.
func (*DeleteKVRequest) Descriptor() ([]byte, []int) {
	return file_kv_svc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteKVRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_kv_svc_proto protoreflect.FileDescriptor

var file_kv_svc_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x76, 0x2d, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x09, 0x4b, 0x56, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2e, 0x92, 0x41, 0x2b, 0x32, 0x29, 0x54, 0x68, 0x65, 0x20, 0x6b, 0x65,
	0x79, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x20, 0x77,
	0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4b, 0x56, 0x20, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x46, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0x92, 0x41, 0x2d, 0x32, 0x2b, 0x54, 0x68,
	0x65, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4b, 0x56,
	0x20, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x48, 0x92, 0x41, 0x45, 0x0a, 0x43, 0x2a, 0x09, 0x4b, 0x56, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x28, 0x41, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x6b, 0x65, 0x79, 0x2d, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0xd2, 0x01, 0x03, 0x6b,
	0x65, 0x79, 0xd2, 0x01, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4b, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x23, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x32, 0xef, 0x03, 0x0a, 0x09, 0x4b, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x86, 0x01, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x12, 0x18, 0x2e, 0x6b,
	0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x56, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x56, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x76,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x32, 0x0a, 0x09, 0x4b, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x4b, 0x56, 0x20, 0x70,
	0x61, 0x69, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x52, 0x50, 0x43, 0x2a, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x56, 0x12, 0x8f, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4b, 0x56, 0x12, 0x1e, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x56, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x76, 0x2f, 0x7b, 0x6b,
	0x65, 0x79, 0x7d, 0x92, 0x41, 0x34, 0x0a, 0x09, 0x4b, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x4b, 0x56, 0x20, 0x70,
	0x61, 0x69, 0x72, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x52, 0x50, 0x43,
	0x2a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x56, 0x12, 0x85, 0x01, 0x0a, 0x05, 0x47,
	0x65, 0x74, 0x4b, 0x56, 0x12, 0x1b, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4b, 0x56, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x76, 0x2f, 0x7b, 0x6b, 0x65, 0x79,
	0x7d, 0x92, 0x41, 0x2e, 0x0a, 0x09, 0x4b, 0x56, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x4b, 0x56, 0x20, 0x70, 0x61, 0x69, 0x72, 0x20, 0x66,
	0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x52, 0x50, 0x43, 0x2a, 0x05, 0x67, 0x65, 0x74,
	0x4b, 0x56, 0x1a, 0x3f, 0x92, 0x41, 0x3c, 0x12, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x4b, 0x65, 0x79, 0x20, 0x2f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x42, 0xf3, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x6f, 0x62, 0x2f, 0x6b, 0x76, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x92, 0x41, 0xcb, 0x01, 0x12, 0xa0,
	0x01, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x41, 0x50, 0x49, 0x12,
	0x29, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x52, 0x45, 0x41, 0x44, 0x2f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x2f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x20, 0x4b, 0x65, 0x79, 0x20, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x20, 0x50, 0x61, 0x69, 0x72, 0x73, 0x2a, 0x5e, 0x0a, 0x14, 0x42, 0x53,
	0x44, 0x20, 0x33, 0x2d, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c,
	0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2e, 0x74, 0x78, 0x74, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e,
	0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kv_svc_proto_rawDescOnce sync.Once
	file_kv_svc_proto_rawDescData = file_kv_svc_proto_rawDesc
)

func file_kv_svc_proto_rawDescGZIP() []byte {
	file_kv_svc_proto_rawDescOnce.Do(func() {
		file_kv_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_svc_proto_rawDescData)
	})
	return file_kv_svc_proto_rawDescData
}

var file_kv_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kv_svc_proto_goTypes = []interface{}{
	(*KVMessage)(nil),       // 0: kv_store_grpc.KVMessage
	(*GetKVRequest)(nil),    // 1: kv_store_grpc.GetKVRequest
	(*DeleteKVRequest)(nil), // 2: kv_store_grpc.DeleteKVRequest
	(*emptypb.Empty)(nil),   // 3: google.protobuf.Empty
}
var file_kv_svc_proto_depIdxs = []int32{
	0, // 0: kv_store_grpc.KVService.CreateKV:input_type -> kv_store_grpc.KVMessage
	2, // 1: kv_store_grpc.KVService.DeleteKV:input_type -> kv_store_grpc.DeleteKVRequest
	1, // 2: kv_store_grpc.KVService.GetKV:input_type -> kv_store_grpc.GetKVRequest
	0, // 3: kv_store_grpc.KVService.CreateKV:output_type -> kv_store_grpc.KVMessage
	3, // 4: kv_store_grpc.KVService.DeleteKV:output_type -> google.protobuf.Empty
	0, // 5: kv_store_grpc.KVService.GetKV:output_type -> kv_store_grpc.KVMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kv_svc_proto_init() }
func file_kv_svc_proto_init() {
	if File_kv_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVMessage); i {
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
		file_kv_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKVRequest); i {
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
		file_kv_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKVRequest); i {
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
			RawDescriptor: file_kv_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_svc_proto_goTypes,
		DependencyIndexes: file_kv_svc_proto_depIdxs,
		MessageInfos:      file_kv_svc_proto_msgTypes,
	}.Build()
	File_kv_svc_proto = out.File
	file_kv_svc_proto_rawDesc = nil
	file_kv_svc_proto_goTypes = nil
	file_kv_svc_proto_depIdxs = nil
}
