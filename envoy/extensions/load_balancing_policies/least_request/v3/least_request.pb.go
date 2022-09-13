// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: envoy/extensions/load_balancing_policies/least_request/v3/least_request.proto

package least_requestv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This configuration allows the built-in LEAST_REQUEST LB policy to be configured via the LB policy
// extension point. See the :ref:`load balancing architecture overview
// <arch_overview_load_balancing_types>` for more information.
// [#extension: envoy.clusters.lb_policy]
type LeastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of random healthy hosts from which the host with the fewest active requests will
	// be chosen. Defaults to 2 so that we perform two-choice selection if the field is not set.
	ChoiceCount *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
	// The following formula is used to calculate the dynamic weights when hosts have different load
	// balancing weights:
	//
	// `weight = load_balancing_weight / (active_requests + 1)^active_request_bias`
	//
	// The larger the active request bias is, the more aggressively active requests will lower the
	// effective weight when all host weights are not equal.
	//
	// `active_request_bias` must be greater than or equal to 0.0.
	//
	// When `active_request_bias == 0.0` the Least Request Load Balancer doesn't consider the number
	// of active requests at the time it picks a host and behaves like the Round Robin Load
	// Balancer.
	//
	// When `active_request_bias > 0.0` the Least Request Load Balancer scales the load balancing
	// weight by the number of active requests at the time it does a pick.
	//
	// The value is cached for performance reasons and refreshed whenever one of the Load Balancer's
	// host sets changes, e.g., whenever there is a host membership update or a host load balancing
	// weight change.
	//
	// .. note::
	//   This setting only takes effect if all host weights are not equal.
	ActiveRequestBias *v3.RuntimeDouble `protobuf:"bytes,2,opt,name=active_request_bias,json=activeRequestBias,proto3" json:"active_request_bias,omitempty"`
	// Configuration for slow start mode.
	// If this configuration is not set, slow start will not be not enabled.
	SlowStartConfig *v31.Cluster_SlowStartConfig `protobuf:"bytes,3,opt,name=slow_start_config,json=slowStartConfig,proto3" json:"slow_start_config,omitempty"`
}

func (x *LeastRequest) Reset() {
	*x = LeastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeastRequest) ProtoMessage() {}

func (x *LeastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeastRequest.ProtoReflect.Descriptor instead.
func (*LeastRequest) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescGZIP(), []int{0}
}

func (x *LeastRequest) GetChoiceCount() *wrappers.UInt32Value {
	if x != nil {
		return x.ChoiceCount
	}
	return nil
}

func (x *LeastRequest) GetActiveRequestBias() *v3.RuntimeDouble {
	if x != nil {
		return x.ActiveRequestBias
	}
	return nil
}

func (x *LeastRequest) GetSlowStartConfig() *v31.Cluster_SlowStartConfig {
	if x != nil {
		return x.SlowStartConfig
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x0c, 0x4c,
	0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0c, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x69, 0x61, 0x73, 0x12, 0x5c, 0x0a, 0x11, 0x73, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x73, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xd8, 0x01, 0x0a, 0x47, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x70, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData = file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_goTypes = []interface{}{
	(*LeastRequest)(nil),                // 0: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest
	(*wrappers.UInt32Value)(nil),        // 1: google.protobuf.UInt32Value
	(*v3.RuntimeDouble)(nil),            // 2: envoy.config.core.v3.RuntimeDouble
	(*v31.Cluster_SlowStartConfig)(nil), // 3: envoy.config.cluster.v3.Cluster.SlowStartConfig
}
var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.choice_count:type_name -> google.protobuf.UInt32Value
	2, // 1: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.active_request_bias:type_name -> envoy.config.core.v3.RuntimeDouble
	3, // 2: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.slow_start_config:type_name -> envoy.config.cluster.v3.Cluster.SlowStartConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_init()
}
func file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_init() {
	if File_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeastRequest); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto = out.File
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_depIdxs = nil
}
