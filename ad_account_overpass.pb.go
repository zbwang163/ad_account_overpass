// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ad_account_overpass.proto

package ad_account_overpass

import (
	ad_base_overpass "github.com/zbwang163/ad_base_overpass"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameOrEmail string                        `protobuf:"bytes,1,opt,name=name_or_email,json=nameOrEmail,proto3" json:"name_or_email,omitempty"` //用户名或邮箱
	Password    string                        `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`                            //登录密码
	Captcha     string                        `protobuf:"bytes,3,opt,name=captcha,proto3" json:"captcha,omitempty"`                              //验证码
	Base        *ad_base_overpass.BaseRequest `protobuf:"bytes,15,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_account_overpass_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ad_account_overpass_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_ad_account_overpass_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetNameOrEmail() string {
	if x != nil {
		return x.NameOrEmail
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

func (x *LoginRequest) GetBase() *ad_base_overpass.BaseRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *LoginResponse_LoginData       `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Base *ad_base_overpass.BaseResponse `protobuf:"bytes,15,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_account_overpass_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ad_account_overpass_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_ad_account_overpass_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetData() *LoginResponse_LoginData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginResponse) GetBase() *ad_base_overpass.BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

type LoginResponse_LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                            //用户名
	AvatarUrl string `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` //头像url
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"` //session id
}

func (x *LoginResponse_LoginData) Reset() {
	*x = LoginResponse_LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_account_overpass_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse_LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse_LoginData) ProtoMessage() {}

func (x *LoginResponse_LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_ad_account_overpass_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse_LoginData.ProtoReflect.Descriptor instead.
func (*LoginResponse_LoginData) Descriptor() ([]byte, []int) {
	return file_ad_account_overpass_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LoginResponse_LoginData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginResponse_LoginData) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *LoginResponse_LoginData) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

var File_ad_account_overpass_proto protoreflect.FileDescriptor

var file_ad_account_overpass_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x64, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x72, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d,
	0x65, 0x4f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x20,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x22, 0xbf, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x1a, 0x5d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x32, 0x3a, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x62, 0x77,
	0x61, 0x6e, 0x67, 0x31, 0x36, 0x33, 0x2f, 0x61, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x70, 0x61, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ad_account_overpass_proto_rawDescOnce sync.Once
	file_ad_account_overpass_proto_rawDescData = file_ad_account_overpass_proto_rawDesc
)

func file_ad_account_overpass_proto_rawDescGZIP() []byte {
	file_ad_account_overpass_proto_rawDescOnce.Do(func() {
		file_ad_account_overpass_proto_rawDescData = protoimpl.X.CompressGZIP(file_ad_account_overpass_proto_rawDescData)
	})
	return file_ad_account_overpass_proto_rawDescData
}

var file_ad_account_overpass_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ad_account_overpass_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),                  // 0: LoginRequest
	(*LoginResponse)(nil),                 // 1: LoginResponse
	(*LoginResponse_LoginData)(nil),       // 2: LoginResponse.LoginData
	(*ad_base_overpass.BaseRequest)(nil),  // 3: BaseRequest
	(*ad_base_overpass.BaseResponse)(nil), // 4: BaseResponse
}
var file_ad_account_overpass_proto_depIdxs = []int32{
	3, // 0: LoginRequest.base:type_name -> BaseRequest
	2, // 1: LoginResponse.data:type_name -> LoginResponse.LoginData
	4, // 2: LoginResponse.base:type_name -> BaseResponse
	0, // 3: AccountService.Login:input_type -> LoginRequest
	1, // 4: AccountService.Login:output_type -> LoginResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ad_account_overpass_proto_init() }
func file_ad_account_overpass_proto_init() {
	if File_ad_account_overpass_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ad_account_overpass_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_ad_account_overpass_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_ad_account_overpass_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse_LoginData); i {
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
			RawDescriptor: file_ad_account_overpass_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ad_account_overpass_proto_goTypes,
		DependencyIndexes: file_ad_account_overpass_proto_depIdxs,
		MessageInfos:      file_ad_account_overpass_proto_msgTypes,
	}.Build()
	File_ad_account_overpass_proto = out.File
	file_ad_account_overpass_proto_rawDesc = nil
	file_ad_account_overpass_proto_goTypes = nil
	file_ad_account_overpass_proto_depIdxs = nil
}
