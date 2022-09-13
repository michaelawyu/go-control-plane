// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: envoy/config/core/v3/socket_option.proto

package corev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type SocketOption_SocketState int32

const (
	// Socket options are applied after socket creation but before binding the socket to a port
	SocketOption_STATE_PREBIND SocketOption_SocketState = 0
	// Socket options are applied after binding the socket to a port but before calling listen()
	SocketOption_STATE_BOUND SocketOption_SocketState = 1
	// Socket options are applied after calling listen()
	SocketOption_STATE_LISTENING SocketOption_SocketState = 2
)

// Enum value maps for SocketOption_SocketState.
var (
	SocketOption_SocketState_name = map[int32]string{
		0: "STATE_PREBIND",
		1: "STATE_BOUND",
		2: "STATE_LISTENING",
	}
	SocketOption_SocketState_value = map[string]int32{
		"STATE_PREBIND":   0,
		"STATE_BOUND":     1,
		"STATE_LISTENING": 2,
	}
)

func (x SocketOption_SocketState) Enum() *SocketOption_SocketState {
	p := new(SocketOption_SocketState)
	*p = x
	return p
}

func (x SocketOption_SocketState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketOption_SocketState) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_core_v3_socket_option_proto_enumTypes[0].Descriptor()
}

func (SocketOption_SocketState) Type() protoreflect.EnumType {
	return &file_envoy_config_core_v3_socket_option_proto_enumTypes[0]
}

func (x SocketOption_SocketState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketOption_SocketState.Descriptor instead.
func (SocketOption_SocketState) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_socket_option_proto_rawDescGZIP(), []int{0, 0}
}

// Generic socket option message. This would be used to set socket options that
// might not exist in upstream kernels or precompiled Envoy binaries.
//
// For example:
//
// .. code-block:: json
//
//  {
//    "description": "support tcp keep alive",
//    "state": 0,
//    "level": 1,
//    "name": 9,
//    "int_value": 1,
//  }
//
// 1 means SOL_SOCKET and 9 means SO_KEEPALIVE on Linux.
// With the above configuration, `TCP Keep-Alives <https://www.freesoft.org/CIE/RFC/1122/114.htm>`_
// can be enabled in socket with Linux, which can be used in
// :ref:`listener's<envoy_v3_api_field_config.listener.v3.Listener.socket_options>` or
// :ref:`admin's <envoy_v3_api_field_config.bootstrap.v3.Admin.socket_options>` socket_options etc.
//
// It should be noted that the name or level may have different values on different platforms.
// [#next-free-field: 7]
type SocketOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional name to give this socket option for debugging, etc.
	// Uniqueness is not required and no special meaning is assumed.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Corresponding to the level value passed to setsockopt, such as IPPROTO_TCP
	Level int64 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	// The numeric name as passed to setsockopt
	Name int64 `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//	*SocketOption_IntValue
	//	*SocketOption_BufValue
	Value isSocketOption_Value `protobuf_oneof:"value"`
	// The state in which the option will be applied. When used in BindConfig
	// STATE_PREBIND is currently the only valid value.
	State SocketOption_SocketState `protobuf:"varint,6,opt,name=state,proto3,enum=envoy.config.core.v3.SocketOption_SocketState" json:"state,omitempty"`
}

func (x *SocketOption) Reset() {
	*x = SocketOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_socket_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketOption) ProtoMessage() {}

func (x *SocketOption) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_socket_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketOption.ProtoReflect.Descriptor instead.
func (*SocketOption) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_socket_option_proto_rawDescGZIP(), []int{0}
}

func (x *SocketOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SocketOption) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SocketOption) GetName() int64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (m *SocketOption) GetValue() isSocketOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SocketOption) GetIntValue() int64 {
	if x, ok := x.GetValue().(*SocketOption_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *SocketOption) GetBufValue() []byte {
	if x, ok := x.GetValue().(*SocketOption_BufValue); ok {
		return x.BufValue
	}
	return nil
}

func (x *SocketOption) GetState() SocketOption_SocketState {
	if x != nil {
		return x.State
	}
	return SocketOption_STATE_PREBIND
}

type isSocketOption_Value interface {
	isSocketOption_Value()
}

type SocketOption_IntValue struct {
	// Because many sockopts take an int value.
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof"`
}

type SocketOption_BufValue struct {
	// Otherwise it's a byte buffer.
	BufValue []byte `protobuf:"bytes,5,opt,name=buf_value,json=bufValue,proto3,oneof"`
}

func (*SocketOption_IntValue) isSocketOption_Value() {}

func (*SocketOption_BufValue) isSocketOption_Value() {}

var File_envoy_config_core_v3_socket_option_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_socket_option_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x0c,
	0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x62, 0x75,
	0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x0b, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x52, 0x45, 0x42, 0x49, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x3a, 0x25,
	0x9a, 0xc5, 0x88, 0x1e, 0x20, 0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x42, 0x85, 0x01, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x72,
	0x65, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_socket_option_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_socket_option_proto_rawDescData = file_envoy_config_core_v3_socket_option_proto_rawDesc
)

func file_envoy_config_core_v3_socket_option_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_socket_option_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_socket_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_socket_option_proto_rawDescData)
	})
	return file_envoy_config_core_v3_socket_option_proto_rawDescData
}

var file_envoy_config_core_v3_socket_option_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_core_v3_socket_option_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v3_socket_option_proto_goTypes = []interface{}{
	(SocketOption_SocketState)(0), // 0: envoy.config.core.v3.SocketOption.SocketState
	(*SocketOption)(nil),          // 1: envoy.config.core.v3.SocketOption
}
var file_envoy_config_core_v3_socket_option_proto_depIdxs = []int32{
	0, // 0: envoy.config.core.v3.SocketOption.state:type_name -> envoy.config.core.v3.SocketOption.SocketState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_socket_option_proto_init() }
func file_envoy_config_core_v3_socket_option_proto_init() {
	if File_envoy_config_core_v3_socket_option_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_socket_option_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketOption); i {
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
	file_envoy_config_core_v3_socket_option_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SocketOption_IntValue)(nil),
		(*SocketOption_BufValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_socket_option_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_socket_option_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_socket_option_proto_depIdxs,
		EnumInfos:         file_envoy_config_core_v3_socket_option_proto_enumTypes,
		MessageInfos:      file_envoy_config_core_v3_socket_option_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_socket_option_proto = out.File
	file_envoy_config_core_v3_socket_option_proto_rawDesc = nil
	file_envoy_config_core_v3_socket_option_proto_goTypes = nil
	file_envoy_config_core_v3_socket_option_proto_depIdxs = nil
}
