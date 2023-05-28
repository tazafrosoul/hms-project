// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.4
// source: usermgt.proto

package usersgrpc

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

// UserMgtClient is the client API for UserMgt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMgtClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userMgtClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMgtClient(cc grpc.ClientConnInterface) UserMgtClient {
	return &userMgtClient{cc}
}

func (c *userMgtClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/UserMgt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMgtServer is the server API for UserMgt service.
// All implementations must embed UnimplementedUserMgtServer
// for forward compatibility
type UserMgtServer interface {
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserMgtServer()
}

// UnimplementedUserMgtServer must be embedded to have forward compatible implementations.
type UnimplementedUserMgtServer struct {
}

func (UnimplementedUserMgtServer) GetUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserMgtServer) mustEmbedUnimplementedUserMgtServer() {}

// UnsafeUserMgtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMgtServer will
// result in compilation errors.
type UnsafeUserMgtServer interface {
	mustEmbedUnimplementedUserMgtServer()
}

func RegisterUserMgtServer(s grpc.ServiceRegistrar, srv UserMgtServer) {
	s.RegisterService(&UserMgt_ServiceDesc, srv)
}

func _UserMgt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserMgt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgtServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMgt_ServiceDesc is the grpc.ServiceDesc for UserMgt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMgt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserMgt",
	HandlerType: (*UserMgtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserMgt_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermgt.proto",
}