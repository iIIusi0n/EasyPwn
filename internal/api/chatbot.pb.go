// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: api/chatbot.proto

package api

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

type GetResponseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponseRequest) Reset() {
	*x = GetResponseRequest{}
	mi := &file_api_chatbot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponseRequest) ProtoMessage() {}

func (x *GetResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatbot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponseRequest.ProtoReflect.Descriptor instead.
func (*GetResponseRequest) Descriptor() ([]byte, []int) {
	return file_api_chatbot_proto_rawDescGZIP(), []int{0}
}

func (x *GetResponseRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GetResponseRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetResponseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponseResponse) Reset() {
	*x = GetResponseResponse{}
	mi := &file_api_chatbot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponseResponse) ProtoMessage() {}

func (x *GetResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chatbot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponseResponse.ProtoReflect.Descriptor instead.
func (*GetResponseResponse) Descriptor() ([]byte, []int) {
	return file_api_chatbot_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponseResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_api_chatbot_proto protoreflect.FileDescriptor

var file_api_chatbot_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x61, 0x73, 0x79, 0x70, 0x77, 0x6e, 0x22, 0x4f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x55, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x12, 0x4a, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x65, 0x61, 0x73,
	0x79, 0x70, 0x77, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x61, 0x73, 0x79, 0x70, 0x77,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x61, 0x73, 0x79, 0x70,
	0x77, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chatbot_proto_rawDescOnce sync.Once
	file_api_chatbot_proto_rawDescData = file_api_chatbot_proto_rawDesc
)

func file_api_chatbot_proto_rawDescGZIP() []byte {
	file_api_chatbot_proto_rawDescOnce.Do(func() {
		file_api_chatbot_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chatbot_proto_rawDescData)
	})
	return file_api_chatbot_proto_rawDescData
}

var file_api_chatbot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_chatbot_proto_goTypes = []any{
	(*GetResponseRequest)(nil),  // 0: easypwn.GetResponseRequest
	(*GetResponseResponse)(nil), // 1: easypwn.GetResponseResponse
}
var file_api_chatbot_proto_depIdxs = []int32{
	0, // 0: easypwn.Chatbot.GetResponse:input_type -> easypwn.GetResponseRequest
	1, // 1: easypwn.Chatbot.GetResponse:output_type -> easypwn.GetResponseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_chatbot_proto_init() }
func file_api_chatbot_proto_init() {
	if File_api_chatbot_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_chatbot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chatbot_proto_goTypes,
		DependencyIndexes: file_api_chatbot_proto_depIdxs,
		MessageInfos:      file_api_chatbot_proto_msgTypes,
	}.Build()
	File_api_chatbot_proto = out.File
	file_api_chatbot_proto_rawDesc = nil
	file_api_chatbot_proto_goTypes = nil
	file_api_chatbot_proto_depIdxs = nil
}
