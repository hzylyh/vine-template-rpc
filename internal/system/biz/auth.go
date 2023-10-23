/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/23 16:33
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/23 16:33
 */

package biz

import "github.com/go-kratos/kratos/v2/log"

type AuthRepo interface {
	Login()
	Logout()
	Register()
}

type AuthBiz struct {
	repo AuthRepo
	log  *log.Helper
}

func NewAuthBiz(repo AuthRepo, logger log.Logger) *AuthBiz {
	return &AuthBiz{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/auth")),
	}
}

func (a *AuthBiz) Login() {

}
