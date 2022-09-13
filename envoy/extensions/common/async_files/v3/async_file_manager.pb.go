// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: envoy/extensions/common/async_files/v3/async_file_manager.proto

package async_filesv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

// Configuration to instantiate or select a singleton ``AsyncFileManager``.
type AsyncFileManagerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional identifier for the manager. An empty string is a valid identifier
	// for a common, default ``AsyncFileManager``.
	//
	// Reusing the same id with different configurations in the same envoy instance
	// is an error.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to ManagerType:
	//	*AsyncFileManagerConfig_ThreadPool_
	ManagerType isAsyncFileManagerConfig_ManagerType `protobuf_oneof:"manager_type"`
}

func (x *AsyncFileManagerConfig) Reset() {
	*x = AsyncFileManagerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncFileManagerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncFileManagerConfig) ProtoMessage() {}

func (x *AsyncFileManagerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncFileManagerConfig.ProtoReflect.Descriptor instead.
func (*AsyncFileManagerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *AsyncFileManagerConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *AsyncFileManagerConfig) GetManagerType() isAsyncFileManagerConfig_ManagerType {
	if m != nil {
		return m.ManagerType
	}
	return nil
}

func (x *AsyncFileManagerConfig) GetThreadPool() *AsyncFileManagerConfig_ThreadPool {
	if x, ok := x.GetManagerType().(*AsyncFileManagerConfig_ThreadPool_); ok {
		return x.ThreadPool
	}
	return nil
}

type isAsyncFileManagerConfig_ManagerType interface {
	isAsyncFileManagerConfig_ManagerType()
}

type AsyncFileManagerConfig_ThreadPool_ struct {
	// Configuration for a thread-pool based async file manager.
	ThreadPool *AsyncFileManagerConfig_ThreadPool `protobuf:"bytes,2,opt,name=thread_pool,json=threadPool,proto3,oneof"`
}

func (*AsyncFileManagerConfig_ThreadPool_) isAsyncFileManagerConfig_ManagerType() {}

type AsyncFileManagerConfig_ThreadPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of threads to use. If unset or zero, will default to the number
	// of concurrent threads the hardware supports. This default is subject to
	// change if performance analysis suggests it.
	ThreadCount uint32 `protobuf:"varint,1,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
}

func (x *AsyncFileManagerConfig_ThreadPool) Reset() {
	*x = AsyncFileManagerConfig_ThreadPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncFileManagerConfig_ThreadPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncFileManagerConfig_ThreadPool) ProtoMessage() {}

func (x *AsyncFileManagerConfig_ThreadPool) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncFileManagerConfig_ThreadPool.ProtoReflect.Descriptor instead.
func (*AsyncFileManagerConfig_ThreadPool) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AsyncFileManagerConfig_ThreadPool) GetThreadCount() uint32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

var File_envoy_extensions_common_async_files_v3_async_file_manager_proto protoreflect.FileDescriptor

var file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x16, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x6c, 0x0a,
	0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x49, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x73, 0x79, 0x6e,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x00, 0x52,
	0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x1a, 0x2f, 0x0a, 0x0a, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x0c,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42,
	0x01, 0x42, 0xbc, 0x01, 0x0a, 0x34, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x15, 0x41, 0x73, 0x79, 0x6e,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x76, 0x33, 0x3b, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescOnce sync.Once
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescData = file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDesc
)

func file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescGZIP() []byte {
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescData)
	})
	return file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDescData
}

var file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_common_async_files_v3_async_file_manager_proto_goTypes = []interface{}{
	(*AsyncFileManagerConfig)(nil),            // 0: envoy.extensions.common.async_files.v3.AsyncFileManagerConfig
	(*AsyncFileManagerConfig_ThreadPool)(nil), // 1: envoy.extensions.common.async_files.v3.AsyncFileManagerConfig.ThreadPool
}
var file_envoy_extensions_common_async_files_v3_async_file_manager_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.common.async_files.v3.AsyncFileManagerConfig.thread_pool:type_name -> envoy.extensions.common.async_files.v3.AsyncFileManagerConfig.ThreadPool
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_common_async_files_v3_async_file_manager_proto_init() }
func file_envoy_extensions_common_async_files_v3_async_file_manager_proto_init() {
	if File_envoy_extensions_common_async_files_v3_async_file_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncFileManagerConfig); i {
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
		file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncFileManagerConfig_ThreadPool); i {
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
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AsyncFileManagerConfig_ThreadPool_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_common_async_files_v3_async_file_manager_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_common_async_files_v3_async_file_manager_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_common_async_files_v3_async_file_manager_proto_msgTypes,
	}.Build()
	File_envoy_extensions_common_async_files_v3_async_file_manager_proto = out.File
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_rawDesc = nil
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_goTypes = nil
	file_envoy_extensions_common_async_files_v3_async_file_manager_proto_depIdxs = nil
}
