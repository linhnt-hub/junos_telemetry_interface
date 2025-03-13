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
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

//
// Nitin Kumar, 2016
//
// This defines the gpb message format used by the npu utilization sensor
//
// The top level message is NetworkProcessorUtilization
//
//
//
// Version 1.0
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: npu_utilization.proto

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

// Top level message
type NetworkProcessorUtilization struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NpuUtilStats  []*Utilization         `protobuf:"bytes,1,rep,name=npu_util_stats,json=npuUtilStats" json:"npu_util_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkProcessorUtilization) Reset() {
	*x = NetworkProcessorUtilization{}
	mi := &file_npu_utilization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkProcessorUtilization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkProcessorUtilization) ProtoMessage() {}

func (x *NetworkProcessorUtilization) ProtoReflect() protoreflect.Message {
	mi := &file_npu_utilization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkProcessorUtilization.ProtoReflect.Descriptor instead.
func (*NetworkProcessorUtilization) Descriptor() ([]byte, []int) {
	return file_npu_utilization_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkProcessorUtilization) GetNpuUtilStats() []*Utilization {
	if x != nil {
		return x.NpuUtilStats
	}
	return nil
}

// Utilization per Network Processor
type Utilization struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Globally unique identifier for an NPU. This is of the form
	// FPCX:NPUY, where X is the slot number of the linecard and Y
	// is the index of the NPU on the linecard
	Identifier *string `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	// A number on a scale of 0-100 that indicates the busyness of
	// an NPU. This is an approximation as the utilization depends
	// on the busyness of several internal components of the NPU
	Utilization *uint32 `protobuf:"varint,2,opt,name=utilization" json:"utilization,omitempty"`
	// Offered Packet load on the NPU.
	Packets []*PacketLoad `protobuf:"bytes,3,rep,name=packets" json:"packets,omitempty"`
	// Utilization of various internal memories of the NPU.
	Memory        []*MemoryLoad `protobuf:"bytes,4,rep,name=memory" json:"memory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Utilization) Reset() {
	*x = Utilization{}
	mi := &file_npu_utilization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Utilization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utilization) ProtoMessage() {}

func (x *Utilization) ProtoReflect() protoreflect.Message {
	mi := &file_npu_utilization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utilization.ProtoReflect.Descriptor instead.
func (*Utilization) Descriptor() ([]byte, []int) {
	return file_npu_utilization_proto_rawDescGZIP(), []int{1}
}

func (x *Utilization) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *Utilization) GetUtilization() uint32 {
	if x != nil && x.Utilization != nil {
		return *x.Utilization
	}
	return 0
}

func (x *Utilization) GetPackets() []*PacketLoad {
	if x != nil {
		return x.Packets
	}
	return nil
}

func (x *Utilization) GetMemory() []*MemoryLoad {
	if x != nil {
		return x.Memory
	}
	return nil
}

// Load on a memory subsystem of the NPU
type MemoryLoad struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A name string to identify the particular memory subsystem
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Various memory utilization metrics
	AverageUtil *uint32 `protobuf:"varint,2,opt,name=average_util,json=averageUtil" json:"average_util,omitempty"`
	HighestUtil *uint32 `protobuf:"varint,3,opt,name=highest_util,json=highestUtil" json:"highest_util,omitempty"`
	LowestUtil  *uint32 `protobuf:"varint,4,opt,name=lowest_util,json=lowestUtil" json:"lowest_util,omitempty"`
	// Each memory is front ended by a cache. The following metrics
	// indicate how these caches are working
	AverageCacheHitRate *uint32 `protobuf:"varint,5,opt,name=average_cache_hit_rate,json=averageCacheHitRate" json:"average_cache_hit_rate,omitempty"`
	HighestCacheHitRate *uint32 `protobuf:"varint,6,opt,name=highest_cache_hit_rate,json=highestCacheHitRate" json:"highest_cache_hit_rate,omitempty"`
	LowestCacheHitRate  *uint32 `protobuf:"varint,7,opt,name=lowest_cache_hit_rate,json=lowestCacheHitRate" json:"lowest_cache_hit_rate,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MemoryLoad) Reset() {
	*x = MemoryLoad{}
	mi := &file_npu_utilization_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemoryLoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLoad) ProtoMessage() {}

func (x *MemoryLoad) ProtoReflect() protoreflect.Message {
	mi := &file_npu_utilization_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLoad.ProtoReflect.Descriptor instead.
func (*MemoryLoad) Descriptor() ([]byte, []int) {
	return file_npu_utilization_proto_rawDescGZIP(), []int{2}
}

func (x *MemoryLoad) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *MemoryLoad) GetAverageUtil() uint32 {
	if x != nil && x.AverageUtil != nil {
		return *x.AverageUtil
	}
	return 0
}

func (x *MemoryLoad) GetHighestUtil() uint32 {
	if x != nil && x.HighestUtil != nil {
		return *x.HighestUtil
	}
	return 0
}

func (x *MemoryLoad) GetLowestUtil() uint32 {
	if x != nil && x.LowestUtil != nil {
		return *x.LowestUtil
	}
	return 0
}

func (x *MemoryLoad) GetAverageCacheHitRate() uint32 {
	if x != nil && x.AverageCacheHitRate != nil {
		return *x.AverageCacheHitRate
	}
	return 0
}

func (x *MemoryLoad) GetHighestCacheHitRate() uint32 {
	if x != nil && x.HighestCacheHitRate != nil {
		return *x.HighestCacheHitRate
	}
	return 0
}

func (x *MemoryLoad) GetLowestCacheHitRate() uint32 {
	if x != nil && x.LowestCacheHitRate != nil {
		return *x.LowestCacheHitRate
	}
	return 0
}

// Offered packet load on an internal subsystem of the NPU
type PacketLoad struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Each internal subsystem of the NPU has a name
	Identifier *string `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	// Rate of packets received
	Rate *uint64 `protobuf:"varint,2,opt,name=rate" json:"rate,omitempty"`
	// The following metrics indicate the compute load on the NPU
	AverageInstructionsPerPacket *uint32 `protobuf:"varint,3,opt,name=average_instructions_per_packet,json=averageInstructionsPerPacket" json:"average_instructions_per_packet,omitempty"`
	AverageWaitCyclesPerPacket   *uint32 `protobuf:"varint,4,opt,name=average_wait_cycles_per_packet,json=averageWaitCyclesPerPacket" json:"average_wait_cycles_per_packet,omitempty"`
	AverageCyclesPerPacket       *uint32 `protobuf:"varint,5,opt,name=average_cycles_per_packet,json=averageCyclesPerPacket" json:"average_cycles_per_packet,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *PacketLoad) Reset() {
	*x = PacketLoad{}
	mi := &file_npu_utilization_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketLoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketLoad) ProtoMessage() {}

func (x *PacketLoad) ProtoReflect() protoreflect.Message {
	mi := &file_npu_utilization_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketLoad.ProtoReflect.Descriptor instead.
func (*PacketLoad) Descriptor() ([]byte, []int) {
	return file_npu_utilization_proto_rawDescGZIP(), []int{3}
}

func (x *PacketLoad) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *PacketLoad) GetRate() uint64 {
	if x != nil && x.Rate != nil {
		return *x.Rate
	}
	return 0
}

func (x *PacketLoad) GetAverageInstructionsPerPacket() uint32 {
	if x != nil && x.AverageInstructionsPerPacket != nil {
		return *x.AverageInstructionsPerPacket
	}
	return 0
}

func (x *PacketLoad) GetAverageWaitCyclesPerPacket() uint32 {
	if x != nil && x.AverageWaitCyclesPerPacket != nil {
		return *x.AverageWaitCyclesPerPacket
	}
	return 0
}

func (x *PacketLoad) GetAverageCyclesPerPacket() uint32 {
	if x != nil && x.AverageCyclesPerPacket != nil {
		return *x.AverageCyclesPerPacket
	}
	return 0
}

var file_npu_utilization_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*NetworkProcessorUtilization)(nil),
		Field:         12,
		Name:          "jnpr_npu_utilization_ext",
		Tag:           "bytes,12,opt,name=jnpr_npu_utilization_ext",
		Filename:      "npu_utilization.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional NetworkProcessorUtilization jnpr_npu_utilization_ext = 12;
	E_JnprNpuUtilizationExt = &file_npu_utilization_proto_extTypes[0]
)

var File_npu_utilization_proto protoreflect.FileDescriptor

var file_npu_utilization_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6e, 0x70, 0x75, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x1b,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0e, 0x6e,
	0x70, 0x75, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x6e, 0x70, 0x75, 0x55, 0x74, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22,
	0x9b, 0x01, 0x0a, 0x0b, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x52,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0xa4, 0x02,
	0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x74, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x55,
	0x74, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x65,
	0x73, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74,
	0x5f, 0x75, 0x74, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x6f, 0x77,
	0x65, 0x73, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x16,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x68, 0x69,
	0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x68, 0x69,
	0x67, 0x68, 0x65, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x69, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x31, 0x0a, 0x15, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x69, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c,
	0x6f, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x1f, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x1c, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x42,
	0x0a, 0x1e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x57,
	0x61, 0x69, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x6e, 0x0a,
	0x18, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6e, 0x70, 0x75, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69,
	0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x55, 0x74, 0x69, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x6a, 0x6e, 0x70, 0x72, 0x4e, 0x70, 0x75, 0x55,
	0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a,
	0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_npu_utilization_proto_rawDescOnce sync.Once
	file_npu_utilization_proto_rawDescData []byte
)

func file_npu_utilization_proto_rawDescGZIP() []byte {
	file_npu_utilization_proto_rawDescOnce.Do(func() {
		file_npu_utilization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_npu_utilization_proto_rawDesc), len(file_npu_utilization_proto_rawDesc)))
	})
	return file_npu_utilization_proto_rawDescData
}

var file_npu_utilization_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npu_utilization_proto_goTypes = []any{
	(*NetworkProcessorUtilization)(nil), // 0: NetworkProcessorUtilization
	(*Utilization)(nil),                 // 1: Utilization
	(*MemoryLoad)(nil),                  // 2: MemoryLoad
	(*PacketLoad)(nil),                  // 3: PacketLoad
	(*JuniperNetworksSensors)(nil),      // 4: JuniperNetworksSensors
}
var file_npu_utilization_proto_depIdxs = []int32{
	1, // 0: NetworkProcessorUtilization.npu_util_stats:type_name -> Utilization
	3, // 1: Utilization.packets:type_name -> PacketLoad
	2, // 2: Utilization.memory:type_name -> MemoryLoad
	4, // 3: jnpr_npu_utilization_ext:extendee -> JuniperNetworksSensors
	0, // 4: jnpr_npu_utilization_ext:type_name -> NetworkProcessorUtilization
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npu_utilization_proto_init() }
func file_npu_utilization_proto_init() {
	if File_npu_utilization_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_npu_utilization_proto_rawDesc), len(file_npu_utilization_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_npu_utilization_proto_goTypes,
		DependencyIndexes: file_npu_utilization_proto_depIdxs,
		MessageInfos:      file_npu_utilization_proto_msgTypes,
		ExtensionInfos:    file_npu_utilization_proto_extTypes,
	}.Build()
	File_npu_utilization_proto = out.File
	file_npu_utilization_proto_goTypes = nil
	file_npu_utilization_proto_depIdxs = nil
}
