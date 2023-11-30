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
	"vine-template-rpc/internal/emonitor/data/ent/site"
)

type siteRepo struct {
	data *Data
	log  *log.Helper
}

func (s siteRepo) Add(ctx context.Context, site *biz.Site) (*biz.Site, error) {
	save, err := s.data.db.Site.
		Create().
		SetName(site.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Site{
		Id: save.ID,
	}, nil
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

func (s siteRepo) List(ctx context.Context, siteInfo *biz.Site) ([]*biz.Site, error) {
	var siteList []*biz.Site
	err := s.data.db.Site.
		Query().
		Where().
		Select(site.FieldName).
		Scan(ctx, &siteList)
	return siteList, err
}

func NewSiteRepo(data *Data, logger log.Logger) biz.SiteRepo {
	return &siteRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/site")),
	}
}
