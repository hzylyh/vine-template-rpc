/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/16 14:09
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/16 14:09
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/system/v1"
)

func (s *SystemService) CreateSystem(ctx context.Context, request *v1.CreateSystemRequest) (*v1.CreateSystemReply, error) {
	return &v1.CreateSystemReply{Name: "d"}, nil
}
