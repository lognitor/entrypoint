// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/entrypoint.proto

package grpc

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
	Entrypoint_WriteLogSync_FullMethodName  = "/grpc.Entrypoint/WriteLogSync"
	Entrypoint_WriteLogAsync_FullMethodName = "/grpc.Entrypoint/WriteLogAsync"
)

// EntrypointClient is the client API for Entrypoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntrypointClient interface {
	WriteLogSync(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*PayloadReply, error)
	WriteLogAsync(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*PayloadReply, error)
}

type entrypointClient struct {
	cc grpc.ClientConnInterface
}

func NewEntrypointClient(cc grpc.ClientConnInterface) EntrypointClient {
	return &entrypointClient{cc}
}

func (c *entrypointClient) WriteLogSync(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*PayloadReply, error) {
	out := new(PayloadReply)
	err := c.cc.Invoke(ctx, Entrypoint_WriteLogSync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entrypointClient) WriteLogAsync(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*PayloadReply, error) {
	out := new(PayloadReply)
	err := c.cc.Invoke(ctx, Entrypoint_WriteLogAsync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntrypointServer is the server API for Entrypoint service.
// All implementations must embed UnimplementedEntrypointServer
// for forward compatibility
type EntrypointServer interface {
	WriteLogSync(context.Context, *PayloadRequest) (*PayloadReply, error)
	WriteLogAsync(context.Context, *PayloadRequest) (*PayloadReply, error)
	mustEmbedUnimplementedEntrypointServer()
}

// UnimplementedEntrypointServer must be embedded to have forward compatible implementations.
type UnimplementedEntrypointServer struct {
}

func (UnimplementedEntrypointServer) WriteLogSync(context.Context, *PayloadRequest) (*PayloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLogSync not implemented")
}
func (UnimplementedEntrypointServer) WriteLogAsync(context.Context, *PayloadRequest) (*PayloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLogAsync not implemented")
}
func (UnimplementedEntrypointServer) mustEmbedUnimplementedEntrypointServer() {}

// UnsafeEntrypointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntrypointServer will
// result in compilation errors.
type UnsafeEntrypointServer interface {
	mustEmbedUnimplementedEntrypointServer()
}

func RegisterEntrypointServer(s grpc.ServiceRegistrar, srv EntrypointServer) {
	s.RegisterService(&Entrypoint_ServiceDesc, srv)
}

func _Entrypoint_WriteLogSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntrypointServer).WriteLogSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Entrypoint_WriteLogSync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntrypointServer).WriteLogSync(ctx, req.(*PayloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entrypoint_WriteLogAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntrypointServer).WriteLogAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Entrypoint_WriteLogAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntrypointServer).WriteLogAsync(ctx, req.(*PayloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Entrypoint_ServiceDesc is the grpc.ServiceDesc for Entrypoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Entrypoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Entrypoint",
	HandlerType: (*EntrypointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteLogSync",
			Handler:    _Entrypoint_WriteLogSync_Handler,
		},
		{
			MethodName: "WriteLogAsync",
			Handler:    _Entrypoint_WriteLogAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/entrypoint.proto",
}
