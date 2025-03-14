//
// Copyrights (c)  2016, Juniper Networks, Inc.
// All rights reserved.
//

//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless optional by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

//
// Balagopal Ambalakkat,Sumanth Kamatala, 2016
//
// This defines the gpb message format used exporting
// cmerror configuration information.
//
// The top level message is Cmerror
// The message is a collection of error items.
//
//
// Version 1.0
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: cmerror.proto

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

// Juniper Error Item information
type Error struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier that uniquely identifies the source of
	// the error.
	// e.g.
	//
	// junos/system/linecard/0/pcie/0/lane/0/pcie_cmerror_uncorrectable_major
	Identifier *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Name of the error
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Instance id of the associated resource
	ComponentId *uint32 `protobuf:"varint,3,opt,name=component_id,json=componentId" json:"component_id,omitempty"`
	// Fru information
	FruType *string `protobuf:"bytes,4,opt,name=fru_type,json=fruType" json:"fru_type,omitempty"`
	FruSlot *uint32 `protobuf:"varint,5,opt,name=fru_slot,json=fruSlot" json:"fru_slot,omitempty"`
	// Scope,Category,Severity
	// in which this error belong to.
	Scope    *string `protobuf:"bytes,6,opt,name=scope" json:"scope,omitempty"`
	Category *string `protobuf:"bytes,7,opt,name=category" json:"category,omitempty"`
	Severity *string `protobuf:"bytes,8,opt,name=severity" json:"severity,omitempty"`
	// Thresholds and action configured for this
	// error.
	Threshold         *uint32 `protobuf:"varint,9,opt,name=threshold" json:"threshold,omitempty"`
	Limit             *uint32 `protobuf:"varint,10,opt,name=limit" json:"limit,omitempty"`
	RaisingThreshold  *uint32 `protobuf:"varint,11,opt,name=raising_threshold,json=raisingThreshold" json:"raising_threshold,omitempty"`
	ClearingThreshold *uint32 `protobuf:"varint,12,opt,name=clearing_threshold,json=clearingThreshold" json:"clearing_threshold,omitempty"`
	Action            *uint32 `protobuf:"varint,13,opt,name=action" json:"action,omitempty"`
	// local/global/both
	ActionHandlingType *uint32 `protobuf:"varint,14,opt,name=action_handling_type,json=actionHandlingType" json:"action_handling_type,omitempty"`
	// user configured thresholds and limits for this error.
	ConfiguredThreshold   *uint32 `protobuf:"varint,15,opt,name=configured_threshold,json=configuredThreshold" json:"configured_threshold,omitempty"`
	ConfiguredLimit       *uint32 `protobuf:"varint,16,opt,name=configured_limit,json=configuredLimit" json:"configured_limit,omitempty"`
	ConfiguredAction      *uint32 `protobuf:"varint,17,opt,name=configured_action,json=configuredAction" json:"configured_action,omitempty"`
	ConfiguredClearAction *uint32 `protobuf:"varint,18,opt,name=configured_clear_action,json=configuredClearAction" json:"configured_clear_action,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_cmerror_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_cmerror_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_cmerror_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *Error) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Error) GetComponentId() uint32 {
	if x != nil && x.ComponentId != nil {
		return *x.ComponentId
	}
	return 0
}

func (x *Error) GetFruType() string {
	if x != nil && x.FruType != nil {
		return *x.FruType
	}
	return ""
}

func (x *Error) GetFruSlot() uint32 {
	if x != nil && x.FruSlot != nil {
		return *x.FruSlot
	}
	return 0
}

func (x *Error) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

func (x *Error) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *Error) GetSeverity() string {
	if x != nil && x.Severity != nil {
		return *x.Severity
	}
	return ""
}

func (x *Error) GetThreshold() uint32 {
	if x != nil && x.Threshold != nil {
		return *x.Threshold
	}
	return 0
}

func (x *Error) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *Error) GetRaisingThreshold() uint32 {
	if x != nil && x.RaisingThreshold != nil {
		return *x.RaisingThreshold
	}
	return 0
}

func (x *Error) GetClearingThreshold() uint32 {
	if x != nil && x.ClearingThreshold != nil {
		return *x.ClearingThreshold
	}
	return 0
}

func (x *Error) GetAction() uint32 {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return 0
}

func (x *Error) GetActionHandlingType() uint32 {
	if x != nil && x.ActionHandlingType != nil {
		return *x.ActionHandlingType
	}
	return 0
}

func (x *Error) GetConfiguredThreshold() uint32 {
	if x != nil && x.ConfiguredThreshold != nil {
		return *x.ConfiguredThreshold
	}
	return 0
}

func (x *Error) GetConfiguredLimit() uint32 {
	if x != nil && x.ConfiguredLimit != nil {
		return *x.ConfiguredLimit
	}
	return 0
}

func (x *Error) GetConfiguredAction() uint32 {
	if x != nil && x.ConfiguredAction != nil {
		return *x.ConfiguredAction
	}
	return 0
}

func (x *Error) GetConfiguredClearAction() uint32 {
	if x != nil && x.ConfiguredClearAction != nil {
		return *x.ConfiguredClearAction
	}
	return 0
}

type GlobalErrorConfiguration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// configuration bucket identifier
	Scope    *string `protobuf:"bytes,1,opt,name=scope" json:"scope,omitempty"`
	Category *string `protobuf:"bytes,2,opt,name=category" json:"category,omitempty"`
	Severity *string `protobuf:"bytes,3,opt,name=severity" json:"severity,omitempty"`
	// configured parameters for this bucket.
	Threshold     *uint32 `protobuf:"varint,4,opt,name=threshold" json:"threshold,omitempty"`
	Action        *uint32 `protobuf:"varint,5,opt,name=action" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GlobalErrorConfiguration) Reset() {
	*x = GlobalErrorConfiguration{}
	mi := &file_cmerror_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlobalErrorConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalErrorConfiguration) ProtoMessage() {}

func (x *GlobalErrorConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_cmerror_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalErrorConfiguration.ProtoReflect.Descriptor instead.
func (*GlobalErrorConfiguration) Descriptor() ([]byte, []int) {
	return file_cmerror_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalErrorConfiguration) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

func (x *GlobalErrorConfiguration) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *GlobalErrorConfiguration) GetSeverity() string {
	if x != nil && x.Severity != nil {
		return *x.Severity
	}
	return ""
}

func (x *GlobalErrorConfiguration) GetThreshold() uint32 {
	if x != nil && x.Threshold != nil {
		return *x.Threshold
	}
	return 0
}

func (x *GlobalErrorConfiguration) GetAction() uint32 {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return 0
}

// Top-level Cmerror message
type Cmerror struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// collection of error items
	ErrorItem []*Error `protobuf:"bytes,1,rep,name=error_item,json=errorItem" json:"error_item,omitempty"`
	// last configuration change for cmerror.
	LastConfigurationChange *uint64 `protobuf:"varint,2,opt,name=last_configuration_change,json=lastConfigurationChange" json:"last_configuration_change,omitempty"`
	// This will toggle at start of every wrap cycle
	WrapCycleMarker *bool `protobuf:"varint,3,opt,name=wrap_cycle_marker,json=wrapCycleMarker" json:"wrap_cycle_marker,omitempty"`
	// Fru slot identifier
	FruSlot *uint32 `protobuf:"varint,4,opt,name=fru_slot,json=fruSlot" json:"fru_slot,omitempty"`
	FruType *string `protobuf:"bytes,5,opt,name=fru_type,json=fruType" json:"fru_type,omitempty"`
	// Collection of global configuration items
	GlobalConfigurationItem []*GlobalErrorConfiguration `protobuf:"bytes,6,rep,name=global_configuration_item,json=globalConfigurationItem" json:"global_configuration_item,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Cmerror) Reset() {
	*x = Cmerror{}
	mi := &file_cmerror_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cmerror) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmerror) ProtoMessage() {}

func (x *Cmerror) ProtoReflect() protoreflect.Message {
	mi := &file_cmerror_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmerror.ProtoReflect.Descriptor instead.
func (*Cmerror) Descriptor() ([]byte, []int) {
	return file_cmerror_proto_rawDescGZIP(), []int{2}
}

func (x *Cmerror) GetErrorItem() []*Error {
	if x != nil {
		return x.ErrorItem
	}
	return nil
}

func (x *Cmerror) GetLastConfigurationChange() uint64 {
	if x != nil && x.LastConfigurationChange != nil {
		return *x.LastConfigurationChange
	}
	return 0
}

func (x *Cmerror) GetWrapCycleMarker() bool {
	if x != nil && x.WrapCycleMarker != nil {
		return *x.WrapCycleMarker
	}
	return false
}

func (x *Cmerror) GetFruSlot() uint32 {
	if x != nil && x.FruSlot != nil {
		return *x.FruSlot
	}
	return 0
}

func (x *Cmerror) GetFruType() string {
	if x != nil && x.FruType != nil {
		return *x.FruType
	}
	return ""
}

func (x *Cmerror) GetGlobalConfigurationItem() []*GlobalErrorConfiguration {
	if x != nil {
		return x.GlobalConfigurationItem
	}
	return nil
}

var file_cmerror_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*Cmerror)(nil),
		Field:         20,
		Name:          "jnpr_cmerror_ext",
		Tag:           "bytes,20,opt,name=jnpr_cmerror_ext",
		Filename:      "cmerror.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional Cmerror jnpr_cmerror_ext = 20;
	E_JnprCmerrorExt = &file_cmerror_proto_extTypes[0]
)

var File_cmerror_proto protoreflect.FileDescriptor

var file_cmerror_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x63, 0x6d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x72, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x72, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x75, 0x5f, 0x73,
	0x6c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x75, 0x53, 0x6c,
	0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63,
	0x6c, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x64, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9e, 0x01,
	0x0a, 0x18, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa5,
	0x02, 0x0a, 0x07, 0x43, 0x6d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x3a, 0x0a, 0x19, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x77, 0x72, 0x61, 0x70, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x77, 0x72, 0x61, 0x70, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x75,
	0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x75,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x55, 0x0a, 0x19, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x3a, 0x4b, 0x0a, 0x10, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x63,
	0x6d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e,
	0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x0e, 0x6a, 0x6e, 0x70, 0x72, 0x43, 0x6d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65,
})

var (
	file_cmerror_proto_rawDescOnce sync.Once
	file_cmerror_proto_rawDescData []byte
)

func file_cmerror_proto_rawDescGZIP() []byte {
	file_cmerror_proto_rawDescOnce.Do(func() {
		file_cmerror_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmerror_proto_rawDesc), len(file_cmerror_proto_rawDesc)))
	})
	return file_cmerror_proto_rawDescData
}

var file_cmerror_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmerror_proto_goTypes = []any{
	(*Error)(nil),                    // 0: Error
	(*GlobalErrorConfiguration)(nil), // 1: GlobalErrorConfiguration
	(*Cmerror)(nil),                  // 2: Cmerror
	(*JuniperNetworksSensors)(nil),   // 3: JuniperNetworksSensors
}
var file_cmerror_proto_depIdxs = []int32{
	0, // 0: Cmerror.error_item:type_name -> Error
	1, // 1: Cmerror.global_configuration_item:type_name -> GlobalErrorConfiguration
	3, // 2: jnpr_cmerror_ext:extendee -> JuniperNetworksSensors
	2, // 3: jnpr_cmerror_ext:type_name -> Cmerror
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmerror_proto_init() }
func file_cmerror_proto_init() {
	if File_cmerror_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmerror_proto_rawDesc), len(file_cmerror_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cmerror_proto_goTypes,
		DependencyIndexes: file_cmerror_proto_depIdxs,
		MessageInfos:      file_cmerror_proto_msgTypes,
		ExtensionInfos:    file_cmerror_proto_extTypes,
	}.Build()
	File_cmerror_proto = out.File
	file_cmerror_proto_goTypes = nil
	file_cmerror_proto_depIdxs = nil
}
