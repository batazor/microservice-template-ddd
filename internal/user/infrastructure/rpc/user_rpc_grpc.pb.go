// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: user_rpc.proto

package user_rpc

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

// UserRPCClient is the client API for UserRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRPCClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type userRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRPCClient(cc grpc.ClientConnInterface) UserRPCClient {
	return &userRPCClient{cc}
}

func (c *userRPCClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_rpc.UserRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRPCServer is the server API for UserRPC service.
// All implementations must embed UnimplementedUserRPCServer
// for forward compatibility
type UserRPCServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedUserRPCServer()
}

// UnimplementedUserRPCServer must be embedded to have forward compatible implementations.
type UnimplementedUserRPCServer struct {
}

func (UnimplementedUserRPCServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserRPCServer) mustEmbedUnimplementedUserRPCServer() {}

// UnsafeUserRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRPCServer will
// result in compilation errors.
type UnsafeUserRPCServer interface {
	mustEmbedUnimplementedUserRPCServer()
}

func RegisterUserRPCServer(s grpc.ServiceRegistrar, srv UserRPCServer) {
	s.RegisterService(&UserRPC_ServiceDesc, srv)
}

func _UserRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_rpc.UserRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRPC_ServiceDesc is the grpc.ServiceDesc for UserRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_rpc.UserRPC",
	HandlerType: (*UserRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserRPC_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_rpc.proto",
}
