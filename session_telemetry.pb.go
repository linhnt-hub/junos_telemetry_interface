// Copyright (c) 2016, 2017 Juniper Networks, Inc.
// All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//
// Version 1.1
//

//
// Shibu Piriyath 2017-09-11
//
// This file defines the messages in Protocol Buffers used by
// the service pic sessions sensor. The top-level message is ServicesSession.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: session_telemetry.proto

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

type ServicesSession struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SvcsSessionInfo []*SessionInfo         `protobuf:"bytes,1,rep,name=svcs_session_info,json=svcsSessionInfo" json:"svcs_session_info,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ServicesSession) Reset() {
	*x = ServicesSession{}
	mi := &file_session_telemetry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServicesSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicesSession) ProtoMessage() {}

func (x *ServicesSession) ProtoReflect() protoreflect.Message {
	mi := &file_session_telemetry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicesSession.ProtoReflect.Descriptor instead.
func (*ServicesSession) Descriptor() ([]byte, []int) {
	return file_session_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *ServicesSession) GetSvcsSessionInfo() []*SessionInfo {
	if x != nil {
		return x.SvcsSessionInfo
	}
	return nil
}

type SessionInfo struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	SnFlowInfo        []*FlowInfo            `protobuf:"bytes,1,rep,name=sn_flow_info,json=snFlowInfo" json:"sn_flow_info,omitempty"`
	SnSvcSetName      *string                `protobuf:"bytes,2,opt,name=sn_svc_set_name,json=snSvcSetName" json:"sn_svc_set_name,omitempty"`
	SnSvcSetId        *uint64                `protobuf:"varint,3,opt,name=sn_svc_set_id,json=snSvcSetId" json:"sn_svc_set_id,omitempty"`
	SnId              *uint64                `protobuf:"varint,4,opt,name=sn_id,json=snId" json:"sn_id,omitempty"`
	SnFlags           *uint64                `protobuf:"varint,5,opt,name=sn_flags,json=snFlags" json:"sn_flags,omitempty"`
	SnAlgId           *uint64                `protobuf:"varint,6,opt,name=sn_alg_id,json=snAlgId" json:"sn_alg_id,omitempty"`
	SnRoutingPathType *uint64                `protobuf:"varint,7,opt,name=sn_routing_path_type,json=snRoutingPathType" json:"sn_routing_path_type,omitempty"`
	SnSessionTimeout  *uint64                `protobuf:"varint,8,opt,name=sn_session_timeout,json=snSessionTimeout" json:"sn_session_timeout,omitempty"`
	SnState           *uint64                `protobuf:"varint,9,opt,name=sn_state,json=snState" json:"sn_state,omitempty"`
	SnSwSessionId     *uint64                `protobuf:"varint,10,opt,name=sn_sw_session_id,json=snSwSessionId" json:"sn_sw_session_id,omitempty"`
	SnSvcId           *uint64                `protobuf:"varint,11,opt,name=sn_svc_id,json=snSvcId" json:"sn_svc_id,omitempty"`
	SnOffload         *bool                  `protobuf:"varint,12,opt,name=sn_offload,json=snOffload" json:"sn_offload,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *SessionInfo) Reset() {
	*x = SessionInfo{}
	mi := &file_session_telemetry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfo) ProtoMessage() {}

func (x *SessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_session_telemetry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfo.ProtoReflect.Descriptor instead.
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return file_session_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *SessionInfo) GetSnFlowInfo() []*FlowInfo {
	if x != nil {
		return x.SnFlowInfo
	}
	return nil
}

func (x *SessionInfo) GetSnSvcSetName() string {
	if x != nil && x.SnSvcSetName != nil {
		return *x.SnSvcSetName
	}
	return ""
}

func (x *SessionInfo) GetSnSvcSetId() uint64 {
	if x != nil && x.SnSvcSetId != nil {
		return *x.SnSvcSetId
	}
	return 0
}

func (x *SessionInfo) GetSnId() uint64 {
	if x != nil && x.SnId != nil {
		return *x.SnId
	}
	return 0
}

func (x *SessionInfo) GetSnFlags() uint64 {
	if x != nil && x.SnFlags != nil {
		return *x.SnFlags
	}
	return 0
}

func (x *SessionInfo) GetSnAlgId() uint64 {
	if x != nil && x.SnAlgId != nil {
		return *x.SnAlgId
	}
	return 0
}

func (x *SessionInfo) GetSnRoutingPathType() uint64 {
	if x != nil && x.SnRoutingPathType != nil {
		return *x.SnRoutingPathType
	}
	return 0
}

func (x *SessionInfo) GetSnSessionTimeout() uint64 {
	if x != nil && x.SnSessionTimeout != nil {
		return *x.SnSessionTimeout
	}
	return 0
}

func (x *SessionInfo) GetSnState() uint64 {
	if x != nil && x.SnState != nil {
		return *x.SnState
	}
	return 0
}

func (x *SessionInfo) GetSnSwSessionId() uint64 {
	if x != nil && x.SnSwSessionId != nil {
		return *x.SnSwSessionId
	}
	return 0
}

func (x *SessionInfo) GetSnSvcId() uint64 {
	if x != nil && x.SnSvcId != nil {
		return *x.SnSvcId
	}
	return 0
}

func (x *SessionInfo) GetSnOffload() bool {
	if x != nil && x.SnOffload != nil {
		return *x.SnOffload
	}
	return false
}

// --------------------------
type FlowInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	SrcAddr        *string                `protobuf:"bytes,1,opt,name=src_addr,json=srcAddr" json:"src_addr,omitempty"`
	DstAddr        *string                `protobuf:"bytes,2,opt,name=dst_addr,json=dstAddr" json:"dst_addr,omitempty"`
	FlowDirection  *uint32                `protobuf:"varint,3,opt,name=flow_direction,json=flowDirection" json:"flow_direction,omitempty"`
	SrcPort        *uint32                `protobuf:"varint,4,opt,name=src_port,json=srcPort" json:"src_port,omitempty"`
	DstPort        *uint32                `protobuf:"varint,5,opt,name=dst_port,json=dstPort" json:"dst_port,omitempty"`
	SrcIsV6        *bool                  `protobuf:"varint,6,opt,name=src_is_v6,json=srcIsV6" json:"src_is_v6,omitempty"`
	DstIsV6        *bool                  `protobuf:"varint,7,opt,name=dst_is_v6,json=dstIsV6" json:"dst_is_v6,omitempty"`
	FlowType       *uint32                `protobuf:"varint,8,opt,name=flow_type,json=flowType" json:"flow_type,omitempty"`
	FlowFlags      *uint32                `protobuf:"varint,9,opt,name=flow_flags,json=flowFlags" json:"flow_flags,omitempty"`
	IpProto        *uint32                `protobuf:"varint,10,opt,name=ip_proto,json=ipProto" json:"ip_proto,omitempty"`
	IdleTimeout    *uint64                `protobuf:"varint,11,opt,name=idle_timeout,json=idleTimeout" json:"idle_timeout,omitempty"`
	FlowPacketsIn  *uint64                `protobuf:"varint,12,opt,name=flow_packets_in,json=flowPacketsIn" json:"flow_packets_in,omitempty"`
	FlowPacketsOut *uint64                `protobuf:"varint,13,opt,name=flow_packets_out,json=flowPacketsOut" json:"flow_packets_out,omitempty"`
	FlowBytesIn    *uint64                `protobuf:"varint,14,opt,name=flow_bytes_in,json=flowBytesIn" json:"flow_bytes_in,omitempty"`
	FlowBytesOut   *uint64                `protobuf:"varint,15,opt,name=flow_bytes_out,json=flowBytesOut" json:"flow_bytes_out,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FlowInfo) Reset() {
	*x = FlowInfo{}
	mi := &file_session_telemetry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowInfo) ProtoMessage() {}

func (x *FlowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_session_telemetry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowInfo.ProtoReflect.Descriptor instead.
func (*FlowInfo) Descriptor() ([]byte, []int) {
	return file_session_telemetry_proto_rawDescGZIP(), []int{2}
}

func (x *FlowInfo) GetSrcAddr() string {
	if x != nil && x.SrcAddr != nil {
		return *x.SrcAddr
	}
	return ""
}

func (x *FlowInfo) GetDstAddr() string {
	if x != nil && x.DstAddr != nil {
		return *x.DstAddr
	}
	return ""
}

func (x *FlowInfo) GetFlowDirection() uint32 {
	if x != nil && x.FlowDirection != nil {
		return *x.FlowDirection
	}
	return 0
}

func (x *FlowInfo) GetSrcPort() uint32 {
	if x != nil && x.SrcPort != nil {
		return *x.SrcPort
	}
	return 0
}

func (x *FlowInfo) GetDstPort() uint32 {
	if x != nil && x.DstPort != nil {
		return *x.DstPort
	}
	return 0
}

func (x *FlowInfo) GetSrcIsV6() bool {
	if x != nil && x.SrcIsV6 != nil {
		return *x.SrcIsV6
	}
	return false
}

func (x *FlowInfo) GetDstIsV6() bool {
	if x != nil && x.DstIsV6 != nil {
		return *x.DstIsV6
	}
	return false
}

func (x *FlowInfo) GetFlowType() uint32 {
	if x != nil && x.FlowType != nil {
		return *x.FlowType
	}
	return 0
}

func (x *FlowInfo) GetFlowFlags() uint32 {
	if x != nil && x.FlowFlags != nil {
		return *x.FlowFlags
	}
	return 0
}

func (x *FlowInfo) GetIpProto() uint32 {
	if x != nil && x.IpProto != nil {
		return *x.IpProto
	}
	return 0
}

func (x *FlowInfo) GetIdleTimeout() uint64 {
	if x != nil && x.IdleTimeout != nil {
		return *x.IdleTimeout
	}
	return 0
}

func (x *FlowInfo) GetFlowPacketsIn() uint64 {
	if x != nil && x.FlowPacketsIn != nil {
		return *x.FlowPacketsIn
	}
	return 0
}

func (x *FlowInfo) GetFlowPacketsOut() uint64 {
	if x != nil && x.FlowPacketsOut != nil {
		return *x.FlowPacketsOut
	}
	return 0
}

func (x *FlowInfo) GetFlowBytesIn() uint64 {
	if x != nil && x.FlowBytesIn != nil {
		return *x.FlowBytesIn
	}
	return 0
}

func (x *FlowInfo) GetFlowBytesOut() uint64 {
	if x != nil && x.FlowBytesOut != nil {
		return *x.FlowBytesOut
	}
	return 0
}

var file_session_telemetry_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*ServicesSession)(nil),
		Field:         79,
		Name:          "jnprScvsSessionExt",
		Tag:           "bytes,79,opt,name=jnprScvsSessionExt",
		Filename:      "session_telemetry.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional ServicesSession jnprScvsSessionExt = 79;
	E_JnprScvsSessionExt = &file_session_telemetry_proto_extTypes[0]
)

var File_session_telemetry_proto protoreflect.FileDescriptor

var file_session_telemetry_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x11, 0x73, 0x76, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x73, 0x76, 0x63, 0x73,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb5, 0x03, 0x0a, 0x0b,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x0c, 0x73,
	0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x6e,
	0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x6e, 0x5f, 0x73,
	0x76, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x73, 0x6e, 0x53, 0x76, 0x63, 0x53,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0d, 0x73, 0x6e, 0x5f, 0x73, 0x76, 0x63,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73,
	0x6e, 0x53, 0x76, 0x63, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x73, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x6e, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x73, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x09, 0x73, 0x6e, 0x5f,
	0x61, 0x6c, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x6e,
	0x41, 0x6c, 0x67, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x11, 0x73, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6e, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x10, 0x73, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x0a, 0x10, 0x73, 0x6e, 0x5f, 0x73, 0x77, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x6e, 0x53, 0x77, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x73, 0x6e, 0x5f, 0x73,
	0x76, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x6e, 0x53,
	0x76, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x6e, 0x4f, 0x66, 0x66, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0xf9, 0x03, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x07, 0x73, 0x72, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x20, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x07, 0x64, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x6c,
	0x6f, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1a, 0x0a, 0x09, 0x73, 0x72, 0x63, 0x5f, 0x69, 0x73, 0x5f, 0x76, 0x36, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x72, 0x63, 0x49, 0x73, 0x56, 0x36, 0x12, 0x1a, 0x0a,
	0x09, 0x64, 0x73, 0x74, 0x5f, 0x69, 0x73, 0x5f, 0x76, 0x36, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x64, 0x73, 0x74, 0x49, 0x73, 0x56, 0x36, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x6c,
	0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x6c, 0x6f, 0x77,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x66, 0x6c,
	0x6f, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x66, 0x6c,
	0x6f, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x3a,
	0x59, 0x0a, 0x12, 0x6a, 0x6e, 0x70, 0x72, 0x53, 0x63, 0x76, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x4f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6a, 0x6e, 0x70, 0x72, 0x53, 0x63, 0x76, 0x73,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f,
	0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_session_telemetry_proto_rawDescOnce sync.Once
	file_session_telemetry_proto_rawDescData []byte
)

func file_session_telemetry_proto_rawDescGZIP() []byte {
	file_session_telemetry_proto_rawDescOnce.Do(func() {
		file_session_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_session_telemetry_proto_rawDesc), len(file_session_telemetry_proto_rawDesc)))
	})
	return file_session_telemetry_proto_rawDescData
}

var file_session_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_session_telemetry_proto_goTypes = []any{
	(*ServicesSession)(nil),        // 0: ServicesSession
	(*SessionInfo)(nil),            // 1: SessionInfo
	(*FlowInfo)(nil),               // 2: FlowInfo
	(*JuniperNetworksSensors)(nil), // 3: JuniperNetworksSensors
}
var file_session_telemetry_proto_depIdxs = []int32{
	1, // 0: ServicesSession.svcs_session_info:type_name -> SessionInfo
	2, // 1: SessionInfo.sn_flow_info:type_name -> FlowInfo
	3, // 2: jnprScvsSessionExt:extendee -> JuniperNetworksSensors
	0, // 3: jnprScvsSessionExt:type_name -> ServicesSession
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_session_telemetry_proto_init() }
func file_session_telemetry_proto_init() {
	if File_session_telemetry_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_session_telemetry_proto_rawDesc), len(file_session_telemetry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_session_telemetry_proto_goTypes,
		DependencyIndexes: file_session_telemetry_proto_depIdxs,
		MessageInfos:      file_session_telemetry_proto_msgTypes,
		ExtensionInfos:    file_session_telemetry_proto_extTypes,
	}.Build()
	File_session_telemetry_proto = out.File
	file_session_telemetry_proto_goTypes = nil
	file_session_telemetry_proto_depIdxs = nil
}
