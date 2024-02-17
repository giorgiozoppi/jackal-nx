// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	// CreateUser creates a new user given a username and password.
	//
	// Return status codes (https://github.com/grpc/grpc/blob/master/doc/statuscodes.md):
	// - ALREADY_EXISTS(6):  When a users already exists.
	// - INTERNAL(13): When an internal problem happens.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// ChangeUserPassword updates the password of an existing user.
	//
	// Return status codes (https://github.com/grpc/grpc/blob/master/doc/statuscodes.md):
	// - NOT_FOUND(5):  When user does not exist.
	// - INTERNAL(13): When an internal problem happens.
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error)
	// DeleteUser removes a previously registered user.
	//
	// Return status codes (https://github.com/grpc/grpc/blob/master/doc/statuscodes.md):
	// - NOT_FOUND(5):  When user does not exist.
	// - INTERNAL(13): When an internal problem happens.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/admin.v1.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error) {
	out := new(ChangeUserPasswordResponse)
	err := c.cc.Invoke(ctx, "/admin.v1.Users/ChangeUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/admin.v1.Users/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	// CreateUser creates a new user given a username and password.
	//
	// Return status codes (https://github.com/grpc/grpc/blob/master/doc/statuscodes.md):
	// - ALREADY_EXISTS(6):  When a users already exists.
	// - INTERNAL(13): When an internal problem happens.
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// ChangeUserPassword updates the password of an existing user.
	//
	// Return status codes (https://github.com/grpc/grpc/blob/master/doc/statuscodes.md):
	// - NOT_FOUND(5):  When user does not exist.
	// - INTERNAL(13): When an internal problem happens.
	ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error)
	// DeleteUser removes a previously registered user.
	//
	// Return status codes (https://github.com/grpc/grpc/blob/master/doc/statuscodes.md):
	// - NOT_FOUND(5):  When user does not exist.
	// - INTERNAL(13): When an internal problem happens.
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServer) ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (UnimplementedUsersServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.v1.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.v1.Users/ChangeUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.v1.Users/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _Users_ChangeUserPassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/admin/v1/users.proto",
}
