// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code_service.proto

package proto // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SaveQrTemplateRequest struct {
	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// * 模板标题
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title"`
	// * 背景图片
	BgImage string `protobuf:"bytes,3,opt,name=BgImage,proto3" json:"BgImage"`
	// * 垂直偏离量
	OffsetX int32 `protobuf:"varint,4,opt,name=OffsetX,proto3" json:"OffsetX"`
	// * 垂直偏移量
	OffsetY int32 `protobuf:"varint,5,opt,name=OffsetY,proto3" json:"OffsetY"`
	// * 二维码模板文本
	Comment string `protobuf:"bytes,6,opt,name=Comment,proto3" json:"Comment"`
	// * 回调地址
	CallbackUrl string `protobuf:"bytes,7,opt,name=CallbackUrl,proto3" json:"CallbackUrl"`
	// * 是否启用
	Enabled              int32    `protobuf:"varint,8,opt,name=Enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveQrTemplateRequest) Reset()         { *m = SaveQrTemplateRequest{} }
func (m *SaveQrTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*SaveQrTemplateRequest) ProtoMessage()    {}
func (*SaveQrTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_4adaab270932ac42, []int{0}
}
func (m *SaveQrTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveQrTemplateRequest.Unmarshal(m, b)
}
func (m *SaveQrTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveQrTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *SaveQrTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveQrTemplateRequest.Merge(dst, src)
}
func (m *SaveQrTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_SaveQrTemplateRequest.Size(m)
}
func (m *SaveQrTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveQrTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveQrTemplateRequest proto.InternalMessageInfo

func (m *SaveQrTemplateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaveQrTemplateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetBgImage() string {
	if m != nil {
		return m.BgImage
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetOffsetX() int32 {
	if m != nil {
		return m.OffsetX
	}
	return 0
}

func (m *SaveQrTemplateRequest) GetOffsetY() int32 {
	if m != nil {
		return m.OffsetY
	}
	return 0
}

func (m *SaveQrTemplateRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetCallbackUrl() string {
	if m != nil {
		return m.CallbackUrl
	}
	return ""
}

func (m *SaveQrTemplateRequest) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

type SaveQrTemplateResponse struct {
	ErrCode              int64    `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg"`
	Id                   int64    `protobuf:"varint,3,opt,name=Id,proto3" json:"Id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveQrTemplateResponse) Reset()         { *m = SaveQrTemplateResponse{} }
func (m *SaveQrTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*SaveQrTemplateResponse) ProtoMessage()    {}
func (*SaveQrTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_4adaab270932ac42, []int{1}
}
func (m *SaveQrTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveQrTemplateResponse.Unmarshal(m, b)
}
func (m *SaveQrTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveQrTemplateResponse.Marshal(b, m, deterministic)
}
func (dst *SaveQrTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveQrTemplateResponse.Merge(dst, src)
}
func (m *SaveQrTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_SaveQrTemplateResponse.Size(m)
}
func (m *SaveQrTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveQrTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveQrTemplateResponse proto.InternalMessageInfo

func (m *SaveQrTemplateResponse) GetErrCode() int64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SaveQrTemplateResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SaveQrTemplateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CommQrTemplateId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommQrTemplateId) Reset()         { *m = CommQrTemplateId{} }
func (m *CommQrTemplateId) String() string { return proto.CompactTextString(m) }
func (*CommQrTemplateId) ProtoMessage()    {}
func (*CommQrTemplateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_4adaab270932ac42, []int{2}
}
func (m *CommQrTemplateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommQrTemplateId.Unmarshal(m, b)
}
func (m *CommQrTemplateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommQrTemplateId.Marshal(b, m, deterministic)
}
func (dst *CommQrTemplateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommQrTemplateId.Merge(dst, src)
}
func (m *CommQrTemplateId) XXX_Size() int {
	return xxx_messageInfo_CommQrTemplateId.Size(m)
}
func (m *CommQrTemplateId) XXX_DiscardUnknown() {
	xxx_messageInfo_CommQrTemplateId.DiscardUnknown(m)
}

var xxx_messageInfo_CommQrTemplateId proto.InternalMessageInfo

func (m *CommQrTemplateId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SQrTemplate struct {
	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// * 模板标题
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title"`
	// * 背景图片
	BgImage string `protobuf:"bytes,3,opt,name=BgImage,proto3" json:"BgImage"`
	// * 垂直偏离量
	OffsetX int32 `protobuf:"varint,4,opt,name=OffsetX,proto3" json:"OffsetX"`
	// * 垂直偏移量
	OffsetY int32 `protobuf:"varint,5,opt,name=OffsetY,proto3" json:"OffsetY"`
	// * 二维码模板文本
	Comment string `protobuf:"bytes,6,opt,name=Comment,proto3" json:"Comment"`
	// * 回调地址
	CallbackUrl string `protobuf:"bytes,7,opt,name=CallbackUrl,proto3" json:"CallbackUrl"`
	// * 是否启用
	Enabled              int32    `protobuf:"varint,8,opt,name=Enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SQrTemplate) Reset()         { *m = SQrTemplate{} }
func (m *SQrTemplate) String() string { return proto.CompactTextString(m) }
func (*SQrTemplate) ProtoMessage()    {}
func (*SQrTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_service_4adaab270932ac42, []int{3}
}
func (m *SQrTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SQrTemplate.Unmarshal(m, b)
}
func (m *SQrTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SQrTemplate.Marshal(b, m, deterministic)
}
func (dst *SQrTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SQrTemplate.Merge(dst, src)
}
func (m *SQrTemplate) XXX_Size() int {
	return xxx_messageInfo_SQrTemplate.Size(m)
}
func (m *SQrTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SQrTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SQrTemplate proto.InternalMessageInfo

func (m *SQrTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SQrTemplate) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SQrTemplate) GetBgImage() string {
	if m != nil {
		return m.BgImage
	}
	return ""
}

func (m *SQrTemplate) GetOffsetX() int32 {
	if m != nil {
		return m.OffsetX
	}
	return 0
}

func (m *SQrTemplate) GetOffsetY() int32 {
	if m != nil {
		return m.OffsetY
	}
	return 0
}

func (m *SQrTemplate) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *SQrTemplate) GetCallbackUrl() string {
	if m != nil {
		return m.CallbackUrl
	}
	return ""
}

func (m *SQrTemplate) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func init() {
	proto.RegisterType((*SaveQrTemplateRequest)(nil), "SaveQrTemplateRequest")
	proto.RegisterType((*SaveQrTemplateResponse)(nil), "SaveQrTemplateResponse")
	proto.RegisterType((*CommQrTemplateId)(nil), "CommQrTemplateId")
	proto.RegisterType((*SQrTemplate)(nil), "SQrTemplate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CodeServiceClient is the client API for CodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodeServiceClient interface {
	// 保存CommQrTemplate
	SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error)
}

type codeServiceClient struct {
	cc *grpc.ClientConn
}

func NewCodeServiceClient(cc *grpc.ClientConn) CodeServiceClient {
	return &codeServiceClient{cc}
}

func (c *codeServiceClient) SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error) {
	out := new(SaveQrTemplateResponse)
	err := c.cc.Invoke(ctx, "/CodeService/SaveQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error) {
	out := new(SQrTemplate)
	err := c.cc.Invoke(ctx, "/CodeService/GetQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CodeService/DeleteQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServiceServer is the server API for CodeService service.
type CodeServiceServer interface {
	// 保存CommQrTemplate
	SaveQrTemplate(context.Context, *SaveQrTemplateRequest) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(context.Context, *CommQrTemplateId) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(context.Context, *CommQrTemplateId) (*Result, error)
}

func RegisterCodeServiceServer(s *grpc.Server, srv CodeServiceServer) {
	s.RegisterService(&_CodeService_serviceDesc, srv)
}

func _CodeService_SaveQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQrTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/SaveQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, req.(*SaveQrTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_GetQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/GetQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_DeleteQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/DeleteQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CodeService",
	HandlerType: (*CodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveQrTemplate",
			Handler:    _CodeService_SaveQrTemplate_Handler,
		},
		{
			MethodName: "GetQrTemplate",
			Handler:    _CodeService_GetQrTemplate_Handler,
		},
		{
			MethodName: "DeleteQrTemplate",
			Handler:    _CodeService_DeleteQrTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_service.proto",
}

func init() { proto.RegisterFile("code_service.proto", fileDescriptor_code_service_4adaab270932ac42) }

var fileDescriptor_code_service_4adaab270932ac42 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xcf, 0x4e, 0xc2, 0x30,
	0x1c, 0xc7, 0x19, 0xc8, 0x1f, 0x0b, 0x12, 0xac, 0x8a, 0x0d, 0x17, 0x97, 0x9d, 0x76, 0xaa, 0x06,
	0x8f, 0xde, 0x98, 0xc4, 0x70, 0x30, 0xc6, 0x81, 0x46, 0xb8, 0x98, 0x6e, 0xfb, 0x31, 0x89, 0x1d,
	0x9d, 0x5d, 0xc7, 0xb3, 0xf9, 0x26, 0x9e, 0x7d, 0x13, 0xb3, 0x7f, 0x61, 0x12, 0xe2, 0x03, 0x78,
	0x5a, 0x3e, 0xfd, 0xae, 0xdf, 0xf4, 0xf7, 0x59, 0x87, 0xb0, 0x2b, 0x3c, 0x78, 0x8d, 0x40, 0x6e,
	0x56, 0x2e, 0xd0, 0x50, 0x0a, 0x25, 0x06, 0x1d, 0x9f, 0x0b, 0x87, 0xf1, 0x8c, 0x8c, 0x6f, 0x0d,
	0x9d, 0x4d, 0xd9, 0x06, 0x1e, 0xe5, 0x0c, 0x82, 0x90, 0x33, 0x05, 0x36, 0x7c, 0xc4, 0x10, 0x29,
	0xdc, 0x45, 0xd5, 0x89, 0x47, 0x34, 0x5d, 0x33, 0x6b, 0x76, 0x75, 0xe2, 0xe1, 0x53, 0x54, 0x9f,
	0xad, 0x14, 0x07, 0x52, 0xd5, 0x35, 0xf3, 0xd0, 0xce, 0x00, 0x13, 0xd4, 0x1c, 0xf9, 0x93, 0x80,
	0xf9, 0x40, 0x6a, 0xe9, 0x7a, 0x81, 0x49, 0xf2, 0xb0, 0x5c, 0x46, 0xa0, 0x5e, 0xc8, 0x81, 0xae,
	0x99, 0x75, 0xbb, 0xc0, 0x6d, 0x32, 0x27, 0xf5, 0x72, 0x32, 0x4f, 0x12, 0x4b, 0x04, 0x01, 0xac,
	0x15, 0x69, 0x64, 0x6d, 0x39, 0x62, 0x1d, 0xb5, 0x2d, 0xc6, 0xb9, 0xc3, 0xdc, 0xf7, 0x27, 0xc9,
	0x49, 0x33, 0x4d, 0xcb, 0x4b, 0xc9, 0xde, 0xf1, 0x9a, 0x39, 0x1c, 0x3c, 0xd2, 0xca, 0x5a, 0x73,
	0x34, 0x16, 0xa8, 0xbf, 0x3b, 0x62, 0x14, 0x8a, 0x75, 0x94, 0x9e, 0x71, 0x2c, 0xa5, 0x25, 0x3c,
	0xc8, 0x07, 0x2d, 0x10, 0xf7, 0x51, 0x63, 0x2c, 0xe5, 0x7d, 0xe4, 0xe7, 0xe3, 0xe6, 0x94, 0x5b,
	0xa9, 0x15, 0x56, 0x0c, 0x13, 0xf5, 0x92, 0x23, 0x6e, 0xbb, 0x33, 0x53, 0xcf, 0x8c, 0xc7, 0x45,
	0x67, 0x06, 0xc6, 0x97, 0x86, 0xda, 0xd3, 0xed, 0x7b, 0xff, 0xc7, 0xef, 0xf0, 0x53, 0x43, 0xed,
	0x44, 0xda, 0x34, 0xbb, 0x67, 0xd8, 0x42, 0xdd, 0xdf, 0xbe, 0x71, 0x9f, 0xee, 0xbd, 0x63, 0x83,
	0x73, 0xba, 0xff, 0xc3, 0x18, 0x15, 0x3c, 0x44, 0x47, 0x77, 0xa0, 0x4a, 0x1d, 0xc7, 0x74, 0x57,
	0xf4, 0xa0, 0x43, 0x4b, 0x42, 0x8d, 0x0a, 0xbe, 0x42, 0xbd, 0x5b, 0xe0, 0xa0, 0xe0, 0xef, 0x6d,
	0x4d, 0x6a, 0x43, 0x14, 0x73, 0x65, 0x54, 0x46, 0x17, 0xe8, 0xc4, 0x15, 0x01, 0xf5, 0x57, 0xea,
	0x2d, 0x76, 0xa8, 0x2f, 0x86, 0x82, 0xca, 0xd0, 0x5d, 0xb4, 0xe8, 0xe5, 0x4d, 0xfa, 0x7f, 0x38,
	0x8d, 0xf4, 0x71, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x12, 0xc5, 0x17, 0x4a, 0x03, 0x00,
	0x00,
}