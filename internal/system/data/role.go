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
	"vine-template-rpc/internal/system/data/ent/role"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *roleRepo) Add(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	save, err := r.data.db.Role.
		Create().
		SetName(role.Name).
		SetDescription(role.Description).
		Save(ctx)
	return &biz.Role{
		Id:          save.ID,
		Name:        save.Name,
		Description: save.Description,
	}, err
}

func (r *roleRepo) Update(ctx context.Context, role *biz.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) Delete(ctx context.Context, role *biz.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleRepo) Get(ctx context.Context, query *biz.Role) (*biz.Role, error) {
	save, err := r.data.db.Role.Query().Where(
		role.IDEQ(query.Id),
	).Only(ctx)
	if err != nil {
		r.log.Errorf("get role error: %v", err)
		return nil, err
	}
	return &biz.Role{
		Id:          save.ID,
		Name:        save.Name,
		Description: save.Description,
	}, nil
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
