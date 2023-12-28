package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"vine-template-rpc/internal/system/biz"
)

const WX_ACCESS_TOKEN = "wx_access_token"

type wechatRepo struct {
	data *Data
	log  *log.Helper
}

func (r *wechatRepo) SetWXAccessToken(token string) error {
	r.log.Info("set wx access token", "token", token)
	return r.data.rdb.Set(context.Background(), WX_ACCESS_TOKEN, token, 2*time.Hour).Err()
}

func (r *wechatRepo) GetWXAccessToken() (string, error) {
	cmd := r.data.rdb.Get(context.Background(), WX_ACCESS_TOKEN)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.Val(), nil
}

func NewWechatRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &wechatRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/role")),
	}
}
