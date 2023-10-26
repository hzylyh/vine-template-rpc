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
	"vine-template-rpc/internal/system/data/ent/user"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) Add(ctx context.Context, user *biz.User) (*biz.User, error) {
	save, err := u.data.db.User.
		Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       save.ID,
		Username: save.Username,
	}, nil
}

func (u *userRepo) Update(ctx context.Context, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Delete(ctx context.Context, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Get(ctx context.Context, query *biz.User) (*biz.User, error) {
	info, err := u.data.db.User.Query().Where(user.Username(query.Username)).Only(ctx)
	if err != nil {
		u.log.Errorf("get user error: %v", err)
		return nil, err
	}
	return &biz.User{
		Id:       info.ID,
		Username: info.Username,
		Password: info.Password,
	}, nil
}

func (u *userRepo) List(ctx context.Context, user *biz.User) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
