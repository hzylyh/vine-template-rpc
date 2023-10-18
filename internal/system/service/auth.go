/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/17 21:04
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/17 21:04
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/system/v1"
)

func (s *SystemService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
}

func (s *SystemService) Logout(ctx context.Context, request *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}

func (s *SystemService) Register(ctx context.Context, request *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return &v1.RegisterReply{}, nil
}
