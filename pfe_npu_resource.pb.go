// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: pfe_npu_resource.proto

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

type JunosPfeNpu struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	NpuMemory     []*JunosPfeNpuNpuMemoryList `protobuf:"bytes,151,rep,name=npu_memory,json=npuMemory" json:"npu_memory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosPfeNpu) Reset() {
	*x = JunosPfeNpu{}
	mi := &file_pfe_npu_resource_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosPfeNpu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosPfeNpu) ProtoMessage() {}

func (x *JunosPfeNpu) ProtoReflect() protoreflect.Message {
	mi := &file_pfe_npu_resource_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosPfeNpu.ProtoReflect.Descriptor instead.
func (*JunosPfeNpu) Descriptor() ([]byte, []int) {
	return file_pfe_npu_resource_proto_rawDescGZIP(), []int{0}
}

func (x *JunosPfeNpu) GetNpuMemory() []*JunosPfeNpuNpuMemoryList {
	if x != nil {
		return x.NpuMemory
	}
	return nil
}

type JunosPfeNpuNpuMemoryList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JunosPfeNpuNpuMemoryList) Reset() {
	*x = JunosPfeNpuNpuMemoryList{}
	mi := &file_pfe_npu_resource_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JunosPfeNpuNpuMemoryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JunosPfeNpuNpuMemoryList) ProtoMessage() {}

func (x *JunosPfeNpuNpuMemoryList) ProtoReflect() protoreflect.Message {
	mi := &file_pfe_npu_resource_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JunosPfeNpuNpuMemoryList.ProtoReflect.Descriptor instead.
func (*JunosPfeNpuNpuMemoryList) Descriptor() ([]byte, []int) {
	return file_pfe_npu_resource_proto_rawDescGZIP(), []int{0, 0}
}

var file_pfe_npu_resource_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*JunosPfeNpu)(nil),
		Field:         59,
		Name:          "jnpr_junos_pfe_npu_ext",
		Tag:           "bytes,59,opt,name=jnpr_junos_pfe_npu_ext",
		Filename:      "pfe_npu_resource.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional junos_pfe_npu jnpr_junos_pfe_npu_ext = 59;
	E_JnprJunosPfeNpuExt = &file_pfe_npu_resource_proto_extTypes[0]
)

var File_pfe_npu_resource_proto protoreflect.FileDescriptor

var file_pfe_npu_resource_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x70, 0x66, 0x65, 0x5f, 0x6e, 0x70, 0x75, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a,
	0x0d, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x70, 0x66, 0x65, 0x5f, 0x6e, 0x70, 0x75, 0x12, 0x3e,
	0x0a, 0x0a, 0x6e, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x97, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x70, 0x66, 0x65, 0x5f,
	0x6e, 0x70, 0x75, 0x2e, 0x6e, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x09, 0x6e, 0x70, 0x75, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x1a, 0x11,
	0x0a, 0x0f, 0x6e, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x3a, 0x5b, 0x0a, 0x16, 0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f,
	0x70, 0x66, 0x65, 0x5f, 0x6e, 0x70, 0x75, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75,
	0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x73, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a, 0x75, 0x6e,
	0x6f, 0x73, 0x5f, 0x70, 0x66, 0x65, 0x5f, 0x6e, 0x70, 0x75, 0x52, 0x12, 0x6a, 0x6e, 0x70, 0x72,
	0x4a, 0x75, 0x6e, 0x6f, 0x73, 0x50, 0x66, 0x65, 0x4e, 0x70, 0x75, 0x45, 0x78, 0x74, 0x42, 0x1d,
	0x5a, 0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_pfe_npu_resource_proto_rawDescOnce sync.Once
	file_pfe_npu_resource_proto_rawDescData []byte
)

func file_pfe_npu_resource_proto_rawDescGZIP() []byte {
	file_pfe_npu_resource_proto_rawDescOnce.Do(func() {
		file_pfe_npu_resource_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pfe_npu_resource_proto_rawDesc), len(file_pfe_npu_resource_proto_rawDesc)))
	})
	return file_pfe_npu_resource_proto_rawDescData
}

var file_pfe_npu_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pfe_npu_resource_proto_goTypes = []any{
	(*JunosPfeNpu)(nil),              // 0: junos_pfe_npu
	(*JunosPfeNpuNpuMemoryList)(nil), // 1: junos_pfe_npu.npu_memory_list
	(*JuniperNetworksSensors)(nil),   // 2: JuniperNetworksSensors
}
var file_pfe_npu_resource_proto_depIdxs = []int32{
	1, // 0: junos_pfe_npu.npu_memory:type_name -> junos_pfe_npu.npu_memory_list
	2, // 1: jnpr_junos_pfe_npu_ext:extendee -> JuniperNetworksSensors
	0, // 2: jnpr_junos_pfe_npu_ext:type_name -> junos_pfe_npu
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pfe_npu_resource_proto_init() }
func file_pfe_npu_resource_proto_init() {
	if File_pfe_npu_resource_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pfe_npu_resource_proto_rawDesc), len(file_pfe_npu_resource_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_pfe_npu_resource_proto_goTypes,
		DependencyIndexes: file_pfe_npu_resource_proto_depIdxs,
		MessageInfos:      file_pfe_npu_resource_proto_msgTypes,
		ExtensionInfos:    file_pfe_npu_resource_proto_extTypes,
	}.Build()
	File_pfe_npu_resource_proto = out.File
	file_pfe_npu_resource_proto_goTypes = nil
	file_pfe_npu_resource_proto_depIdxs = nil
}
