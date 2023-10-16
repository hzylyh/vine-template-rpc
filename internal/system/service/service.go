/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/16 14:08
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/16 14:08
 */

package service

import (
	v1 "vine-cms-rpc/api/system/v1"
	"vine-cms-rpc/internal/system/biz"
)

// SystemService is a greeter service.
type SystemService struct {
	v1.UnimplementedSystemServer
	ub *biz.UserBiz
}

// NewSystemService new a greeter service.
func NewSystemService(ub *biz.UserBiz) *SystemService {
	return &SystemService{
		ub: ub,
	}
}
