// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/associate/v1/associate-api.proto

package associate_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AssociateApiService_ListStimusWords_FullMethodName        = "/user.api.associateApiService/ListStimusWords"
	AssociateApiService_CreateExperiment_FullMethodName       = "/user.api.associateApiService/CreateExperiment"
	AssociateApiService_CreateExperimentResult_FullMethodName = "/user.api.associateApiService/CreateExperimentResult"
	AssociateApiService_GetExperiment_FullMethodName          = "/user.api.associateApiService/GetExperiment"
	AssociateApiService_CreateUser_FullMethodName             = "/user.api.associateApiService/CreateUser"
	AssociateApiService_ListUsers_FullMethodName              = "/user.api.associateApiService/ListUsers"
	AssociateApiService_GetUser_FullMethodName                = "/user.api.associateApiService/GetUser"
	AssociateApiService_Login_FullMethodName                  = "/user.api.associateApiService/Login"
	AssociateApiService_Logout_FullMethodName                 = "/user.api.associateApiService/Logout"
	AssociateApiService_CreateRole_FullMethodName             = "/user.api.associateApiService/CreateRole"
	AssociateApiService_UpdateRole_FullMethodName             = "/user.api.associateApiService/UpdateRole"
	AssociateApiService_DeleteRole_FullMethodName             = "/user.api.associateApiService/DeleteRole"
	AssociateApiService_GetRole_FullMethodName                = "/user.api.associateApiService/GetRole"
	AssociateApiService_ListRoles_FullMethodName              = "/user.api.associateApiService/ListRoles"
	AssociateApiService_CreateUserRole_FullMethodName         = "/user.api.associateApiService/CreateUserRole"
	AssociateApiService_GetUserRole_FullMethodName            = "/user.api.associateApiService/GetUserRole"
	AssociateApiService_DeleteUserRole_FullMethodName         = "/user.api.associateApiService/DeleteUserRole"
	AssociateApiService_ListUserRoles_FullMethodName          = "/user.api.associateApiService/ListUserRoles"
)

// AssociateApiServiceClient is the client API for AssociateApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssociateApiServiceClient interface {
	ListStimusWords(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStimusWordsResponse, error)
	CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*CreateExperimentResponse, error)
	CreateExperimentResult(ctx context.Context, in *CreateExperimentResultRequest, opts ...grpc.CallOption) (*CreateExperimentResultResponse, error)
	GetExperiment(ctx context.Context, in *GetExperimentRequest, opts ...grpc.CallOption) (*GetExperimentResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ListUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error)
	ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRolesResponse, error)
	CreateUserRole(ctx context.Context, in *CreateUserRoleRequest, opts ...grpc.CallOption) (*CreateUserRoleResponse, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
	DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListUserRoles(ctx context.Context, in *ListUserRolesRequest, opts ...grpc.CallOption) (*ListUserRolesResponse, error)
}

type associateApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssociateApiServiceClient(cc grpc.ClientConnInterface) AssociateApiServiceClient {
	return &associateApiServiceClient{cc}
}

func (c *associateApiServiceClient) ListStimusWords(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStimusWordsResponse, error) {
	out := new(ListStimusWordsResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_ListStimusWords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*CreateExperimentResponse, error) {
	out := new(CreateExperimentResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_CreateExperiment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) CreateExperimentResult(ctx context.Context, in *CreateExperimentResultRequest, opts ...grpc.CallOption) (*CreateExperimentResultResponse, error) {
	out := new(CreateExperimentResultResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_CreateExperimentResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) GetExperiment(ctx context.Context, in *GetExperimentRequest, opts ...grpc.CallOption) (*GetExperimentResponse, error) {
	out := new(GetExperimentResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_GetExperiment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) ListUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_ListUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AssociateApiService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AssociateApiService_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error) {
	out := new(GetRoleResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_ListRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) CreateUserRole(ctx context.Context, in *CreateUserRoleRequest, opts ...grpc.CallOption) (*CreateUserRoleResponse, error) {
	out := new(CreateUserRoleResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_CreateUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	out := new(GetUserRoleResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_GetUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AssociateApiService_DeleteUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *associateApiServiceClient) ListUserRoles(ctx context.Context, in *ListUserRolesRequest, opts ...grpc.CallOption) (*ListUserRolesResponse, error) {
	out := new(ListUserRolesResponse)
	err := c.cc.Invoke(ctx, AssociateApiService_ListUserRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssociateApiServiceServer is the server API for AssociateApiService service.
// All implementations must embed UnimplementedAssociateApiServiceServer
// for forward compatibility
type AssociateApiServiceServer interface {
	ListStimusWords(context.Context, *emptypb.Empty) (*ListStimusWordsResponse, error)
	CreateExperiment(context.Context, *CreateExperimentRequest) (*CreateExperimentResponse, error)
	CreateExperimentResult(context.Context, *CreateExperimentResultRequest) (*CreateExperimentResultResponse, error)
	GetExperiment(context.Context, *GetExperimentRequest) (*GetExperimentResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ListUsers(context.Context, *emptypb.Empty) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error)
	GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error)
	ListRoles(context.Context, *emptypb.Empty) (*ListRolesResponse, error)
	CreateUserRole(context.Context, *CreateUserRoleRequest) (*CreateUserRoleResponse, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error)
	DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*emptypb.Empty, error)
	ListUserRoles(context.Context, *ListUserRolesRequest) (*ListUserRolesResponse, error)
	mustEmbedUnimplementedAssociateApiServiceServer()
}

// UnimplementedAssociateApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssociateApiServiceServer struct {
}

func (UnimplementedAssociateApiServiceServer) ListStimusWords(context.Context, *emptypb.Empty) (*ListStimusWordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStimusWords not implemented")
}
func (UnimplementedAssociateApiServiceServer) CreateExperiment(context.Context, *CreateExperimentRequest) (*CreateExperimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExperiment not implemented")
}
func (UnimplementedAssociateApiServiceServer) CreateExperimentResult(context.Context, *CreateExperimentResultRequest) (*CreateExperimentResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExperimentResult not implemented")
}
func (UnimplementedAssociateApiServiceServer) GetExperiment(context.Context, *GetExperimentRequest) (*GetExperimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperiment not implemented")
}
func (UnimplementedAssociateApiServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAssociateApiServiceServer) ListUsers(context.Context, *emptypb.Empty) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedAssociateApiServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAssociateApiServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAssociateApiServiceServer) Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAssociateApiServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) ListRoles(context.Context, *emptypb.Empty) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedAssociateApiServiceServer) CreateUserRole(context.Context, *CreateUserRoleRequest) (*CreateUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRole not implemented")
}
func (UnimplementedAssociateApiServiceServer) ListUserRoles(context.Context, *ListUserRolesRequest) (*ListUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRoles not implemented")
}
func (UnimplementedAssociateApiServiceServer) mustEmbedUnimplementedAssociateApiServiceServer() {}

// UnsafeAssociateApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssociateApiServiceServer will
// result in compilation errors.
type UnsafeAssociateApiServiceServer interface {
	mustEmbedUnimplementedAssociateApiServiceServer()
}

func RegisterAssociateApiServiceServer(s grpc.ServiceRegistrar, srv AssociateApiServiceServer) {
	s.RegisterService(&AssociateApiService_ServiceDesc, srv)
}

func _AssociateApiService_ListStimusWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).ListStimusWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_ListStimusWords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).ListStimusWords(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_CreateExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).CreateExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_CreateExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).CreateExperiment(ctx, req.(*CreateExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_CreateExperimentResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperimentResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).CreateExperimentResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_CreateExperimentResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).CreateExperimentResult(ctx, req.(*CreateExperimentResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_GetExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).GetExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_GetExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).GetExperiment(ctx, req.(*GetExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).ListUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).Logout(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_ListRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).ListRoles(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_CreateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).CreateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_CreateUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).CreateUserRole(ctx, req.(*CreateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_GetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_DeleteUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).DeleteUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_DeleteUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).DeleteUserRole(ctx, req.(*DeleteUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssociateApiService_ListUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssociateApiServiceServer).ListUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssociateApiService_ListUserRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssociateApiServiceServer).ListUserRoles(ctx, req.(*ListUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssociateApiService_ServiceDesc is the grpc.ServiceDesc for AssociateApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssociateApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.associateApiService",
	HandlerType: (*AssociateApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStimusWords",
			Handler:    _AssociateApiService_ListStimusWords_Handler,
		},
		{
			MethodName: "CreateExperiment",
			Handler:    _AssociateApiService_CreateExperiment_Handler,
		},
		{
			MethodName: "CreateExperimentResult",
			Handler:    _AssociateApiService_CreateExperimentResult_Handler,
		},
		{
			MethodName: "GetExperiment",
			Handler:    _AssociateApiService_GetExperiment_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AssociateApiService_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _AssociateApiService_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AssociateApiService_GetUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AssociateApiService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AssociateApiService_Logout_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _AssociateApiService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _AssociateApiService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _AssociateApiService_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _AssociateApiService_GetRole_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _AssociateApiService_ListRoles_Handler,
		},
		{
			MethodName: "CreateUserRole",
			Handler:    _AssociateApiService_CreateUserRole_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _AssociateApiService_GetUserRole_Handler,
		},
		{
			MethodName: "DeleteUserRole",
			Handler:    _AssociateApiService_DeleteUserRole_Handler,
		},
		{
			MethodName: "ListUserRoles",
			Handler:    _AssociateApiService_ListUserRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/associate/v1/associate-api.proto",
}
