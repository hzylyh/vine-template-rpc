/**
 * @Description: 权限相关 biz
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/28 21:25
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/28 21:25
 */

package biz

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	v1 "vine-template-rpc/api/system/v1"
)

type Permission struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//type PermissionRepo interface {
//	Add(permission *Permission) (*Permission, error)
//	Update(permission *Permission) error
//	Delete(permission *Permission) error
//	Get(permission *Permission) (*Permission, error)
//	List(permission *Permission) ([]*Permission, error)
//}

type PermBiz struct {
	//repo PermissionRepo
	enforcer *casbin.Enforcer
	log      *log.Helper
}

// NewPermBiz .
func NewPermBiz(
	enforcer *casbin.Enforcer,
	logger log.Logger,
) *PermBiz {
	return &PermBiz{
		enforcer: enforcer,
		log:      log.NewHelper(log.With(logger, "module", "biz/permission")),
	}
}

// AddPerm .
func (p *PermBiz) AddPerm(ctx context.Context, request *v1.AddPermRequest) (*v1.AddPermReply, error) {
	var permissions = [][]string{{"data1", "read"}, {"data2", "write"}}
	for _, permission := range permissions {
		_, err := p.enforcer.AddPermissionsForUser("alice", permission)
		if err != nil {
			p.log.Errorf("add permission error: %v", err)
			return nil, err
		}
	}
	return &v1.AddPermReply{}, nil
}
