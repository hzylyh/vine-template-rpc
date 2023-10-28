/**
 * @Description: 角色相关 biz
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/27 11:13
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/27 11:13
 */

package biz

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	v1 "vine-template-rpc/api/system/v1"
)

type RoleRepo interface {
	Add(ctx context.Context, role *Role) (*Role, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, role *Role) error
	Get(ctx context.Context, role *Role) (*Role, error)
	List(ctx context.Context, role *Role) ([]*Role, error)
}

type Role struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleBiz struct {
	repo     RoleRepo
	userRepo UserRepo
	enforcer *casbin.Enforcer
	log      *log.Helper
}

// NewRoleBiz .
func NewRoleBiz(
	repo RoleRepo,
	userRepo UserRepo,
	enforcer *casbin.Enforcer,
	logger log.Logger) *RoleBiz {
	return &RoleBiz{
		repo:     repo,
		userRepo: userRepo,
		enforcer: enforcer,
		log:      log.NewHelper(log.With(logger, "module", "biz/role")),
	}
}

func (r *RoleBiz) AddRole(ctx context.Context, request *v1.AddRoleRequest) (*v1.AddRoleReply, error) {
	role, err := r.repo.Add(ctx, &Role{
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		r.log.Errorf("add role error: %v", err)
		return nil, err
	}
	return &v1.AddRoleReply{
		Id:   role.Id,
		Name: role.Name,
	}, nil
}

func (r *RoleBiz) BindUser(ctx context.Context, request *v1.BindUserRequest) (*v1.BindUserReply, error) {
	roleInfo, err := r.repo.Get(ctx, &Role{
		Id: request.RoleId,
	})
	if err != nil {
		r.log.Errorf("get role error: %v", err)
		return nil, err
	}
	for _, username := range request.UsernameList {
		user, err := r.userRepo.Get(ctx, &User{
			Username: username,
		})
		if err != nil {
			r.log.Errorf("get user error: %v", err)
			return nil, err
		}
		//userWithID := user.Username + "_" + strconv.FormatInt(user.Id, 10)
		userId := user.Id
		_, err = r.enforcer.AddRoleForUser(userId.String(), strconv.FormatInt(roleInfo.Id, 10))
		if err != nil {
			r.log.Errorf("bind user error: %v", err)
			return nil, err
		}
	}
	return &v1.BindUserReply{}, nil
}
