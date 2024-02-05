// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: order/v1/order.proto

//

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	page "vine-template-rpc/internal/page"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Order_AddOrder_FullMethodName    = "/api.order.v1.Order/AddOrder"
	Order_UpdateOrder_FullMethodName = "/api.order.v1.Order/UpdateOrder"
	Order_DeleteOrder_FullMethodName = "/api.order.v1.Order/DeleteOrder"
	Order_GetOrder_FullMethodName    = "/api.order.v1.Order/GetOrder"
	Order_ListOrder_FullMethodName   = "/api.order.v1.Order/ListOrder"
	Order_ReviewOrder_FullMethodName = "/api.order.v1.Order/ReviewOrder"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderReply, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderReply, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderReply, error)
	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*page.Page, error)
	// 审核工单
	ReviewOrder(ctx context.Context, in *ReviewOrderRequest, opts ...grpc.CallOption) (*ReviewOrderReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderReply, error) {
	out := new(AddOrderReply)
	err := c.cc.Invoke(ctx, Order_AddOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderReply, error) {
	out := new(UpdateOrderReply)
	err := c.cc.Invoke(ctx, Order_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error) {
	out := new(DeleteOrderReply)
	err := c.cc.Invoke(ctx, Order_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderReply, error) {
	out := new(GetOrderReply)
	err := c.cc.Invoke(ctx, Order_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*page.Page, error) {
	out := new(page.Page)
	err := c.cc.Invoke(ctx, Order_ListOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ReviewOrder(ctx context.Context, in *ReviewOrderRequest, opts ...grpc.CallOption) (*ReviewOrderReply, error) {
	out := new(ReviewOrderReply)
	err := c.cc.Invoke(ctx, Order_ReviewOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	AddOrder(context.Context, *AddOrderRequest) (*AddOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderReply, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderReply, error)
	ListOrder(context.Context, *ListOrderRequest) (*page.Page, error)
	// 审核工单
	ReviewOrder(context.Context, *ReviewOrderRequest) (*ReviewOrderReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) AddOrder(context.Context, *AddOrderRequest) (*AddOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedOrderServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServer) ListOrder(context.Context, *ListOrderRequest) (*page.Page, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServer) ReviewOrder(context.Context, *ReviewOrderRequest) (*ReviewOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_AddOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AddOrder(ctx, req.(*AddOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_ListOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ListOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ReviewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ReviewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_ReviewOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ReviewOrder(ctx, req.(*ReviewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order.v1.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _Order_AddOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Order_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Order_DeleteOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Order_GetOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _Order_ListOrder_Handler,
		},
		{
			MethodName: "ReviewOrder",
			Handler:    _Order_ReviewOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/v1/order.proto",
}
