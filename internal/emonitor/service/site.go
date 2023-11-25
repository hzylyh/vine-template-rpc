/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 11:34
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 11:34
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/emonitor/v1"
	"vine-template-rpc/internal/page"
)

func (em *EMonitorService) AddSite(ctx context.Context, request *v1.AddSiteRequest) (*v1.AddSiteReply, error) {
	return em.sb.AddSite(ctx, request)
}

func (em *EMonitorService) UpdateSite(ctx context.Context, request *v1.UpdateSiteRequest) (*v1.UpdateSiteReply, error) {
	return em.sb.UpdateSite(ctx, request)
}

func (em *EMonitorService) DeleteSite(ctx context.Context, request *v1.DeleteSiteRequest) (*v1.DeleteSiteReply, error) {
	return em.sb.DeleteSite(ctx, request)
}

func (em *EMonitorService) GetSite(ctx context.Context, request *v1.GetSiteRequest) (*v1.GetSiteReply, error) {
	return em.sb.GetSite(ctx, request)
}

func (em *EMonitorService) ListSite(ctx context.Context, request *v1.ListSiteRequest) (*page.Page, error) { //nolint:dupl
	return em.sb.ListSite(ctx, request)
}
