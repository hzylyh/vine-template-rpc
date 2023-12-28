/**
 * @Description: 角色相关 dao
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/27 11:08
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/27 11:08
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vine-template-rpc/internal/system/biz"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *roleRepo) Add(ctx context.Context, role *biz.Role) error {

	return r.data.gdb.Create(role).Error
}

func (r *roleRepo) Update(ctx context.Context, role *biz.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) Delete(ctx context.Context, role *biz.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) Get(ctx context.Context, query *biz.Role) (role *biz.Role, err error) {
	err = r.data.gdb.Where(query).First(role).Error
	if err != nil {
		r.log.Errorf("get role error: %v", err)
		return nil, err
	}
	return role, nil
}

func (r *roleRepo) List(ctx context.Context, role *biz.Role) ([]*biz.Role, error) {
	//TODO implement me
	panic("implement me")
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/role")),
	}
}
