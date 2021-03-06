// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: internal/pb/delivery.proto

package pb

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

// DeliveryServiceClient is the client API for DeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	CreateDelivery(ctx context.Context, in *Order, opts ...grpc.CallOption) (*DeliveryID, error)
	GetDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Delivery, error)
	DeleteDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Complete, error)
	CompleteDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Complete, error)
}

type deliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryServiceClient(cc grpc.ClientConnInterface) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) CreateDelivery(ctx context.Context, in *Order, opts ...grpc.CallOption) (*DeliveryID, error) {
	out := new(DeliveryID)
	err := c.cc.Invoke(ctx, "/delivery.DeliveryService/createDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) GetDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Delivery, error) {
	out := new(Delivery)
	err := c.cc.Invoke(ctx, "/delivery.DeliveryService/getDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) DeleteDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Complete, error) {
	out := new(Complete)
	err := c.cc.Invoke(ctx, "/delivery.DeliveryService/deleteDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) CompleteDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Complete, error) {
	out := new(Complete)
	err := c.cc.Invoke(ctx, "/delivery.DeliveryService/completeDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryServiceServer is the server API for DeliveryService service.
// All implementations must embed UnimplementedDeliveryServiceServer
// for forward compatibility
type DeliveryServiceServer interface {
	CreateDelivery(context.Context, *Order) (*DeliveryID, error)
	GetDelivery(context.Context, *DeliveryID) (*Delivery, error)
	DeleteDelivery(context.Context, *DeliveryID) (*Complete, error)
	CompleteDelivery(context.Context, *DeliveryID) (*Complete, error)
	mustEmbedUnimplementedDeliveryServiceServer()
}

// UnimplementedDeliveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveryServiceServer struct {
}

func (UnimplementedDeliveryServiceServer) CreateDelivery(context.Context, *Order) (*DeliveryID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) GetDelivery(context.Context, *DeliveryID) (*Delivery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) DeleteDelivery(context.Context, *DeliveryID) (*Complete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) CompleteDelivery(context.Context, *DeliveryID) (*Complete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) mustEmbedUnimplementedDeliveryServiceServer() {}

// UnsafeDeliveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryServiceServer will
// result in compilation errors.
type UnsafeDeliveryServiceServer interface {
	mustEmbedUnimplementedDeliveryServiceServer()
}

func RegisterDeliveryServiceServer(s grpc.ServiceRegistrar, srv DeliveryServiceServer) {
	s.RegisterService(&DeliveryService_ServiceDesc, srv)
}

func _DeliveryService_CreateDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).CreateDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delivery.DeliveryService/createDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).CreateDelivery(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_GetDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).GetDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delivery.DeliveryService/getDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).GetDelivery(ctx, req.(*DeliveryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_DeleteDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).DeleteDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delivery.DeliveryService/deleteDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).DeleteDelivery(ctx, req.(*DeliveryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_CompleteDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).CompleteDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delivery.DeliveryService/completeDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).CompleteDelivery(ctx, req.(*DeliveryID))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryService_ServiceDesc is the grpc.ServiceDesc for DeliveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delivery.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDelivery",
			Handler:    _DeliveryService_CreateDelivery_Handler,
		},
		{
			MethodName: "getDelivery",
			Handler:    _DeliveryService_GetDelivery_Handler,
		},
		{
			MethodName: "deleteDelivery",
			Handler:    _DeliveryService_DeleteDelivery_Handler,
		},
		{
			MethodName: "completeDelivery",
			Handler:    _DeliveryService_CompleteDelivery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/delivery.proto",
}
