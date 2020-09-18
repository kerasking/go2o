// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item_service.proto

package proto // import "."

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

// 商品销售类型
type EItemSalesType int32

const (
	EItemSalesType_IT_NORMAL    EItemSalesType = 0
	EItemSalesType_IT_WHOLESALE EItemSalesType = 1
)

var EItemSalesType_name = map[int32]string{
	0: "IT_NORMAL",
	1: "IT_WHOLESALE",
}
var EItemSalesType_value = map[string]int32{
	"IT_NORMAL":    0,
	"IT_WHOLESALE": 1,
}

func (x EItemSalesType) String() string {
	return proto.EnumName(EItemSalesType_name, int32(x))
}
func (EItemSalesType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_item_service_1b4c1e00c9c7fdb9, []int{0}
}

type SkuRequest struct {
	ItemId               int64    `protobuf:"zigzag64,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
	SkuId                int64    `protobuf:"zigzag64,2,opt,name=skuId,proto3" json:"skuId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SkuRequest) Reset()         { *m = SkuRequest{} }
func (m *SkuRequest) String() string { return proto.CompactTextString(m) }
func (*SkuRequest) ProtoMessage()    {}
func (*SkuRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_service_1b4c1e00c9c7fdb9, []int{0}
}
func (m *SkuRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkuRequest.Unmarshal(m, b)
}
func (m *SkuRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkuRequest.Marshal(b, m, deterministic)
}
func (dst *SkuRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkuRequest.Merge(dst, src)
}
func (m *SkuRequest) XXX_Size() int {
	return xxx_messageInfo_SkuRequest.Size(m)
}
func (m *SkuRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SkuRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SkuRequest proto.InternalMessageInfo

func (m *SkuRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *SkuRequest) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

type ItemDetailRequest struct {
	ItemId               int64    `protobuf:"zigzag64,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
	IType                int32    `protobuf:"zigzag32,2,opt,name=iType,proto3" json:"iType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemDetailRequest) Reset()         { *m = ItemDetailRequest{} }
func (m *ItemDetailRequest) String() string { return proto.CompactTextString(m) }
func (*ItemDetailRequest) ProtoMessage()    {}
func (*ItemDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_service_1b4c1e00c9c7fdb9, []int{1}
}
func (m *ItemDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemDetailRequest.Unmarshal(m, b)
}
func (m *ItemDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemDetailRequest.Marshal(b, m, deterministic)
}
func (dst *ItemDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemDetailRequest.Merge(dst, src)
}
func (m *ItemDetailRequest) XXX_Size() int {
	return xxx_messageInfo_ItemDetailRequest.Size(m)
}
func (m *ItemDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemDetailRequest proto.InternalMessageInfo

func (m *ItemDetailRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ItemDetailRequest) GetIType() int32 {
	if m != nil {
		return m.IType
	}
	return 0
}

type PagingGoodsRequest struct {
	ItemType             EItemSalesType `protobuf:"varint,1,opt,name=ItemType,proto3,enum=EItemSalesType" json:"ItemType,omitempty"`
	CategoryId           int64          `protobuf:"varint,2,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Params               *SPagingParams `protobuf:"bytes,3,opt,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PagingGoodsRequest) Reset()         { *m = PagingGoodsRequest{} }
func (m *PagingGoodsRequest) String() string { return proto.CompactTextString(m) }
func (*PagingGoodsRequest) ProtoMessage()    {}
func (*PagingGoodsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_service_1b4c1e00c9c7fdb9, []int{2}
}
func (m *PagingGoodsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingGoodsRequest.Unmarshal(m, b)
}
func (m *PagingGoodsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingGoodsRequest.Marshal(b, m, deterministic)
}
func (dst *PagingGoodsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingGoodsRequest.Merge(dst, src)
}
func (m *PagingGoodsRequest) XXX_Size() int {
	return xxx_messageInfo_PagingGoodsRequest.Size(m)
}
func (m *PagingGoodsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingGoodsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PagingGoodsRequest proto.InternalMessageInfo

func (m *PagingGoodsRequest) GetItemType() EItemSalesType {
	if m != nil {
		return m.ItemType
	}
	return EItemSalesType_IT_NORMAL
}

func (m *PagingGoodsRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *PagingGoodsRequest) GetParams() *SPagingParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type PagingShopGoodsRequest struct {
	ShopId               int64          `protobuf:"varint,1,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	CategoryId           int64          `protobuf:"varint,2,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Params               *SPagingParams `protobuf:"bytes,3,opt,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PagingShopGoodsRequest) Reset()         { *m = PagingShopGoodsRequest{} }
func (m *PagingShopGoodsRequest) String() string { return proto.CompactTextString(m) }
func (*PagingShopGoodsRequest) ProtoMessage()    {}
func (*PagingShopGoodsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_service_1b4c1e00c9c7fdb9, []int{3}
}
func (m *PagingShopGoodsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingShopGoodsRequest.Unmarshal(m, b)
}
func (m *PagingShopGoodsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingShopGoodsRequest.Marshal(b, m, deterministic)
}
func (dst *PagingShopGoodsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingShopGoodsRequest.Merge(dst, src)
}
func (m *PagingShopGoodsRequest) XXX_Size() int {
	return xxx_messageInfo_PagingShopGoodsRequest.Size(m)
}
func (m *PagingShopGoodsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingShopGoodsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PagingShopGoodsRequest proto.InternalMessageInfo

func (m *PagingShopGoodsRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *PagingShopGoodsRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *PagingShopGoodsRequest) GetParams() *SPagingParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*SkuRequest)(nil), "SkuRequest")
	proto.RegisterType((*ItemDetailRequest)(nil), "ItemDetailRequest")
	proto.RegisterType((*PagingGoodsRequest)(nil), "PagingGoodsRequest")
	proto.RegisterType((*PagingShopGoodsRequest)(nil), "PagingShopGoodsRequest")
	proto.RegisterEnum("EItemSalesType", EItemSalesType_name, EItemSalesType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemServiceClient interface {
	// 获取SKU
	GetSku(ctx context.Context, in *SkuRequest, opts ...grpc.CallOption) (*SSku, error)
	// 获取商品的Sku-JSON格式,itemId
	GetItemSkuJson(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*String, error)
	// 获取商品详细数据
	GetItemDetailData(ctx context.Context, in *ItemDetailRequest, opts ...grpc.CallOption) (*String, error)
	// 根据销售标签获取指定数目的商品
	GetValueGoodsBySaleLabel(ctx context.Context, in *GetItemsByLabelRequest, opts ...grpc.CallOption) (*PagingShopGoodsResponse, error)
	// 获取店铺分页上架的商品
	GetShopPagedOnShelvesGoods(ctx context.Context, in *PagingShopGoodsRequest, opts ...grpc.CallOption) (*PagingShopGoodsResponse, error)
	// 获取上架商品数据（分页）
	GetPagedOnShelvesItem(ctx context.Context, in *PagingGoodsRequest, opts ...grpc.CallOption) (*PagingGoodsResponse, error)
}

type itemServiceClient struct {
	cc *grpc.ClientConn
}

func NewItemServiceClient(cc *grpc.ClientConn) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) GetSku(ctx context.Context, in *SkuRequest, opts ...grpc.CallOption) (*SSku, error) {
	out := new(SSku)
	err := c.cc.Invoke(ctx, "/ItemService/GetSku", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItemSkuJson(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/ItemService/GetItemSkuJson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItemDetailData(ctx context.Context, in *ItemDetailRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/ItemService/GetItemDetailData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetValueGoodsBySaleLabel(ctx context.Context, in *GetItemsByLabelRequest, opts ...grpc.CallOption) (*PagingShopGoodsResponse, error) {
	out := new(PagingShopGoodsResponse)
	err := c.cc.Invoke(ctx, "/ItemService/GetValueGoodsBySaleLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetShopPagedOnShelvesGoods(ctx context.Context, in *PagingShopGoodsRequest, opts ...grpc.CallOption) (*PagingShopGoodsResponse, error) {
	out := new(PagingShopGoodsResponse)
	err := c.cc.Invoke(ctx, "/ItemService/GetShopPagedOnShelvesGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetPagedOnShelvesItem(ctx context.Context, in *PagingGoodsRequest, opts ...grpc.CallOption) (*PagingGoodsResponse, error) {
	out := new(PagingGoodsResponse)
	err := c.cc.Invoke(ctx, "/ItemService/GetPagedOnShelvesItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer interface {
	// 获取SKU
	GetSku(context.Context, *SkuRequest) (*SSku, error)
	// 获取商品的Sku-JSON格式,itemId
	GetItemSkuJson(context.Context, *Int64) (*String, error)
	// 获取商品详细数据
	GetItemDetailData(context.Context, *ItemDetailRequest) (*String, error)
	// 根据销售标签获取指定数目的商品
	GetValueGoodsBySaleLabel(context.Context, *GetItemsByLabelRequest) (*PagingShopGoodsResponse, error)
	// 获取店铺分页上架的商品
	GetShopPagedOnShelvesGoods(context.Context, *PagingShopGoodsRequest) (*PagingShopGoodsResponse, error)
	// 获取上架商品数据（分页）
	GetPagedOnShelvesItem(context.Context, *PagingGoodsRequest) (*PagingGoodsResponse, error)
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_GetSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/GetSku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetSku(ctx, req.(*SkuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItemSkuJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItemSkuJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/GetItemSkuJson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItemSkuJson(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItemDetailData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItemDetailData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/GetItemDetailData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItemDetailData(ctx, req.(*ItemDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetValueGoodsBySaleLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsByLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetValueGoodsBySaleLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/GetValueGoodsBySaleLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetValueGoodsBySaleLabel(ctx, req.(*GetItemsByLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetShopPagedOnShelvesGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingShopGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetShopPagedOnShelvesGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/GetShopPagedOnShelvesGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetShopPagedOnShelvesGoods(ctx, req.(*PagingShopGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetPagedOnShelvesItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetPagedOnShelvesItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/GetPagedOnShelvesItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetPagedOnShelvesItem(ctx, req.(*PagingGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSku",
			Handler:    _ItemService_GetSku_Handler,
		},
		{
			MethodName: "GetItemSkuJson",
			Handler:    _ItemService_GetItemSkuJson_Handler,
		},
		{
			MethodName: "GetItemDetailData",
			Handler:    _ItemService_GetItemDetailData_Handler,
		},
		{
			MethodName: "GetValueGoodsBySaleLabel",
			Handler:    _ItemService_GetValueGoodsBySaleLabel_Handler,
		},
		{
			MethodName: "GetShopPagedOnShelvesGoods",
			Handler:    _ItemService_GetShopPagedOnShelvesGoods_Handler,
		},
		{
			MethodName: "GetPagedOnShelvesItem",
			Handler:    _ItemService_GetPagedOnShelvesItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item_service.proto",
}

func init() { proto.RegisterFile("item_service.proto", fileDescriptor_item_service_1b4c1e00c9c7fdb9) }

var fileDescriptor_item_service_1b4c1e00c9c7fdb9 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x8e, 0xd2, 0x4e,
	0x14, 0xc6, 0x61, 0xc9, 0x76, 0xff, 0x7b, 0xd8, 0x3f, 0xc2, 0x71, 0x65, 0x09, 0x17, 0x9b, 0x4d,
	0x4d, 0xcc, 0x46, 0x93, 0x31, 0xa2, 0xf1, 0x42, 0xaf, 0x40, 0x08, 0xd6, 0xa0, 0x60, 0x87, 0x68,
	0xe2, 0x0d, 0x19, 0xe4, 0xa4, 0x34, 0x2d, 0x9d, 0xda, 0x99, 0x6e, 0xe4, 0x11, 0x7c, 0x5b, 0x1f,
	0xc1, 0xcc, 0xb4, 0xe8, 0x12, 0x8c, 0xf1, 0xc2, 0x2b, 0x72, 0x7e, 0xf3, 0x9d, 0x6f, 0xce, 0x99,
	0x8f, 0x02, 0x86, 0x9a, 0x36, 0x0b, 0x45, 0xd9, 0x4d, 0xf8, 0x99, 0x58, 0x9a, 0x49, 0x2d, 0xbb,
	0x67, 0x41, 0x2c, 0x97, 0x22, 0x2e, 0x2b, 0xdc, 0x90, 0x52, 0x22, 0xa0, 0xc7, 0x46, 0x59, 0x30,
	0xf7, 0x05, 0x00, 0x8f, 0x72, 0x9f, 0xbe, 0xe4, 0xa4, 0x34, 0xb6, 0xc1, 0x31, 0x67, 0xde, 0xaa,
	0x53, 0xbd, 0xaa, 0x5e, 0xa3, 0x5f, 0x56, 0x78, 0x0e, 0xc7, 0x2a, 0xca, 0xbd, 0x55, 0xe7, 0xc8,
	0xe2, 0xa2, 0x70, 0xfb, 0xd0, 0xf2, 0x34, 0x6d, 0x86, 0xa4, 0x45, 0x18, 0xff, 0x85, 0x45, 0x38,
	0xdf, 0xa6, 0x64, 0x2d, 0x5a, 0x7e, 0x51, 0xb8, 0xdf, 0xaa, 0x80, 0x33, 0x11, 0x84, 0x49, 0x30,
	0x96, 0x72, 0xa5, 0x76, 0x26, 0x8f, 0xe0, 0x3f, 0xe3, 0x6c, 0xf5, 0xc6, 0xa6, 0xd1, 0xbb, 0xc3,
	0x46, 0x86, 0x70, 0x11, 0x93, 0x32, 0xd8, 0xff, 0x29, 0xc0, 0x4b, 0x80, 0x57, 0x42, 0x53, 0x20,
	0xb3, 0x6d, 0x39, 0x61, 0xcd, 0xbf, 0x45, 0xf0, 0x01, 0x38, 0x33, 0x91, 0x89, 0x8d, 0xea, 0xd4,
	0xae, 0xaa, 0xd7, 0xf5, 0x5e, 0x83, 0xf1, 0xe2, 0xca, 0x82, 0xfa, 0xe5, 0xa9, 0xfb, 0x15, 0xda,
	0x05, 0xe7, 0x6b, 0x99, 0xee, 0x8d, 0xd3, 0x06, 0xc7, 0xb0, 0x72, 0xa7, 0x9a, 0x5f, 0x56, 0xff,
	0xea, 0xe6, 0x87, 0x4f, 0xa0, 0xb1, 0xbf, 0x1d, 0xfe, 0x0f, 0xa7, 0xde, 0x7c, 0xf1, 0x6e, 0xea,
	0xbf, 0xed, 0x4f, 0x9a, 0x15, 0x6c, 0xc2, 0x99, 0x37, 0x5f, 0x7c, 0x7c, 0x3d, 0x9d, 0x8c, 0x78,
	0x7f, 0x32, 0x6a, 0x56, 0x7b, 0xdf, 0x8f, 0xa0, 0x6e, 0x5b, 0x8a, 0xbc, 0xf1, 0x12, 0x9c, 0x31,
	0x69, 0x1e, 0xe5, 0x58, 0x67, 0xbf, 0x02, 0xed, 0x1e, 0x33, 0xce, 0xa3, 0xdc, 0xad, 0xe0, 0x7d,
	0x68, 0x8c, 0x49, 0xdb, 0x8e, 0x28, 0x7f, 0xa3, 0x64, 0x82, 0x0e, 0xf3, 0x12, 0xfd, 0xfc, 0x59,
	0xf7, 0x84, 0x71, 0x9d, 0x85, 0x49, 0xe0, 0x56, 0xb0, 0x07, 0xad, 0x52, 0x54, 0x64, 0x3a, 0x14,
	0x5a, 0x20, 0xb2, 0x83, 0x90, 0x6f, 0xf7, 0x4c, 0xa1, 0x33, 0x26, 0xfd, 0x41, 0xc4, 0x39, 0xd9,
	0x37, 0x1b, 0x6c, 0xcd, 0x16, 0x13, 0xb1, 0xa4, 0x18, 0x2f, 0x58, 0x69, 0xa7, 0x06, 0x5b, 0x4b,
	0x76, 0xfd, 0x1d, 0x76, 0xf0, 0xd2, 0x2a, 0x95, 0x89, 0x22, 0xb7, 0x82, 0xef, 0xa1, 0x6b, 0x36,
	0x59, 0xcb, 0x74, 0x26, 0x02, 0x5a, 0x4d, 0x13, 0xbe, 0xa6, 0xf8, 0x86, 0x94, 0xd5, 0xe1, 0x05,
	0xfb, 0x7d, 0x46, 0x7f, 0xb4, 0x1c, 0xc2, 0xbd, 0x31, 0xe9, 0x7d, 0x3b, 0x33, 0x16, 0xde, 0x65,
	0x87, 0x7f, 0xbe, 0xee, 0xf9, 0x3e, 0xdc, 0xb9, 0x0c, 0x4e, 0x3f, 0x9d, 0xb0, 0x97, 0xf6, 0xab,
	0x59, 0x3a, 0xf6, 0xe7, 0xe9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xb3, 0xdd, 0xbf, 0x74,
	0x03, 0x00, 0x00,
}
