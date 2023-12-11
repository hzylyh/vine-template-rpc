// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: alarm/v1/alarm.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AddAlarmRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AddAlarmRuleRequest) Reset() {
	*x = AddAlarmRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAlarmRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAlarmRuleRequest) ProtoMessage() {}

func (x *AddAlarmRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAlarmRuleRequest.ProtoReflect.Descriptor instead.
func (*AddAlarmRuleRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{0}
}

func (x *AddAlarmRuleRequest) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type AddAlarmRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddAlarmRuleReply) Reset() {
	*x = AddAlarmRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAlarmRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAlarmRuleReply) ProtoMessage() {}

func (x *AddAlarmRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAlarmRuleReply.ProtoReflect.Descriptor instead.
func (*AddAlarmRuleReply) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{1}
}

func (x *AddAlarmRuleReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAlarmRequest) Reset() {
	*x = UpdateAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlarmRequest) ProtoMessage() {}

func (x *UpdateAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlarmRequest.ProtoReflect.Descriptor instead.
func (*UpdateAlarmRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{2}
}

type UpdateAlarmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAlarmReply) Reset() {
	*x = UpdateAlarmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlarmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlarmReply) ProtoMessage() {}

func (x *UpdateAlarmReply) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlarmReply.ProtoReflect.Descriptor instead.
func (*UpdateAlarmReply) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{3}
}

type DeleteAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlarmRequest) Reset() {
	*x = DeleteAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlarmRequest) ProtoMessage() {}

func (x *DeleteAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlarmRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlarmRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{4}
}

type DeleteAlarmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlarmReply) Reset() {
	*x = DeleteAlarmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlarmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlarmReply) ProtoMessage() {}

func (x *DeleteAlarmReply) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlarmReply.ProtoReflect.Descriptor instead.
func (*DeleteAlarmReply) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{5}
}

type GetAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAlarmRequest) Reset() {
	*x = GetAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlarmRequest) ProtoMessage() {}

func (x *GetAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlarmRequest.ProtoReflect.Descriptor instead.
func (*GetAlarmRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{6}
}

type GetAlarmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAlarmReply) Reset() {
	*x = GetAlarmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlarmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlarmReply) ProtoMessage() {}

func (x *GetAlarmReply) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlarmReply.ProtoReflect.Descriptor instead.
func (*GetAlarmReply) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{7}
}

type ListAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAlarmRequest) Reset() {
	*x = ListAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlarmRequest) ProtoMessage() {}

func (x *ListAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlarmRequest.ProtoReflect.Descriptor instead.
func (*ListAlarmRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{8}
}

type ListAlarmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAlarmReply) Reset() {
	*x = ListAlarmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlarmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlarmReply) ProtoMessage() {}

func (x *ListAlarmReply) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlarmReply.ProtoReflect.Descriptor instead.
func (*ListAlarmReply) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{9}
}

type TriggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TriggerRequest) Reset() {
	*x = TriggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerRequest) ProtoMessage() {}

func (x *TriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerRequest.ProtoReflect.Descriptor instead.
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{10}
}

type TriggerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TriggerReply) Reset() {
	*x = TriggerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerReply) ProtoMessage() {}

func (x *TriggerReply) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerReply.ProtoReflect.Descriptor instead.
func (*TriggerReply) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{11}
}

var File_alarm_v1_alarm_proto protoreflect.FileDescriptor

var file_alarm_v1_alarm_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x2d, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x8c, 0x04, 0x0a,
	0x05, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x6e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x72, 0x75,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x4f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x49, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x07, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x42, 0x33, 0x0a, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x21, 0x76,
	0x69, 0x6e, 0x65, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alarm_v1_alarm_proto_rawDescOnce sync.Once
	file_alarm_v1_alarm_proto_rawDescData = file_alarm_v1_alarm_proto_rawDesc
)

func file_alarm_v1_alarm_proto_rawDescGZIP() []byte {
	file_alarm_v1_alarm_proto_rawDescOnce.Do(func() {
		file_alarm_v1_alarm_proto_rawDescData = protoimpl.X.CompressGZIP(file_alarm_v1_alarm_proto_rawDescData)
	})
	return file_alarm_v1_alarm_proto_rawDescData
}

var file_alarm_v1_alarm_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_alarm_v1_alarm_proto_goTypes = []interface{}{
	(*AddAlarmRuleRequest)(nil), // 0: api.alarm.v1.AddAlarmRuleRequest
	(*AddAlarmRuleReply)(nil),   // 1: api.alarm.v1.AddAlarmRuleReply
	(*UpdateAlarmRequest)(nil),  // 2: api.alarm.v1.UpdateAlarmRequest
	(*UpdateAlarmReply)(nil),    // 3: api.alarm.v1.UpdateAlarmReply
	(*DeleteAlarmRequest)(nil),  // 4: api.alarm.v1.DeleteAlarmRequest
	(*DeleteAlarmReply)(nil),    // 5: api.alarm.v1.DeleteAlarmReply
	(*GetAlarmRequest)(nil),     // 6: api.alarm.v1.GetAlarmRequest
	(*GetAlarmReply)(nil),       // 7: api.alarm.v1.GetAlarmReply
	(*ListAlarmRequest)(nil),    // 8: api.alarm.v1.ListAlarmRequest
	(*ListAlarmReply)(nil),      // 9: api.alarm.v1.ListAlarmReply
	(*TriggerRequest)(nil),      // 10: api.alarm.v1.TriggerRequest
	(*TriggerReply)(nil),        // 11: api.alarm.v1.TriggerReply
}
var file_alarm_v1_alarm_proto_depIdxs = []int32{
	0,  // 0: api.alarm.v1.Alarm.AddAlarmRule:input_type -> api.alarm.v1.AddAlarmRuleRequest
	2,  // 1: api.alarm.v1.Alarm.UpdateAlarm:input_type -> api.alarm.v1.UpdateAlarmRequest
	4,  // 2: api.alarm.v1.Alarm.DeleteAlarm:input_type -> api.alarm.v1.DeleteAlarmRequest
	6,  // 3: api.alarm.v1.Alarm.GetAlarm:input_type -> api.alarm.v1.GetAlarmRequest
	8,  // 4: api.alarm.v1.Alarm.ListAlarm:input_type -> api.alarm.v1.ListAlarmRequest
	10, // 5: api.alarm.v1.Alarm.Trigger:input_type -> api.alarm.v1.TriggerRequest
	1,  // 6: api.alarm.v1.Alarm.AddAlarmRule:output_type -> api.alarm.v1.AddAlarmRuleReply
	3,  // 7: api.alarm.v1.Alarm.UpdateAlarm:output_type -> api.alarm.v1.UpdateAlarmReply
	5,  // 8: api.alarm.v1.Alarm.DeleteAlarm:output_type -> api.alarm.v1.DeleteAlarmReply
	7,  // 9: api.alarm.v1.Alarm.GetAlarm:output_type -> api.alarm.v1.GetAlarmReply
	9,  // 10: api.alarm.v1.Alarm.ListAlarm:output_type -> api.alarm.v1.ListAlarmReply
	11, // 11: api.alarm.v1.Alarm.Trigger:output_type -> api.alarm.v1.TriggerReply
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_alarm_v1_alarm_proto_init() }
func file_alarm_v1_alarm_proto_init() {
	if File_alarm_v1_alarm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alarm_v1_alarm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAlarmRuleRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAlarmRuleReply); i {
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
		file_alarm_v1_alarm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlarmRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlarmReply); i {
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
		file_alarm_v1_alarm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlarmRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlarmReply); i {
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
		file_alarm_v1_alarm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlarmRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlarmReply); i {
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
		file_alarm_v1_alarm_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlarmRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlarmReply); i {
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
		file_alarm_v1_alarm_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerReply); i {
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
			RawDescriptor: file_alarm_v1_alarm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alarm_v1_alarm_proto_goTypes,
		DependencyIndexes: file_alarm_v1_alarm_proto_depIdxs,
		MessageInfos:      file_alarm_v1_alarm_proto_msgTypes,
	}.Build()
	File_alarm_v1_alarm_proto = out.File
	file_alarm_v1_alarm_proto_rawDesc = nil
	file_alarm_v1_alarm_proto_goTypes = nil
	file_alarm_v1_alarm_proto_depIdxs = nil
}
