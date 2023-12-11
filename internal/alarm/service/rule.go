/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 15:50
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 15:50
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/alarm/v1"
)

func (as *AlarmService) AddAlarmRule(ctx context.Context, request *v1.AddAlarmRuleRequest) (*v1.AddAlarmRuleReply, error) {
	as.rb.Add(request.Info)
	return &v1.AddAlarmRuleReply{
		Message: "部署成功",
	}, nil
}
