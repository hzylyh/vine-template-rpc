/**
 * @Description: 部门相关 service
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

func (s *SystemService) AddDept(ctx context.Context, request *v1.AddDeptRequest) (*v1.AddDeptReply, error) {
	return &v1.AddDeptReply{}, nil
}

func (s *SystemService) UpdateDept(ctx context.Context, request *v1.UpdateDeptRequest) (*v1.UpdateDeptReply, error) {
	return &v1.UpdateDeptReply{}, nil
}

func (s *SystemService) DeleteDept(ctx context.Context, request *v1.DeleteDeptRequest) (*v1.DeleteDeptReply, error) {
	return &v1.DeleteDeptReply{}, nil
}

func (s *SystemService) GetDept(ctx context.Context, request *v1.GetDeptRequest) (*v1.GetDeptReply, error) {
	return &v1.GetDeptReply{}, nil
}

func (s *SystemService) ListDept(ctx context.Context, request *v1.ListDeptRequest) (*v1.ListDeptReply, error) {
	return &v1.ListDeptReply{}, nil
}
