// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: matcher.proto

package pb

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

type MatchScenarioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPrompt string `protobuf:"bytes,1,opt,name=user_prompt,json=userPrompt,proto3" json:"user_prompt,omitempty"`
}

func (x *MatchScenarioRequest) Reset() {
	*x = MatchScenarioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchScenarioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchScenarioRequest) ProtoMessage() {}

func (x *MatchScenarioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchScenarioRequest.ProtoReflect.Descriptor instead.
func (*MatchScenarioRequest) Descriptor() ([]byte, []int) {
	return file_matcher_proto_rawDescGZIP(), []int{0}
}

func (x *MatchScenarioRequest) GetUserPrompt() string {
	if x != nil {
		return x.UserPrompt
	}
	return ""
}

type MatchScenarioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScenarioName string `protobuf:"bytes,1,opt,name=scenario_name,json=scenarioName,proto3" json:"scenario_name,omitempty"`
	RootId       string `protobuf:"bytes,2,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
}

func (x *MatchScenarioResponse) Reset() {
	*x = MatchScenarioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchScenarioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchScenarioResponse) ProtoMessage() {}

func (x *MatchScenarioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchScenarioResponse.ProtoReflect.Descriptor instead.
func (*MatchScenarioResponse) Descriptor() ([]byte, []int) {
	return file_matcher_proto_rawDescGZIP(), []int{1}
}

func (x *MatchScenarioResponse) GetScenarioName() string {
	if x != nil {
		return x.ScenarioName
	}
	return ""
}

func (x *MatchScenarioResponse) GetRootId() string {
	if x != nil {
		return x.RootId
	}
	return ""
}

var File_matcher_proto protoreflect.FileDescriptor

var file_matcher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x37, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x32,
	0x49, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0d, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x15, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x65, 0x77, 0x2d,
	0x74, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x65, 0x6c, 0x65, 0x73, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matcher_proto_rawDescOnce sync.Once
	file_matcher_proto_rawDescData = file_matcher_proto_rawDesc
)

func file_matcher_proto_rawDescGZIP() []byte {
	file_matcher_proto_rawDescOnce.Do(func() {
		file_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_matcher_proto_rawDescData)
	})
	return file_matcher_proto_rawDescData
}

var file_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_matcher_proto_goTypes = []interface{}{
	(*MatchScenarioRequest)(nil),  // 0: MatchScenarioRequest
	(*MatchScenarioResponse)(nil), // 1: MatchScenarioResponse
}
var file_matcher_proto_depIdxs = []int32{
	0, // 0: Matcher.MatchScenario:input_type -> MatchScenarioRequest
	1, // 1: Matcher.MatchScenario:output_type -> MatchScenarioResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_matcher_proto_init() }
func file_matcher_proto_init() {
	if File_matcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_matcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchScenarioRequest); i {
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
		file_matcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchScenarioResponse); i {
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
			RawDescriptor: file_matcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_matcher_proto_goTypes,
		DependencyIndexes: file_matcher_proto_depIdxs,
		MessageInfos:      file_matcher_proto_msgTypes,
	}.Build()
	File_matcher_proto = out.File
	file_matcher_proto_rawDesc = nil
	file_matcher_proto_goTypes = nil
	file_matcher_proto_depIdxs = nil
}
