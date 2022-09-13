// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: contrib/envoy/extensions/network/connection_balance/dlb/v3alpha/dlb.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// The Dlb is a hardware managed system of queues and arbiters connecting producers and consumers. It is a PCIE device
// in the CPU package. It interacts with software running on cores and potentially other devices. The Dlb implements the
// following balancing features:
//
// -  Lock-free multi-producer/multi-consumer operation.
// -  Multiple priorities for varying traffic types.
// -  Various distribution schemes.
//
// Dlb connection balancer uses Dlb hardware to balance connections, and can significantly reduce latency.
//
// As the Dlb connection balancer provides assistance from dedicated Dlb hardware, it can be used for a proxy with a large number of connections
// (e.g., a gateway).
type Dlb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Dlb hardware, start from 0.
	// If not specified, use the first available device as default.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Dlb) Reset() {
	*x = Dlb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dlb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dlb) ProtoMessage() {}

func (x *Dlb) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dlb.ProtoReflect.Descriptor instead.
func (*Dlb) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescGZIP(), []int{0}
}

func (x *Dlb) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x64, 0x6c, 0x62, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x64, 0x6c, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x6c, 0x62, 0x2e, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x03, 0x44, 0x6c, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x42, 0xbb, 0x01, 0x0a, 0x45, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x6c, 0x62, 0x2e, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x42, 0x08, 0x44, 0x6c, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x64, 0x6c, 0x62, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescData = file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDesc
)

func file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDescData
}

var file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_goTypes = []interface{}{
	(*Dlb)(nil), // 0: envoy.extensions.network.connection_balance.dlb.v3alpha.Dlb
}
var file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_init() }
func file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_init() {
	if File_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dlb); i {
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
			RawDescriptor: file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto = out.File
	file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_rawDesc = nil
	file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_goTypes = nil
	file_contrib_envoy_extensions_network_connection_balance_dlb_v3alpha_dlb_proto_depIdxs = nil
}
