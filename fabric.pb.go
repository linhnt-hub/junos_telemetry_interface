//
// This defines the gpb message format used by the fabric sensor.
//
// The top level message is 'fabric_message'
//
// Version 1.0
//
// Copyrights (c) 2015, 2016, Juniper Networks, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: fabric.proto

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

type FabricMessageSensorLocation int32

const (
	FabricMessage_Linecard      FabricMessageSensorLocation = 1 // LC
	FabricMessage_Switch_Fabric FabricMessageSensorLocation = 2 // FAB
)

// Enum value maps for FabricMessageSensorLocation.
var (
	FabricMessageSensorLocation_name = map[int32]string{
		1: "Linecard",
		2: "Switch_Fabric",
	}
	FabricMessageSensorLocation_value = map[string]int32{
		"Linecard":      1,
		"Switch_Fabric": 2,
	}
)

func (x FabricMessageSensorLocation) Enum() *FabricMessageSensorLocation {
	p := new(FabricMessageSensorLocation)
	*p = x
	return p
}

func (x FabricMessageSensorLocation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FabricMessageSensorLocation) Descriptor() protoreflect.EnumDescriptor {
	return file_fabric_proto_enumTypes[0].Descriptor()
}

func (FabricMessageSensorLocation) Type() protoreflect.EnumType {
	return &file_fabric_proto_enumTypes[0]
}

func (x FabricMessageSensorLocation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FabricMessageSensorLocation) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FabricMessageSensorLocation(num)
	return nil
}

// Deprecated: Use FabricMessageSensorLocation.Descriptor instead.
func (FabricMessageSensorLocation) EnumDescriptor() ([]byte, []int) {
	return file_fabric_proto_rawDescGZIP(), []int{0, 0}
}

type EdgeStatsIdentifierType int32

const (
	EdgeStats_Switch_Fabric EdgeStatsIdentifierType = 1
	EdgeStats_Linecard      EdgeStatsIdentifierType = 2
)

// Enum value maps for EdgeStatsIdentifierType.
var (
	EdgeStatsIdentifierType_name = map[int32]string{
		1: "Switch_Fabric",
		2: "Linecard",
	}
	EdgeStatsIdentifierType_value = map[string]int32{
		"Switch_Fabric": 1,
		"Linecard":      2,
	}
)

func (x EdgeStatsIdentifierType) Enum() *EdgeStatsIdentifierType {
	p := new(EdgeStatsIdentifierType)
	*p = x
	return p
}

func (x EdgeStatsIdentifierType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EdgeStatsIdentifierType) Descriptor() protoreflect.EnumDescriptor {
	return file_fabric_proto_enumTypes[1].Descriptor()
}

func (EdgeStatsIdentifierType) Type() protoreflect.EnumType {
	return &file_fabric_proto_enumTypes[1]
}

func (x EdgeStatsIdentifierType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EdgeStatsIdentifierType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EdgeStatsIdentifierType(num)
	return nil
}

// Deprecated: Use EdgeStatsIdentifierType.Descriptor instead.
func (EdgeStatsIdentifierType) EnumDescriptor() ([]byte, []int) {
	return file_fabric_proto_rawDescGZIP(), []int{1, 0}
}

type FabricMessage struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Edges         []*EdgeStats                 `protobuf:"bytes,1,rep,name=edges" json:"edges,omitempty"`
	Location      *FabricMessageSensorLocation `protobuf:"varint,2,opt,name=location,enum=FabricMessageSensorLocation" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FabricMessage) Reset() {
	*x = FabricMessage{}
	mi := &file_fabric_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FabricMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FabricMessage) ProtoMessage() {}

func (x *FabricMessage) ProtoReflect() protoreflect.Message {
	mi := &file_fabric_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FabricMessage.ProtoReflect.Descriptor instead.
func (*FabricMessage) Descriptor() ([]byte, []int) {
	return file_fabric_proto_rawDescGZIP(), []int{0}
}

func (x *FabricMessage) GetEdges() []*EdgeStats {
	if x != nil {
		return x.Edges
	}
	return nil
}

func (x *FabricMessage) GetLocation() FabricMessageSensorLocation {
	if x != nil && x.Location != nil {
		return *x.Location
	}
	return FabricMessage_Linecard
}

type EdgeStats struct {
	state           protoimpl.MessageState   `protogen:"open.v1"`
	SourceType      *EdgeStatsIdentifierType `protobuf:"varint,1,opt,name=source_type,json=sourceType,enum=EdgeStatsIdentifierType" json:"source_type,omitempty"`
	SourceSlot      *uint32                  `protobuf:"varint,2,opt,name=source_slot,json=sourceSlot" json:"source_slot,omitempty"`
	SourcePfe       *uint32                  `protobuf:"varint,3,opt,name=source_pfe,json=sourcePfe" json:"source_pfe,omitempty"`
	DestinationType *EdgeStatsIdentifierType `protobuf:"varint,4,opt,name=destination_type,json=destinationType,enum=EdgeStatsIdentifierType" json:"destination_type,omitempty"`
	DestinationSlot *uint32                  `protobuf:"varint,5,opt,name=destination_slot,json=destinationSlot" json:"destination_slot,omitempty"`
	DestinationPfe  *uint32                  `protobuf:"varint,6,opt,name=destination_pfe,json=destinationPfe" json:"destination_pfe,omitempty"`
	// stats from src's perspective
	ClassStats    []*ClassStats `protobuf:"bytes,7,rep,name=class_stats,json=classStats" json:"class_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EdgeStats) Reset() {
	*x = EdgeStats{}
	mi := &file_fabric_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EdgeStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeStats) ProtoMessage() {}

func (x *EdgeStats) ProtoReflect() protoreflect.Message {
	mi := &file_fabric_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeStats.ProtoReflect.Descriptor instead.
func (*EdgeStats) Descriptor() ([]byte, []int) {
	return file_fabric_proto_rawDescGZIP(), []int{1}
}

func (x *EdgeStats) GetSourceType() EdgeStatsIdentifierType {
	if x != nil && x.SourceType != nil {
		return *x.SourceType
	}
	return EdgeStats_Switch_Fabric
}

func (x *EdgeStats) GetSourceSlot() uint32 {
	if x != nil && x.SourceSlot != nil {
		return *x.SourceSlot
	}
	return 0
}

func (x *EdgeStats) GetSourcePfe() uint32 {
	if x != nil && x.SourcePfe != nil {
		return *x.SourcePfe
	}
	return 0
}

func (x *EdgeStats) GetDestinationType() EdgeStatsIdentifierType {
	if x != nil && x.DestinationType != nil {
		return *x.DestinationType
	}
	return EdgeStats_Switch_Fabric
}

func (x *EdgeStats) GetDestinationSlot() uint32 {
	if x != nil && x.DestinationSlot != nil {
		return *x.DestinationSlot
	}
	return 0
}

func (x *EdgeStats) GetDestinationPfe() uint32 {
	if x != nil && x.DestinationPfe != nil {
		return *x.DestinationPfe
	}
	return 0
}

func (x *EdgeStats) GetClassStats() []*ClassStats {
	if x != nil {
		return x.ClassStats
	}
	return nil
}

type ClassStats struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// service class, 0 being highest
	Priority *string `protobuf:"bytes,1,opt,name=priority" json:"priority,omitempty"`
	// transmit data from src to dst
	TransmitCounts *Counters `protobuf:"bytes,2,opt,name=transmit_counts,json=transmitCounts" json:"transmit_counts,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ClassStats) Reset() {
	*x = ClassStats{}
	mi := &file_fabric_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassStats) ProtoMessage() {}

func (x *ClassStats) ProtoReflect() protoreflect.Message {
	mi := &file_fabric_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassStats.ProtoReflect.Descriptor instead.
func (*ClassStats) Descriptor() ([]byte, []int) {
	return file_fabric_proto_rawDescGZIP(), []int{2}
}

func (x *ClassStats) GetPriority() string {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return ""
}

func (x *ClassStats) GetTransmitCounts() *Counters {
	if x != nil {
		return x.TransmitCounts
	}
	return nil
}

type Counters struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Counter: total packets/cells from this source->destination
	Packets *uint64 `protobuf:"varint,1,opt,name=packets" json:"packets,omitempty"`
	// Counter: total bytes from this source->destination
	Bytes *uint64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	// Rate: packets/cells per seconds, measured across an interval of 1 second.
	PacketsPerSecond *uint64 `protobuf:"varint,3,opt,name=packets_per_second,json=packetsPerSecond" json:"packets_per_second,omitempty"`
	// Rate: bits per seconds, measured across an interval of 1 second.
	BytesPerSecond *uint64 `protobuf:"varint,4,opt,name=bytes_per_second,json=bytesPerSecond" json:"bytes_per_second,omitempty"`
	// Counter: total number of dropped packets/cells from this source->destination
	DropPackets *uint64 `protobuf:"varint,5,opt,name=drop_packets,json=dropPackets" json:"drop_packets,omitempty"`
	// Counter: total number of dropped bytes from this source->destination
	DropBytes *uint64 `protobuf:"varint,6,opt,name=drop_bytes,json=dropBytes" json:"drop_bytes,omitempty"`
	// Rate: the rate at which packets/cells are dropped, in pps/cps, measured
	// across an interval of 1 second.
	DropPacketsPerSecond *uint64 `protobuf:"varint,7,opt,name=drop_packets_per_second,json=dropPacketsPerSecond" json:"drop_packets_per_second,omitempty"`
	// Rate: the rate at which bytes are dropped, in bytes per sec, measured
	// across an interval of 1 second.
	DropBytesPerSecond *uint64 `protobuf:"varint,8,opt,name=drop_bytes_per_second,json=dropBytesPerSecond" json:"drop_bytes_per_second,omitempty"`
	// Average: avg queue depth,TAQL:time-average-queue-len, in packets
	QueueDepthAverage *uint64 `protobuf:"varint,9,opt,name=queue_depth_average,json=queueDepthAverage" json:"queue_depth_average,omitempty"`
	// Gauge: current queue depth, in packets
	QueueDepthCurrent *uint64 `protobuf:"varint,10,opt,name=queue_depth_current,json=queueDepthCurrent" json:"queue_depth_current,omitempty"`
	// Peak: the max measured queue depth, in packets, across all measurements since boot.
	QueueDepthPeak *uint64 `protobuf:"varint,11,opt,name=queue_depth_peak,json=queueDepthPeak" json:"queue_depth_peak,omitempty"`
	// The configured maximum depth of the queue, in packets
	QueueDepthMaximum *uint64 `protobuf:"varint,12,opt,name=queue_depth_maximum,json=queueDepthMaximum" json:"queue_depth_maximum,omitempty"`
	// Counter: total number of errored packets
	ErrorPackets *uint64 `protobuf:"varint,13,opt,name=error_packets,json=errorPackets" json:"error_packets,omitempty"`
	// Rate: the rate of errored packets, in pps, measured across an interval of 1 second.
	ErrorPacketsPerSecond *uint64 `protobuf:"varint,14,opt,name=error_packets_per_second,json=errorPacketsPerSecond" json:"error_packets_per_second,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *Counters) Reset() {
	*x = Counters{}
	mi := &file_fabric_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Counters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counters) ProtoMessage() {}

func (x *Counters) ProtoReflect() protoreflect.Message {
	mi := &file_fabric_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counters.ProtoReflect.Descriptor instead.
func (*Counters) Descriptor() ([]byte, []int) {
	return file_fabric_proto_rawDescGZIP(), []int{3}
}

func (x *Counters) GetPackets() uint64 {
	if x != nil && x.Packets != nil {
		return *x.Packets
	}
	return 0
}

func (x *Counters) GetBytes() uint64 {
	if x != nil && x.Bytes != nil {
		return *x.Bytes
	}
	return 0
}

func (x *Counters) GetPacketsPerSecond() uint64 {
	if x != nil && x.PacketsPerSecond != nil {
		return *x.PacketsPerSecond
	}
	return 0
}

func (x *Counters) GetBytesPerSecond() uint64 {
	if x != nil && x.BytesPerSecond != nil {
		return *x.BytesPerSecond
	}
	return 0
}

func (x *Counters) GetDropPackets() uint64 {
	if x != nil && x.DropPackets != nil {
		return *x.DropPackets
	}
	return 0
}

func (x *Counters) GetDropBytes() uint64 {
	if x != nil && x.DropBytes != nil {
		return *x.DropBytes
	}
	return 0
}

func (x *Counters) GetDropPacketsPerSecond() uint64 {
	if x != nil && x.DropPacketsPerSecond != nil {
		return *x.DropPacketsPerSecond
	}
	return 0
}

func (x *Counters) GetDropBytesPerSecond() uint64 {
	if x != nil && x.DropBytesPerSecond != nil {
		return *x.DropBytesPerSecond
	}
	return 0
}

func (x *Counters) GetQueueDepthAverage() uint64 {
	if x != nil && x.QueueDepthAverage != nil {
		return *x.QueueDepthAverage
	}
	return 0
}

func (x *Counters) GetQueueDepthCurrent() uint64 {
	if x != nil && x.QueueDepthCurrent != nil {
		return *x.QueueDepthCurrent
	}
	return 0
}

func (x *Counters) GetQueueDepthPeak() uint64 {
	if x != nil && x.QueueDepthPeak != nil {
		return *x.QueueDepthPeak
	}
	return 0
}

func (x *Counters) GetQueueDepthMaximum() uint64 {
	if x != nil && x.QueueDepthMaximum != nil {
		return *x.QueueDepthMaximum
	}
	return 0
}

func (x *Counters) GetErrorPackets() uint64 {
	if x != nil && x.ErrorPackets != nil {
		return *x.ErrorPackets
	}
	return 0
}

func (x *Counters) GetErrorPacketsPerSecond() uint64 {
	if x != nil && x.ErrorPacketsPerSecond != nil {
		return *x.ErrorPacketsPerSecond
	}
	return 0
}

var file_fabric_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*FabricMessage)(nil),
		Field:         2,
		Name:          "fabricMessageExt",
		Tag:           "bytes,2,opt,name=fabricMessageExt",
		Filename:      "fabric.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional fabric_message fabricMessageExt = 2;
	E_FabricMessageExt = &file_fabric_proto_extTypes[0]
)

var File_fabric_proto protoreflect.FileDescriptor

var file_fabric_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x66, 0x61,
	0x62, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x05, 0x82, 0x40,
	0x02, 0x08, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a,
	0x0f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x63, 0x61, 0x72, 0x64, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x46, 0x61, 0x62, 0x72, 0x69, 0x63, 0x10,
	0x02, 0x22, 0xb3, 0x03, 0x0a, 0x0a, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x43, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08,
	0x01, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x24, 0x0a,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x66, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x66, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08,
	0x01, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x30, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x82, 0x40,
	0x02, 0x08, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x66, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x82,
	0x40, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x66, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x5f, 0x46, 0x61, 0x62, 0x72, 0x69, 0x63, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x69, 0x6e,
	0x65, 0x63, 0x61, 0x72, 0x64, 0x10, 0x02, 0x22, 0x64, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xb8, 0x05,
	0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02,
	0x18, 0x01, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x18,
	0x01, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x12, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x10, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x2f, 0x0a,
	0x10, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x0e,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x28,
	0x0a, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x64, 0x72, 0x6f,
	0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0a, 0x64, 0x72, 0x6f, 0x70,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40,
	0x02, 0x18, 0x01, 0x52, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3c,
	0x0a, 0x17, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x14, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x15,
	0x64, 0x72, 0x6f, 0x70, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02,
	0x20, 0x01, 0x52, 0x12, 0x64, 0x72, 0x6f, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x35, 0x0a, 0x13, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x11, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a,
	0x13, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20,
	0x01, 0x52, 0x11, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x64, 0x65,
	0x70, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x61, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05,
	0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x0e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x70, 0x74,
	0x68, 0x50, 0x65, 0x61, 0x6b, 0x12, 0x35, 0x0a, 0x13, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x64,
	0x65, 0x70, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20, 0x01, 0x52, 0x11, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x44, 0x65, 0x70, 0x74, 0x68, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x0d,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x18, 0x01, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x18, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0x82, 0x40, 0x02, 0x20,
	0x01, 0x52, 0x15, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x3a, 0x54, 0x0a, 0x10, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a,
	0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x61,
	0x62, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10, 0x66, 0x61,
	0x62, 0x72, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x42, 0x1d,
	0x5a, 0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_fabric_proto_rawDescOnce sync.Once
	file_fabric_proto_rawDescData []byte
)

func file_fabric_proto_rawDescGZIP() []byte {
	file_fabric_proto_rawDescOnce.Do(func() {
		file_fabric_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fabric_proto_rawDesc), len(file_fabric_proto_rawDesc)))
	})
	return file_fabric_proto_rawDescData
}

var file_fabric_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_fabric_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fabric_proto_goTypes = []any{
	(FabricMessageSensorLocation)(0), // 0: fabric_message.sensor_location
	(EdgeStatsIdentifierType)(0),     // 1: edge_stats.identifier_type
	(*FabricMessage)(nil),            // 2: fabric_message
	(*EdgeStats)(nil),                // 3: edge_stats
	(*ClassStats)(nil),               // 4: class_stats
	(*Counters)(nil),                 // 5: counters
	(*JuniperNetworksSensors)(nil),   // 6: JuniperNetworksSensors
}
var file_fabric_proto_depIdxs = []int32{
	3, // 0: fabric_message.edges:type_name -> edge_stats
	0, // 1: fabric_message.location:type_name -> fabric_message.sensor_location
	1, // 2: edge_stats.source_type:type_name -> edge_stats.identifier_type
	1, // 3: edge_stats.destination_type:type_name -> edge_stats.identifier_type
	4, // 4: edge_stats.class_stats:type_name -> class_stats
	5, // 5: class_stats.transmit_counts:type_name -> counters
	6, // 6: fabricMessageExt:extendee -> JuniperNetworksSensors
	2, // 7: fabricMessageExt:type_name -> fabric_message
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	7, // [7:8] is the sub-list for extension type_name
	6, // [6:7] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_fabric_proto_init() }
func file_fabric_proto_init() {
	if File_fabric_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fabric_proto_rawDesc), len(file_fabric_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_fabric_proto_goTypes,
		DependencyIndexes: file_fabric_proto_depIdxs,
		EnumInfos:         file_fabric_proto_enumTypes,
		MessageInfos:      file_fabric_proto_msgTypes,
		ExtensionInfos:    file_fabric_proto_extTypes,
	}.Build()
	File_fabric_proto = out.File
	file_fabric_proto_goTypes = nil
	file_fabric_proto_depIdxs = nil
}
