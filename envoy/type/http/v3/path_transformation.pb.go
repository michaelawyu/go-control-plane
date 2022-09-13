// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: envoy/type/http/v3/path_transformation.proto

package httpv3

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

type PathTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of operations to apply. Transformations will be performed in the order that they appear.
	Operations []*PathTransformation_Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *PathTransformation) Reset() {
	*x = PathTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathTransformation) ProtoMessage() {}

func (x *PathTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathTransformation.ProtoReflect.Descriptor instead.
func (*PathTransformation) Descriptor() ([]byte, []int) {
	return file_envoy_type_http_v3_path_transformation_proto_rawDescGZIP(), []int{0}
}

func (x *PathTransformation) GetOperations() []*PathTransformation_Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

// A type of operation to alter text.
type PathTransformation_Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OperationSpecifier:
	//	*PathTransformation_Operation_NormalizePathRfc_3986
	//	*PathTransformation_Operation_MergeSlashes_
	OperationSpecifier isPathTransformation_Operation_OperationSpecifier `protobuf_oneof:"operation_specifier"`
}

func (x *PathTransformation_Operation) Reset() {
	*x = PathTransformation_Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathTransformation_Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathTransformation_Operation) ProtoMessage() {}

func (x *PathTransformation_Operation) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathTransformation_Operation.ProtoReflect.Descriptor instead.
func (*PathTransformation_Operation) Descriptor() ([]byte, []int) {
	return file_envoy_type_http_v3_path_transformation_proto_rawDescGZIP(), []int{0, 0}
}

func (m *PathTransformation_Operation) GetOperationSpecifier() isPathTransformation_Operation_OperationSpecifier {
	if m != nil {
		return m.OperationSpecifier
	}
	return nil
}

func (x *PathTransformation_Operation) GetNormalizePathRfc_3986() *PathTransformation_Operation_NormalizePathRFC3986 {
	if x, ok := x.GetOperationSpecifier().(*PathTransformation_Operation_NormalizePathRfc_3986); ok {
		return x.NormalizePathRfc_3986
	}
	return nil
}

func (x *PathTransformation_Operation) GetMergeSlashes() *PathTransformation_Operation_MergeSlashes {
	if x, ok := x.GetOperationSpecifier().(*PathTransformation_Operation_MergeSlashes_); ok {
		return x.MergeSlashes
	}
	return nil
}

type isPathTransformation_Operation_OperationSpecifier interface {
	isPathTransformation_Operation_OperationSpecifier()
}

type PathTransformation_Operation_NormalizePathRfc_3986 struct {
	// Enable path normalization per RFC 3986.
	NormalizePathRfc_3986 *PathTransformation_Operation_NormalizePathRFC3986 `protobuf:"bytes,2,opt,name=normalize_path_rfc_3986,json=normalizePathRfc3986,proto3,oneof"`
}

type PathTransformation_Operation_MergeSlashes_ struct {
	// Enable merging adjacent slashes.
	MergeSlashes *PathTransformation_Operation_MergeSlashes `protobuf:"bytes,3,opt,name=merge_slashes,json=mergeSlashes,proto3,oneof"`
}

func (*PathTransformation_Operation_NormalizePathRfc_3986) isPathTransformation_Operation_OperationSpecifier() {
}

func (*PathTransformation_Operation_MergeSlashes_) isPathTransformation_Operation_OperationSpecifier() {
}

// Should text be normalized according to RFC 3986? This typically is used for path headers
// before any processing of requests by HTTP filters or routing. This applies percent-encoded
// normalization and path segment normalization. Fails on characters disallowed in URLs
// (e.g. NULLs). See `Normalization and Comparison
// <https://tools.ietf.org/html/rfc3986#section-6>`_ for details of normalization. Note that
// this options does not perform `case normalization
// <https://tools.ietf.org/html/rfc3986#section-6.2.2.1>`_
type PathTransformation_Operation_NormalizePathRFC3986 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PathTransformation_Operation_NormalizePathRFC3986) Reset() {
	*x = PathTransformation_Operation_NormalizePathRFC3986{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathTransformation_Operation_NormalizePathRFC3986) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathTransformation_Operation_NormalizePathRFC3986) ProtoMessage() {}

func (x *PathTransformation_Operation_NormalizePathRFC3986) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathTransformation_Operation_NormalizePathRFC3986.ProtoReflect.Descriptor instead.
func (*PathTransformation_Operation_NormalizePathRFC3986) Descriptor() ([]byte, []int) {
	return file_envoy_type_http_v3_path_transformation_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Determines if adjacent slashes are merged into one. A common use case is for a request path
// header. Using this option in ``:ref: PathNormalizationOptions
// <envoy_v3_api_msg_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.PathNormalizationOptions>``
// will allow incoming requests with path ``//dir///file`` to match against route with ``prefix``
// match set to ``/dir``. When using for header transformations, note that slash merging is not
// part of `HTTP spec <https://tools.ietf.org/html/rfc3986>`_ and is provided for convenience.
type PathTransformation_Operation_MergeSlashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PathTransformation_Operation_MergeSlashes) Reset() {
	*x = PathTransformation_Operation_MergeSlashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathTransformation_Operation_MergeSlashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathTransformation_Operation_MergeSlashes) ProtoMessage() {}

func (x *PathTransformation_Operation_MergeSlashes) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_http_v3_path_transformation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathTransformation_Operation_MergeSlashes.ProtoReflect.Descriptor instead.
func (*PathTransformation_Operation_MergeSlashes) Descriptor() ([]byte, []int) {
	return file_envoy_type_http_v3_path_transformation_proto_rawDescGZIP(), []int{0, 0, 1}
}

var File_envoy_type_http_v3_path_transformation_proto protoreflect.FileDescriptor

var file_envoy_type_http_v3_path_transformation_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x03, 0x0a, 0x12, 0x50,
	0x61, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x50, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0xb5, 0x02, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x7e, 0x0a, 0x17, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x72, 0x66, 0x63, 0x5f, 0x33, 0x39, 0x38, 0x36, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x45, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x46, 0x43, 0x33, 0x39, 0x38, 0x36, 0x48, 0x00, 0x52, 0x14, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x66, 0x63, 0x33, 0x39, 0x38,
	0x36, 0x12, 0x64, 0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x1a, 0x16, 0x0a, 0x14, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x46, 0x43, 0x33, 0x39, 0x38, 0x36, 0x1a,
	0x0e, 0x0a, 0x0c, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x42,
	0x1a, 0x0a, 0x13, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x87, 0x01, 0x0a, 0x20,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33,
	0x42, 0x17, 0x50, 0x61, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x76, 0x33, 0x3b, 0x68, 0x74, 0x74, 0x70, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_http_v3_path_transformation_proto_rawDescOnce sync.Once
	file_envoy_type_http_v3_path_transformation_proto_rawDescData = file_envoy_type_http_v3_path_transformation_proto_rawDesc
)

func file_envoy_type_http_v3_path_transformation_proto_rawDescGZIP() []byte {
	file_envoy_type_http_v3_path_transformation_proto_rawDescOnce.Do(func() {
		file_envoy_type_http_v3_path_transformation_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_http_v3_path_transformation_proto_rawDescData)
	})
	return file_envoy_type_http_v3_path_transformation_proto_rawDescData
}

var file_envoy_type_http_v3_path_transformation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_type_http_v3_path_transformation_proto_goTypes = []interface{}{
	(*PathTransformation)(nil),                                // 0: envoy.type.http.v3.PathTransformation
	(*PathTransformation_Operation)(nil),                      // 1: envoy.type.http.v3.PathTransformation.Operation
	(*PathTransformation_Operation_NormalizePathRFC3986)(nil), // 2: envoy.type.http.v3.PathTransformation.Operation.NormalizePathRFC3986
	(*PathTransformation_Operation_MergeSlashes)(nil),         // 3: envoy.type.http.v3.PathTransformation.Operation.MergeSlashes
}
var file_envoy_type_http_v3_path_transformation_proto_depIdxs = []int32{
	1, // 0: envoy.type.http.v3.PathTransformation.operations:type_name -> envoy.type.http.v3.PathTransformation.Operation
	2, // 1: envoy.type.http.v3.PathTransformation.Operation.normalize_path_rfc_3986:type_name -> envoy.type.http.v3.PathTransformation.Operation.NormalizePathRFC3986
	3, // 2: envoy.type.http.v3.PathTransformation.Operation.merge_slashes:type_name -> envoy.type.http.v3.PathTransformation.Operation.MergeSlashes
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_type_http_v3_path_transformation_proto_init() }
func file_envoy_type_http_v3_path_transformation_proto_init() {
	if File_envoy_type_http_v3_path_transformation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_http_v3_path_transformation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathTransformation); i {
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
		file_envoy_type_http_v3_path_transformation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathTransformation_Operation); i {
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
		file_envoy_type_http_v3_path_transformation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathTransformation_Operation_NormalizePathRFC3986); i {
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
		file_envoy_type_http_v3_path_transformation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathTransformation_Operation_MergeSlashes); i {
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
	file_envoy_type_http_v3_path_transformation_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PathTransformation_Operation_NormalizePathRfc_3986)(nil),
		(*PathTransformation_Operation_MergeSlashes_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_http_v3_path_transformation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_http_v3_path_transformation_proto_goTypes,
		DependencyIndexes: file_envoy_type_http_v3_path_transformation_proto_depIdxs,
		MessageInfos:      file_envoy_type_http_v3_path_transformation_proto_msgTypes,
	}.Build()
	File_envoy_type_http_v3_path_transformation_proto = out.File
	file_envoy_type_http_v3_path_transformation_proto_rawDesc = nil
	file_envoy_type_http_v3_path_transformation_proto_goTypes = nil
	file_envoy_type_http_v3_path_transformation_proto_depIdxs = nil
}
