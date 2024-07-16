// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: book_rpc.proto

package book_rpc

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

// BookRPCClient is the client API for BookRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookRPCClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Rent(ctx context.Context, in *RentRequest, opts ...grpc.CallOption) (*RentResponse, error)
}

type bookRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewBookRPCClient(cc grpc.ClientConnInterface) BookRPCClient {
	return &bookRPCClient{cc}
}

func (c *bookRPCClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/book_rpc.BookRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookRPCClient) Rent(ctx context.Context, in *RentRequest, opts ...grpc.CallOption) (*RentResponse, error) {
	out := new(RentResponse)
	err := c.cc.Invoke(ctx, "/book_rpc.BookRPC/Rent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookRPCServer is the server API for BookRPC service.
// All implementations must embed UnimplementedBookRPCServer
// for forward compatibility
type BookRPCServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Rent(context.Context, *RentRequest) (*RentResponse, error)
	mustEmbedUnimplementedBookRPCServer()
}

// UnimplementedBookRPCServer must be embedded to have forward compatible implementations.
type UnimplementedBookRPCServer struct {
}

func (UnimplementedBookRPCServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookRPCServer) Rent(context.Context, *RentRequest) (*RentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rent not implemented")
}
func (UnimplementedBookRPCServer) mustEmbedUnimplementedBookRPCServer() {}

// UnsafeBookRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookRPCServer will
// result in compilation errors.
type UnsafeBookRPCServer interface {
	mustEmbedUnimplementedBookRPCServer()
}

func RegisterBookRPCServer(s grpc.ServiceRegistrar, srv BookRPCServer) {
	s.RegisterService(&BookRPC_ServiceDesc, srv)
}

func _BookRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book_rpc.BookRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRPCServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookRPC_Rent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRPCServer).Rent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book_rpc.BookRPC/Rent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRPCServer).Rent(ctx, req.(*RentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookRPC_ServiceDesc is the grpc.ServiceDesc for BookRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_rpc.BookRPC",
	HandlerType: (*BookRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BookRPC_Get_Handler,
		},
		{
			MethodName: "Rent",
			Handler:    _BookRPC_Rent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book_rpc.proto",
}
