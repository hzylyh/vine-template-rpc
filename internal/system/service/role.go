/**
 * @Description: 角色相关 service
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/18 15:17
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/18 15:17
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/system/v1"
)

func (s *SystemService) AddRole(ctx context.Context, request *v1.AddRoleRequest) (*v1.AddRoleReply, error) {
	return s.rb.AddRole(ctx, request)
}

func (s *SystemService) UpdateRole(ctx context.Context, request *v1.UpdateRoleRequest) (*v1.UpdateRoleReply, error) {
	return &v1.UpdateRoleReply{}, nil
}

func (s *SystemService) DeleteRole(ctx context.Context, request *v1.DeleteRoleRequest) (*v1.DeleteRoleReply, error) {
	return &v1.DeleteRoleReply{}, nil
}

func (s *SystemService) GetRole(ctx context.Context, request *v1.GetRoleRequest) (*v1.GetRoleReply, error) {
	return &v1.GetRoleReply{}, nil
}

func (s *SystemService) ListRole(ctx context.Context, request *v1.ListRoleRequest) (*v1.ListRoleReply, error) {
	return &v1.ListRoleReply{}, nil
}

func (s *SystemService) BindUser(ctx context.Context, request *v1.BindUserRequest) (*v1.BindUserReply, error) {
	return s.rb.BindUser(ctx, request)
}
