// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package yelp

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// YelpServiceClient is the client API for YelpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YelpServiceClient interface {
	// Sends a greeting
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateBusiness(ctx context.Context, in *CreateBusinessRequest, opts ...grpc.CallOption) (*Business, error)
	UpdateBusiness(ctx context.Context, in *UpdateBusinessRequest, opts ...grpc.CallOption) (*Business, error)
	DeleteBusiness(ctx context.Context, in *DeleteBusinessRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*Business, error)
	ListBusinesses(ctx context.Context, in *ListBusinessesRequest, opts ...grpc.CallOption) (*ListBusinessesResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type yelpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYelpServiceClient(cc grpc.ClientConnInterface) YelpServiceClient {
	return &yelpServiceClient{cc}
}

func (c *yelpServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) CreateBusiness(ctx context.Context, in *CreateBusinessRequest, opts ...grpc.CallOption) (*Business, error) {
	out := new(Business)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/CreateBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) UpdateBusiness(ctx context.Context, in *UpdateBusinessRequest, opts ...grpc.CallOption) (*Business, error) {
	out := new(Business)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/UpdateBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) DeleteBusiness(ctx context.Context, in *DeleteBusinessRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/DeleteBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*Business, error) {
	out := new(Business)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/GetBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) ListBusinesses(ctx context.Context, in *ListBusinessesRequest, opts ...grpc.CallOption) (*ListBusinessesResponse, error) {
	out := new(ListBusinessesResponse)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/ListBusinesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yelpServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/yelp.YelpService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YelpServiceServer is the server API for YelpService service.
// All implementations must embed UnimplementedYelpServiceServer
// for forward compatibility
type YelpServiceServer interface {
	// Sends a greeting
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateBusiness(context.Context, *CreateBusinessRequest) (*Business, error)
	UpdateBusiness(context.Context, *UpdateBusinessRequest) (*Business, error)
	DeleteBusiness(context.Context, *DeleteBusinessRequest) (*empty.Empty, error)
	GetBusiness(context.Context, *GetBusinessRequest) (*Business, error)
	ListBusinesses(context.Context, *ListBusinessesRequest) (*ListBusinessesResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	mustEmbedUnimplementedYelpServiceServer()
}

// UnimplementedYelpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYelpServiceServer struct {
}

func (*UnimplementedYelpServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedYelpServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedYelpServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedYelpServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedYelpServiceServer) CreateBusiness(context.Context, *CreateBusinessRequest) (*Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusiness not implemented")
}
func (*UnimplementedYelpServiceServer) UpdateBusiness(context.Context, *UpdateBusinessRequest) (*Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusiness not implemented")
}
func (*UnimplementedYelpServiceServer) DeleteBusiness(context.Context, *DeleteBusinessRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBusiness not implemented")
}
func (*UnimplementedYelpServiceServer) GetBusiness(context.Context, *GetBusinessRequest) (*Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusiness not implemented")
}
func (*UnimplementedYelpServiceServer) ListBusinesses(context.Context, *ListBusinessesRequest) (*ListBusinessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBusinesses not implemented")
}
func (*UnimplementedYelpServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedYelpServiceServer) mustEmbedUnimplementedYelpServiceServer() {}

func RegisterYelpServiceServer(s *grpc.Server, srv YelpServiceServer) {
	s.RegisterService(&_YelpService_serviceDesc, srv)
}

func _YelpService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_CreateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).CreateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/CreateBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).CreateBusiness(ctx, req.(*CreateBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_UpdateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).UpdateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/UpdateBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).UpdateBusiness(ctx, req.(*UpdateBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_DeleteBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).DeleteBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/DeleteBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).DeleteBusiness(ctx, req.(*DeleteBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_GetBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).GetBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/GetBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).GetBusiness(ctx, req.(*GetBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_ListBusinesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusinessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).ListBusinesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/ListBusinesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).ListBusinesses(ctx, req.(*ListBusinessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YelpService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YelpServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yelp.YelpService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YelpServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _YelpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yelp.YelpService",
	HandlerType: (*YelpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _YelpService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _YelpService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _YelpService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _YelpService_GetUser_Handler,
		},
		{
			MethodName: "CreateBusiness",
			Handler:    _YelpService_CreateBusiness_Handler,
		},
		{
			MethodName: "UpdateBusiness",
			Handler:    _YelpService_UpdateBusiness_Handler,
		},
		{
			MethodName: "DeleteBusiness",
			Handler:    _YelpService_DeleteBusiness_Handler,
		},
		{
			MethodName: "GetBusiness",
			Handler:    _YelpService_GetBusiness_Handler,
		},
		{
			MethodName: "ListBusinesses",
			Handler:    _YelpService_ListBusinesses_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _YelpService_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yelp.proto",
}
