// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: bbe-smgd_rsmon_stats_oc.proto

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

type JunosRsmonStats struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	System        *JunosRsmonStatsSystemType `protobuf:"bytes,151,opt,name=system" json:"system,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosRsmonStats) Reset() {
	*x = JunosRsmonStats{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStats) ProtoMessage() {}

func (x *JunosRsmonStats) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStats.ProtoReflect.Descriptor instead.
func (*JunosRsmonStats) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0}
}

func (x *JunosRsmonStats) GetSystem() *JunosRsmonStatsSystemType {
	if x != nil {
		return x.System
	}
	return nil
}

type JunosRsmonStatsSystemType struct {
	state                protoimpl.MessageState                             `protogen:"open.v1"`
	SubscriberManagement *JunosRsmonStatsSystemTypeSubscriberManagementType `protobuf:"bytes,151,opt,name=subscriber_management,json=subscriberManagement" json:"subscriber_management,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemType) Reset() {
	*x = JunosRsmonStatsSystemType{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemType) ProtoMessage() {}

func (x *JunosRsmonStatsSystemType) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemType.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemType) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0}
}

func (x *JunosRsmonStatsSystemType) GetSubscriberManagement() *JunosRsmonStatsSystemTypeSubscriberManagementType {
	if x != nil {
		return x.SubscriberManagement
	}
	return nil
}

type JunosRsmonStatsSystemTypeSubscriberManagementType struct {
	state         protoimpl.MessageState                                      `protogen:"open.v1"`
	Infra         *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType `protobuf:"bytes,151,opt,name=infra" json:"infra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementType) Reset() {
	*x = JunosRsmonStatsSystemTypeSubscriberManagementType{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemTypeSubscriberManagementType) ProtoMessage() {}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementType) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemTypeSubscriberManagementType.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemTypeSubscriberManagementType) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementType) GetInfra() *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType {
	if x != nil {
		return x.Infra
	}
	return nil
}

type JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType struct {
	state           protoimpl.MessageState                                                         `protogen:"open.v1"`
	ResourceMonitor *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType `protobuf:"bytes,151,opt,name=resource_monitor,json=resourceMonitor" json:"resource_monitor,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType) Reset() {
	*x = JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType) ProtoMessage() {}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType) GetResourceMonitor() *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType {
	if x != nil {
		return x.ResourceMonitor
	}
	return nil
}

type JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType struct {
	state         protoimpl.MessageState                                                                       `protogen:"open.v1"`
	RsmonInfra    *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType `protobuf:"bytes,151,opt,name=rsmon_infra,json=rsmonInfra" json:"rsmon_infra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType) Reset() {
	*x = JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType) ProtoMessage() {
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0}
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType) GetRsmonInfra() *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType {
	if x != nil {
		return x.RsmonInfra
	}
	return nil
}

type JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType struct {
	state         protoimpl.MessageState                                                                               `protogen:"open.v1"`
	Fpcs          *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType `protobuf:"bytes,151,opt,name=fpcs" json:"fpcs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType) Reset() {
	*x = JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType) ProtoMessage() {
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0, 0}
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType) GetFpcs() *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType {
	if x != nil {
		return x.Fpcs
	}
	return nil
}

type JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType struct {
	state         protoimpl.MessageState                                                                                        `protogen:"open.v1"`
	Fpc           []*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList `protobuf:"bytes,151,rep,name=fpc" json:"fpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType) Reset() {
	*x = JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType) ProtoMessage() {
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0, 0, 0}
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType) GetFpc() []*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList {
	if x != nil {
		return x.Fpc
	}
	return nil
}

type JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	Slot                   *uint32                `protobuf:"varint,51,opt,name=slot" json:"slot,omitempty"`
	DelayRoundTripExceeded *uint32                `protobuf:"varint,61,opt,name=delay_round_trip_exceeded,json=delayRoundTripExceeded" json:"delay_round_trip_exceeded,omitempty"`
	DelayRoundTripNominal  *uint32                `protobuf:"varint,62,opt,name=delay_round_trip_nominal,json=delayRoundTripNominal" json:"delay_round_trip_nominal,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) Reset() {
	*x = JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList{}
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) ProtoMessage() {
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) ProtoReflect() protoreflect.Message {
	mi := &file_bbe_smgd_rsmon_stats_oc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList.ProtoReflect.Descriptor instead.
func (*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) Descriptor() ([]byte, []int) {
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0, 0, 0, 0}
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) GetSlot() uint32 {
	if x != nil && x.Slot != nil {
		return *x.Slot
	}
	return 0
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) GetDelayRoundTripExceeded() uint32 {
	if x != nil && x.DelayRoundTripExceeded != nil {
		return *x.DelayRoundTripExceeded
	}
	return 0
}

func (x *JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList) GetDelayRoundTripNominal() uint32 {
	if x != nil && x.DelayRoundTripNominal != nil {
		return *x.DelayRoundTripNominal
	}
	return 0
}

var file_bbe_smgd_rsmon_stats_oc_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*JunosRsmonStats)(nil),
		Field:         37,
		Name:          "jnpr_junos_rsmon_stats_ext",
		Tag:           "bytes,37,opt,name=jnpr_junos_rsmon_stats_ext",
		Filename:      "bbe-smgd_rsmon_stats_oc.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional junos_rsmon_stats jnpr_junos_rsmon_stats_ext = 37;
	E_JnprJunosRsmonStatsExt = &file_bbe_smgd_rsmon_stats_oc_proto_extTypes[0]
)

var File_bbe_smgd_rsmon_stats_oc_proto protoreflect.FileDescriptor

var file_bbe_smgd_rsmon_stats_oc_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x62, 0x62, 0x65, 0x2d, 0x73, 0x6d, 0x67, 0x64, 0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x08, 0x0a, 0x11, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x72,
	0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x75,
	0x6e, 0x6f, 0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x1a, 0x9c, 0x08, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x6f, 0x0a, 0x15, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x97, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x14,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x9b, 0x07, 0x0a, 0x1a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x5b, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x18, 0x97, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x1a, 0x9f, 0x06, 0x0a, 0x0a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x86, 0x01, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x6a, 0x75,
	0x6e, 0x6f, 0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x87, 0x05, 0x0a, 0x15, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0b, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6b, 0x2e, 0x6a, 0x75, 0x6e, 0x6f,
	0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x1a, 0xdd, 0x03, 0x0a, 0x10, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x04, 0x66, 0x70, 0x63, 0x73,
	0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x75, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f,
	0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x66, 0x70, 0x63, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x66, 0x70, 0x63, 0x73, 0x1a, 0xbb, 0x02, 0x0a, 0x09, 0x66, 0x70, 0x63, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x03, 0x66, 0x70, 0x63, 0x18, 0x97, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x7e, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x72, 0x73,
	0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x66,
	0x70, 0x63, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x66, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x03, 0x66, 0x70, 0x63, 0x1a, 0x99, 0x01, 0x0a, 0x08, 0x66, 0x70, 0x63, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x39,
	0x0a, 0x19, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72,
	0x69, 0x70, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x3d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x16, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x69,
	0x70, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x6e, 0x6f,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x69, 0x70, 0x4e, 0x6f, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x3a, 0x67, 0x0a, 0x1a, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6a, 0x75, 0x6e, 0x6f, 0x73,
	0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x65, 0x78, 0x74,
	0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x25, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x72, 0x73, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x16, 0x6a, 0x6e, 0x70, 0x72, 0x4a, 0x75, 0x6e, 0x6f, 0x73, 0x52, 0x73,
	0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x2e,
	0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_bbe_smgd_rsmon_stats_oc_proto_rawDescOnce sync.Once
	file_bbe_smgd_rsmon_stats_oc_proto_rawDescData []byte
)

func file_bbe_smgd_rsmon_stats_oc_proto_rawDescGZIP() []byte {
	file_bbe_smgd_rsmon_stats_oc_proto_rawDescOnce.Do(func() {
		file_bbe_smgd_rsmon_stats_oc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bbe_smgd_rsmon_stats_oc_proto_rawDesc), len(file_bbe_smgd_rsmon_stats_oc_proto_rawDesc)))
	})
	return file_bbe_smgd_rsmon_stats_oc_proto_rawDescData
}

var file_bbe_smgd_rsmon_stats_oc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bbe_smgd_rsmon_stats_oc_proto_goTypes = []any{
	(*JunosRsmonStats)(nil),                                                                                            // 0: junos_rsmon_stats
	(*JunosRsmonStatsSystemType)(nil),                                                                                  // 1: junos_rsmon_stats.system_type
	(*JunosRsmonStatsSystemTypeSubscriberManagementType)(nil),                                                          // 2: junos_rsmon_stats.system_type.subscriber_management_type
	(*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraType)(nil),                                                 // 3: junos_rsmon_stats.system_type.subscriber_management_type.infra_type
	(*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorType)(nil),                              // 4: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type
	(*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraType)(nil),                // 5: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type
	(*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsType)(nil),        // 6: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type.fpcs_type
	(*JunosRsmonStatsSystemTypeSubscriberManagementTypeInfraTypeResourceMonitorTypeRsmonInfraTypeFpcsTypeFpcList)(nil), // 7: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type.fpcs_type.fpc_list
	(*JuniperNetworksSensors)(nil),                                                                                     // 8: JuniperNetworksSensors
}
var file_bbe_smgd_rsmon_stats_oc_proto_depIdxs = []int32{
	1, // 0: junos_rsmon_stats.system:type_name -> junos_rsmon_stats.system_type
	2, // 1: junos_rsmon_stats.system_type.subscriber_management:type_name -> junos_rsmon_stats.system_type.subscriber_management_type
	3, // 2: junos_rsmon_stats.system_type.subscriber_management_type.infra:type_name -> junos_rsmon_stats.system_type.subscriber_management_type.infra_type
	4, // 3: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor:type_name -> junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type
	5, // 4: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra:type_name -> junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type
	6, // 5: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type.fpcs:type_name -> junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type.fpcs_type
	7, // 6: junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type.fpcs_type.fpc:type_name -> junos_rsmon_stats.system_type.subscriber_management_type.infra_type.resource_monitor_type.rsmon_infra_type.fpcs_type.fpc_list
	8, // 7: jnpr_junos_rsmon_stats_ext:extendee -> JuniperNetworksSensors
	0, // 8: jnpr_junos_rsmon_stats_ext:type_name -> junos_rsmon_stats
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	8, // [8:9] is the sub-list for extension type_name
	7, // [7:8] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_bbe_smgd_rsmon_stats_oc_proto_init() }
func file_bbe_smgd_rsmon_stats_oc_proto_init() {
	if File_bbe_smgd_rsmon_stats_oc_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bbe_smgd_rsmon_stats_oc_proto_rawDesc), len(file_bbe_smgd_rsmon_stats_oc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_bbe_smgd_rsmon_stats_oc_proto_goTypes,
		DependencyIndexes: file_bbe_smgd_rsmon_stats_oc_proto_depIdxs,
		MessageInfos:      file_bbe_smgd_rsmon_stats_oc_proto_msgTypes,
		ExtensionInfos:    file_bbe_smgd_rsmon_stats_oc_proto_extTypes,
	}.Build()
	File_bbe_smgd_rsmon_stats_oc_proto = out.File
	file_bbe_smgd_rsmon_stats_oc_proto_goTypes = nil
	file_bbe_smgd_rsmon_stats_oc_proto_depIdxs = nil
}
