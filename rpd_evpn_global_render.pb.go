// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: rpd_evpn_global_render.proto

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

type JunosEvpn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Evpn          *JunosEvpnEvpnType     `protobuf:"bytes,171,opt,name=evpn" json:"evpn,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEvpn) Reset() {
	*x = JunosEvpn{}
	mi := &file_rpd_evpn_global_render_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEvpn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEvpn) ProtoMessage() {}

func (x *JunosEvpn) ProtoReflect() protoreflect.Message {
	mi := &file_rpd_evpn_global_render_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEvpn.ProtoReflect.Descriptor instead.
func (*JunosEvpn) Descriptor() ([]byte, []int) {
	return file_rpd_evpn_global_render_proto_rawDescGZIP(), []int{0}
}

func (x *JunosEvpn) GetEvpn() *JunosEvpnEvpnType {
	if x != nil {
		return x.Evpn
	}
	return nil
}

type JunosEvpnEvpnType struct {
	state              protoimpl.MessageState                   `protogen:"open.v1"`
	EvpnSmetForwarding *JunosEvpnEvpnTypeEvpnSmetForwardingType `protobuf:"bytes,171,opt,name=evpn_smet_forwarding,json=evpnSmetForwarding" json:"evpn_smet_forwarding,omitempty"`
	L3Context          []*JunosEvpnEvpnTypeL3ContextList        `protobuf:"bytes,180,rep,name=l3_context,json=l3Context" json:"l3_context,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *JunosEvpnEvpnType) Reset() {
	*x = JunosEvpnEvpnType{}
	mi := &file_rpd_evpn_global_render_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEvpnEvpnType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEvpnEvpnType) ProtoMessage() {}

func (x *JunosEvpnEvpnType) ProtoReflect() protoreflect.Message {
	mi := &file_rpd_evpn_global_render_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEvpnEvpnType.ProtoReflect.Descriptor instead.
func (*JunosEvpnEvpnType) Descriptor() ([]byte, []int) {
	return file_rpd_evpn_global_render_proto_rawDescGZIP(), []int{0, 0}
}

func (x *JunosEvpnEvpnType) GetEvpnSmetForwarding() *JunosEvpnEvpnTypeEvpnSmetForwardingType {
	if x != nil {
		return x.EvpnSmetForwarding
	}
	return nil
}

func (x *JunosEvpnEvpnType) GetL3Context() []*JunosEvpnEvpnTypeL3ContextList {
	if x != nil {
		return x.L3Context
	}
	return nil
}

type JunosEvpnEvpnTypeEvpnSmetForwardingType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enabled       *bool                  `protobuf:"varint,71,opt,name=enabled" json:"enabled,omitempty"`
	NexthopLimit  *uint32                `protobuf:"varint,72,opt,name=nexthop_limit,json=nexthopLimit" json:"nexthop_limit,omitempty"`
	NexthopUsage  *uint32                `protobuf:"varint,73,opt,name=nexthop_usage,json=nexthopUsage" json:"nexthop_usage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEvpnEvpnTypeEvpnSmetForwardingType) Reset() {
	*x = JunosEvpnEvpnTypeEvpnSmetForwardingType{}
	mi := &file_rpd_evpn_global_render_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEvpnEvpnTypeEvpnSmetForwardingType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEvpnEvpnTypeEvpnSmetForwardingType) ProtoMessage() {}

func (x *JunosEvpnEvpnTypeEvpnSmetForwardingType) ProtoReflect() protoreflect.Message {
	mi := &file_rpd_evpn_global_render_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEvpnEvpnTypeEvpnSmetForwardingType.ProtoReflect.Descriptor instead.
func (*JunosEvpnEvpnTypeEvpnSmetForwardingType) Descriptor() ([]byte, []int) {
	return file_rpd_evpn_global_render_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *JunosEvpnEvpnTypeEvpnSmetForwardingType) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *JunosEvpnEvpnTypeEvpnSmetForwardingType) GetNexthopLimit() uint32 {
	if x != nil && x.NexthopLimit != nil {
		return *x.NexthopLimit
	}
	return 0
}

func (x *JunosEvpnEvpnTypeEvpnSmetForwardingType) GetNexthopUsage() uint32 {
	if x != nil && x.NexthopUsage != nil {
		return *x.NexthopUsage
	}
	return 0
}

type JunosEvpnEvpnTypeL3ContextList struct {
	state                protoimpl.MessageState                                `protogen:"open.v1"`
	ContextName          *string                                               `protobuf:"bytes,71,opt,name=context_name,json=contextName" json:"context_name,omitempty"`
	Encapsulation        *string                                               `protobuf:"bytes,72,opt,name=encapsulation" json:"encapsulation,omitempty"`
	AdvertisementMode    *string                                               `protobuf:"bytes,73,opt,name=advertisement_mode,json=advertisementMode" json:"advertisement_mode,omitempty"`
	MulticastRoutingMode *string                                               `protobuf:"bytes,74,opt,name=multicast_routing_mode,json=multicastRoutingMode" json:"multicast_routing_mode,omitempty"`
	IpPrefixDatabase     []*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList `protobuf:"bytes,170,rep,name=ip_prefix_database,json=ipPrefixDatabase" json:"ip_prefix_database,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *JunosEvpnEvpnTypeL3ContextList) Reset() {
	*x = JunosEvpnEvpnTypeL3ContextList{}
	mi := &file_rpd_evpn_global_render_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEvpnEvpnTypeL3ContextList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEvpnEvpnTypeL3ContextList) ProtoMessage() {}

func (x *JunosEvpnEvpnTypeL3ContextList) ProtoReflect() protoreflect.Message {
	mi := &file_rpd_evpn_global_render_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEvpnEvpnTypeL3ContextList.ProtoReflect.Descriptor instead.
func (*JunosEvpnEvpnTypeL3ContextList) Descriptor() ([]byte, []int) {
	return file_rpd_evpn_global_render_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *JunosEvpnEvpnTypeL3ContextList) GetContextName() string {
	if x != nil && x.ContextName != nil {
		return *x.ContextName
	}
	return ""
}

func (x *JunosEvpnEvpnTypeL3ContextList) GetEncapsulation() string {
	if x != nil && x.Encapsulation != nil {
		return *x.Encapsulation
	}
	return ""
}

func (x *JunosEvpnEvpnTypeL3ContextList) GetAdvertisementMode() string {
	if x != nil && x.AdvertisementMode != nil {
		return *x.AdvertisementMode
	}
	return ""
}

func (x *JunosEvpnEvpnTypeL3ContextList) GetMulticastRoutingMode() string {
	if x != nil && x.MulticastRoutingMode != nil {
		return *x.MulticastRoutingMode
	}
	return ""
}

func (x *JunosEvpnEvpnTypeL3ContextList) GetIpPrefixDatabase() []*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList {
	if x != nil {
		return x.IpPrefixDatabase
	}
	return nil
}

type JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IpPrefix      *string                `protobuf:"bytes,71,opt,name=ip_prefix,json=ipPrefix" json:"ip_prefix,omitempty"`
	RouteStatus   *string                `protobuf:"bytes,72,opt,name=route_status,json=routeStatus" json:"route_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) Reset() {
	*x = JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList{}
	mi := &file_rpd_evpn_global_render_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) ProtoMessage() {}

func (x *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) ProtoReflect() protoreflect.Message {
	mi := &file_rpd_evpn_global_render_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList.ProtoReflect.Descriptor instead.
func (*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) Descriptor() ([]byte, []int) {
	return file_rpd_evpn_global_render_proto_rawDescGZIP(), []int{0, 0, 1, 0}
}

func (x *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) GetIpPrefix() string {
	if x != nil && x.IpPrefix != nil {
		return *x.IpPrefix
	}
	return ""
}

func (x *JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList) GetRouteStatus() string {
	if x != nil && x.RouteStatus != nil {
		return *x.RouteStatus
	}
	return ""
}

var file_rpd_evpn_global_render_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*JunosEvpn)(nil),
		Field:         113,
		Name:          "jnpr_junos_evpn_ext",
		Tag:           "bytes,113,opt,name=jnpr_junos_evpn_ext",
		Filename:      "rpd_evpn_global_render.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional junos_evpn jnpr_junos_evpn_ext = 113;
	E_JnprJunosEvpnExt = &file_rpd_evpn_global_render_proto_extTypes[0]
)

var File_rpd_evpn_global_render_proto protoreflect.FileDescriptor

var file_rpd_evpn_global_render_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x72, 0x70, 0x64, 0x5f, 0x65, 0x76, 0x70, 0x6e, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x06, 0x0a, 0x0a, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76,
	0x70, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x65, 0x76, 0x70, 0x6e, 0x18, 0xab, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x70, 0x6e, 0x2e, 0x65,
	0x76, 0x70, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x65, 0x76, 0x70, 0x6e, 0x1a, 0xd0,
	0x05, 0x0a, 0x09, 0x65, 0x76, 0x70, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x62, 0x0a, 0x14,
	0x65, 0x76, 0x70, 0x6e, 0x5f, 0x73, 0x6d, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0xab, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6a, 0x75,
	0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x70, 0x6e, 0x2e, 0x65, 0x76, 0x70, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x65, 0x76, 0x70, 0x6e, 0x5f, 0x73, 0x6d, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x12, 0x65, 0x76,
	0x70, 0x6e, 0x53, 0x6d, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x45, 0x0a, 0x0a, 0x6c, 0x33, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0xb4,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76,
	0x70, 0x6e, 0x2e, 0x65, 0x76, 0x70, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6c, 0x33, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x6c, 0x33,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x7f, 0x0a, 0x19, 0x65, 0x76, 0x70, 0x6e, 0x5f,
	0x73, 0x6d, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x47, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x48, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x49, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74,
	0x68, 0x6f, 0x70, 0x55, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x96, 0x03, 0x0a, 0x0f, 0x6c, 0x33, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x47, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x48, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x6e, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x49, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x4a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x6c, 0x0a, 0x12, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0xaa, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x65, 0x76, 0x70, 0x6e, 0x2e, 0x65, 0x76, 0x70, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6c, 0x33, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x10, 0x69,
	0x70, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a,
	0x60, 0x0a, 0x17, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x09, 0x69, 0x70,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82,
	0x40, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69, 0x70, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x48,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x3a, 0x53, 0x0a, 0x13, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f,
	0x65, 0x76, 0x70, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69, 0x70,
	0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x73, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f,
	0x65, 0x76, 0x70, 0x6e, 0x52, 0x10, 0x6a, 0x6e, 0x70, 0x72, 0x4a, 0x75, 0x6e, 0x6f, 0x73, 0x45,
	0x76, 0x70, 0x6e, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f,
	0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_rpd_evpn_global_render_proto_rawDescOnce sync.Once
	file_rpd_evpn_global_render_proto_rawDescData []byte
)

func file_rpd_evpn_global_render_proto_rawDescGZIP() []byte {
	file_rpd_evpn_global_render_proto_rawDescOnce.Do(func() {
		file_rpd_evpn_global_render_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpd_evpn_global_render_proto_rawDesc), len(file_rpd_evpn_global_render_proto_rawDesc)))
	})
	return file_rpd_evpn_global_render_proto_rawDescData
}

var file_rpd_evpn_global_render_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpd_evpn_global_render_proto_goTypes = []any{
	(*JunosEvpn)(nil),                                          // 0: junos_evpn
	(*JunosEvpnEvpnType)(nil),                                  // 1: junos_evpn.evpn_type
	(*JunosEvpnEvpnTypeEvpnSmetForwardingType)(nil),            // 2: junos_evpn.evpn_type.evpn_smet_forwarding_type
	(*JunosEvpnEvpnTypeL3ContextList)(nil),                     // 3: junos_evpn.evpn_type.l3_context_list
	(*JunosEvpnEvpnTypeL3ContextListIpPrefixDatabaseList)(nil), // 4: junos_evpn.evpn_type.l3_context_list.ip_prefix_database_list
	(*JuniperNetworksSensors)(nil),                             // 5: JuniperNetworksSensors
}
var file_rpd_evpn_global_render_proto_depIdxs = []int32{
	1, // 0: junos_evpn.evpn:type_name -> junos_evpn.evpn_type
	2, // 1: junos_evpn.evpn_type.evpn_smet_forwarding:type_name -> junos_evpn.evpn_type.evpn_smet_forwarding_type
	3, // 2: junos_evpn.evpn_type.l3_context:type_name -> junos_evpn.evpn_type.l3_context_list
	4, // 3: junos_evpn.evpn_type.l3_context_list.ip_prefix_database:type_name -> junos_evpn.evpn_type.l3_context_list.ip_prefix_database_list
	5, // 4: jnpr_junos_evpn_ext:extendee -> JuniperNetworksSensors
	0, // 5: jnpr_junos_evpn_ext:type_name -> junos_evpn
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	5, // [5:6] is the sub-list for extension type_name
	4, // [4:5] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpd_evpn_global_render_proto_init() }
func file_rpd_evpn_global_render_proto_init() {
	if File_rpd_evpn_global_render_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpd_evpn_global_render_proto_rawDesc), len(file_rpd_evpn_global_render_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_rpd_evpn_global_render_proto_goTypes,
		DependencyIndexes: file_rpd_evpn_global_render_proto_depIdxs,
		MessageInfos:      file_rpd_evpn_global_render_proto_msgTypes,
		ExtensionInfos:    file_rpd_evpn_global_render_proto_extTypes,
	}.Build()
	File_rpd_evpn_global_render_proto = out.File
	file_rpd_evpn_global_render_proto_goTypes = nil
	file_rpd_evpn_global_render_proto_depIdxs = nil
}
