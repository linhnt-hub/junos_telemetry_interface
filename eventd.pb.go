// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: eventd.proto

package junos_telemetry_interface

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JunosEvents struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Events        *JunosEventsEventsType `protobuf:"bytes,151,opt,name=events" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEvents) Reset() {
	*x = JunosEvents{}
	mi := &file_eventd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEvents) ProtoMessage() {}

func (x *JunosEvents) ProtoReflect() protoreflect.Message {
	mi := &file_eventd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEvents.ProtoReflect.Descriptor instead.
func (*JunosEvents) Descriptor() ([]byte, []int) {
	return file_eventd_proto_rawDescGZIP(), []int{0}
}

func (x *JunosEvents) GetEvents() *JunosEventsEventsType {
	if x != nil {
		return x.Events
	}
	return nil
}

type JunosEventsEventsType struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Event         []*JunosEventsEventsTypeEventList `protobuf:"bytes,151,rep,name=event" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEventsEventsType) Reset() {
	*x = JunosEventsEventsType{}
	mi := &file_eventd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEventsEventsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEventsEventsType) ProtoMessage() {}

func (x *JunosEventsEventsType) ProtoReflect() protoreflect.Message {
	mi := &file_eventd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEventsEventsType.ProtoReflect.Descriptor instead.
func (*JunosEventsEventsType) Descriptor() ([]byte, []int) {
	return file_eventd_proto_rawDescGZIP(), []int{0, 0}
}

func (x *JunosEventsEventsType) GetEvent() []*JunosEventsEventsTypeEventList {
	if x != nil {
		return x.Event
	}
	return nil
}

type JunosEventsEventsTypeEventList struct {
	state         protoimpl.MessageState                          `protogen:"open.v1"`
	Id            *string                                         `protobuf:"bytes,51,opt,name=id" json:"id,omitempty"`
	Type          *string                                         `protobuf:"bytes,52,opt,name=type" json:"type,omitempty"`
	Timestamp     *JunosEventsEventsTypeEventListTimestampType    `protobuf:"bytes,151,opt,name=timestamp" json:"timestamp,omitempty"`
	Priority      *string                                         `protobuf:"bytes,53,opt,name=priority" json:"priority,omitempty"`
	Facility      *string                                         `protobuf:"bytes,54,opt,name=facility" json:"facility,omitempty"`
	Pid           *uint32                                         `protobuf:"varint,55,opt,name=pid" json:"pid,omitempty"`
	Message       *string                                         `protobuf:"bytes,56,opt,name=message" json:"message,omitempty"`
	Daemon        *string                                         `protobuf:"bytes,57,opt,name=daemon" json:"daemon,omitempty"`
	Hostname      *string                                         `protobuf:"bytes,58,opt,name=hostname" json:"hostname,omitempty"`
	Lsname        *string                                         `protobuf:"bytes,59,opt,name=lsname" json:"lsname,omitempty"`
	Attributes    []*JunosEventsEventsTypeEventListAttributesList `protobuf:"bytes,152,rep,name=attributes" json:"attributes,omitempty"`
	Logoptions    *int32                                          `protobuf:"varint,60,opt,name=logoptions" json:"logoptions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEventsEventsTypeEventList) Reset() {
	*x = JunosEventsEventsTypeEventList{}
	mi := &file_eventd_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEventsEventsTypeEventList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEventsEventsTypeEventList) ProtoMessage() {}

func (x *JunosEventsEventsTypeEventList) ProtoReflect() protoreflect.Message {
	mi := &file_eventd_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEventsEventsTypeEventList.ProtoReflect.Descriptor instead.
func (*JunosEventsEventsTypeEventList) Descriptor() ([]byte, []int) {
	return file_eventd_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *JunosEventsEventsTypeEventList) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetTimestamp() *JunosEventsEventsTypeEventListTimestampType {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *JunosEventsEventsTypeEventList) GetPriority() string {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetFacility() string {
	if x != nil && x.Facility != nil {
		return *x.Facility
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetPid() uint32 {
	if x != nil && x.Pid != nil {
		return *x.Pid
	}
	return 0
}

func (x *JunosEventsEventsTypeEventList) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetDaemon() string {
	if x != nil && x.Daemon != nil {
		return *x.Daemon
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetLsname() string {
	if x != nil && x.Lsname != nil {
		return *x.Lsname
	}
	return ""
}

func (x *JunosEventsEventsTypeEventList) GetAttributes() []*JunosEventsEventsTypeEventListAttributesList {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *JunosEventsEventsTypeEventList) GetLogoptions() int32 {
	if x != nil && x.Logoptions != nil {
		return *x.Logoptions
	}
	return 0
}

type JunosEventsEventsTypeEventListTimestampType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Seconds       *uint32                `protobuf:"varint,51,opt,name=seconds" json:"seconds,omitempty"`
	Microseconds  *uint32                `protobuf:"varint,52,opt,name=microseconds" json:"microseconds,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEventsEventsTypeEventListTimestampType) Reset() {
	*x = JunosEventsEventsTypeEventListTimestampType{}
	mi := &file_eventd_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEventsEventsTypeEventListTimestampType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEventsEventsTypeEventListTimestampType) ProtoMessage() {}

func (x *JunosEventsEventsTypeEventListTimestampType) ProtoReflect() protoreflect.Message {
	mi := &file_eventd_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEventsEventsTypeEventListTimestampType.ProtoReflect.Descriptor instead.
func (*JunosEventsEventsTypeEventListTimestampType) Descriptor() ([]byte, []int) {
	return file_eventd_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *JunosEventsEventsTypeEventListTimestampType) GetSeconds() uint32 {
	if x != nil && x.Seconds != nil {
		return *x.Seconds
	}
	return 0
}

func (x *JunosEventsEventsTypeEventListTimestampType) GetMicroseconds() uint32 {
	if x != nil && x.Microseconds != nil {
		return *x.Microseconds
	}
	return 0
}

type JunosEventsEventsTypeEventListAttributesList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           *string                `protobuf:"bytes,51,opt,name=key" json:"key,omitempty"`
	Value         *string                `protobuf:"bytes,52,opt,name=value" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEventsEventsTypeEventListAttributesList) Reset() {
	*x = JunosEventsEventsTypeEventListAttributesList{}
	mi := &file_eventd_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEventsEventsTypeEventListAttributesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEventsEventsTypeEventListAttributesList) ProtoMessage() {}

func (x *JunosEventsEventsTypeEventListAttributesList) ProtoReflect() protoreflect.Message {
	mi := &file_eventd_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEventsEventsTypeEventListAttributesList.ProtoReflect.Descriptor instead.
func (*JunosEventsEventsTypeEventListAttributesList) Descriptor() ([]byte, []int) {
	return file_eventd_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *JunosEventsEventsTypeEventListAttributesList) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *JunosEventsEventsTypeEventListAttributesList) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var file_eventd_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*JunosEvents)(nil),
		Field:         42,
		Name:          "jnpr_junos_events_ext",
		Tag:           "bytes,42,opt,name=jnpr_junos_events_ext",
		Filename:      "eventd.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional junos_events jnpr_junos_events_ext = 42;
	E_JnprJunosEventsExt = &file_eventd_proto_extTypes[0]
)

var File_eventd_proto protoreflect.FileDescriptor

var file_eventd_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x05, 0x0a, 0x0c, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x97,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x8a, 0x05, 0x0a, 0x0b, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x97, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0xbd, 0x04, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x34, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6a,
	0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x37, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x18, 0x39, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x3b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x98, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x4e, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x1a, 0x40, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x59, 0x0a, 0x15, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6a, 0x75,
	0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17,
	0x2e, 0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x12, 0x6a, 0x6e,
	0x70, 0x72, 0x4a, 0x75, 0x6e, 0x6f, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x78, 0x74,
	0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_eventd_proto_rawDescOnce sync.Once
	file_eventd_proto_rawDescData []byte
)

func file_eventd_proto_rawDescGZIP() []byte {
	file_eventd_proto_rawDescOnce.Do(func() {
		file_eventd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eventd_proto_rawDesc), len(file_eventd_proto_rawDesc)))
	})
	return file_eventd_proto_rawDescData
}

var file_eventd_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eventd_proto_goTypes = []any{
	(*JunosEvents)(nil),                                  // 0: junos_events
	(*JunosEventsEventsType)(nil),                        // 1: junos_events.events_type
	(*JunosEventsEventsTypeEventList)(nil),               // 2: junos_events.events_type.event_list
	(*JunosEventsEventsTypeEventListTimestampType)(nil),  // 3: junos_events.events_type.event_list.timestamp_type
	(*JunosEventsEventsTypeEventListAttributesList)(nil), // 4: junos_events.events_type.event_list.attributes_list
	(*JuniperNetworksSensors)(nil),                       // 5: JuniperNetworksSensors
}
var file_eventd_proto_depIdxs = []int32{
	1, // 0: junos_events.events:type_name -> junos_events.events_type
	2, // 1: junos_events.events_type.event:type_name -> junos_events.events_type.event_list
	3, // 2: junos_events.events_type.event_list.timestamp:type_name -> junos_events.events_type.event_list.timestamp_type
	4, // 3: junos_events.events_type.event_list.attributes:type_name -> junos_events.events_type.event_list.attributes_list
	5, // 4: jnpr_junos_events_ext:extendee -> JuniperNetworksSensors
	0, // 5: jnpr_junos_events_ext:type_name -> junos_events
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	5, // [5:6] is the sub-list for extension type_name
	4, // [4:5] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eventd_proto_init() }
func file_eventd_proto_init() {
	if File_eventd_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eventd_proto_rawDesc), len(file_eventd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_eventd_proto_goTypes,
		DependencyIndexes: file_eventd_proto_depIdxs,
		MessageInfos:      file_eventd_proto_msgTypes,
		ExtensionInfos:    file_eventd_proto_extTypes,
	}.Build()
	File_eventd_proto = out.File
	file_eventd_proto_goTypes = nil
	file_eventd_proto_depIdxs = nil
}
