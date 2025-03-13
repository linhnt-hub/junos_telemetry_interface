// Copyright (c) 2015, 2016 Juniper Networks, Inc.
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
// the IPSec sensor. The top-level message is IPsecVPN.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: ipsec_telemetry.proto

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

type IPsecVPN struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	IpsecGlobalInfo []*IPsecGlobalInfo     `protobuf:"bytes,1,rep,name=ipsec_global_info,json=ipsecGlobalInfo" json:"ipsec_global_info,omitempty"`
	IpsecSvcsetInfo []*IPsecPerSvcsetInfo  `protobuf:"bytes,2,rep,name=ipsec_svcset_info,json=ipsecSvcsetInfo" json:"ipsec_svcset_info,omitempty"`
	IpsecTunnelInfo []*IPsecPerTunnelInfo  `protobuf:"bytes,3,rep,name=ipsec_tunnel_info,json=ipsecTunnelInfo" json:"ipsec_tunnel_info,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *IPsecVPN) Reset() {
	*x = IPsecVPN{}
	mi := &file_ipsec_telemetry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPsecVPN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPsecVPN) ProtoMessage() {}

func (x *IPsecVPN) ProtoReflect() protoreflect.Message {
	mi := &file_ipsec_telemetry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPsecVPN.ProtoReflect.Descriptor instead.
func (*IPsecVPN) Descriptor() ([]byte, []int) {
	return file_ipsec_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *IPsecVPN) GetIpsecGlobalInfo() []*IPsecGlobalInfo {
	if x != nil {
		return x.IpsecGlobalInfo
	}
	return nil
}

func (x *IPsecVPN) GetIpsecSvcsetInfo() []*IPsecPerSvcsetInfo {
	if x != nil {
		return x.IpsecSvcsetInfo
	}
	return nil
}

func (x *IPsecVPN) GetIpsecTunnelInfo() []*IPsecPerTunnelInfo {
	if x != nil {
		return x.IpsecTunnelInfo
	}
	return nil
}

type IPsecGlobalInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// RE-Pconn Connect Retry
	RePconnConnect *uint64 `protobuf:"varint,1,opt,name=re_pconn_connect,json=rePconnConnect" json:"re_pconn_connect,omitempty"`
	// RE-Pconn Status
	PconnStatus *uint64 `protobuf:"varint,2,opt,name=pconn_status,json=pconnStatus" json:"pconn_status,omitempty"`
	// Request Enq Succ
	SaTriggerEnqSuccess *uint64 `protobuf:"varint,3,opt,name=sa_trigger_enq_success,json=saTriggerEnqSuccess" json:"sa_trigger_enq_success,omitempty"`
	// Request Enq Fail
	SaTriggerEnqFail *uint64 `protobuf:"varint,4,opt,name=sa_trigger_enq_fail,json=saTriggerEnqFail" json:"sa_trigger_enq_fail,omitempty"`
	// Retry Enq Succ
	SaTriggerRetrySuccess *uint64 `protobuf:"varint,5,opt,name=sa_trigger_retry_success,json=saTriggerRetrySuccess" json:"sa_trigger_retry_success,omitempty"`
	// Retry Enq Fail
	SaTriggerRetryFail *uint64 `protobuf:"varint,6,opt,name=sa_trigger_retry_fail,json=saTriggerRetryFail" json:"sa_trigger_retry_fail,omitempty"`
	// Trigger Send Succ
	SaTriggerSent *uint64 `protobuf:"varint,7,opt,name=sa_trigger_sent,json=saTriggerSent" json:"sa_trigger_sent,omitempty"`
	// Trigger Send Fail
	SaTriggerFail *uint64 `protobuf:"varint,8,opt,name=sa_trigger_fail,json=saTriggerFail" json:"sa_trigger_fail,omitempty"`
	// Trigger Alloc
	SaTriggerAlloc *uint64 `protobuf:"varint,9,opt,name=sa_trigger_alloc,json=saTriggerAlloc" json:"sa_trigger_alloc,omitempty"`
	// Alloc Fail
	SaTriggerAllocFail *uint64 `protobuf:"varint,10,opt,name=sa_trigger_alloc_fail,json=saTriggerAllocFail" json:"sa_trigger_alloc_fail,omitempty"`
	// Trigger Free
	SaTriggerFree *uint64 `protobuf:"varint,11,opt,name=sa_trigger_free,json=saTriggerFree" json:"sa_trigger_free,omitempty"`
	// Outstanding Trig Msg in Queue
	SaTrigEnqCount *uint64 `protobuf:"varint,12,opt,name=sa_trig_enq_count,json=saTrigEnqCount" json:"sa_trig_enq_count,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IPsecGlobalInfo) Reset() {
	*x = IPsecGlobalInfo{}
	mi := &file_ipsec_telemetry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPsecGlobalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPsecGlobalInfo) ProtoMessage() {}

func (x *IPsecGlobalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ipsec_telemetry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPsecGlobalInfo.ProtoReflect.Descriptor instead.
func (*IPsecGlobalInfo) Descriptor() ([]byte, []int) {
	return file_ipsec_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *IPsecGlobalInfo) GetRePconnConnect() uint64 {
	if x != nil && x.RePconnConnect != nil {
		return *x.RePconnConnect
	}
	return 0
}

func (x *IPsecGlobalInfo) GetPconnStatus() uint64 {
	if x != nil && x.PconnStatus != nil {
		return *x.PconnStatus
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerEnqSuccess() uint64 {
	if x != nil && x.SaTriggerEnqSuccess != nil {
		return *x.SaTriggerEnqSuccess
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerEnqFail() uint64 {
	if x != nil && x.SaTriggerEnqFail != nil {
		return *x.SaTriggerEnqFail
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerRetrySuccess() uint64 {
	if x != nil && x.SaTriggerRetrySuccess != nil {
		return *x.SaTriggerRetrySuccess
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerRetryFail() uint64 {
	if x != nil && x.SaTriggerRetryFail != nil {
		return *x.SaTriggerRetryFail
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerSent() uint64 {
	if x != nil && x.SaTriggerSent != nil {
		return *x.SaTriggerSent
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerFail() uint64 {
	if x != nil && x.SaTriggerFail != nil {
		return *x.SaTriggerFail
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerAlloc() uint64 {
	if x != nil && x.SaTriggerAlloc != nil {
		return *x.SaTriggerAlloc
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerAllocFail() uint64 {
	if x != nil && x.SaTriggerAllocFail != nil {
		return *x.SaTriggerAllocFail
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTriggerFree() uint64 {
	if x != nil && x.SaTriggerFree != nil {
		return *x.SaTriggerFree
	}
	return 0
}

func (x *IPsecGlobalInfo) GetSaTrigEnqCount() uint64 {
	if x != nil && x.SaTrigEnqCount != nil {
		return *x.SaTrigEnqCount
	}
	return 0
}

type IPsecPerSvcsetInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Service Set ID
	SvcsetId *uint64 `protobuf:"varint,1,req,name=svcset_id,json=svcsetId" json:"svcset_id,omitempty"`
	// Rule Lookup Fail
	RuleLookupFailed *uint64 `protobuf:"varint,2,opt,name=rule_lookup_failed,json=ruleLookupFailed" json:"rule_lookup_failed,omitempty"`
	// SA Lookup Fail
	SaLookupFailed *uint64 `protobuf:"varint,3,opt,name=sa_lookup_failed,json=saLookupFailed" json:"sa_lookup_failed,omitempty"`
	// Pre Frags
	ExceedsTunnelMtu *uint64 `protobuf:"varint,4,opt,name=exceeds_tunnel_mtu,json=exceedsTunnelMtu" json:"exceeds_tunnel_mtu,omitempty"`
	// Clear Pkts Recvd for encap
	ClearPktReceived *uint64 `protobuf:"varint,5,opt,name=clear_pkt_received,json=clearPktReceived" json:"clear_pkt_received,omitempty"`
	// ESP Pkts Recvd
	EspPktReceived *uint64 `protobuf:"varint,6,opt,name=esp_pkt_received,json=espPktReceived" json:"esp_pkt_received,omitempty"`
	// Encap Pkts From Crypto
	EncapCallback *uint64 `protobuf:"varint,7,opt,name=encap_callback,json=encapCallback" json:"encap_callback,omitempty"`
	// Decap Pkts From Crypto
	DecapCallback *uint64 `protobuf:"varint,8,opt,name=decap_callback,json=decapCallback" json:"decap_callback,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPsecPerSvcsetInfo) Reset() {
	*x = IPsecPerSvcsetInfo{}
	mi := &file_ipsec_telemetry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPsecPerSvcsetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPsecPerSvcsetInfo) ProtoMessage() {}

func (x *IPsecPerSvcsetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ipsec_telemetry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPsecPerSvcsetInfo.ProtoReflect.Descriptor instead.
func (*IPsecPerSvcsetInfo) Descriptor() ([]byte, []int) {
	return file_ipsec_telemetry_proto_rawDescGZIP(), []int{2}
}

func (x *IPsecPerSvcsetInfo) GetSvcsetId() uint64 {
	if x != nil && x.SvcsetId != nil {
		return *x.SvcsetId
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetRuleLookupFailed() uint64 {
	if x != nil && x.RuleLookupFailed != nil {
		return *x.RuleLookupFailed
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetSaLookupFailed() uint64 {
	if x != nil && x.SaLookupFailed != nil {
		return *x.SaLookupFailed
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetExceedsTunnelMtu() uint64 {
	if x != nil && x.ExceedsTunnelMtu != nil {
		return *x.ExceedsTunnelMtu
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetClearPktReceived() uint64 {
	if x != nil && x.ClearPktReceived != nil {
		return *x.ClearPktReceived
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetEspPktReceived() uint64 {
	if x != nil && x.EspPktReceived != nil {
		return *x.EspPktReceived
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetEncapCallback() uint64 {
	if x != nil && x.EncapCallback != nil {
		return *x.EncapCallback
	}
	return 0
}

func (x *IPsecPerSvcsetInfo) GetDecapCallback() uint64 {
	if x != nil && x.DecapCallback != nil {
		return *x.DecapCallback
	}
	return 0
}

type IPsecPerTunnelInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Tunnel IDs on PIC
	TunnelId *uint64 `protobuf:"varint,1,req,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Seq Num Zero
	EspRplzero *uint64 `protobuf:"varint,2,opt,name=esp_rplzero,json=espRplzero" json:"esp_rplzero,omitempty"`
	// Bad Pad
	IpsecBadHeaders *uint64 `protobuf:"varint,3,opt,name=ipsec_bad_headers,json=ipsecBadHeaders" json:"ipsec_bad_headers,omitempty"`
	// ESP Tail Err
	EspBadTrailers *uint64 `protobuf:"varint,4,opt,name=esp_bad_trailers,json=espBadTrailers" json:"esp_bad_trailers,omitempty"`
	// Next Proto Err
	DecapNxtProtoErr *uint64 `protobuf:"varint,5,opt,name=decap_nxt_proto_err,json=decapNxtProtoErr" json:"decap_nxt_proto_err,omitempty"`
	// Inner Len Err
	DecapInnerLenErr *uint64 `protobuf:"varint,6,opt,name=decap_inner_len_err,json=decapInnerLenErr" json:"decap_inner_len_err,omitempty"`
	// Outer Hdr Err
	DecapHdrErr *uint64 `protobuf:"varint,7,opt,name=decap_hdr_err,json=decapHdrErr" json:"decap_hdr_err,omitempty"`
	// Inner Saddr Err
	DecapInnerSaddrErr *uint64 `protobuf:"varint,8,opt,name=decap_inner_saddr_err,json=decapInnerSaddrErr" json:"decap_inner_saddr_err,omitempty"`
	// Inner Daddr Err
	DecapInnerDaddrErr *uint64 `protobuf:"varint,9,opt,name=decap_inner_daddr_err,json=decapInnerDaddrErr" json:"decap_inner_daddr_err,omitempty"`
	// Sn Alloc Fail
	DecapSnAllocFail *uint64 `protobuf:"varint,10,opt,name=decap_sn_alloc_fail,json=decapSnAllocFail" json:"decap_sn_alloc_fail,omitempty"`
	// Sn Ext Fail
	DecapSnExtFail *uint64 `protobuf:"varint,11,opt,name=decap_sn_ext_fail,json=decapSnExtFail" json:"decap_sn_ext_fail,omitempty"`
	// Auth Fail
	EspAuthFailed *uint64 `protobuf:"varint,12,opt,name=esp_auth_failed,json=espAuthFailed" json:"esp_auth_failed,omitempty"`
	// Reinject Fail
	DecapReinjectFail *uint64 `protobuf:"varint,13,opt,name=decap_reinject_fail,json=decapReinjectFail" json:"decap_reinject_fail,omitempty"`
	// Session Transient Drop
	DecapSnTransientDrop *uint64 `protobuf:"varint,14,opt,name=decap_sn_transient_drop,json=decapSnTransientDrop" json:"decap_sn_transient_drop,omitempty"`
	// Replay Before Win
	EspRplbeforewindow *uint64 `protobuf:"varint,15,opt,name=esp_rplbeforewindow,json=espRplbeforewindow" json:"esp_rplbeforewindow,omitempty"`
	// Replayed Pkts
	EspRplduplicate *uint64 `protobuf:"varint,16,opt,name=esp_rplduplicate,json=espRplduplicate" json:"esp_rplduplicate,omitempty"`
	// Encrypted Bytes
	EspProtectedBytesSent *uint64 `protobuf:"varint,17,opt,name=esp_protected_bytes_sent,json=espProtectedBytesSent" json:"esp_protected_bytes_sent,omitempty"`
	// Decrypted Bytes
	EspProtectedBytesRecvd *uint64 `protobuf:"varint,18,opt,name=esp_protected_bytes_recvd,json=espProtectedBytesRecvd" json:"esp_protected_bytes_recvd,omitempty"`
	// Encrypted Packets
	Encrypts *uint64 `protobuf:"varint,19,opt,name=encrypts" json:"encrypts,omitempty"`
	// Decrypted Packets
	Decrypts      *uint64 `protobuf:"varint,20,opt,name=decrypts" json:"decrypts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPsecPerTunnelInfo) Reset() {
	*x = IPsecPerTunnelInfo{}
	mi := &file_ipsec_telemetry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPsecPerTunnelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPsecPerTunnelInfo) ProtoMessage() {}

func (x *IPsecPerTunnelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ipsec_telemetry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPsecPerTunnelInfo.ProtoReflect.Descriptor instead.
func (*IPsecPerTunnelInfo) Descriptor() ([]byte, []int) {
	return file_ipsec_telemetry_proto_rawDescGZIP(), []int{3}
}

func (x *IPsecPerTunnelInfo) GetTunnelId() uint64 {
	if x != nil && x.TunnelId != nil {
		return *x.TunnelId
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspRplzero() uint64 {
	if x != nil && x.EspRplzero != nil {
		return *x.EspRplzero
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetIpsecBadHeaders() uint64 {
	if x != nil && x.IpsecBadHeaders != nil {
		return *x.IpsecBadHeaders
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspBadTrailers() uint64 {
	if x != nil && x.EspBadTrailers != nil {
		return *x.EspBadTrailers
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapNxtProtoErr() uint64 {
	if x != nil && x.DecapNxtProtoErr != nil {
		return *x.DecapNxtProtoErr
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapInnerLenErr() uint64 {
	if x != nil && x.DecapInnerLenErr != nil {
		return *x.DecapInnerLenErr
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapHdrErr() uint64 {
	if x != nil && x.DecapHdrErr != nil {
		return *x.DecapHdrErr
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapInnerSaddrErr() uint64 {
	if x != nil && x.DecapInnerSaddrErr != nil {
		return *x.DecapInnerSaddrErr
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapInnerDaddrErr() uint64 {
	if x != nil && x.DecapInnerDaddrErr != nil {
		return *x.DecapInnerDaddrErr
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapSnAllocFail() uint64 {
	if x != nil && x.DecapSnAllocFail != nil {
		return *x.DecapSnAllocFail
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapSnExtFail() uint64 {
	if x != nil && x.DecapSnExtFail != nil {
		return *x.DecapSnExtFail
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspAuthFailed() uint64 {
	if x != nil && x.EspAuthFailed != nil {
		return *x.EspAuthFailed
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapReinjectFail() uint64 {
	if x != nil && x.DecapReinjectFail != nil {
		return *x.DecapReinjectFail
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecapSnTransientDrop() uint64 {
	if x != nil && x.DecapSnTransientDrop != nil {
		return *x.DecapSnTransientDrop
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspRplbeforewindow() uint64 {
	if x != nil && x.EspRplbeforewindow != nil {
		return *x.EspRplbeforewindow
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspRplduplicate() uint64 {
	if x != nil && x.EspRplduplicate != nil {
		return *x.EspRplduplicate
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspProtectedBytesSent() uint64 {
	if x != nil && x.EspProtectedBytesSent != nil {
		return *x.EspProtectedBytesSent
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEspProtectedBytesRecvd() uint64 {
	if x != nil && x.EspProtectedBytesRecvd != nil {
		return *x.EspProtectedBytesRecvd
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetEncrypts() uint64 {
	if x != nil && x.Encrypts != nil {
		return *x.Encrypts
	}
	return 0
}

func (x *IPsecPerTunnelInfo) GetDecrypts() uint64 {
	if x != nil && x.Decrypts != nil {
		return *x.Decrypts
	}
	return 0
}

var file_ipsec_telemetry_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*IPsecVPN)(nil),
		Field:         77,
		Name:          "jnprIPsecVPNExt",
		Tag:           "bytes,77,opt,name=jnprIPsecVPNExt",
		Filename:      "ipsec_telemetry.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional IPsecVPN jnprIPsecVPNExt = 77;
	E_JnprIPsecVPNExt = &file_ipsec_telemetry_proto_extTypes[0]
)

var File_ipsec_telemetry_proto protoreflect.FileDescriptor

var file_ipsec_telemetry_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x69, 0x70, 0x73, 0x65, 0x63, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a,
	0x08, 0x49, 0x50, 0x73, 0x65, 0x63, 0x56, 0x50, 0x4e, 0x12, 0x3c, 0x0a, 0x11, 0x69, 0x70, 0x73,
	0x65, 0x63, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x49, 0x50, 0x73, 0x65, 0x63, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x69, 0x70, 0x73, 0x65, 0x63, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x11, 0x69, 0x70, 0x73, 0x65, 0x63,
	0x5f, 0x73, 0x76, 0x63, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x49, 0x50, 0x73, 0x65, 0x63, 0x50, 0x65, 0x72, 0x53, 0x76, 0x63,
	0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x69, 0x70, 0x73, 0x65, 0x63, 0x53, 0x76,
	0x63, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x11, 0x69, 0x70, 0x73, 0x65,
	0x63, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x49, 0x50, 0x73, 0x65, 0x63, 0x50, 0x65, 0x72, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x69, 0x70, 0x73, 0x65, 0x63, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xae, 0x04, 0x0a, 0x0f, 0x49, 0x50,
	0x73, 0x65, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a,
	0x10, 0x72, 0x65, 0x5f, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x50, 0x63, 0x6f, 0x6e, 0x6e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x63, 0x6f, 0x6e, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70,
	0x63, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x61,
	0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x71, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x73, 0x61, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x71, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x2d, 0x0a, 0x13, 0x73, 0x61, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x71, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x61,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x71, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x37,
	0x0a, 0x18, 0x73, 0x61, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x15, 0x73, 0x61, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x73, 0x61, 0x5f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x73, 0x61, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x74, 0x72, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x61,
	0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x61, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x61, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x61, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x61,
	0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x61, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x31, 0x0a, 0x15, 0x73, 0x61, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x73, 0x61, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x61, 0x5f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x73, 0x61, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x12,
	0x29, 0x0a, 0x11, 0x73, 0x61, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x5f, 0x65, 0x6e, 0x71, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x61, 0x54, 0x72,
	0x69, 0x67, 0x45, 0x6e, 0x71, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xdd, 0x02, 0x0a, 0x12, 0x49,
	0x50, 0x73, 0x65, 0x63, 0x50, 0x65, 0x72, 0x53, 0x76, 0x63, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x76, 0x63, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x08, 0x73, 0x76, 0x63, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x72, 0x75, 0x6c, 0x65,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x61, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x61, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64,
	0x73, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6d, 0x74, 0x75, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x10, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x73, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x4d, 0x74, 0x75, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x70, 0x6b,
	0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x50, 0x6b, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x73, 0x70, 0x5f, 0x70, 0x6b, 0x74, 0x5f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x73,
	0x70, 0x50, 0x6b, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6e, 0x63, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x61, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x65, 0x63,
	0x61, 0x70, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x81, 0x07, 0x0a, 0x12, 0x49,
	0x50, 0x73, 0x65, 0x63, 0x50, 0x65, 0x72, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x73, 0x70, 0x5f, 0x72, 0x70, 0x6c, 0x7a, 0x65, 0x72, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x73, 0x70, 0x52, 0x70, 0x6c, 0x7a, 0x65, 0x72, 0x6f, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x70, 0x73, 0x65, 0x63, 0x5f, 0x62, 0x61, 0x64, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x69, 0x70, 0x73, 0x65,
	0x63, 0x42, 0x61, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x65,
	0x73, 0x70, 0x5f, 0x62, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x73, 0x70, 0x42, 0x61, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x6e,
	0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x10, 0x64, 0x65, 0x63, 0x61, 0x70, 0x4e, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x45, 0x72, 0x72, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x64, 0x65, 0x63, 0x61, 0x70, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x65, 0x6e,
	0x45, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x68, 0x64, 0x72,
	0x5f, 0x65, 0x72, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x63, 0x61,
	0x70, 0x48, 0x64, 0x72, 0x45, 0x72, 0x72, 0x12, 0x31, 0x0a, 0x15, 0x64, 0x65, 0x63, 0x61, 0x70,
	0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x65, 0x72, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x64, 0x65, 0x63, 0x61, 0x70, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x53, 0x61, 0x64, 0x64, 0x72, 0x45, 0x72, 0x72, 0x12, 0x31, 0x0a, 0x15, 0x64, 0x65,
	0x63, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x64, 0x64, 0x72, 0x5f,
	0x65, 0x72, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x64, 0x65, 0x63, 0x61, 0x70,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x61, 0x64, 0x64, 0x72, 0x45, 0x72, 0x72, 0x12, 0x2d, 0x0a,
	0x13, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x73, 0x6e, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x64, 0x65, 0x63, 0x61,
	0x70, 0x53, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x11,
	0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x73, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x61, 0x70, 0x53, 0x6e,
	0x45, 0x78, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x73, 0x70, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x65, 0x73, 0x70, 0x41, 0x75, 0x74, 0x68, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12,
	0x2e, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x72, 0x65, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x64, 0x65,
	0x63, 0x61, 0x70, 0x52, 0x65, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x12,
	0x35, 0x0a, 0x17, 0x64, 0x65, 0x63, 0x61, 0x70, 0x5f, 0x73, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x14, 0x64, 0x65, 0x63, 0x61, 0x70, 0x53, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65,
	0x6e, 0x74, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x73, 0x70, 0x5f, 0x72, 0x70,
	0x6c, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x65, 0x73, 0x70, 0x52, 0x70, 0x6c, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x73, 0x70, 0x5f, 0x72,
	0x70, 0x6c, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0f, 0x65, 0x73, 0x70, 0x52, 0x70, 0x6c, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x73, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x65, 0x73, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x65,
	0x73, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16,
	0x65, 0x73, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x63, 0x76, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x73, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x73, 0x3a, 0x4c,
	0x0a, 0x0f, 0x6a, 0x6e, 0x70, 0x72, 0x49, 0x50, 0x73, 0x65, 0x63, 0x56, 0x50, 0x4e, 0x45, 0x78,
	0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x4d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x49, 0x50, 0x73, 0x65, 0x63, 0x56, 0x50, 0x4e, 0x52, 0x0f, 0x6a, 0x6e, 0x70,
	0x72, 0x49, 0x50, 0x73, 0x65, 0x63, 0x56, 0x50, 0x4e, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a, 0x1b,
	0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_ipsec_telemetry_proto_rawDescOnce sync.Once
	file_ipsec_telemetry_proto_rawDescData []byte
)

func file_ipsec_telemetry_proto_rawDescGZIP() []byte {
	file_ipsec_telemetry_proto_rawDescOnce.Do(func() {
		file_ipsec_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ipsec_telemetry_proto_rawDesc), len(file_ipsec_telemetry_proto_rawDesc)))
	})
	return file_ipsec_telemetry_proto_rawDescData
}

var file_ipsec_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ipsec_telemetry_proto_goTypes = []any{
	(*IPsecVPN)(nil),               // 0: IPsecVPN
	(*IPsecGlobalInfo)(nil),        // 1: IPsecGlobalInfo
	(*IPsecPerSvcsetInfo)(nil),     // 2: IPsecPerSvcsetInfo
	(*IPsecPerTunnelInfo)(nil),     // 3: IPsecPerTunnelInfo
	(*JuniperNetworksSensors)(nil), // 4: JuniperNetworksSensors
}
var file_ipsec_telemetry_proto_depIdxs = []int32{
	1, // 0: IPsecVPN.ipsec_global_info:type_name -> IPsecGlobalInfo
	2, // 1: IPsecVPN.ipsec_svcset_info:type_name -> IPsecPerSvcsetInfo
	3, // 2: IPsecVPN.ipsec_tunnel_info:type_name -> IPsecPerTunnelInfo
	4, // 3: jnprIPsecVPNExt:extendee -> JuniperNetworksSensors
	0, // 4: jnprIPsecVPNExt:type_name -> IPsecVPN
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ipsec_telemetry_proto_init() }
func file_ipsec_telemetry_proto_init() {
	if File_ipsec_telemetry_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ipsec_telemetry_proto_rawDesc), len(file_ipsec_telemetry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_ipsec_telemetry_proto_goTypes,
		DependencyIndexes: file_ipsec_telemetry_proto_depIdxs,
		MessageInfos:      file_ipsec_telemetry_proto_msgTypes,
		ExtensionInfos:    file_ipsec_telemetry_proto_extTypes,
	}.Build()
	File_ipsec_telemetry_proto = out.File
	file_ipsec_telemetry_proto_goTypes = nil
	file_ipsec_telemetry_proto_depIdxs = nil
}
