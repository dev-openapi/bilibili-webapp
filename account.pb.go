// B站开放接口 - 用户管理 bilibili-webapp

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: bilibili-webapp/account.proto

package bilibili_webapp

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetAccountInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetAccountInfoReq) Reset() {
	*x = GetAccountInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountInfoReq) ProtoMessage() {}

func (x *GetAccountInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountInfoReq.ProtoReflect.Descriptor instead.
func (*GetAccountInfoReq) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_account_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountInfoReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetAccountInfoReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetAccountInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Ttl     int32                   `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Data    *GetAccountInfoRes_Data `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAccountInfoRes) Reset() {
	*x = GetAccountInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountInfoRes) ProtoMessage() {}

func (x *GetAccountInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountInfoRes.ProtoReflect.Descriptor instead.
func (*GetAccountInfoRes) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountInfoRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAccountInfoRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAccountInfoRes) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *GetAccountInfoRes) GetData() *GetAccountInfoRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAccountInfoRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Face   string `protobuf:"bytes,2,opt,name=face,proto3" json:"face,omitempty"`
	Openid string `protobuf:"bytes,3,opt,name=openid,proto3" json:"openid,omitempty"`
}

func (x *GetAccountInfoRes_Data) Reset() {
	*x = GetAccountInfoRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountInfoRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountInfoRes_Data) ProtoMessage() {}

func (x *GetAccountInfoRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountInfoRes_Data.ProtoReflect.Descriptor instead.
func (*GetAccountInfoRes_Data) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_account_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAccountInfoRes_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAccountInfoRes_Data) GetFace() string {
	if x != nil {
		return x.Face
	}
	return ""
}

func (x *GetAccountInfoRes_Data) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

var File_bilibili_webapp_account_proto protoreflect.FileDescriptor

var file_bilibili_webapp_account_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x77, 0x65, 0x62, 0x61, 0x70,
	0x70, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x3f, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x46,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x61, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x32, 0x9a, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x72, 0x63, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x66,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x65, 0x76, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69,
	0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_webapp_account_proto_rawDescOnce sync.Once
	file_bilibili_webapp_account_proto_rawDescData = file_bilibili_webapp_account_proto_rawDesc
)

func file_bilibili_webapp_account_proto_rawDescGZIP() []byte {
	file_bilibili_webapp_account_proto_rawDescOnce.Do(func() {
		file_bilibili_webapp_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_webapp_account_proto_rawDescData)
	})
	return file_bilibili_webapp_account_proto_rawDescData
}

var file_bilibili_webapp_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bilibili_webapp_account_proto_goTypes = []interface{}{
	(*GetAccountInfoReq)(nil),      // 0: member.bilibili.com.GetAccountInfoReq
	(*GetAccountInfoRes)(nil),      // 1: member.bilibili.com.GetAccountInfoRes
	(*GetAccountInfoRes_Data)(nil), // 2: member.bilibili.com.GetAccountInfoRes.Data
}
var file_bilibili_webapp_account_proto_depIdxs = []int32{
	2, // 0: member.bilibili.com.GetAccountInfoRes.data:type_name -> member.bilibili.com.GetAccountInfoRes.Data
	0, // 1: member.bilibili.com.AccountService.GetAccountInfo:input_type -> member.bilibili.com.GetAccountInfoReq
	1, // 2: member.bilibili.com.AccountService.GetAccountInfo:output_type -> member.bilibili.com.GetAccountInfoRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_webapp_account_proto_init() }
func file_bilibili_webapp_account_proto_init() {
	if File_bilibili_webapp_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_webapp_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountInfoReq); i {
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
		file_bilibili_webapp_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountInfoRes); i {
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
		file_bilibili_webapp_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountInfoRes_Data); i {
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
			RawDescriptor: file_bilibili_webapp_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bilibili_webapp_account_proto_goTypes,
		DependencyIndexes: file_bilibili_webapp_account_proto_depIdxs,
		MessageInfos:      file_bilibili_webapp_account_proto_msgTypes,
	}.Build()
	File_bilibili_webapp_account_proto = out.File
	file_bilibili_webapp_account_proto_rawDesc = nil
	file_bilibili_webapp_account_proto_goTypes = nil
	file_bilibili_webapp_account_proto_depIdxs = nil
}
