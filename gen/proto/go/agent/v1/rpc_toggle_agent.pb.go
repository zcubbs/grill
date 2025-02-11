// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: agent/v1/rpc_toggle_agent.proto

package agentv1

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

type ToggleAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsActive bool   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *ToggleAgentRequest) Reset() {
	*x = ToggleAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_v1_rpc_toggle_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleAgentRequest) ProtoMessage() {}

func (x *ToggleAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_v1_rpc_toggle_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleAgentRequest.ProtoReflect.Descriptor instead.
func (*ToggleAgentRequest) Descriptor() ([]byte, []int) {
	return file_agent_v1_rpc_toggle_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ToggleAgentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToggleAgentRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ToggleAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ToggleAgentResponse) Reset() {
	*x = ToggleAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_v1_rpc_toggle_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleAgentResponse) ProtoMessage() {}

func (x *ToggleAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_v1_rpc_toggle_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleAgentResponse.ProtoReflect.Descriptor instead.
func (*ToggleAgentResponse) Descriptor() ([]byte, []int) {
	return file_agent_v1_rpc_toggle_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ToggleAgentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_agent_v1_rpc_toggle_agent_proto protoreflect.FileDescriptor

var file_agent_v1_rpc_toggle_agent_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x74,
	0x6f, 0x67, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x41, 0x0a, 0x12, 0x54,
	0x6f, 0x67, 0x67, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x2d,
	0x0a, 0x13, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x76, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x52,
	0x70, 0x63, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_v1_rpc_toggle_agent_proto_rawDescOnce sync.Once
	file_agent_v1_rpc_toggle_agent_proto_rawDescData = file_agent_v1_rpc_toggle_agent_proto_rawDesc
)

func file_agent_v1_rpc_toggle_agent_proto_rawDescGZIP() []byte {
	file_agent_v1_rpc_toggle_agent_proto_rawDescOnce.Do(func() {
		file_agent_v1_rpc_toggle_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_v1_rpc_toggle_agent_proto_rawDescData)
	})
	return file_agent_v1_rpc_toggle_agent_proto_rawDescData
}

var file_agent_v1_rpc_toggle_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agent_v1_rpc_toggle_agent_proto_goTypes = []interface{}{
	(*ToggleAgentRequest)(nil),  // 0: agent.v1.ToggleAgentRequest
	(*ToggleAgentResponse)(nil), // 1: agent.v1.ToggleAgentResponse
}
var file_agent_v1_rpc_toggle_agent_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_v1_rpc_toggle_agent_proto_init() }
func file_agent_v1_rpc_toggle_agent_proto_init() {
	if File_agent_v1_rpc_toggle_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_v1_rpc_toggle_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleAgentRequest); i {
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
		file_agent_v1_rpc_toggle_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleAgentResponse); i {
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
			RawDescriptor: file_agent_v1_rpc_toggle_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agent_v1_rpc_toggle_agent_proto_goTypes,
		DependencyIndexes: file_agent_v1_rpc_toggle_agent_proto_depIdxs,
		MessageInfos:      file_agent_v1_rpc_toggle_agent_proto_msgTypes,
	}.Build()
	File_agent_v1_rpc_toggle_agent_proto = out.File
	file_agent_v1_rpc_toggle_agent_proto_rawDesc = nil
	file_agent_v1_rpc_toggle_agent_proto_goTypes = nil
	file_agent_v1_rpc_toggle_agent_proto_depIdxs = nil
}
