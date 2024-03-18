// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: goods.proto

package goods

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Goods_GetGood_FullMethodName       = "/goods.Goods/GetGood"
	Goods_GetGoodByName_FullMethodName = "/goods.Goods/GetGoodByName"
	Goods_GetGoods_FullMethodName      = "/goods.Goods/GetGoods"
	Goods_CreateGood_FullMethodName    = "/goods.Goods/CreateGood"
	Goods_DeleteGood_FullMethodName    = "/goods.Goods/DeleteGood"
	Goods_UpdateGood_FullMethodName    = "/goods.Goods/UpdateGood"
	Goods_UpdateStock_FullMethodName   = "/goods.Goods/UpdateStock"
)

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
	GetGoodByName(ctx context.Context, in *GetGoodsByNameRequest, opts ...grpc.CallOption) (*GetGoodsByNameResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error)
	DeleteGood(ctx context.Context, in *DeleteGoodRequest, opts ...grpc.CallOption) (*DeleteGoodResponse, error)
	UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodResponse, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := c.cc.Invoke(ctx, Goods_GetGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodByName(ctx context.Context, in *GetGoodsByNameRequest, opts ...grpc.CallOption) (*GetGoodsByNameResponse, error) {
	out := new(GetGoodsByNameResponse)
	err := c.cc.Invoke(ctx, Goods_GetGoodByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, Goods_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error) {
	out := new(CreateGoodResponse)
	err := c.cc.Invoke(ctx, Goods_CreateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteGood(ctx context.Context, in *DeleteGoodRequest, opts ...grpc.CallOption) (*DeleteGoodResponse, error) {
	out := new(DeleteGoodResponse)
	err := c.cc.Invoke(ctx, Goods_DeleteGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodResponse, error) {
	out := new(UpdateGoodResponse)
	err := c.cc.Invoke(ctx, Goods_UpdateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error) {
	out := new(UpdateStockResponse)
	err := c.cc.Invoke(ctx, Goods_UpdateStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
	GetGoodByName(context.Context, *GetGoodsByNameRequest) (*GetGoodsByNameResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error)
	DeleteGood(context.Context, *DeleteGoodRequest) (*DeleteGoodResponse, error)
	UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodResponse, error)
	UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGood not implemented")
}
func (UnimplementedGoodsServer) GetGoodByName(context.Context, *GetGoodsByNameRequest) (*GetGoodsByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodByName not implemented")
}
func (UnimplementedGoodsServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServer) CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGood not implemented")
}
func (UnimplementedGoodsServer) DeleteGood(context.Context, *DeleteGoodRequest) (*DeleteGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGood not implemented")
}
func (UnimplementedGoodsServer) UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGood not implemented")
}
func (UnimplementedGoodsServer) UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetGoodByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodByName(ctx, req.(*GetGoodsByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_CreateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGood(ctx, req.(*CreateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_DeleteGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteGood(ctx, req.(*DeleteGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_UpdateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateGood(ctx, req.(*UpdateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGood",
			Handler:    _Goods_GetGood_Handler,
		},
		{
			MethodName: "GetGoodByName",
			Handler:    _Goods_GetGoodByName_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _Goods_GetGoods_Handler,
		},
		{
			MethodName: "CreateGood",
			Handler:    _Goods_CreateGood_Handler,
		},
		{
			MethodName: "DeleteGood",
			Handler:    _Goods_DeleteGood_Handler,
		},
		{
			MethodName: "UpdateGood",
			Handler:    _Goods_UpdateGood_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _Goods_UpdateStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}