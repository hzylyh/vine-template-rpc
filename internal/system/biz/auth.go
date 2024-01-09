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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	"net/url"
	"time"
	v1 "vine-template-rpc/api/system/v1"
	"vine-template-rpc/pkg/vjwt"
)

type AuthRepo interface {
	SetWXAccessToken(token string) error
	GetWXAccessToken() (string, error)
}

type AuthBiz struct {
	authRepo AuthRepo
	userRepo UserRepo
	log      *log.Helper
}

func NewAuthBiz(
	authRepo AuthRepo,
	userRepo UserRepo,
	logger log.Logger,
) *AuthBiz {

	authBiz := &AuthBiz{
		authRepo: authRepo,
		userRepo: userRepo,
		log:      log.NewHelper(log.With(logger, "module", "biz/auth")),
	}

	ticker := time.NewTicker(1 * time.Second) // 设置每隔一秒执行一次

	// 创建一个 goroutine 来处理定时任务
	go func() {
		for {
			select {
			case <-ticker.C: // 从 ticker 接收数据
				authBiz.getAccessToken() // 执行任务
			}
		}
	}()

	return authBiz
}

func (a *AuthBiz) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginReply, error) {
	//if request.Type == "1" { // 1: 微信登录 2: 账号密码登录
	//	unionID := getUnionID(request.Code)
	//	fmt.Println(unionID)
	//}
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
	token, err := vjwt.GenerateJwtToken(userInfo.ID, userInfo.Username)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("系统内部错误，请联系管理员, %v", err))
	}
	return &v1.LoginReply{
		Token: token,
	}, nil
}

func getUnionID(code string) string {
	baseURL := "https://api.weixin.qq.com/sns/jscode2session"

	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	params := url.Values{}
	params.Add("appid", "wx50625797ea3c80eb")
	params.Add("secret", "8fd6bd2b6fe4864acaaca37f2061abd3")
	params.Add("js_code", code)
	params.Add("grant_type", "authorization_code")

	u.RawQuery = params.Encode()

	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println(result)
		return result["unionid"].(string)
	}

	return ""
}

// 获取access_token
func (a *AuthBiz) getAccessToken() {
	accessToken, err := a.authRepo.GetWXAccessToken()
	if err == nil && accessToken != "" {
		return
	}
	baseURL := "https://api.weixin.qq.com/cgi-bin/token"

	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
	}

	params := url.Values{}
	params.Add("appid", "wx50625797ea3c80eb")
	params.Add("secret", "8fd6bd2b6fe4864acaaca37f2061abd3")
	params.Add("grant_type", "client_credential")

	u.RawQuery = params.Encode()

	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		accessToken := result["access_token"].(string)
		err = a.authRepo.SetWXAccessToken(accessToken)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}
