// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: l2ald_oc_intf.proto

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

type InterfacesL2AlInterface struct {
	state         protoimpl.MessageState                  `protogen:"open.v1"`
	Interface     []*InterfacesL2AlInterfaceInterfaceList `protobuf:"bytes,151,rep,name=interface" json:"interface,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterface) Reset() {
	*x = InterfacesL2AlInterface{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterface) ProtoMessage() {}

func (x *InterfacesL2AlInterface) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterface.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterface) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0}
}

func (x *InterfacesL2AlInterface) GetInterface() []*InterfacesL2AlInterfaceInterfaceList {
	if x != nil {
		return x.Interface
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceList struct {
	state         protoimpl.MessageState                                 `protogen:"open.v1"`
	Name          *string                                                `protobuf:"bytes,51,opt,name=name" json:"name,omitempty"`
	Subinterfaces *InterfacesL2AlInterfaceInterfaceListSubinterfacesType `protobuf:"bytes,151,opt,name=subinterfaces" json:"subinterfaces,omitempty"`
	Ethernet      *InterfacesL2AlInterfaceInterfaceListEthernetType      `protobuf:"bytes,152,opt,name=ethernet" json:"ethernet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceList) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceList{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceList) ProtoMessage() {}

func (x *InterfacesL2AlInterfaceInterfaceList) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceList.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceList) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceList) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *InterfacesL2AlInterfaceInterfaceList) GetSubinterfaces() *InterfacesL2AlInterfaceInterfaceListSubinterfacesType {
	if x != nil {
		return x.Subinterfaces
	}
	return nil
}

func (x *InterfacesL2AlInterfaceInterfaceList) GetEthernet() *InterfacesL2AlInterfaceInterfaceListEthernetType {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesType struct {
	state         protoimpl.MessageState                                                   `protogen:"open.v1"`
	Subinterface  []*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList `protobuf:"bytes,151,rep,name=subinterface" json:"subinterface,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListSubinterfacesType{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesType) ProtoMessage() {}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListSubinterfacesType.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesType) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesType) GetSubinterface() []*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList {
	if x != nil {
		return x.Subinterface
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListEthernetType struct {
	state         protoimpl.MessageState                                            `protogen:"open.v1"`
	SwitchedVlan  *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType `protobuf:"bytes,151,opt,name=switched_vlan,json=switchedVlan" json:"switched_vlan,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetType) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListEthernetType{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListEthernetType) ProtoMessage() {}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetType) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListEthernetType.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListEthernetType) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetType) GetSwitchedVlan() *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType {
	if x != nil {
		return x.SwitchedVlan
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList struct {
	state         protoimpl.MessageState                                                         `protogen:"open.v1"`
	Index         *uint32                                                                        `protobuf:"varint,51,opt,name=index" json:"index,omitempty"`
	Vlan          *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType `protobuf:"bytes,151,opt,name=vlan" json:"vlan,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) ProtoMessage() {}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList) GetVlan() *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType {
	if x != nil {
		return x.Vlan
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType struct {
	state         protoimpl.MessageState                                                                  `protogen:"open.v1"`
	State         *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType `protobuf:"bytes,152,opt,name=state" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) ProtoMessage() {
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType) GetState() *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType {
	if x != nil {
		return x.State
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VlanId        *uint32                `protobuf:"varint,51,opt,name=vlan_id,json=vlanId" json:"vlan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) ProtoMessage() {
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType) GetVlanId() uint32 {
	if x != nil && x.VlanId != nil {
		return *x.VlanId
	}
	return 0
}

type InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType struct {
	state         protoimpl.MessageState                                                     `protogen:"open.v1"`
	State         *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) ProtoMessage() {}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 1, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType) GetState() *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType {
	if x != nil {
		return x.State
	}
	return nil
}

type InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InterfaceMode *string                `protobuf:"bytes,51,opt,name=interface_mode,json=interfaceMode" json:"interface_mode,omitempty"`
	NativeVlan    *uint32                `protobuf:"varint,52,opt,name=native_vlan,json=nativeVlan" json:"native_vlan,omitempty"`
	AccessVlan    *uint32                `protobuf:"varint,53,opt,name=access_vlan,json=accessVlan" json:"access_vlan,omitempty"`
	TrunkVlans    []string               `protobuf:"bytes,54,rep,name=trunk_vlans,json=trunkVlans" json:"trunk_vlans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) Reset() {
	*x = InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType{}
	mi := &file_l2ald_oc_intf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) ProtoMessage() {}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) ProtoReflect() protoreflect.Message {
	mi := &file_l2ald_oc_intf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType.ProtoReflect.Descriptor instead.
func (*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) Descriptor() ([]byte, []int) {
	return file_l2ald_oc_intf_proto_rawDescGZIP(), []int{0, 0, 1, 0, 0}
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetInterfaceMode() string {
	if x != nil && x.InterfaceMode != nil {
		return *x.InterfaceMode
	}
	return ""
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetNativeVlan() uint32 {
	if x != nil && x.NativeVlan != nil {
		return *x.NativeVlan
	}
	return 0
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetAccessVlan() uint32 {
	if x != nil && x.AccessVlan != nil {
		return *x.AccessVlan
	}
	return 0
}

func (x *InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType) GetTrunkVlans() []string {
	if x != nil {
		return x.TrunkVlans
	}
	return nil
}

var file_l2ald_oc_intf_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*JuniperNetworksSensors)(nil),
		ExtensionType: (*InterfacesL2AlInterface)(nil),
		Field:         49,
		Name:          "jnpr_interfaces_l2al_interface_ext",
		Tag:           "bytes,49,opt,name=jnpr_interfaces_l2al_interface_ext",
		Filename:      "l2ald_oc_intf.proto",
	},
}

// Extension fields to JuniperNetworksSensors.
var (
	// optional interfaces_l2al_interface jnpr_interfaces_l2al_interface_ext = 49;
	E_JnprInterfacesL2AlInterfaceExt = &file_l2ald_oc_intf_proto_extTypes[0]
)

var File_l2ald_oc_intf_proto protoreflect.FileDescriptor

var file_l2ald_oc_intf_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x6c, 0x32, 0x61, 0x6c, 0x64, 0x5f, 0x6f, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x5f, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x09, 0x0a, 0x19, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x97, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x1a, 0xe7, 0x08, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0x82, 0x40, 0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x63, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x18, 0x98, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x1a, 0xdc, 0x03, 0x0a, 0x12,
	0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x73, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x18, 0x97, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x1a, 0xd0, 0x02, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x82, 0x40,
	0x02, 0x08, 0x01, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x6d, 0x0a, 0x04, 0x76, 0x6c,
	0x61, 0x6e, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x1a, 0xae, 0x01, 0x0a, 0x09, 0x76, 0x6c,
	0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x7a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x98, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x25, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x1a, 0x9f, 0x03, 0x0a, 0x0d, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x70, 0x0a, 0x0d,
	0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x97, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x56, 0x6c, 0x61, 0x6e, 0x1a, 0x9b,
	0x02, 0x0a, 0x12, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x76, 0x6c, 0x61, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x6c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x97,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x5f, 0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x73,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x96, 0x01, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x56, 0x6c, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x56, 0x6c, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x72, 0x75, 0x6e, 0x6b, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x36, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x72, 0x75, 0x6e, 0x6b, 0x56, 0x6c, 0x61, 0x6e, 0x73, 0x3a, 0x7f, 0x0a, 0x22,
	0x6a, 0x6e, 0x70, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f,
	0x6c, 0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x65,
	0x78, 0x74, 0x12, 0x17, 0x2e, 0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x31, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x6c,
	0x32, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x1e, 0x6a,
	0x6e, 0x70, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x4c, 0x32, 0x61,
	0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x45, 0x78, 0x74, 0x42, 0x1d, 0x5a,
	0x1b, 0x2e, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
})

var (
	file_l2ald_oc_intf_proto_rawDescOnce sync.Once
	file_l2ald_oc_intf_proto_rawDescData []byte
)

func file_l2ald_oc_intf_proto_rawDescGZIP() []byte {
	file_l2ald_oc_intf_proto_rawDescOnce.Do(func() {
		file_l2ald_oc_intf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_l2ald_oc_intf_proto_rawDesc), len(file_l2ald_oc_intf_proto_rawDesc)))
	})
	return file_l2ald_oc_intf_proto_rawDescData
}

var file_l2ald_oc_intf_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_l2ald_oc_intf_proto_goTypes = []any{
	(*InterfacesL2AlInterface)(nil),                                                                // 0: interfaces_l2al_interface
	(*InterfacesL2AlInterfaceInterfaceList)(nil),                                                   // 1: interfaces_l2al_interface.interface_list
	(*InterfacesL2AlInterfaceInterfaceListSubinterfacesType)(nil),                                  // 2: interfaces_l2al_interface.interface_list.subinterfaces_type
	(*InterfacesL2AlInterfaceInterfaceListEthernetType)(nil),                                       // 3: interfaces_l2al_interface.interface_list.ethernet_type
	(*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceList)(nil),                  // 4: interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list
	(*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanType)(nil),          // 5: interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type
	(*InterfacesL2AlInterfaceInterfaceListSubinterfacesTypeSubinterfaceListVlanTypeStateType)(nil), // 6: interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type.state_type
	(*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanType)(nil),                       // 7: interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type
	(*InterfacesL2AlInterfaceInterfaceListEthernetTypeSwitchedVlanTypeStateType)(nil),              // 8: interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type.state_type
	(*JuniperNetworksSensors)(nil),                                                                 // 9: JuniperNetworksSensors
}
var file_l2ald_oc_intf_proto_depIdxs = []int32{
	1,  // 0: interfaces_l2al_interface.interface:type_name -> interfaces_l2al_interface.interface_list
	2,  // 1: interfaces_l2al_interface.interface_list.subinterfaces:type_name -> interfaces_l2al_interface.interface_list.subinterfaces_type
	3,  // 2: interfaces_l2al_interface.interface_list.ethernet:type_name -> interfaces_l2al_interface.interface_list.ethernet_type
	4,  // 3: interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface:type_name -> interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list
	7,  // 4: interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan:type_name -> interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type
	5,  // 5: interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan:type_name -> interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type
	6,  // 6: interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type.state:type_name -> interfaces_l2al_interface.interface_list.subinterfaces_type.subinterface_list.vlan_type.state_type
	8,  // 7: interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type.state:type_name -> interfaces_l2al_interface.interface_list.ethernet_type.switched_vlan_type.state_type
	9,  // 8: jnpr_interfaces_l2al_interface_ext:extendee -> JuniperNetworksSensors
	0,  // 9: jnpr_interfaces_l2al_interface_ext:type_name -> interfaces_l2al_interface
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	9,  // [9:10] is the sub-list for extension type_name
	8,  // [8:9] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_l2ald_oc_intf_proto_init() }
func file_l2ald_oc_intf_proto_init() {
	if File_l2ald_oc_intf_proto != nil {
		return
	}
	file_telemetry_top_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_l2ald_oc_intf_proto_rawDesc), len(file_l2ald_oc_intf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_l2ald_oc_intf_proto_goTypes,
		DependencyIndexes: file_l2ald_oc_intf_proto_depIdxs,
		MessageInfos:      file_l2ald_oc_intf_proto_msgTypes,
		ExtensionInfos:    file_l2ald_oc_intf_proto_extTypes,
	}.Build()
	File_l2ald_oc_intf_proto = out.File
	file_l2ald_oc_intf_proto_goTypes = nil
	file_l2ald_oc_intf_proto_depIdxs = nil
}
