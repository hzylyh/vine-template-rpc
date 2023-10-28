// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: system/v1/system.proto

package v1

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
	System_Login_FullMethodName      = "/api.system.v1.System/Login"
	System_Logout_FullMethodName     = "/api.system.v1.System/Logout"
	System_Register_FullMethodName   = "/api.system.v1.System/Register"
	System_AddUser_FullMethodName    = "/api.system.v1.System/AddUser"
	System_UpdateUser_FullMethodName = "/api.system.v1.System/UpdateUser"
	System_DeleteUser_FullMethodName = "/api.system.v1.System/DeleteUser"
	System_GetUser_FullMethodName    = "/api.system.v1.System/GetUser"
	System_ListUser_FullMethodName   = "/api.system.v1.System/ListUser"
	System_AddRole_FullMethodName    = "/api.system.v1.System/AddRole"
	System_UpdateRole_FullMethodName = "/api.system.v1.System/UpdateRole"
	System_DeleteRole_FullMethodName = "/api.system.v1.System/DeleteRole"
	System_GetRole_FullMethodName    = "/api.system.v1.System/GetRole"
	System_ListRole_FullMethodName   = "/api.system.v1.System/ListRole"
	System_BindUser_FullMethodName   = "/api.system.v1.System/BindUser"
	System_AddDept_FullMethodName    = "/api.system.v1.System/AddDept"
	System_UpdateDept_FullMethodName = "/api.system.v1.System/UpdateDept"
	System_DeleteDept_FullMethodName = "/api.system.v1.System/DeleteDept"
	System_GetDept_FullMethodName    = "/api.system.v1.System/GetDept"
	System_ListDept_FullMethodName   = "/api.system.v1.System/ListDept"
	System_AddPerm_FullMethodName    = "/api.system.v1.System/AddPerm"
	System_UpdatePerm_FullMethodName = "/api.system.v1.System/UpdatePerm"
	System_DeletePerm_FullMethodName = "/api.system.v1.System/DeletePerm"
	System_GetPerm_FullMethodName    = "/api.system.v1.System/GetPerm"
	System_ListPerm_FullMethodName   = "/api.system.v1.System/ListPerm"
)

// SystemClient is the client API for System service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemClient interface {
	// -------- auth 认证相关 --------
	// ---- login 登陆 ----
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// ---- logout 登出 ----
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	// ---- register 注册 ----
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	// -------- user 用户相关 --------
	// ---- add 新增用户信息 ----
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserReply, error)
	// ---- update 修改用户信息 ----
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	// ---- delete 删除用户信息 ----
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error)
	// ---- get 获取用户详情 ----
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	// ---- list 获取用户列表 ----
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	// -------- role 角色相关 --------
	// ---- add 新增角色信息 ----
	AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleReply, error)
	// ---- update 修改角色信息 ----
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleReply, error)
	// ---- delete 删除角色信息 ----
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleReply, error)
	// ---- get 获取角色详情 ----
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleReply, error)
	// ---- list 获取角色列表 ----
	ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleReply, error)
	// ---- bind_user 绑定用户 ----
	BindUser(ctx context.Context, in *BindUserRequest, opts ...grpc.CallOption) (*BindUserReply, error)
	// -------- dept 部门相关 --------
	// ---- add 新增部门信息 ----
	AddDept(ctx context.Context, in *AddDeptRequest, opts ...grpc.CallOption) (*AddDeptReply, error)
	// ---- update 修改部门信息 ----
	UpdateDept(ctx context.Context, in *UpdateDeptRequest, opts ...grpc.CallOption) (*UpdateDeptReply, error)
	// ---- delete 删除部门信息 ----
	DeleteDept(ctx context.Context, in *DeleteDeptRequest, opts ...grpc.CallOption) (*DeleteDeptReply, error)
	// ---- get 获取部门详情 ----
	GetDept(ctx context.Context, in *GetDeptRequest, opts ...grpc.CallOption) (*GetDeptReply, error)
	// ---- list 获取部门列表 ----
	ListDept(ctx context.Context, in *ListDeptRequest, opts ...grpc.CallOption) (*ListDeptReply, error)
	// -------- permission --------
	// ---- add 权限新增 ----
	AddPerm(ctx context.Context, in *AddPermRequest, opts ...grpc.CallOption) (*AddPermReply, error)
	// ---- update 权限修改 ----
	UpdatePerm(ctx context.Context, in *UpdatePermRequest, opts ...grpc.CallOption) (*UpdatePermReply, error)
	// ---- delete 权限删除 ----
	DeletePerm(ctx context.Context, in *DeletePermRequest, opts ...grpc.CallOption) (*DeletePermReply, error)
	// ---- get 权限详情 ----
	GetPerm(ctx context.Context, in *GetPermRequest, opts ...grpc.CallOption) (*GetPermReply, error)
	// ---- list 权限列表 ----
	ListPerm(ctx context.Context, in *ListPermRequest, opts ...grpc.CallOption) (*ListPermReply, error)
}

type systemClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemClient(cc grpc.ClientConnInterface) SystemClient {
	return &systemClient{cc}
}

func (c *systemClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, System_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, System_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, System_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserReply, error) {
	out := new(AddUserReply)
	err := c.cc.Invoke(ctx, System_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, System_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, System_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, System_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, System_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleReply, error) {
	out := new(AddRoleReply)
	err := c.cc.Invoke(ctx, System_AddRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleReply, error) {
	out := new(UpdateRoleReply)
	err := c.cc.Invoke(ctx, System_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleReply, error) {
	out := new(DeleteRoleReply)
	err := c.cc.Invoke(ctx, System_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleReply, error) {
	out := new(GetRoleReply)
	err := c.cc.Invoke(ctx, System_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleReply, error) {
	out := new(ListRoleReply)
	err := c.cc.Invoke(ctx, System_ListRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) BindUser(ctx context.Context, in *BindUserRequest, opts ...grpc.CallOption) (*BindUserReply, error) {
	out := new(BindUserReply)
	err := c.cc.Invoke(ctx, System_BindUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) AddDept(ctx context.Context, in *AddDeptRequest, opts ...grpc.CallOption) (*AddDeptReply, error) {
	out := new(AddDeptReply)
	err := c.cc.Invoke(ctx, System_AddDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UpdateDept(ctx context.Context, in *UpdateDeptRequest, opts ...grpc.CallOption) (*UpdateDeptReply, error) {
	out := new(UpdateDeptReply)
	err := c.cc.Invoke(ctx, System_UpdateDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) DeleteDept(ctx context.Context, in *DeleteDeptRequest, opts ...grpc.CallOption) (*DeleteDeptReply, error) {
	out := new(DeleteDeptReply)
	err := c.cc.Invoke(ctx, System_DeleteDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) GetDept(ctx context.Context, in *GetDeptRequest, opts ...grpc.CallOption) (*GetDeptReply, error) {
	out := new(GetDeptReply)
	err := c.cc.Invoke(ctx, System_GetDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) ListDept(ctx context.Context, in *ListDeptRequest, opts ...grpc.CallOption) (*ListDeptReply, error) {
	out := new(ListDeptReply)
	err := c.cc.Invoke(ctx, System_ListDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) AddPerm(ctx context.Context, in *AddPermRequest, opts ...grpc.CallOption) (*AddPermReply, error) {
	out := new(AddPermReply)
	err := c.cc.Invoke(ctx, System_AddPerm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UpdatePerm(ctx context.Context, in *UpdatePermRequest, opts ...grpc.CallOption) (*UpdatePermReply, error) {
	out := new(UpdatePermReply)
	err := c.cc.Invoke(ctx, System_UpdatePerm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) DeletePerm(ctx context.Context, in *DeletePermRequest, opts ...grpc.CallOption) (*DeletePermReply, error) {
	out := new(DeletePermReply)
	err := c.cc.Invoke(ctx, System_DeletePerm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) GetPerm(ctx context.Context, in *GetPermRequest, opts ...grpc.CallOption) (*GetPermReply, error) {
	out := new(GetPermReply)
	err := c.cc.Invoke(ctx, System_GetPerm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) ListPerm(ctx context.Context, in *ListPermRequest, opts ...grpc.CallOption) (*ListPermReply, error) {
	out := new(ListPermReply)
	err := c.cc.Invoke(ctx, System_ListPerm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServer is the server API for System service.
// All implementations must embed UnimplementedSystemServer
// for forward compatibility
type SystemServer interface {
	// -------- auth 认证相关 --------
	// ---- login 登陆 ----
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// ---- logout 登出 ----
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// ---- register 注册 ----
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	// -------- user 用户相关 --------
	// ---- add 新增用户信息 ----
	AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
	// ---- update 修改用户信息 ----
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	// ---- delete 删除用户信息 ----
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// ---- get 获取用户详情 ----
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	// ---- list 获取用户列表 ----
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// -------- role 角色相关 --------
	// ---- add 新增角色信息 ----
	AddRole(context.Context, *AddRoleRequest) (*AddRoleReply, error)
	// ---- update 修改角色信息 ----
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleReply, error)
	// ---- delete 删除角色信息 ----
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error)
	// ---- get 获取角色详情 ----
	GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error)
	// ---- list 获取角色列表 ----
	ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error)
	// ---- bind_user 绑定用户 ----
	BindUser(context.Context, *BindUserRequest) (*BindUserReply, error)
	// -------- dept 部门相关 --------
	// ---- add 新增部门信息 ----
	AddDept(context.Context, *AddDeptRequest) (*AddDeptReply, error)
	// ---- update 修改部门信息 ----
	UpdateDept(context.Context, *UpdateDeptRequest) (*UpdateDeptReply, error)
	// ---- delete 删除部门信息 ----
	DeleteDept(context.Context, *DeleteDeptRequest) (*DeleteDeptReply, error)
	// ---- get 获取部门详情 ----
	GetDept(context.Context, *GetDeptRequest) (*GetDeptReply, error)
	// ---- list 获取部门列表 ----
	ListDept(context.Context, *ListDeptRequest) (*ListDeptReply, error)
	// -------- permission --------
	// ---- add 权限新增 ----
	AddPerm(context.Context, *AddPermRequest) (*AddPermReply, error)
	// ---- update 权限修改 ----
	UpdatePerm(context.Context, *UpdatePermRequest) (*UpdatePermReply, error)
	// ---- delete 权限删除 ----
	DeletePerm(context.Context, *DeletePermRequest) (*DeletePermReply, error)
	// ---- get 权限详情 ----
	GetPerm(context.Context, *GetPermRequest) (*GetPermReply, error)
	// ---- list 权限列表 ----
	ListPerm(context.Context, *ListPermRequest) (*ListPermReply, error)
	mustEmbedUnimplementedSystemServer()
}

// UnimplementedSystemServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServer struct {
}

func (UnimplementedSystemServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSystemServer) Logout(context.Context, *LogoutRequest) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSystemServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSystemServer) AddUser(context.Context, *AddUserRequest) (*AddUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedSystemServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedSystemServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedSystemServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedSystemServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedSystemServer) AddRole(context.Context, *AddRoleRequest) (*AddRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedSystemServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedSystemServer) DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedSystemServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedSystemServer) ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedSystemServer) BindUser(context.Context, *BindUserRequest) (*BindUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindUser not implemented")
}
func (UnimplementedSystemServer) AddDept(context.Context, *AddDeptRequest) (*AddDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDept not implemented")
}
func (UnimplementedSystemServer) UpdateDept(context.Context, *UpdateDeptRequest) (*UpdateDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDept not implemented")
}
func (UnimplementedSystemServer) DeleteDept(context.Context, *DeleteDeptRequest) (*DeleteDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDept not implemented")
}
func (UnimplementedSystemServer) GetDept(context.Context, *GetDeptRequest) (*GetDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDept not implemented")
}
func (UnimplementedSystemServer) ListDept(context.Context, *ListDeptRequest) (*ListDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDept not implemented")
}
func (UnimplementedSystemServer) AddPerm(context.Context, *AddPermRequest) (*AddPermReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerm not implemented")
}
func (UnimplementedSystemServer) UpdatePerm(context.Context, *UpdatePermRequest) (*UpdatePermReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerm not implemented")
}
func (UnimplementedSystemServer) DeletePerm(context.Context, *DeletePermRequest) (*DeletePermReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerm not implemented")
}
func (UnimplementedSystemServer) GetPerm(context.Context, *GetPermRequest) (*GetPermReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerm not implemented")
}
func (UnimplementedSystemServer) ListPerm(context.Context, *ListPermRequest) (*ListPermReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerm not implemented")
}
func (UnimplementedSystemServer) mustEmbedUnimplementedSystemServer() {}

// UnsafeSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServer will
// result in compilation errors.
type UnsafeSystemServer interface {
	mustEmbedUnimplementedSystemServer()
}

func RegisterSystemServer(s grpc.ServiceRegistrar, srv SystemServer) {
	s.RegisterService(&System_ServiceDesc, srv)
}

func _System_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_AddRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).AddRole(ctx, req.(*AddRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).ListRole(ctx, req.(*ListRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_BindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).BindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_BindUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).BindUser(ctx, req.(*BindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_AddDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).AddDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_AddDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).AddDept(ctx, req.(*AddDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UpdateDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UpdateDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UpdateDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UpdateDept(ctx, req.(*UpdateDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_DeleteDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).DeleteDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_DeleteDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).DeleteDept(ctx, req.(*DeleteDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_GetDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_GetDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetDept(ctx, req.(*GetDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_ListDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).ListDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_ListDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).ListDept(ctx, req.(*ListDeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_AddPerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).AddPerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_AddPerm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).AddPerm(ctx, req.(*AddPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UpdatePerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UpdatePerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UpdatePerm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UpdatePerm(ctx, req.(*UpdatePermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_DeletePerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).DeletePerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_DeletePerm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).DeletePerm(ctx, req.(*DeletePermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_GetPerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).GetPerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_GetPerm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).GetPerm(ctx, req.(*GetPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_ListPerm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).ListPerm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_ListPerm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).ListPerm(ctx, req.(*ListPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// System_ServiceDesc is the grpc.ServiceDesc for System service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var System_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.v1.System",
	HandlerType: (*SystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _System_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _System_Logout_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _System_Register_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _System_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _System_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _System_DeleteUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _System_GetUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _System_ListUser_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _System_AddRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _System_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _System_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _System_GetRole_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _System_ListRole_Handler,
		},
		{
			MethodName: "BindUser",
			Handler:    _System_BindUser_Handler,
		},
		{
			MethodName: "AddDept",
			Handler:    _System_AddDept_Handler,
		},
		{
			MethodName: "UpdateDept",
			Handler:    _System_UpdateDept_Handler,
		},
		{
			MethodName: "DeleteDept",
			Handler:    _System_DeleteDept_Handler,
		},
		{
			MethodName: "GetDept",
			Handler:    _System_GetDept_Handler,
		},
		{
			MethodName: "ListDept",
			Handler:    _System_ListDept_Handler,
		},
		{
			MethodName: "AddPerm",
			Handler:    _System_AddPerm_Handler,
		},
		{
			MethodName: "UpdatePerm",
			Handler:    _System_UpdatePerm_Handler,
		},
		{
			MethodName: "DeletePerm",
			Handler:    _System_DeletePerm_Handler,
		},
		{
			MethodName: "GetPerm",
			Handler:    _System_GetPerm_Handler,
		},
		{
			MethodName: "ListPerm",
			Handler:    _System_ListPerm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/v1/system.proto",
}
