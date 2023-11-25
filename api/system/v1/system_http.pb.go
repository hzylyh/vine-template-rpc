// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: system/v1/system.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	page "vine-template-rpc/internal/page"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSystemAddDept = "/api.system.v1.System/AddDept"
const OperationSystemAddPerm = "/api.system.v1.System/AddPerm"
const OperationSystemAddRole = "/api.system.v1.System/AddRole"
const OperationSystemAddUser = "/api.system.v1.System/AddUser"
const OperationSystemBindUser = "/api.system.v1.System/BindUser"
const OperationSystemDeleteDept = "/api.system.v1.System/DeleteDept"
const OperationSystemDeletePerm = "/api.system.v1.System/DeletePerm"
const OperationSystemDeleteRole = "/api.system.v1.System/DeleteRole"
const OperationSystemDeleteUser = "/api.system.v1.System/DeleteUser"
const OperationSystemGetDept = "/api.system.v1.System/GetDept"
const OperationSystemGetPerm = "/api.system.v1.System/GetPerm"
const OperationSystemGetRole = "/api.system.v1.System/GetRole"
const OperationSystemGetUser = "/api.system.v1.System/GetUser"
const OperationSystemListDept = "/api.system.v1.System/ListDept"
const OperationSystemListPerm = "/api.system.v1.System/ListPerm"
const OperationSystemListRole = "/api.system.v1.System/ListRole"
const OperationSystemListUser = "/api.system.v1.System/ListUser"
const OperationSystemLogin = "/api.system.v1.System/Login"
const OperationSystemLogout = "/api.system.v1.System/Logout"
const OperationSystemRegister = "/api.system.v1.System/Register"
const OperationSystemUpdateDept = "/api.system.v1.System/UpdateDept"
const OperationSystemUpdatePerm = "/api.system.v1.System/UpdatePerm"
const OperationSystemUpdateRole = "/api.system.v1.System/UpdateRole"
const OperationSystemUpdateUser = "/api.system.v1.System/UpdateUser"

type SystemHTTPServer interface {
	// AddDept -------- dept 部门相关 --------
	// ---- add 新增部门信息 ----
	AddDept(context.Context, *AddDeptRequest) (*AddDeptReply, error)
	// AddPerm -------- permission --------
	// ---- add 权限新增 ----
	AddPerm(context.Context, *AddPermRequest) (*AddPermReply, error)
	// AddRole -------- role 角色相关 --------
	// ---- add 新增角色信息 ----
	AddRole(context.Context, *AddRoleRequest) (*AddRoleReply, error)
	// AddUser -------- user 用户相关 --------
	// ---- add 新增用户信息 ----
	AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
	// BindUser ---- bind_user 绑定用户 ----
	BindUser(context.Context, *BindUserRequest) (*BindUserReply, error)
	// DeleteDept ---- delete 删除部门信息 ----
	DeleteDept(context.Context, *DeleteDeptRequest) (*DeleteDeptReply, error)
	// DeletePerm ---- delete 权限删除 ----
	DeletePerm(context.Context, *DeletePermRequest) (*DeletePermReply, error)
	// DeleteRole ---- delete 删除角色信息 ----
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error)
	// DeleteUser ---- delete 删除用户信息 ----
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// GetDept ---- get 获取部门详情 ----
	GetDept(context.Context, *GetDeptRequest) (*GetDeptReply, error)
	// GetPerm ---- get 权限详情 ----
	GetPerm(context.Context, *GetPermRequest) (*GetPermReply, error)
	// GetRole ---- get 获取角色详情 ----
	GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error)
	// GetUser ---- get 获取用户详情 ----
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	// ListDept ---- list 获取部门列表 ----
	ListDept(context.Context, *ListDeptRequest) (*ListDeptReply, error)
	// ListPerm ---- list 权限列表 ----
	ListPerm(context.Context, *ListPermRequest) (*ListPermReply, error)
	// ListRole ---- list 获取角色列表 ----
	ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error)
	// ListUser ---- list 获取用户列表 ----
	ListUser(context.Context, *ListUserRequest) (*page.Page, error)
	// Login -------- auth 认证相关 --------
	// ---- login 登陆 ----
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// Logout ---- logout 登出 ----
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// Register ---- register 注册 ----
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	// UpdateDept ---- update 修改部门信息 ----
	UpdateDept(context.Context, *UpdateDeptRequest) (*UpdateDeptReply, error)
	// UpdatePerm ---- update 权限修改 ----
	UpdatePerm(context.Context, *UpdatePermRequest) (*UpdatePermReply, error)
	// UpdateRole ---- update 修改角色信息 ----
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleReply, error)
	// UpdateUser ---- update 修改用户信息 ----
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterSystemHTTPServer(s *http.Server, srv SystemHTTPServer) {
	r := s.Route("/")
	r.POST("/auth/login", _System_Login0_HTTP_Handler(srv))
	r.GET("/auth/logout", _System_Logout0_HTTP_Handler(srv))
	r.GET("/auth/register", _System_Register0_HTTP_Handler(srv))
	r.POST("/user/add", _System_AddUser0_HTTP_Handler(srv))
	r.GET("/user/update", _System_UpdateUser0_HTTP_Handler(srv))
	r.GET("/user/delete", _System_DeleteUser0_HTTP_Handler(srv))
	r.POST("/user/detail", _System_GetUser0_HTTP_Handler(srv))
	r.POST("/user/list", _System_ListUser0_HTTP_Handler(srv))
	r.POST("/role/add", _System_AddRole0_HTTP_Handler(srv))
	r.GET("/role/update", _System_UpdateRole0_HTTP_Handler(srv))
	r.GET("/role/delete", _System_DeleteRole0_HTTP_Handler(srv))
	r.GET("/role/detail", _System_GetRole0_HTTP_Handler(srv))
	r.GET("/role/list", _System_ListRole0_HTTP_Handler(srv))
	r.POST("/role/bindUser", _System_BindUser0_HTTP_Handler(srv))
	r.GET("/dept/add", _System_AddDept0_HTTP_Handler(srv))
	r.GET("/dept/update", _System_UpdateDept0_HTTP_Handler(srv))
	r.GET("/dept/delete", _System_DeleteDept0_HTTP_Handler(srv))
	r.GET("/dept/detail", _System_GetDept0_HTTP_Handler(srv))
	r.GET("/dept/list", _System_ListDept0_HTTP_Handler(srv))
	r.POST("/perm/add", _System_AddPerm0_HTTP_Handler(srv))
	r.GET("/perm/update", _System_UpdatePerm0_HTTP_Handler(srv))
	r.GET("/perm/delete", _System_DeletePerm0_HTTP_Handler(srv))
	r.GET("/perm/detail", _System_GetPerm0_HTTP_Handler(srv))
	r.GET("/perm/list", _System_ListPerm0_HTTP_Handler(srv))
}

func _System_Login0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _System_Logout0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _System_Register0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _System_AddUser0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAddUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddUser(ctx, req.(*AddUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddUserReply)
		return ctx.Result(200, reply)
	}
}

func _System_UpdateUser0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _System_DeleteUser0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

func _System_GetUser0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _System_ListUser0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*page.Page)
		return ctx.Result(200, reply)
	}
}

func _System_AddRole0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAddRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddRole(ctx, req.(*AddRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddRoleReply)
		return ctx.Result(200, reply)
	}
}

func _System_UpdateRole0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemUpdateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRole(ctx, req.(*UpdateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateRoleReply)
		return ctx.Result(200, reply)
	}
}

func _System_DeleteRole0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemDeleteRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRole(ctx, req.(*DeleteRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteRoleReply)
		return ctx.Result(200, reply)
	}
}

func _System_GetRole0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRole(ctx, req.(*GetRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRoleReply)
		return ctx.Result(200, reply)
	}
}

func _System_ListRole0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemListRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRole(ctx, req.(*ListRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRoleReply)
		return ctx.Result(200, reply)
	}
}

func _System_BindUser0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemBindUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BindUser(ctx, req.(*BindUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BindUserReply)
		return ctx.Result(200, reply)
	}
}

func _System_AddDept0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAddDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddDept(ctx, req.(*AddDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddDeptReply)
		return ctx.Result(200, reply)
	}
}

func _System_UpdateDept0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemUpdateDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDept(ctx, req.(*UpdateDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDeptReply)
		return ctx.Result(200, reply)
	}
}

func _System_DeleteDept0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemDeleteDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDept(ctx, req.(*DeleteDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDeptReply)
		return ctx.Result(200, reply)
	}
}

func _System_GetDept0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDept(ctx, req.(*GetDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDeptReply)
		return ctx.Result(200, reply)
	}
}

func _System_ListDept0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDeptRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemListDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDept(ctx, req.(*ListDeptRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDeptReply)
		return ctx.Result(200, reply)
	}
}

func _System_AddPerm0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddPermRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemAddPerm)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddPerm(ctx, req.(*AddPermRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddPermReply)
		return ctx.Result(200, reply)
	}
}

func _System_UpdatePerm0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePermRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemUpdatePerm)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePerm(ctx, req.(*UpdatePermRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePermReply)
		return ctx.Result(200, reply)
	}
}

func _System_DeletePerm0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePermRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemDeletePerm)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePerm(ctx, req.(*DeletePermRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePermReply)
		return ctx.Result(200, reply)
	}
}

func _System_GetPerm0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPermRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemGetPerm)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPerm(ctx, req.(*GetPermRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPermReply)
		return ctx.Result(200, reply)
	}
}

func _System_ListPerm0_HTTP_Handler(srv SystemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPermRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemListPerm)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPerm(ctx, req.(*ListPermRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPermReply)
		return ctx.Result(200, reply)
	}
}

type SystemHTTPClient interface {
	AddDept(ctx context.Context, req *AddDeptRequest, opts ...http.CallOption) (rsp *AddDeptReply, err error)
	AddPerm(ctx context.Context, req *AddPermRequest, opts ...http.CallOption) (rsp *AddPermReply, err error)
	AddRole(ctx context.Context, req *AddRoleRequest, opts ...http.CallOption) (rsp *AddRoleReply, err error)
	AddUser(ctx context.Context, req *AddUserRequest, opts ...http.CallOption) (rsp *AddUserReply, err error)
	BindUser(ctx context.Context, req *BindUserRequest, opts ...http.CallOption) (rsp *BindUserReply, err error)
	DeleteDept(ctx context.Context, req *DeleteDeptRequest, opts ...http.CallOption) (rsp *DeleteDeptReply, err error)
	DeletePerm(ctx context.Context, req *DeletePermRequest, opts ...http.CallOption) (rsp *DeletePermReply, err error)
	DeleteRole(ctx context.Context, req *DeleteRoleRequest, opts ...http.CallOption) (rsp *DeleteRoleReply, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	GetDept(ctx context.Context, req *GetDeptRequest, opts ...http.CallOption) (rsp *GetDeptReply, err error)
	GetPerm(ctx context.Context, req *GetPermRequest, opts ...http.CallOption) (rsp *GetPermReply, err error)
	GetRole(ctx context.Context, req *GetRoleRequest, opts ...http.CallOption) (rsp *GetRoleReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListDept(ctx context.Context, req *ListDeptRequest, opts ...http.CallOption) (rsp *ListDeptReply, err error)
	ListPerm(ctx context.Context, req *ListPermRequest, opts ...http.CallOption) (rsp *ListPermReply, err error)
	ListRole(ctx context.Context, req *ListRoleRequest, opts ...http.CallOption) (rsp *ListRoleReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *page.Page, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	UpdateDept(ctx context.Context, req *UpdateDeptRequest, opts ...http.CallOption) (rsp *UpdateDeptReply, err error)
	UpdatePerm(ctx context.Context, req *UpdatePermRequest, opts ...http.CallOption) (rsp *UpdatePermReply, err error)
	UpdateRole(ctx context.Context, req *UpdateRoleRequest, opts ...http.CallOption) (rsp *UpdateRoleReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
}

type SystemHTTPClientImpl struct {
	cc *http.Client
}

func NewSystemHTTPClient(client *http.Client) SystemHTTPClient {
	return &SystemHTTPClientImpl{client}
}

func (c *SystemHTTPClientImpl) AddDept(ctx context.Context, in *AddDeptRequest, opts ...http.CallOption) (*AddDeptReply, error) {
	var out AddDeptReply
	pattern := "/dept/add"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemAddDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AddPerm(ctx context.Context, in *AddPermRequest, opts ...http.CallOption) (*AddPermReply, error) {
	var out AddPermReply
	pattern := "/perm/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAddPerm))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AddRole(ctx context.Context, in *AddRoleRequest, opts ...http.CallOption) (*AddRoleReply, error) {
	var out AddRoleReply
	pattern := "/role/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAddRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) AddUser(ctx context.Context, in *AddUserRequest, opts ...http.CallOption) (*AddUserReply, error) {
	var out AddUserReply
	pattern := "/user/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemAddUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) BindUser(ctx context.Context, in *BindUserRequest, opts ...http.CallOption) (*BindUserReply, error) {
	var out BindUserReply
	pattern := "/role/bindUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemBindUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) DeleteDept(ctx context.Context, in *DeleteDeptRequest, opts ...http.CallOption) (*DeleteDeptReply, error) {
	var out DeleteDeptReply
	pattern := "/dept/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemDeleteDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) DeletePerm(ctx context.Context, in *DeletePermRequest, opts ...http.CallOption) (*DeletePermReply, error) {
	var out DeletePermReply
	pattern := "/perm/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemDeletePerm))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...http.CallOption) (*DeleteRoleReply, error) {
	var out DeleteRoleReply
	pattern := "/role/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemDeleteRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/user/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetDept(ctx context.Context, in *GetDeptRequest, opts ...http.CallOption) (*GetDeptReply, error) {
	var out GetDeptReply
	pattern := "/dept/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemGetDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetPerm(ctx context.Context, in *GetPermRequest, opts ...http.CallOption) (*GetPermReply, error) {
	var out GetPermReply
	pattern := "/perm/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemGetPerm))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetRole(ctx context.Context, in *GetRoleRequest, opts ...http.CallOption) (*GetRoleReply, error) {
	var out GetRoleReply
	pattern := "/role/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemGetRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/user/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) ListDept(ctx context.Context, in *ListDeptRequest, opts ...http.CallOption) (*ListDeptReply, error) {
	var out ListDeptReply
	pattern := "/dept/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemListDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) ListPerm(ctx context.Context, in *ListPermRequest, opts ...http.CallOption) (*ListPermReply, error) {
	var out ListPermReply
	pattern := "/perm/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemListPerm))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) ListRole(ctx context.Context, in *ListRoleRequest, opts ...http.CallOption) (*ListRoleReply, error) {
	var out ListRoleReply
	pattern := "/role/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemListRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*page.Page, error) {
	var out page.Page
	pattern := "/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/auth/logout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/auth/register"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) UpdateDept(ctx context.Context, in *UpdateDeptRequest, opts ...http.CallOption) (*UpdateDeptReply, error) {
	var out UpdateDeptReply
	pattern := "/dept/update"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemUpdateDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) UpdatePerm(ctx context.Context, in *UpdatePermRequest, opts ...http.CallOption) (*UpdatePermReply, error) {
	var out UpdatePermReply
	pattern := "/perm/update"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemUpdatePerm))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...http.CallOption) (*UpdateRoleReply, error) {
	var out UpdateRoleReply
	pattern := "/role/update"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemUpdateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SystemHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/user/update"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
