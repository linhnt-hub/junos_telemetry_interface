// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: sr_te_per_lsp_transit_stats.proto

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

type Mpls_SrTePerLspTransitStats struct {
	state              protoimpl.MessageState                             `protogen:"open.v1"`
	SignalingProtocols *Mpls_SrTePerLspTransitStatsSignalingProtocolsType `protobuf:"bytes,151,opt,name=signaling_protocols,json=signalingProtocols" json:"signaling_protocols,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Mpls_SrTePerLspTransitStats) Reset() {
	*x = Mpls_SrTePerLspTransitStats{}
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mpls_SrTePerLspTransitStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mpls_SrTePerLspTransitStats) ProtoMessage() {}

func (x *Mpls_SrTePerLspTransitStats) ProtoReflect() protoreflect.Message {
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mpls_SrTePerLspTransitStats.ProtoReflect.Descriptor instead.
func (*Mpls_SrTePerLspTransitStats) Descriptor() ([]byte, []int) {
	return file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Mpls_SrTePerLspTransitStats) GetSignalingProtocols() *Mpls_SrTePerLspTransitStatsSignalingProtocolsType {
	if x != nil {
		return x.SignalingProtocols
	}
	return nil
}

type Mpls_SrTePerLspTransitStatsSignalingProtocolsType struct {
	state          protoimpl.MessageState                                               `protogen:"open.v1"`
	SegmentRouting *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType `protobuf:"bytes,151,opt,name=segment_routing,json=segmentRouting" json:"segment_routing,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsType) Reset() {
	*x = Mpls_SrTePerLspTransitStatsSignalingProtocolsType{}
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsType) ProtoMessage() {}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsType) ProtoReflect() protoreflect.Message {
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mpls_SrTePerLspTransitStatsSignalingProtocolsType.ProtoReflect.Descriptor instead.
func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsType) Descriptor() ([]byte, []int) {
	return file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsType) GetSegmentRouting() *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType {
	if x != nil {
		return x.SegmentRouting
	}
	return nil
}

type Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType struct {
	state                     protoimpl.MessageState                                                                            `protogen:"open.v1"`
	SrTePerLspTransitPolicies *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType `protobuf:"bytes,151,opt,name=sr_te_per_lsp_transit_policies,json=srTePerLspTransitPolicies" json:"sr_te_per_lsp_transit_policies,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType) Reset() {
	*x = Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType{}
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType) ProtoMessage() {}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType) ProtoReflect() protoreflect.Message {
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType.ProtoReflect.Descriptor instead.
func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType) Descriptor() ([]byte, []int) {
	return file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType) GetSrTePerLspTransitPolicies() *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType {
	if x != nil {
		return x.SrTePerLspTransitPolicies
	}
	return nil
}

type Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType struct {
	state                   protoimpl.MessageState                                                                                                         `protogen:"open.v1"`
	SrTePerLspTransitPolicy []*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList `protobuf:"bytes,151,rep,name=sr_te_per_lsp_transit_policy,json=srTePerLspTransitPolicy" json:"sr_te_per_lsp_transit_policy,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType) Reset() {
	*x = Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType{}
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType) ProtoMessage() {
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType) ProtoReflect() protoreflect.Message {
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType.ProtoReflect.Descriptor instead.
func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType) Descriptor() ([]byte, []int) {
	return file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType) GetSrTePerLspTransitPolicy() []*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList {
	if x != nil {
		return x.SrTePerLspTransitPolicy
	}
	return nil
}

type Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList struct {
	state         protoimpl.MessageState                                                                                                                   `protogen:"open.v1"`
	Record        []*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList `protobuf:"bytes,151,rep,name=record" json:"record,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList) Reset() {
	*x = Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList{}
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList) ProtoMessage() {
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList) ProtoReflect() protoreflect.Message {
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList.ProtoReflect.Descriptor instead.
func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList) Descriptor() ([]byte, []int) {
	return file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0}
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList) GetRecord() []*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList {
	if x != nil {
		return x.Record
	}
	return nil
}

type Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	TunnelName         *string                `protobuf:"bytes,51,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	Source             *string                `protobuf:"bytes,52,opt,name=source" json:"source,omitempty"`
	Origin             *string                `protobuf:"bytes,53,opt,name=origin" json:"origin,omitempty"`
	Distinguisher      *string                `protobuf:"bytes,54,opt,name=distinguisher" json:"distinguisher,omitempty"`
	LspName            *string                `protobuf:"bytes,55,opt,name=lsp_name,json=lspName" json:"lsp_name,omitempty"`
	CounterName        *string                `protobuf:"bytes,56,opt,name=counter_name,json=counterName" json:"counter_name,omitempty"`
	InstanceIdentifier *uint32                `protobuf:"varint,57,opt,name=instance_identifier,json=instanceIdentifier" json:"instance_identifier,omitempty"`
	Packets            *uint64                `protobuf:"varint,58,opt,name=packets" json:"packets,omitempty"`
	Bytes              *uint64                `protobuf:"varint,59,opt,name=bytes" json:"bytes,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) Reset() {
	*x = Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList{}
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) ProtoMessage() {
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) ProtoReflect() protoreflect.Message {
	mi := &file_sr_te_per_lsp_transit_stats_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList.ProtoReflect.Descriptor instead.
func (*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) Descriptor() ([]byte, []int) {
	return file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0, 0}
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetTunnelName() string {
	if x != nil && x.TunnelName != nil {
		return *x.TunnelName
	}
	return ""
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetOrigin() string {
	if x != nil && x.Origin != nil {
		return *x.Origin
	}
	return ""
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetDistinguisher() string {
	if x != nil && x.Distinguisher != nil {
		return *x.Distinguisher
	}
	return ""
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetLspName() string {
	if x != nil && x.LspName != nil {
		return *x.LspName
	}
	return ""
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetCounterName() string {
	if x != nil && x.CounterName != nil {
		return *x.CounterName
	}
	return ""
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetInstanceIdentifier() uint32 {
	if x != nil && x.InstanceIdentifier != nil {
		return *x.InstanceIdentifier
	}
	return 0
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetPackets() uint64 {
	if x != nil && x.Packets != nil {
		return *x.Packets
	}
	return 0
}

func (x *Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList) GetBytes() uint64 {
	if x != nil && x.Bytes != nil {
		return *x.Bytes
	}
	return 0
}

var file_sr_te_per_lsp_transit_stats_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*Mpls_SrTePerLspTransitStats)(nil),
		Field:         104,
		Name:          "jnpr_mpls_SrTePerLspTransitStats_ext",
		Tag:           "bytes,104,opt,name=jnpr_mpls_SrTePerLspTransitStats_ext",
		Filename:      "sr_te_per_lsp_transit_stats.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional mpls_SrTePerLspTransitStats jnpr_mpls_SrTePerLspTransitStats_ext = 104;
	E_JnprMpls_SrTePerLspTransitStatsExt = &file_sr_te_per_lsp_transit_stats_proto_extTypes[0]
)

var File_sr_te_per_lsp_transit_stats_proto protoreflect.FileDescriptor

var file_sr_te_per_lsp_transit_stats_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x70, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x0a, 0x0a, 0x1b, 0x6d, 0x70, 0x6c,
	0x73, 0x5f, 0x53, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x67, 0x0a, 0x13, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18,
	0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x53, 0x72,
	0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x12, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x1a, 0x90, 0x09, 0x0a, 0x18, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x74,
	0x0a, 0x0f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x5f,
	0x53, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x1a, 0xfd, 0x07, 0x0a, 0x14, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0xb2, 0x01,
	0x0a, 0x1e, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x70, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6e, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x53,
	0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73,
	0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x19, 0x73, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72,
	0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x1a, 0xaf, 0x06, 0x0a, 0x23, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x6c, 0x73, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0xd1, 0x01, 0x0a, 0x1c, 0x73,
	0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x70, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x97, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x90, 0x01, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x53, 0x72, 0x54, 0x65, 0x50,
	0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x70, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73,
	0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x17, 0x73, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73,
	0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0xb3,
	0x04, 0x0a, 0x21, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x70,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0xb6, 0x01, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x97, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x9c, 0x01, 0x2e, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x53,
	0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x73,
	0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x72, 0x5f, 0x74, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x6c, 0x73, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0xd4, 0x02,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0b, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x35,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x75, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x18, 0x36, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08,
	0x01, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x08, 0x6c, 0x73, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x37, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x07, 0x6c, 0x73, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x39, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x18, 0x01, 0x52, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x3b, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x18, 0x01, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x3a, 0x86, 0x01, 0x0a, 0x24, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6d, 0x70,
	0x6c, 0x73, 0x5f, 0x53, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e,
	0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d,
	0x70, 0x6c, 0x73, 0x5f, 0x53, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x21, 0x6a, 0x6e, 0x70, 0x72,
	0x4d, 0x70, 0x6c, 0x73, 0x53, 0x72, 0x54, 0x65, 0x50, 0x65, 0x72, 0x4c, 0x73, 0x70, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a,
	0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_sr_te_per_lsp_transit_stats_proto_rawDescOnce sync.Once
	file_sr_te_per_lsp_transit_stats_proto_rawDescData []byte
)

func file_sr_te_per_lsp_transit_stats_proto_rawDescGZIP() []byte {
	file_sr_te_per_lsp_transit_stats_proto_rawDescOnce.Do(func() {
		file_sr_te_per_lsp_transit_stats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sr_te_per_lsp_transit_stats_proto_rawDesc), len(file_sr_te_per_lsp_transit_stats_proto_rawDesc)))
	})
	return file_sr_te_per_lsp_transit_stats_proto_rawDescData
}

var file_sr_te_per_lsp_transit_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sr_te_per_lsp_transit_stats_proto_goTypes = []any{
	(*Mpls_SrTePerLspTransitStats)(nil),                                                                                                           // 0: mpls_SrTePerLspTransitStats
	(*Mpls_SrTePerLspTransitStatsSignalingProtocolsType)(nil),                                                                                     // 1: mpls_SrTePerLspTransitStats.signaling_protocols_type
	(*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingType)(nil),                                                                   // 2: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type
	(*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesType)(nil),                                      // 3: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type
	(*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyList)(nil),           // 4: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type.sr_te_per_lsp_transit_policy_list
	(*Mpls_SrTePerLspTransitStatsSignalingProtocolsTypeSegmentRoutingTypeSrTePerLspTransitPoliciesTypeSrTePerLspTransitPolicyListRecordList)(nil), // 5: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type.sr_te_per_lsp_transit_policy_list.record_list
	(*JuniperNetworksSensors)(nil), // 6: JuniperNetworksSensors
}
var file_sr_te_per_lsp_transit_stats_proto_depIdxs = []int32{
	1, // 0: mpls_SrTePerLspTransitStats.signaling_protocols:type_name -> mpls_SrTePerLspTransitStats.signaling_protocols_type
	2, // 1: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing:type_name -> mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type
	3, // 2: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies:type_name -> mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type
	4, // 3: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type.sr_te_per_lsp_transit_policy:type_name -> mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type.sr_te_per_lsp_transit_policy_list
	5, // 4: mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type.sr_te_per_lsp_transit_policy_list.record:type_name -> mpls_SrTePerLspTransitStats.signaling_protocols_type.segment_routing_type.sr_te_per_lsp_transit_policies_type.sr_te_per_lsp_transit_policy_list.record_list
	6, // 5: jnpr_mpls_SrTePerLspTransitStats_ext:extendee -> JuniperNetworksSensors
	0, // 6: jnpr_mpls_SrTePerLspTransitStats_ext:type_name -> mpls_SrTePerLspTransitStats
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	6, // [6:7] is the sub-list for extension type_name
	5, // [5:6] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sr_te_per_lsp_transit_stats_proto_init() }
func file_sr_te_per_lsp_transit_stats_proto_init() {
	if File_sr_te_per_lsp_transit_stats_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sr_te_per_lsp_transit_stats_proto_rawDesc), len(file_sr_te_per_lsp_transit_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_sr_te_per_lsp_transit_stats_proto_goTypes,
		DependencyIndexes: file_sr_te_per_lsp_transit_stats_proto_depIdxs,
		MessageInfos:      file_sr_te_per_lsp_transit_stats_proto_msgTypes,
		ExtensionInfos:    file_sr_te_per_lsp_transit_stats_proto_extTypes,
	}.Build()
	File_sr_te_per_lsp_transit_stats_proto = out.File
	file_sr_te_per_lsp_transit_stats_proto_goTypes = nil
	file_sr_te_per_lsp_transit_stats_proto_depIdxs = nil
}
