/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/23 16:33
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/23 16:33
 */

package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "vine-template-rpc/api/system/v1"
	"vine-template-rpc/pkg/vjwt"
)

type AuthBiz struct {
	userRepo UserRepo
	log      *log.Helper
}

func NewAuthBiz(userRepo UserRepo, logger log.Logger) *AuthBiz {
	return &AuthBiz{
		userRepo: userRepo,
		log:      log.NewHelper(log.With(logger, "module", "biz/auth")),
	}
}

func (a *AuthBiz) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginReply, error) {
	userInfo, err := a.userRepo.Get(ctx, &User{
		Username: request.Username,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("系统内部错误，请联系管理员, %v", err))
	}
	if userInfo == nil {
		return nil, errors.New("用户不存在")
	}
	if userInfo.Password != request.Password {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := vjwt.GenerateJwtToken(userInfo.Id.String(), userInfo.Username)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("系统内部错误，请联系管理员, %v", err))
	}
	return &v1.LoginReply{
		Token: token,
	}, nil
}
