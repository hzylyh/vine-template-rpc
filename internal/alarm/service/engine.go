/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/21 09:31
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/21 09:31
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/alarm/v1"
)

func (as *AlarmService) Trigger(ctx context.Context, request *v1.TriggerRequest) (*v1.TriggerReply, error) {
	as.eb.Trigger("", "")
	return nil, nil
}
