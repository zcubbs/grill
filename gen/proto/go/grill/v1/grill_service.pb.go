// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: grill/v1/grill_service.proto

package grillv1

import (
	_ "github.com/zcubbs/grill/google/api"
	_ "github.com/zcubbs/grill/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grill_v1_grill_service_proto protoreflect.FileDescriptor

var file_grill_v1_grill_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x69, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x69, 0x6c, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x67, 0x72, 0x69, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x67, 0x72, 0x69, 0x6c, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x81, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x71, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x69, 0x6c, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x72, 0x69, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x92, 0x41, 0x27, 0x0a, 0x03, 0x4f, 0x70,
	0x73, 0x12, 0x0f, 0x50, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x1a, 0x0f, 0x50, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x69, 0x6e, 0x67, 0x42, 0x8c, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x69, 0x6c,
	0x6c, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x47, 0x72, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x63, 0x75, 0x62, 0x62, 0x73, 0x2f, 0x67, 0x72, 0x69,
	0x6c, 0x6c, 0x2f, 0x67, 0x72, 0x69, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x69, 0x6c,
	0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x47, 0x72, 0x69, 0x6c,
	0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x47, 0x72, 0x69, 0x6c, 0x6c, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x14, 0x47, 0x72, 0x69, 0x6c, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x47, 0x72, 0x69, 0x6c, 0x6c, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_grill_v1_grill_service_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: grill.v1.PingRequest
	(*PingResponse)(nil), // 1: grill.v1.PingResponse
}
var file_grill_v1_grill_service_proto_depIdxs = []int32{
	0, // 0: grill.v1.GrillService.Ping:input_type -> grill.v1.PingRequest
	1, // 1: grill.v1.GrillService.Ping:output_type -> grill.v1.PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grill_v1_grill_service_proto_init() }
func file_grill_v1_grill_service_proto_init() {
	if File_grill_v1_grill_service_proto != nil {
		return
	}
	file_grill_v1_rpc_ping_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grill_v1_grill_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grill_v1_grill_service_proto_goTypes,
		DependencyIndexes: file_grill_v1_grill_service_proto_depIdxs,
	}.Build()
	File_grill_v1_grill_service_proto = out.File
	file_grill_v1_grill_service_proto_rawDesc = nil
	file_grill_v1_grill_service_proto_goTypes = nil
	file_grill_v1_grill_service_proto_depIdxs = nil
}
