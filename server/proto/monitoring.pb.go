// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: proto/monitoring.proto

package monitor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the usage data.
type UsageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUsage    float32      `protobuf:"fixed32,1,opt,name=cpuUsage,proto3" json:"cpuUsage,omitempty"`
	RamUsage    *MemoryUsage `protobuf:"bytes,2,opt,name=ramUsage,proto3" json:"ramUsage,omitempty"`
	HostDetails *HostDetails `protobuf:"bytes,3,opt,name=hostDetails,proto3" json:"hostDetails,omitempty"`
}

func (x *UsageData) Reset() {
	*x = UsageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageData) ProtoMessage() {}

func (x *UsageData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageData.ProtoReflect.Descriptor instead.
func (*UsageData) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *UsageData) GetCpuUsage() float32 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *UsageData) GetRamUsage() *MemoryUsage {
	if x != nil {
		return x.RamUsage
	}
	return nil
}

func (x *UsageData) GetHostDetails() *HostDetails {
	if x != nil {
		return x.HostDetails
	}
	return nil
}

// Nested message for RAM usage details.
type MemoryUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Free        float32 `protobuf:"fixed32,1,opt,name=free,proto3" json:"free,omitempty"`
	UsedPercent float32 `protobuf:"fixed32,2,opt,name=usedPercent,proto3" json:"usedPercent,omitempty"`
	Total       float32 `protobuf:"fixed32,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *MemoryUsage) Reset() {
	*x = MemoryUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryUsage) ProtoMessage() {}

func (x *MemoryUsage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryUsage.ProtoReflect.Descriptor instead.
func (*MemoryUsage) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *MemoryUsage) GetFree() float32 {
	if x != nil {
		return x.Free
	}
	return 0
}

func (x *MemoryUsage) GetUsedPercent() float32 {
	if x != nil {
		return x.UsedPercent
	}
	return 0
}

func (x *MemoryUsage) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type HostDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname      string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Os            string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	KernelVersion string `protobuf:"bytes,3,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	KernelArch    string `protobuf:"bytes,4,opt,name=kernelArch,proto3" json:"kernelArch,omitempty"`
	BootTime      uint64 `protobuf:"varint,5,opt,name=bootTime,proto3" json:"bootTime,omitempty"`
	Uptime        uint64 `protobuf:"varint,6,opt,name=uptime,proto3" json:"uptime,omitempty"`
	HostId        string `protobuf:"bytes,7,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *HostDetails) Reset() {
	*x = HostDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostDetails) ProtoMessage() {}

func (x *HostDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostDetails.ProtoReflect.Descriptor instead.
func (*HostDetails) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{2}
}

func (x *HostDetails) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *HostDetails) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *HostDetails) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *HostDetails) GetKernelArch() string {
	if x != nil {
		return x.KernelArch
	}
	return ""
}

func (x *HostDetails) GetBootTime() uint64 {
	if x != nil {
		return x.BootTime
	}
	return 0
}

func (x *HostDetails) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *HostDetails) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

// The response message containing the server's response.
type UsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsageResponse) Reset() {
	*x = UsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_monitoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageResponse) ProtoMessage() {}

func (x *UsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_monitoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageResponse.ProtoReflect.Descriptor instead.
func (*UsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_monitoring_proto_rawDescGZIP(), []int{3}
}

func (x *UsageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_monitoring_proto protoreflect.FileDescriptor

var file_proto_monitoring_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33,
	0x0a, 0x08, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x72, 0x61, 0x6d, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x59,
	0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x66, 0x72, 0x65,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x48, 0x6f,
	0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x41, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x6f, 0x6f, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62,
	0x6f, 0x6f, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x51, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x19, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x66, 0x72, 0x65, 0x64, 0x6f, 0x73, 0x61, 0x2f, 0x67, 0x6f,
	0x2d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_monitoring_proto_rawDescOnce sync.Once
	file_proto_monitoring_proto_rawDescData = file_proto_monitoring_proto_rawDesc
)

func file_proto_monitoring_proto_rawDescGZIP() []byte {
	file_proto_monitoring_proto_rawDescOnce.Do(func() {
		file_proto_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_monitoring_proto_rawDescData)
	})
	return file_proto_monitoring_proto_rawDescData
}

var file_proto_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_monitoring_proto_goTypes = []interface{}{
	(*UsageData)(nil),     // 0: monitoring.UsageData
	(*MemoryUsage)(nil),   // 1: monitoring.MemoryUsage
	(*HostDetails)(nil),   // 2: monitoring.HostDetails
	(*UsageResponse)(nil), // 3: monitoring.UsageResponse
}
var file_proto_monitoring_proto_depIdxs = []int32{
	1, // 0: monitoring.UsageData.ramUsage:type_name -> monitoring.MemoryUsage
	2, // 1: monitoring.UsageData.hostDetails:type_name -> monitoring.HostDetails
	0, // 2: monitoring.MonitorService.SendUsage:input_type -> monitoring.UsageData
	3, // 3: monitoring.MonitorService.SendUsage:output_type -> monitoring.UsageResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_monitoring_proto_init() }
func file_proto_monitoring_proto_init() {
	if File_proto_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryUsage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_monitoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_monitoring_proto_goTypes,
		DependencyIndexes: file_proto_monitoring_proto_depIdxs,
		MessageInfos:      file_proto_monitoring_proto_msgTypes,
	}.Build()
	File_proto_monitoring_proto = out.File
	file_proto_monitoring_proto_rawDesc = nil
	file_proto_monitoring_proto_goTypes = nil
	file_proto_monitoring_proto_depIdxs = nil
}
