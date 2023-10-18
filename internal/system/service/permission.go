/**
 * @Description: 权限相关 service
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

func (s *SystemService) AddPerm(ctx context.Context, request *v1.AddPermRequest) (*v1.AddPermReply, error) {
	return &v1.AddPermReply{}, nil
}

func (s *SystemService) UpdatePerm(ctx context.Context, request *v1.UpdatePermRequest) (*v1.UpdatePermReply, error) {
	return &v1.UpdatePermReply{}, nil
}

func (s *SystemService) DeletePerm(ctx context.Context, request *v1.DeletePermRequest) (*v1.DeletePermReply, error) {
	return &v1.DeletePermReply{}, nil
}

func (s *SystemService) GetPerm(ctx context.Context, request *v1.GetPermRequest) (*v1.GetPermReply, error) {
	return &v1.GetPermReply{}, nil
}

func (s *SystemService) ListPerm(ctx context.Context, request *v1.ListPermRequest) (*v1.ListPermReply, error) {
	return &v1.ListPermReply{}, nil
}
