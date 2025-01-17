// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: api/mailer.proto

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

type SendConfirmationEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendConfirmationEmailRequest) Reset() {
	*x = SendConfirmationEmailRequest{}
	mi := &file_api_mailer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendConfirmationEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendConfirmationEmailRequest) ProtoMessage() {}

func (x *SendConfirmationEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_mailer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendConfirmationEmailRequest.ProtoReflect.Descriptor instead.
func (*SendConfirmationEmailRequest) Descriptor() ([]byte, []int) {
	return file_api_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *SendConfirmationEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendConfirmationEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendConfirmationEmailResponse) Reset() {
	*x = SendConfirmationEmailResponse{}
	mi := &file_api_mailer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendConfirmationEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendConfirmationEmailResponse) ProtoMessage() {}

func (x *SendConfirmationEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_mailer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendConfirmationEmailResponse.ProtoReflect.Descriptor instead.
func (*SendConfirmationEmailResponse) Descriptor() ([]byte, []int) {
	return file_api_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *SendConfirmationEmailResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetConfirmationCodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfirmationCodeRequest) Reset() {
	*x = GetConfirmationCodeRequest{}
	mi := &file_api_mailer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfirmationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfirmationCodeRequest) ProtoMessage() {}

func (x *GetConfirmationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_mailer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfirmationCodeRequest.ProtoReflect.Descriptor instead.
func (*GetConfirmationCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_mailer_proto_rawDescGZIP(), []int{2}
}

func (x *GetConfirmationCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetConfirmationCodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfirmationCodeResponse) Reset() {
	*x = GetConfirmationCodeResponse{}
	mi := &file_api_mailer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfirmationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfirmationCodeResponse) ProtoMessage() {}

func (x *GetConfirmationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_mailer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfirmationCodeResponse.ProtoReflect.Descriptor instead.
func (*GetConfirmationCodeResponse) Descriptor() ([]byte, []int) {
	return file_api_mailer_proto_rawDescGZIP(), []int{3}
}

func (x *GetConfirmationCodeResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_api_mailer_proto protoreflect.FileDescriptor

var file_api_mailer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x65, 0x61, 0x73, 0x79, 0x70, 0x77, 0x6e, 0x22, 0x34, 0x0a, 0x1c, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x33, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x31, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xd2, 0x01,
	0x0a, 0x06, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x25, 0x2e, 0x65, 0x61, 0x73, 0x79, 0x70, 0x77, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x61, 0x73, 0x79, 0x70,
	0x77, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x60, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x61, 0x73, 0x79, 0x70, 0x77,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65,
	0x61, 0x73, 0x79, 0x70, 0x77, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x61, 0x73, 0x79, 0x70, 0x77, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_mailer_proto_rawDescOnce sync.Once
	file_api_mailer_proto_rawDescData = file_api_mailer_proto_rawDesc
)

func file_api_mailer_proto_rawDescGZIP() []byte {
	file_api_mailer_proto_rawDescOnce.Do(func() {
		file_api_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mailer_proto_rawDescData)
	})
	return file_api_mailer_proto_rawDescData
}

var file_api_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_mailer_proto_goTypes = []any{
	(*SendConfirmationEmailRequest)(nil),  // 0: easypwn.SendConfirmationEmailRequest
	(*SendConfirmationEmailResponse)(nil), // 1: easypwn.SendConfirmationEmailResponse
	(*GetConfirmationCodeRequest)(nil),    // 2: easypwn.GetConfirmationCodeRequest
	(*GetConfirmationCodeResponse)(nil),   // 3: easypwn.GetConfirmationCodeResponse
}
var file_api_mailer_proto_depIdxs = []int32{
	0, // 0: easypwn.Mailer.SendConfirmationEmail:input_type -> easypwn.SendConfirmationEmailRequest
	2, // 1: easypwn.Mailer.GetConfirmationCode:input_type -> easypwn.GetConfirmationCodeRequest
	1, // 2: easypwn.Mailer.SendConfirmationEmail:output_type -> easypwn.SendConfirmationEmailResponse
	3, // 3: easypwn.Mailer.GetConfirmationCode:output_type -> easypwn.GetConfirmationCodeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_mailer_proto_init() }
func file_api_mailer_proto_init() {
	if File_api_mailer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_mailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mailer_proto_goTypes,
		DependencyIndexes: file_api_mailer_proto_depIdxs,
		MessageInfos:      file_api_mailer_proto_msgTypes,
	}.Build()
	File_api_mailer_proto = out.File
	file_api_mailer_proto_rawDesc = nil
	file_api_mailer_proto_goTypes = nil
	file_api_mailer_proto_depIdxs = nil
}
