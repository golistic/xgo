// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_aaa.proto

package services

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
	AAAService_Method1_FullMethodName = "/services.AAAService/Method1"
)

// AAAServiceClient is the client API for AAAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AAAServiceClient interface {
	Method1(ctx context.Context, in *Method1Request, opts ...grpc.CallOption) (*Method1Reply, error)
}

type aAAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAAAServiceClient(cc grpc.ClientConnInterface) AAAServiceClient {
	return &aAAServiceClient{cc}
}

func (c *aAAServiceClient) Method1(ctx context.Context, in *Method1Request, opts ...grpc.CallOption) (*Method1Reply, error) {
	out := new(Method1Reply)
	err := c.cc.Invoke(ctx, AAAService_Method1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AAAServiceServer is the server API for AAAService service.
// All implementations must embed UnimplementedAAAServiceServer
// for forward compatibility
type AAAServiceServer interface {
	Method1(context.Context, *Method1Request) (*Method1Reply, error)
	mustEmbedUnimplementedAAAServiceServer()
}

// UnimplementedAAAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAAAServiceServer struct {
}

func (UnimplementedAAAServiceServer) Method1(context.Context, *Method1Request) (*Method1Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method1 not implemented")
}
func (UnimplementedAAAServiceServer) mustEmbedUnimplementedAAAServiceServer() {}

// UnsafeAAAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AAAServiceServer will
// result in compilation errors.
type UnsafeAAAServiceServer interface {
	mustEmbedUnimplementedAAAServiceServer()
}

func RegisterAAAServiceServer(s grpc.ServiceRegistrar, srv AAAServiceServer) {
	s.RegisterService(&AAAService_ServiceDesc, srv)
}

func _AAAService_Method1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Method1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AAAServiceServer).Method1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AAAService_Method1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AAAServiceServer).Method1(ctx, req.(*Method1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AAAService_ServiceDesc is the grpc.ServiceDesc for AAAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AAAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.AAAService",
	HandlerType: (*AAAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method1",
			Handler:    _AAAService_Method1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_aaa.proto",
}
