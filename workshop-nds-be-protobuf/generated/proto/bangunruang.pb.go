// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: bangunruang.proto

package proto

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

type RequestKubus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sisi int32 `protobuf:"varint,1,opt,name=sisi,proto3" json:"sisi,omitempty"`
}

func (x *RequestKubus) Reset() {
	*x = RequestKubus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestKubus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestKubus) ProtoMessage() {}

func (x *RequestKubus) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestKubus.ProtoReflect.Descriptor instead.
func (*RequestKubus) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{0}
}

func (x *RequestKubus) GetSisi() int32 {
	if x != nil {
		return x.Sisi
	}
	return 0
}

type ResponseVolumeKubus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume int32 `protobuf:"varint,1,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *ResponseVolumeKubus) Reset() {
	*x = ResponseVolumeKubus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseVolumeKubus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVolumeKubus) ProtoMessage() {}

func (x *ResponseVolumeKubus) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseVolumeKubus.ProtoReflect.Descriptor instead.
func (*ResponseVolumeKubus) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseVolumeKubus) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type RequestKelilingkubus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Panjang int32 `protobuf:"varint,1,opt,name=panjang,proto3" json:"panjang,omitempty"`
	Lebar   int32 `protobuf:"varint,2,opt,name=lebar,proto3" json:"lebar,omitempty"`
}

func (x *RequestKelilingkubus) Reset() {
	*x = RequestKelilingkubus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestKelilingkubus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestKelilingkubus) ProtoMessage() {}

func (x *RequestKelilingkubus) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestKelilingkubus.ProtoReflect.Descriptor instead.
func (*RequestKelilingkubus) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{2}
}

func (x *RequestKelilingkubus) GetPanjang() int32 {
	if x != nil {
		return x.Panjang
	}
	return 0
}

func (x *RequestKelilingkubus) GetLebar() int32 {
	if x != nil {
		return x.Lebar
	}
	return 0
}

type ResponseKelilingKubus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keliling int32 `protobuf:"varint,1,opt,name=keliling,proto3" json:"keliling,omitempty"`
}

func (x *ResponseKelilingKubus) Reset() {
	*x = ResponseKelilingKubus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseKelilingKubus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseKelilingKubus) ProtoMessage() {}

func (x *ResponseKelilingKubus) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseKelilingKubus.ProtoReflect.Descriptor instead.
func (*ResponseKelilingKubus) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseKelilingKubus) GetKeliling() int32 {
	if x != nil {
		return x.Keliling
	}
	return 0
}

type RequestLingkaran struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Radius float32 `protobuf:"fixed32,1,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *RequestLingkaran) Reset() {
	*x = RequestLingkaran{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLingkaran) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLingkaran) ProtoMessage() {}

func (x *RequestLingkaran) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLingkaran.ProtoReflect.Descriptor instead.
func (*RequestLingkaran) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{4}
}

func (x *RequestLingkaran) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type ResponseLingkaran struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Luaslingkaran     float32 `protobuf:"fixed32,1,opt,name=luaslingkaran,proto3" json:"luaslingkaran,omitempty"`
	Kelilinglingkaran float32 `protobuf:"fixed32,2,opt,name=kelilinglingkaran,proto3" json:"kelilinglingkaran,omitempty"`
}

func (x *ResponseLingkaran) Reset() {
	*x = ResponseLingkaran{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLingkaran) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLingkaran) ProtoMessage() {}

func (x *ResponseLingkaran) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLingkaran.ProtoReflect.Descriptor instead.
func (*ResponseLingkaran) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseLingkaran) GetLuaslingkaran() float32 {
	if x != nil {
		return x.Luaslingkaran
	}
	return 0
}

func (x *ResponseLingkaran) GetKelilinglingkaran() float32 {
	if x != nil {
		return x.Kelilinglingkaran
	}
	return 0
}

type RequestSegitiga struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alas   int32 `protobuf:"varint,1,opt,name=alas,proto3" json:"alas,omitempty"`
	Tinggi int32 `protobuf:"varint,2,opt,name=tinggi,proto3" json:"tinggi,omitempty"`
	Sisi1  int32 `protobuf:"varint,3,opt,name=sisi1,proto3" json:"sisi1,omitempty"`
	Sisi2  int32 `protobuf:"varint,4,opt,name=sisi2,proto3" json:"sisi2,omitempty"`
	Sisi3  int32 `protobuf:"varint,5,opt,name=sisi3,proto3" json:"sisi3,omitempty"`
}

func (x *RequestSegitiga) Reset() {
	*x = RequestSegitiga{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSegitiga) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSegitiga) ProtoMessage() {}

func (x *RequestSegitiga) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSegitiga.ProtoReflect.Descriptor instead.
func (*RequestSegitiga) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{6}
}

func (x *RequestSegitiga) GetAlas() int32 {
	if x != nil {
		return x.Alas
	}
	return 0
}

func (x *RequestSegitiga) GetTinggi() int32 {
	if x != nil {
		return x.Tinggi
	}
	return 0
}

func (x *RequestSegitiga) GetSisi1() int32 {
	if x != nil {
		return x.Sisi1
	}
	return 0
}

func (x *RequestSegitiga) GetSisi2() int32 {
	if x != nil {
		return x.Sisi2
	}
	return 0
}

func (x *RequestSegitiga) GetSisi3() int32 {
	if x != nil {
		return x.Sisi3
	}
	return 0
}

type ResponseSegitiga struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Luassegitiga     int32 `protobuf:"varint,1,opt,name=luassegitiga,proto3" json:"luassegitiga,omitempty"`
	Kelilingsegitiga int32 `protobuf:"varint,2,opt,name=kelilingsegitiga,proto3" json:"kelilingsegitiga,omitempty"`
}

func (x *ResponseSegitiga) Reset() {
	*x = ResponseSegitiga{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bangunruang_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSegitiga) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSegitiga) ProtoMessage() {}

func (x *ResponseSegitiga) ProtoReflect() protoreflect.Message {
	mi := &file_bangunruang_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSegitiga.ProtoReflect.Descriptor instead.
func (*ResponseSegitiga) Descriptor() ([]byte, []int) {
	return file_bangunruang_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseSegitiga) GetLuassegitiga() int32 {
	if x != nil {
		return x.Luassegitiga
	}
	return 0
}

func (x *ResponseSegitiga) GetKelilingsegitiga() int32 {
	if x != nil {
		return x.Kelilingsegitiga
	}
	return 0
}

var File_bangunruang_proto protoreflect.FileDescriptor

var file_bangunruang_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x61, 0x6e, 0x67, 0x75, 0x6e, 0x72, 0x75, 0x61, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x73, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x73, 0x69, 0x22, 0x2d,
	0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x4b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x46, 0x0a,
	0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x6b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x65, 0x62, 0x61, 0x72, 0x22, 0x33, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x67, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6c,
	0x75, 0x61, 0x73, 0x6c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x6c, 0x75, 0x61, 0x73, 0x6c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61,
	0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x6b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x6e,
	0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x6b, 0x65,
	0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x22,
	0x7f, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x67, 0x69, 0x74, 0x69,
	0x67, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6c, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x61, 0x6c, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x6e, 0x67, 0x67, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x69, 0x6e, 0x67, 0x67, 0x69, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x69, 0x73, 0x69, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x69, 0x73, 0x69, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x73, 0x69, 0x32, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x69, 0x73, 0x69, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69,
	0x73, 0x69, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x69, 0x73, 0x69, 0x33,
	0x22, 0x62, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x67, 0x69,
	0x74, 0x69, 0x67, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x75, 0x61, 0x73, 0x73, 0x65, 0x67, 0x69,
	0x74, 0x69, 0x67, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x75, 0x61, 0x73,
	0x73, 0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x12, 0x2a, 0x0a, 0x10, 0x6b, 0x65, 0x6c, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x6b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x67, 0x69,
	0x74, 0x69, 0x67, 0x61, 0x32, 0xb0, 0x02, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x67, 0x75, 0x6e, 0x52,
	0x75, 0x61, 0x6e, 0x67, 0x12, 0x44, 0x0a, 0x11, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x4b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x75, 0x73, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x52, 0x0a, 0x15, 0x48, 0x69,
	0x74, 0x75, 0x6e, 0x67, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x73,
	0x65, 0x67, 0x69, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x6b, 0x75, 0x62, 0x75, 0x73,
	0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4b, 0x65, 0x6c, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4b, 0x75, 0x62, 0x75, 0x73, 0x12, 0x44,
	0x0a, 0x0f, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x72, 0x61, 0x6e, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x6e, 0x67, 0x6b,
	0x61, 0x72, 0x61, 0x6e, 0x12, 0x41, 0x0a, 0x0e, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x53, 0x65,
	0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x65, 0x67, 0x69, 0x74, 0x69, 0x67, 0x61, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bangunruang_proto_rawDescOnce sync.Once
	file_bangunruang_proto_rawDescData = file_bangunruang_proto_rawDesc
)

func file_bangunruang_proto_rawDescGZIP() []byte {
	file_bangunruang_proto_rawDescOnce.Do(func() {
		file_bangunruang_proto_rawDescData = protoimpl.X.CompressGZIP(file_bangunruang_proto_rawDescData)
	})
	return file_bangunruang_proto_rawDescData
}

var file_bangunruang_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bangunruang_proto_goTypes = []interface{}{
	(*RequestKubus)(nil),          // 0: proto.RequestKubus
	(*ResponseVolumeKubus)(nil),   // 1: proto.ResponseVolumeKubus
	(*RequestKelilingkubus)(nil),  // 2: proto.RequestKelilingkubus
	(*ResponseKelilingKubus)(nil), // 3: proto.ResponseKelilingKubus
	(*RequestLingkaran)(nil),      // 4: proto.RequestLingkaran
	(*ResponseLingkaran)(nil),     // 5: proto.ResponseLingkaran
	(*RequestSegitiga)(nil),       // 6: proto.RequestSegitiga
	(*ResponseSegitiga)(nil),      // 7: proto.ResponseSegitiga
}
var file_bangunruang_proto_depIdxs = []int32{
	0, // 0: proto.BangunRuang.HitungVolumeKubus:input_type -> proto.RequestKubus
	2, // 1: proto.BangunRuang.HitungKelilingPersegi:input_type -> proto.RequestKelilingkubus
	4, // 2: proto.BangunRuang.HitungLingkaran:input_type -> proto.RequestLingkaran
	6, // 3: proto.BangunRuang.HitungSegitiga:input_type -> proto.RequestSegitiga
	1, // 4: proto.BangunRuang.HitungVolumeKubus:output_type -> proto.ResponseVolumeKubus
	3, // 5: proto.BangunRuang.HitungKelilingPersegi:output_type -> proto.ResponseKelilingKubus
	5, // 6: proto.BangunRuang.HitungLingkaran:output_type -> proto.ResponseLingkaran
	7, // 7: proto.BangunRuang.HitungSegitiga:output_type -> proto.ResponseSegitiga
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bangunruang_proto_init() }
func file_bangunruang_proto_init() {
	if File_bangunruang_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bangunruang_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestKubus); i {
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
		file_bangunruang_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseVolumeKubus); i {
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
		file_bangunruang_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestKelilingkubus); i {
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
		file_bangunruang_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseKelilingKubus); i {
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
		file_bangunruang_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLingkaran); i {
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
		file_bangunruang_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLingkaran); i {
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
		file_bangunruang_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSegitiga); i {
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
		file_bangunruang_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSegitiga); i {
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
			RawDescriptor: file_bangunruang_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bangunruang_proto_goTypes,
		DependencyIndexes: file_bangunruang_proto_depIdxs,
		MessageInfos:      file_bangunruang_proto_msgTypes,
	}.Build()
	File_bangunruang_proto = out.File
	file_bangunruang_proto_rawDesc = nil
	file_bangunruang_proto_goTypes = nil
	file_bangunruang_proto_depIdxs = nil
}
