// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: redis.proto

package redis

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

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ttl   int32  `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{0}
}

func (x *String) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *String) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *String) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Fields []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Ttl    int32    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{1}
}

func (x *Hash) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Hash) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Hash) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ttl   int32  `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{2}
}

func (x *List) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *List) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *List) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Member []string `protobuf:"bytes,2,rep,name=member,proto3" json:"member,omitempty"`
	Ttl    int32    `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{3}
}

func (x *Set) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Set) GetMember() []string {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *Set) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type ZSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Members []*ZMember `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Ttl     int32      `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *ZSet) Reset() {
	*x = ZSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZSet) ProtoMessage() {}

func (x *ZSet) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZSet.ProtoReflect.Descriptor instead.
func (*ZSet) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{4}
}

func (x *ZSet) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ZSet) GetMembers() []*ZMember {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *ZSet) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Messages []*XMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Ttl      int32       `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{5}
}

func (x *Stream) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Stream) GetMessages() []*XMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Stream) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{6}
}

func (x *Field) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Field) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ZMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member string  `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	Scores float64 `protobuf:"fixed64,2,opt,name=scores,proto3" json:"scores,omitempty"`
}

func (x *ZMember) Reset() {
	*x = ZMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZMember) ProtoMessage() {}

func (x *ZMember) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZMember.ProtoReflect.Descriptor instead.
func (*ZMember) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{7}
}

func (x *ZMember) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *ZMember) GetScores() float64 {
	if x != nil {
		return x.Scores
	}
	return 0
}

type XMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Field string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *XMessage) Reset() {
	*x = XMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMessage) ProtoMessage() {}

func (x *XMessage) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMessage.ProtoReflect.Descriptor instead.
func (*XMessage) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{8}
}

func (x *XMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *XMessage) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *XMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_redis_proto protoreflect.FileDescriptor

var file_redis_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x22, 0x42, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x50, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x40, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x41, 0x0a, 0x03,
	0x53, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22,
	0x54, 0x0a, 0x04, 0x5a, 0x53, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x5a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x59, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x58, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x22, 0x33, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x07, 0x5a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x22, 0x46, 0x0a, 0x08, 0x58, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redis_proto_rawDescOnce sync.Once
	file_redis_proto_rawDescData = file_redis_proto_rawDesc
)

func file_redis_proto_rawDescGZIP() []byte {
	file_redis_proto_rawDescOnce.Do(func() {
		file_redis_proto_rawDescData = protoimpl.X.CompressGZIP(file_redis_proto_rawDescData)
	})
	return file_redis_proto_rawDescData
}

var file_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_redis_proto_goTypes = []interface{}{
	(*String)(nil),   // 0: redis.String
	(*Hash)(nil),     // 1: redis.Hash
	(*List)(nil),     // 2: redis.List
	(*Set)(nil),      // 3: redis.Set
	(*ZSet)(nil),     // 4: redis.ZSet
	(*Stream)(nil),   // 5: redis.Stream
	(*Field)(nil),    // 6: redis.Field
	(*ZMember)(nil),  // 7: redis.ZMember
	(*XMessage)(nil), // 8: redis.XMessage
}
var file_redis_proto_depIdxs = []int32{
	6, // 0: redis.Hash.fields:type_name -> redis.Field
	7, // 1: redis.ZSet.members:type_name -> redis.ZMember
	8, // 2: redis.Stream.messages:type_name -> redis.XMessage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_redis_proto_init() }
func file_redis_proto_init() {
	if File_redis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
		file_redis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_redis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
		file_redis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_redis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZSet); i {
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
		file_redis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_redis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_redis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZMember); i {
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
		file_redis_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMessage); i {
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
			RawDescriptor: file_redis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redis_proto_goTypes,
		DependencyIndexes: file_redis_proto_depIdxs,
		MessageInfos:      file_redis_proto_msgTypes,
	}.Build()
	File_redis_proto = out.File
	file_redis_proto_rawDesc = nil
	file_redis_proto_goTypes = nil
	file_redis_proto_depIdxs = nil
}