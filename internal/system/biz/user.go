/**
 * @Description: 用户相关 biz
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/16 14:06
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/16 14:06
 */

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "vine-template-rpc/api/system/v1"
	"vine-template-rpc/internal/system/data/ent"
)

type UserRepo interface {
	Add(ctx context.Context, user *ent.User) error
	Update(ctx context.Context, user *ent.User) error
	Delete(ctx context.Context, user *ent.User) error
	Get(ctx context.Context, user *ent.User) (*ent.User, error)
	List(ctx context.Context, user *ent.User) ([]*ent.User, error)
}

type UserBiz struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserBiz {
	return &UserBiz{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/user")),
	}
}

func (u *UserBiz) AddUser(ctx context.Context, request *v1.AddUserRequest) (*v1.AddUserReply, error) {
	return &v1.AddUserReply{}, nil
}

func (u *UserBiz) UpdateUser(ctx context.Context, request *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}

func (u *UserBiz) DeleteUser(ctx context.Context, request *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}

func (u *UserBiz) GetUser(ctx context.Context, request *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return &v1.GetUserReply{}, nil
}

func (u *UserBiz) ListUser(ctx context.Context, request *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return &v1.ListUserReply{}, nil
}
