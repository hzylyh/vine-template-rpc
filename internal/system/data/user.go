/**
 * @Description: 用户相关 dao
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/23 14:31
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/23 14:31
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vine-template-rpc/internal/system/biz"
	"vine-template-rpc/internal/system/data/ent"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) Add(ctx context.Context, user *ent.User) error {
	//TODO implement me
	u.data.db.User.Create()
	return nil
}

func (u *userRepo) Update(ctx context.Context, user *ent.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Delete(ctx context.Context, user *ent.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Get(ctx context.Context, user *ent.User) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) List(ctx context.Context, user *ent.User) ([]*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
