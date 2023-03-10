// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: storpc/message.proto

package stoRPC

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

// StorpcClient is the client API for Storpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorpcClient interface {
	PutValue(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*PutResponse, error)
	GetValue(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetResponse, error)
	DeleteValue(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DelResponse, error)
}

type storpcClient struct {
	cc grpc.ClientConnInterface
}

func NewStorpcClient(cc grpc.ClientConnInterface) StorpcClient {
	return &storpcClient{cc}
}

func (c *storpcClient) PutValue(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/storpc.Storpc/PutValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storpcClient) GetValue(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/storpc.Storpc/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storpcClient) DeleteValue(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/storpc.Storpc/DeleteValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorpcServer is the server API for Storpc service.
// All implementations must embed UnimplementedStorpcServer
// for forward compatibility
type StorpcServer interface {
	PutValue(context.Context, *Pair) (*PutResponse, error)
	GetValue(context.Context, *Key) (*GetResponse, error)
	DeleteValue(context.Context, *Key) (*DelResponse, error)
	mustEmbedUnimplementedStorpcServer()
}

// UnimplementedStorpcServer must be embedded to have forward compatible implementations.
type UnimplementedStorpcServer struct {
}

func (UnimplementedStorpcServer) PutValue(context.Context, *Pair) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutValue not implemented")
}
func (UnimplementedStorpcServer) GetValue(context.Context, *Key) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedStorpcServer) DeleteValue(context.Context, *Key) (*DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteValue not implemented")
}
func (UnimplementedStorpcServer) mustEmbedUnimplementedStorpcServer() {}

// UnsafeStorpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorpcServer will
// result in compilation errors.
type UnsafeStorpcServer interface {
	mustEmbedUnimplementedStorpcServer()
}

func RegisterStorpcServer(s grpc.ServiceRegistrar, srv StorpcServer) {
	s.RegisterService(&Storpc_ServiceDesc, srv)
}

func _Storpc_PutValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorpcServer).PutValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storpc.Storpc/PutValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorpcServer).PutValue(ctx, req.(*Pair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storpc_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorpcServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storpc.Storpc/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorpcServer).GetValue(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storpc_DeleteValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorpcServer).DeleteValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storpc.Storpc/DeleteValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorpcServer).DeleteValue(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// Storpc_ServiceDesc is the grpc.ServiceDesc for Storpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storpc.Storpc",
	HandlerType: (*StorpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutValue",
			Handler:    _Storpc_PutValue_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _Storpc_GetValue_Handler,
		},
		{
			MethodName: "DeleteValue",
			Handler:    _Storpc_DeleteValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storpc/message.proto",
}
