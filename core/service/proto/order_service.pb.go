// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order_service.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	// 提交订单
	SubmitOrderV1(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*StringMap, error)
	// 预生成订单
	PrepareOrder(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*PrepareOrderResponse, error)
	// 提交普通订单
	SubmitNormalOrder_(ctx context.Context, in *SubmitNormalOrderV2Request, opts ...grpc.CallOption) (*NormalOrderSubmitResponse, error)
	//
	// 获取订单信息
	GetParentOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SParentOrder, error)
	// 获取子订单,orderId
	GetOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SSingleOrder, error)
	// 提交交易订单
	SubmitTradeOrder(ctx context.Context, in *TradeOrderSubmitRequest, opts ...grpc.CallOption) (*Result, error)
	// 交易单现金支付,orderId
	TradeOrderCashPay(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 上传交易单发票
	TradeOrderUpdateTicket(ctx context.Context, in *TradeOrderTicketRequest, opts ...grpc.CallOption) (*Result, error)
	// 预生成订单，使用优惠券
	PrepareOrderWithCoupon_(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*StringMap, error)
	// 取消订单
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Result, error)
	// 确定订单
	ConfirmOrder(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 备货完成
	PickUp(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 订单发货,并记录配送服务商编号及单号
	Ship(ctx context.Context, in *OrderShipmentRequest, opts ...grpc.CallOption) (*Result, error)
	// 买家收货
	BuyerReceived(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 获取订单日志
	LogBytes(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*String, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) SubmitOrderV1(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitOrderV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PrepareOrder(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*PrepareOrderResponse, error) {
	out := new(PrepareOrderResponse)
	err := c.cc.Invoke(ctx, "/OrderService/PrepareOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SubmitNormalOrder_(ctx context.Context, in *SubmitNormalOrderV2Request, opts ...grpc.CallOption) (*NormalOrderSubmitResponse, error) {
	out := new(NormalOrderSubmitResponse)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitNormalOrder_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetParentOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SParentOrder, error) {
	out := new(SParentOrder)
	err := c.cc.Invoke(ctx, "/OrderService/GetParentOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SSingleOrder, error) {
	out := new(SSingleOrder)
	err := c.cc.Invoke(ctx, "/OrderService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SubmitTradeOrder(ctx context.Context, in *TradeOrderSubmitRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitTradeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TradeOrderCashPay(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/TradeOrderCashPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TradeOrderUpdateTicket(ctx context.Context, in *TradeOrderTicketRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/TradeOrderUpdateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PrepareOrderWithCoupon_(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/OrderService/PrepareOrderWithCoupon_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ConfirmOrder(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/ConfirmOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PickUp(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/PickUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Ship(ctx context.Context, in *OrderShipmentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/Ship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) BuyerReceived(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/BuyerReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) LogBytes(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/OrderService/LogBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	// 提交订单
	SubmitOrderV1(context.Context, *SubmitOrderRequest) (*StringMap, error)
	// 预生成订单
	PrepareOrder(context.Context, *PrepareOrderRequest) (*PrepareOrderResponse, error)
	// 提交普通订单
	SubmitNormalOrder_(context.Context, *SubmitNormalOrderV2Request) (*NormalOrderSubmitResponse, error)
	//
	// 获取订单信息
	GetParentOrder(context.Context, *OrderNoV2) (*SParentOrder, error)
	// 获取子订单,orderId
	GetOrder(context.Context, *OrderNoV2) (*SSingleOrder, error)
	// 提交交易订单
	SubmitTradeOrder(context.Context, *TradeOrderSubmitRequest) (*Result, error)
	// 交易单现金支付,orderId
	TradeOrderCashPay(context.Context, *Int64) (*Result, error)
	// 上传交易单发票
	TradeOrderUpdateTicket(context.Context, *TradeOrderTicketRequest) (*Result, error)
	// 预生成订单，使用优惠券
	PrepareOrderWithCoupon_(context.Context, *PrepareOrderRequest) (*StringMap, error)
	// 取消订单
	CancelOrder(context.Context, *CancelOrderRequest) (*Result, error)
	// 确定订单
	ConfirmOrder(context.Context, *OrderNo) (*Result, error)
	// 备货完成
	PickUp(context.Context, *OrderNo) (*Result, error)
	// 订单发货,并记录配送服务商编号及单号
	Ship(context.Context, *OrderShipmentRequest) (*Result, error)
	// 买家收货
	BuyerReceived(context.Context, *OrderNo) (*Result, error)
	// 获取订单日志
	LogBytes(context.Context, *OrderNo) (*String, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_SubmitOrderV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitOrderV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitOrderV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitOrderV1(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PrepareOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PrepareOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PrepareOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PrepareOrder(ctx, req.(*PrepareOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SubmitNormalOrder__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitNormalOrderV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitNormalOrder_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitNormalOrder_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitNormalOrder_(ctx, req.(*SubmitNormalOrderV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetParentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNoV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetParentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetParentOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetParentOrder(ctx, req.(*OrderNoV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNoV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderNoV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SubmitTradeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeOrderSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitTradeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitTradeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitTradeOrder(ctx, req.(*TradeOrderSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TradeOrderCashPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).TradeOrderCashPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/TradeOrderCashPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).TradeOrderCashPay(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TradeOrderUpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeOrderTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).TradeOrderUpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/TradeOrderUpdateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).TradeOrderUpdateTicket(ctx, req.(*TradeOrderTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PrepareOrderWithCoupon__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PrepareOrderWithCoupon_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PrepareOrderWithCoupon_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PrepareOrderWithCoupon_(ctx, req.(*PrepareOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ConfirmOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/ConfirmOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PickUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PickUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PickUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PickUp(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Ship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Ship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Ship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Ship(ctx, req.(*OrderShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_BuyerReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).BuyerReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/BuyerReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).BuyerReceived(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_LogBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).LogBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/LogBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).LogBytes(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrderV1",
			Handler:    _OrderService_SubmitOrderV1_Handler,
		},
		{
			MethodName: "PrepareOrder",
			Handler:    _OrderService_PrepareOrder_Handler,
		},
		{
			MethodName: "SubmitNormalOrder_",
			Handler:    _OrderService_SubmitNormalOrder__Handler,
		},
		{
			MethodName: "GetParentOrder",
			Handler:    _OrderService_GetParentOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "SubmitTradeOrder",
			Handler:    _OrderService_SubmitTradeOrder_Handler,
		},
		{
			MethodName: "TradeOrderCashPay",
			Handler:    _OrderService_TradeOrderCashPay_Handler,
		},
		{
			MethodName: "TradeOrderUpdateTicket",
			Handler:    _OrderService_TradeOrderUpdateTicket_Handler,
		},
		{
			MethodName: "PrepareOrderWithCoupon_",
			Handler:    _OrderService_PrepareOrderWithCoupon__Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "ConfirmOrder",
			Handler:    _OrderService_ConfirmOrder_Handler,
		},
		{
			MethodName: "PickUp",
			Handler:    _OrderService_PickUp_Handler,
		},
		{
			MethodName: "Ship",
			Handler:    _OrderService_Ship_Handler,
		},
		{
			MethodName: "BuyerReceived",
			Handler:    _OrderService_BuyerReceived_Handler,
		},
		{
			MethodName: "LogBytes",
			Handler:    _OrderService_LogBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_service.proto",
}

func init() { proto.RegisterFile("order_service.proto", fileDescriptor_order_service_d0fb5b614428c77e) }

var fileDescriptor_order_service_d0fb5b614428c77e = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xf3, 0x00, 0xa5, 0x3a, 0x5a, 0x04, 0x2e, 0x63, 0xa8, 0x08, 0x4d, 0x0a, 0x08, 0x78,
	0xd9, 0x4d, 0x04, 0x04, 0x0f, 0xa8, 0x2f, 0xed, 0xc3, 0x84, 0x04, 0x5b, 0xb5, 0x6c, 0x45, 0xe2,
	0xa5, 0x72, 0x93, 0x23, 0xb5, 0x96, 0xd8, 0xc6, 0x76, 0x26, 0xf5, 0x43, 0xf3, 0x1d, 0x50, 0xe3,
	0x45, 0x75, 0x21, 0x7b, 0x72, 0xee, 0x7f, 0x3f, 0xdf, 0xff, 0xee, 0xac, 0xc0, 0x48, 0x99, 0x9c,
	0xcc, 0xd2, 0x92, 0xb9, 0x11, 0x19, 0xa1, 0x36, 0xca, 0xa9, 0xf1, 0xa0, 0x28, 0xd5, 0x8a, 0x97,
	0xb7, 0xd1, 0x61, 0x45, 0xd6, 0xf2, 0x82, 0x4e, 0x3c, 0x9a, 0x3b, 0xe5, 0x13, 0xc9, 0x9f, 0xfb,
	0x30, 0x38, 0xdf, 0x6a, 0xa9, 0xbf, 0xcd, 0x12, 0x18, 0xa6, 0xf5, 0xaa, 0x12, 0xae, 0x51, 0x17,
	0xef, 0xd9, 0x08, 0x83, 0xf8, 0x82, 0x7e, 0xd7, 0x64, 0xdd, 0x18, 0x30, 0x75, 0x46, 0xc8, 0xe2,
	0x3b, 0xd7, 0x71, 0xc4, 0x26, 0x30, 0x98, 0x1b, 0xd2, 0xdc, 0x50, 0x03, 0xb1, 0xa7, 0x18, 0x86,
	0xed, 0x9d, 0x83, 0x7f, 0x54, 0xab, 0x95, 0xb4, 0x14, 0x47, 0xec, 0x1c, 0x98, 0xb7, 0x38, 0x53,
	0xa6, 0xe2, 0x65, 0x93, 0x5e, 0xb2, 0x17, 0xf8, 0x9f, 0xb8, 0x48, 0xda, 0x5a, 0x63, 0x0c, 0x64,
	0xcf, 0x05, 0x05, 0x8f, 0xe1, 0xd1, 0x29, 0xb9, 0x39, 0x37, 0x24, 0x7d, 0xdb, 0x0c, 0xb0, 0x39,
	0xcf, 0xd4, 0x22, 0x19, 0x0f, 0x31, 0x0d, 0x52, 0x71, 0xc4, 0xde, 0x42, 0xff, 0x94, 0xba, 0xc1,
	0x54, 0xc8, 0xa2, 0xa4, 0x16, 0xfc, 0x0c, 0x8f, 0xbd, 0xd7, 0xa5, 0xe1, 0xf9, 0xed, 0xac, 0xcf,
	0x71, 0x17, 0xb4, 0x8d, 0xf8, 0x1e, 0x1f, 0xe0, 0x05, 0xd9, 0xba, 0x74, 0x71, 0xc4, 0xde, 0xc0,
	0x93, 0x1d, 0x35, 0xe3, 0x76, 0x3d, 0xe7, 0x1b, 0xd6, 0xc3, 0xaf, 0xd2, 0x7d, 0xfa, 0x18, 0x72,
	0x13, 0x78, 0xb6, 0xe3, 0xae, 0x74, 0xce, 0x1d, 0x5d, 0x8a, 0xec, 0x9a, 0xdc, 0x9e, 0x8d, 0x97,
	0x3a, 0x6c, 0x26, 0x70, 0x18, 0xae, 0xf8, 0x87, 0x70, 0xeb, 0x99, 0xaa, 0xb5, 0x92, 0xcb, 0x3b,
	0x9e, 0x64, 0xff, 0x19, 0x8f, 0xe1, 0xe1, 0x8c, 0xcb, 0x8c, 0xfc, 0x56, 0xd9, 0x08, 0x83, 0xa8,
	0xc3, 0xed, 0x15, 0x0c, 0x66, 0x4a, 0xfe, 0x12, 0xa6, 0xf2, 0x7c, 0xbf, 0x5d, 0x5d, 0x08, 0xbd,
	0x84, 0xde, 0x5c, 0x64, 0xd7, 0x57, 0xba, 0x3b, 0xfd, 0x0e, 0xee, 0xa5, 0x6b, 0xa1, 0xd9, 0x81,
	0x4f, 0x6e, 0xbf, 0x2b, 0x92, 0x5d, 0xb3, 0xbd, 0x86, 0xe1, 0xb4, 0xde, 0x6c, 0x1b, 0xc9, 0x48,
	0xdc, 0x50, 0xde, 0x5d, 0xef, 0x08, 0xfa, 0xdf, 0x54, 0x31, 0xdd, 0x38, 0xb2, 0x7b, 0x80, 0x1f,
	0x33, 0x8e, 0xa6, 0x47, 0x30, 0xca, 0x54, 0x85, 0x85, 0x70, 0xeb, 0x7a, 0x85, 0x85, 0x4a, 0x14,
	0x1a, 0x9d, 0xfd, 0xec, 0xe3, 0xc9, 0x97, 0xe6, 0x87, 0x58, 0xf5, 0x9a, 0xe3, 0xc3, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x67, 0x3d, 0x79, 0x8e, 0x55, 0x03, 0x00, 0x00,
}
