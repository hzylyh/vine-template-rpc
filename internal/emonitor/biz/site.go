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
	"strconv"
	emonitorV1 "vine-template-rpc/api/emonitor/v1"
	"vine-template-rpc/internal/emonitor/data/schema"
	"vine-template-rpc/internal/page"
	"vine-template-rpc/pkg/pagehelper"
)

type SiteRepo interface {
	Add(ctx context.Context, site *schema.Site) error
	Update(ctx context.Context, site *Site) error
	Delete(ctx context.Context, site *Site) error
	Get(ctx context.Context, site *Site) (*Site, error)
	List(ctx context.Context, site *schema.Site) ([]*schema.Site, error)
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
	siteInfo := &schema.Site{
		Name:            request.Name,
		Type:            request.Type,
		Dept:            request.Dept,
		Owner:           request.Owner,
		Remark:          request.Remark,
		Lon:             request.Lon,
		Lat:             request.Lat,
		Status:          request.Status,
		Address:         request.Address,
		LastServiceTime: request.LastServiceTime,
		Creator:         request.Creator,
	}
	err := e.repo.Add(ctx, siteInfo)
	if err != nil {
		return nil, err
	}
	return &emonitorV1.AddSiteReply{
		Id: strconv.Itoa(int(siteInfo.ID)),
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

func (e *SiteBiz) ListSite(ctx context.Context, request *emonitorV1.ListSiteRequest) (*page.Page, error) {
	list, err := e.repo.List(ctx, &schema.Site{})
	if err != nil {
		return nil, err
	}
	// TODO: 改为数据库分页
	pageClient := pagehelper.NewMemPage(list)
	resWithPage := pageClient.Paginator(request.PageNum, request.PageSize)
	return resWithPage, nil
}

// NewSiteBiz .
func NewSiteBiz(repo SiteRepo, logger log.Logger) *SiteBiz {
	return &SiteBiz{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/site")),
	}
}
