/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 10:40
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 10:40
 */

package service

import (
	v1 "vine-template-rpc/api/emonitor/v1"
	"vine-template-rpc/internal/emonitor/biz"
)

// EMonitorService is a monitor service.
type EMonitorService struct {
	v1.UnimplementedEmonitorServer
	sb *biz.SiteBiz
}

// NewEMonitorService new a monitor service.
func NewEMonitorService(
	sb *biz.SiteBiz,
) *EMonitorService {
	return &EMonitorService{
		sb: sb,
	}
}
