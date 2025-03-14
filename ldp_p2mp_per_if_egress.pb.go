// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: ldp_p2mp_per_if_egress.proto

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

// Top-level message
type LdpP2MpPerIfEgress struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of LDP P2mp stats per IF egress records
	PerIfRecords  []*LdpP2MpInterfaceRecord `protobuf:"bytes,1,rep,name=per_if_records,json=perIfRecords" json:"per_if_records,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LdpP2MpPerIfEgress) Reset() {
	*x = LdpP2MpPerIfEgress{}
	mi := &file_ldp_p2mp_per_if_egress_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LdpP2MpPerIfEgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LdpP2MpPerIfEgress) ProtoMessage() {}

func (x *LdpP2MpPerIfEgress) ProtoReflect() protoreflect.Message {
	mi := &file_ldp_p2mp_per_if_egress_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LdpP2MpPerIfEgress.ProtoReflect.Descriptor instead.
func (*LdpP2MpPerIfEgress) Descriptor() ([]byte, []int) {
	return file_ldp_p2mp_per_if_egress_proto_rawDescGZIP(), []int{0}
}

func (x *LdpP2MpPerIfEgress) GetPerIfRecords() []*LdpP2MpInterfaceRecord {
	if x != nil {
		return x.PerIfRecords
	}
	return nil
}

// SR statistics record
type LdpP2MpInterfaceRecord struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Interface name, e.g., xe-0/0/0
	IfName *string `protobuf:"bytes,1,req,name=if_name,json=ifName" json:"if_name,omitempty"`
	// Name of the counter. This is useful when an interface has multiple counters.
	// for some scenarios, it is possible that a new counter is
	// created in the hardware.
	CounterName *string `protobuf:"bytes,2,opt,name=counter_name,json=counterName" json:"counter_name,omitempty"`
	// Traffic statistics
	IngressStats  *LabelDistributionProtocolP2MpIfStats `protobuf:"bytes,3,opt,name=ingress_stats,json=ingressStats" json:"ingress_stats,omitempty"`
	EgressStats   *LabelDistributionProtocolP2MpIfStats `protobuf:"bytes,4,opt,name=egress_stats,json=egressStats" json:"egress_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LdpP2MpInterfaceRecord) Reset() {
	*x = LdpP2MpInterfaceRecord{}
	mi := &file_ldp_p2mp_per_if_egress_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LdpP2MpInterfaceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LdpP2MpInterfaceRecord) ProtoMessage() {}

func (x *LdpP2MpInterfaceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_ldp_p2mp_per_if_egress_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LdpP2MpInterfaceRecord.ProtoReflect.Descriptor instead.
func (*LdpP2MpInterfaceRecord) Descriptor() ([]byte, []int) {
	return file_ldp_p2mp_per_if_egress_proto_rawDescGZIP(), []int{1}
}

func (x *LdpP2MpInterfaceRecord) GetIfName() string {
	if x != nil && x.IfName != nil {
		return *x.IfName
	}
	return ""
}

func (x *LdpP2MpInterfaceRecord) GetCounterName() string {
	if x != nil && x.CounterName != nil {
		return *x.CounterName
	}
	return ""
}

func (x *LdpP2MpInterfaceRecord) GetIngressStats() *LabelDistributionProtocolP2MpIfStats {
	if x != nil {
		return x.IngressStats
	}
	return nil
}

func (x *LdpP2MpInterfaceRecord) GetEgressStats() *LabelDistributionProtocolP2MpIfStats {
	if x != nil {
		return x.EgressStats
	}
	return nil
}

type LabelDistributionProtocolP2MpIfStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Packet and Byte statistics
	Packets *uint64 `protobuf:"varint,1,opt,name=packets" json:"packets,omitempty"`
	Bytes   *uint64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	// Rates of the above counters.
	PacketRate    *uint64 `protobuf:"varint,3,opt,name=packet_rate,json=packetRate" json:"packet_rate,omitempty"`
	ByteRate      *uint64 `protobuf:"varint,4,opt,name=byte_rate,json=byteRate" json:"byte_rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelDistributionProtocolP2MpIfStats) Reset() {
	*x = LabelDistributionProtocolP2MpIfStats{}
	mi := &file_ldp_p2mp_per_if_egress_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelDistributionProtocolP2MpIfStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelDistributionProtocolP2MpIfStats) ProtoMessage() {}

func (x *LabelDistributionProtocolP2MpIfStats) ProtoReflect() protoreflect.Message {
	mi := &file_ldp_p2mp_per_if_egress_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelDistributionProtocolP2MpIfStats.ProtoReflect.Descriptor instead.
func (*LabelDistributionProtocolP2MpIfStats) Descriptor() ([]byte, []int) {
	return file_ldp_p2mp_per_if_egress_proto_rawDescGZIP(), []int{2}
}

func (x *LabelDistributionProtocolP2MpIfStats) GetPackets() uint64 {
	if x != nil && x.Packets != nil {
		return *x.Packets
	}
	return 0
}

func (x *LabelDistributionProtocolP2MpIfStats) GetBytes() uint64 {
	if x != nil && x.Bytes != nil {
		return *x.Bytes
	}
	return 0
}

func (x *LabelDistributionProtocolP2MpIfStats) GetPacketRate() uint64 {
	if x != nil && x.PacketRate != nil {
		return *x.PacketRate
	}
	return 0
}

func (x *LabelDistributionProtocolP2MpIfStats) GetByteRate() uint64 {
	if x != nil && x.ByteRate != nil {
		return *x.ByteRate
	}
	return 0
}

var file_ldp_p2mp_per_if_egress_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*LdpP2MpPerIfEgress)(nil),
		Field:         156,
		Name:          "jnpr_ldp_p2mp_per_if_egress_ext",
		Tag:           "bytes,156,opt,name=jnpr_ldp_p2mp_per_if_egress_ext",
		Filename:      "ldp_p2mp_per_if_egress.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional LdpP2mpPerIfEgress jnpr_ldp_p2mp_per_if_egress_ext = 156;
	E_JnprLdpP2MpPerIfEgressExt = &file_ldp_p2mp_per_if_egress_proto_extTypes[0]
)

var File_ldp_p2mp_per_if_egress_proto protoreflect.FileDescriptor

var file_ldp_p2mp_per_if_egress_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x6c, 0x64, 0x70, 0x5f, 0x70, 0x32, 0x6d, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x69,
	0x66, 0x5f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x12, 0x4c, 0x64, 0x70, 0x50, 0x32, 0x6d, 0x70, 0x50, 0x65,
	0x72, 0x49, 0x66, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x70, 0x65, 0x72,
	0x5f, 0x69, 0x66, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x4c, 0x64, 0x70, 0x50, 0x32, 0x6d, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x49,
	0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xf8, 0x01, 0x0a, 0x16, 0x4c, 0x64, 0x70,
	0x50, 0x32, 0x6d, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x07, 0x69, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x06, 0x69, 0x66, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x0d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x50, 0x32, 0x6d, 0x70, 0x49, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0c, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x32, 0x6d, 0x70, 0x49,
	0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x24, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x50, 0x32, 0x6d, 0x70, 0x49, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82,
	0x40, 0x02, 0x18, 0x01, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x1b, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40,
	0x02, 0x18, 0x01, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x08, 0x62, 0x79,
	0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x3a, 0x71, 0x0a, 0x1f, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6c,
	0x64, 0x70, 0x5f, 0x70, 0x32, 0x6d, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x66, 0x5f, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69,
	0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x73, 0x18, 0x9c, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4c, 0x64, 0x70, 0x50,
	0x32, 0x6d, 0x70, 0x50, 0x65, 0x72, 0x49, 0x66, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x19,
	0x6a, 0x6e, 0x70, 0x72, 0x4c, 0x64, 0x70, 0x50, 0x32, 0x6d, 0x70, 0x50, 0x65, 0x72, 0x49, 0x66,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x6a,
	0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_ldp_p2mp_per_if_egress_proto_rawDescOnce sync.Once
	file_ldp_p2mp_per_if_egress_proto_rawDescData []byte
)

func file_ldp_p2mp_per_if_egress_proto_rawDescGZIP() []byte {
	file_ldp_p2mp_per_if_egress_proto_rawDescOnce.Do(func() {
		file_ldp_p2mp_per_if_egress_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ldp_p2mp_per_if_egress_proto_rawDesc), len(file_ldp_p2mp_per_if_egress_proto_rawDesc)))
	})
	return file_ldp_p2mp_per_if_egress_proto_rawDescData
}

var file_ldp_p2mp_per_if_egress_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ldp_p2mp_per_if_egress_proto_goTypes = []any{
	(*LdpP2MpPerIfEgress)(nil),                   // 0: LdpP2mpPerIfEgress
	(*LdpP2MpInterfaceRecord)(nil),               // 1: LdpP2mpInterfaceRecord
	(*LabelDistributionProtocolP2MpIfStats)(nil), // 2: LabelDistributionProtocolP2mpIfStats
	(*JuniperNetworksSensors)(nil),               // 3: JuniperNetworksSensors
}
var file_ldp_p2mp_per_if_egress_proto_depIdxs = []int32{
	1, // 0: LdpP2mpPerIfEgress.per_if_records:type_name -> LdpP2mpInterfaceRecord
	2, // 1: LdpP2mpInterfaceRecord.ingress_stats:type_name -> LabelDistributionProtocolP2mpIfStats
	2, // 2: LdpP2mpInterfaceRecord.egress_stats:type_name -> LabelDistributionProtocolP2mpIfStats
	3, // 3: jnpr_ldp_p2mp_per_if_egress_ext:extendee -> JuniperNetworksSensors
	0, // 4: jnpr_ldp_p2mp_per_if_egress_ext:type_name -> LdpP2mpPerIfEgress
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ldp_p2mp_per_if_egress_proto_init() }
func file_ldp_p2mp_per_if_egress_proto_init() {
	if File_ldp_p2mp_per_if_egress_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ldp_p2mp_per_if_egress_proto_rawDesc), len(file_ldp_p2mp_per_if_egress_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_ldp_p2mp_per_if_egress_proto_goTypes,
		DependencyIndexes: file_ldp_p2mp_per_if_egress_proto_depIdxs,
		MessageInfos:      file_ldp_p2mp_per_if_egress_proto_msgTypes,
		ExtensionInfos:    file_ldp_p2mp_per_if_egress_proto_extTypes,
	}.Build()
	File_ldp_p2mp_per_if_egress_proto = out.File
	file_ldp_p2mp_per_if_egress_proto_goTypes = nil
	file_ldp_p2mp_per_if_egress_proto_depIdxs = nil
}
