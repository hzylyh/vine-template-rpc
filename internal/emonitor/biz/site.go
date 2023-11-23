/**
 * @Description: 站点管理
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 11:17
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 11:17
 */

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	emonitorV1 "vine-template-rpc/api/emonitor/v1"
)

type SiteRepo interface {
	Add(ctx context.Context, site *Site) (*Site, error)
	Update(ctx context.Context, site *Site) error
	Delete(ctx context.Context, site *Site) error
	Get(ctx context.Context, site *Site) (*Site, error)
	List(ctx context.Context, site *Site) ([]*Site, error)
}

type Site struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type SiteBiz struct {
	repo SiteRepo
	log  *log.Helper
}

func (e *SiteBiz) AddSite(ctx context.Context, request *emonitorV1.AddSiteRequest) (*emonitorV1.AddSiteReply, error) {
	add, err := e.repo.Add(ctx, &Site{
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}
	return &emonitorV1.AddSiteReply{
		Id: add.Id.String(),
	}, nil
}

func (e *SiteBiz) UpdateSite(ctx context.Context, request *emonitorV1.UpdateSiteRequest) (*emonitorV1.UpdateSiteReply, error) {
	return nil, nil
}

func (e *SiteBiz) DeleteSite(ctx context.Context, request *emonitorV1.DeleteSiteRequest) (*emonitorV1.DeleteSiteReply, error) {
	return nil, nil
}

func (e *SiteBiz) GetSite(ctx context.Context, request *emonitorV1.GetSiteRequest) (*emonitorV1.GetSiteReply, error) {
	return nil, nil
}

func (e *SiteBiz) ListSite(ctx context.Context, request *emonitorV1.ListSiteRequest) (*emonitorV1.ListSiteReply, error) {
	return nil, nil
}

// NewSiteBiz .
func NewSiteBiz(repo SiteRepo, logger log.Logger) *SiteBiz {
	return &SiteBiz{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/site")),
	}
}
