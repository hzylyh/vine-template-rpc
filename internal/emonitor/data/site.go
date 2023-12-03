/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 11:19
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 11:19
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vine-template-rpc/internal/emonitor/biz"
	"vine-template-rpc/internal/emonitor/data/schema"
)

type siteRepo struct {
	data *Data
	log  *log.Helper
}

func (s siteRepo) Add(ctx context.Context, site *schema.Site) error {
	return s.data.gdb.Create(site).Error
}

func (s siteRepo) Update(ctx context.Context, site *biz.Site) error {
	//TODO implement me
	panic("implement me")
}

func (s siteRepo) Delete(ctx context.Context, site *biz.Site) error {
	//TODO implement me
	panic("implement me")
}

func (s siteRepo) Get(ctx context.Context, site *biz.Site) (*biz.Site, error) {
	//TODO implement me
	panic("implement me")
}

func (s siteRepo) List(ctx context.Context, site *schema.Site) ([]*schema.Site, error) {
	sites := make([]*schema.Site, 0)
	err := s.data.gdb.Where(site).Find(&sites).Error
	return sites, err
}

func NewSiteRepo(data *Data, logger log.Logger) biz.SiteRepo {
	return &siteRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/site")),
	}
}
