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
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) Add(ctx context.Context, user *biz.User) error {
	return u.data.gdb.Create(user).Error
}

func (u *userRepo) Update(ctx context.Context, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Delete(ctx context.Context, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Get(ctx context.Context, query *biz.User) (user *biz.User, err error) {

	err = u.data.gdb.Where(query).First(user).Error
	if err != nil {
		u.log.Errorf("get user error: %v", err)
		return nil, err
	}
	return user, nil
}

func (u *userRepo) List(ctx context.Context, user *biz.User) (users []*biz.User, err error) {
	err = u.data.gdb.Where(user).Find(users).Error
	if err != nil {
		u.log.Errorf("get user error: %v", err)
		return nil, err
	}
	for _, user := range users {
		user.Password = ""
	}

	return users, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
