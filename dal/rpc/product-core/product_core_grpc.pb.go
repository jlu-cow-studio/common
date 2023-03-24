// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: product_core.proto

package product_core

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

// ProductCoreServiceClient is the client API for ProductCoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCoreServiceClient interface {
	// 添加商品方法
	AddItem(ctx context.Context, in *AddItemReq, opts ...grpc.CallOption) (*AddItemRes, error)
	// 删除商品方法
	DeleteItem(ctx context.Context, in *DeleteItemReq, opts ...grpc.CallOption) (*DeleteItemRes, error)
}

type productCoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCoreServiceClient(cc grpc.ClientConnInterface) ProductCoreServiceClient {
	return &productCoreServiceClient{cc}
}

func (c *productCoreServiceClient) AddItem(ctx context.Context, in *AddItemReq, opts ...grpc.CallOption) (*AddItemRes, error) {
	out := new(AddItemRes)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.product_core.ProductCoreService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCoreServiceClient) DeleteItem(ctx context.Context, in *DeleteItemReq, opts ...grpc.CallOption) (*DeleteItemRes, error) {
	out := new(DeleteItemRes)
	err := c.cc.Invoke(ctx, "/jlu_cow_studio.product_core.ProductCoreService/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCoreServiceServer is the server API for ProductCoreService service.
// All implementations must embed UnimplementedProductCoreServiceServer
// for forward compatibility
type ProductCoreServiceServer interface {
	// 添加商品方法
	AddItem(context.Context, *AddItemReq) (*AddItemRes, error)
	// 删除商品方法
	DeleteItem(context.Context, *DeleteItemReq) (*DeleteItemRes, error)
	mustEmbedUnimplementedProductCoreServiceServer()
}

// UnimplementedProductCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductCoreServiceServer struct {
}

func (UnimplementedProductCoreServiceServer) AddItem(context.Context, *AddItemReq) (*AddItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedProductCoreServiceServer) DeleteItem(context.Context, *DeleteItemReq) (*DeleteItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedProductCoreServiceServer) mustEmbedUnimplementedProductCoreServiceServer() {}

// UnsafeProductCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCoreServiceServer will
// result in compilation errors.
type UnsafeProductCoreServiceServer interface {
	mustEmbedUnimplementedProductCoreServiceServer()
}

func RegisterProductCoreServiceServer(s grpc.ServiceRegistrar, srv ProductCoreServiceServer) {
	s.RegisterService(&ProductCoreService_ServiceDesc, srv)
}

func _ProductCoreService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCoreServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.product_core.ProductCoreService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCoreServiceServer).AddItem(ctx, req.(*AddItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCoreService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCoreServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jlu_cow_studio.product_core.ProductCoreService/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCoreServiceServer).DeleteItem(ctx, req.(*DeleteItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCoreService_ServiceDesc is the grpc.ServiceDesc for ProductCoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jlu_cow_studio.product_core.ProductCoreService",
	HandlerType: (*ProductCoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _ProductCoreService_AddItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ProductCoreService_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_core.proto",
}
