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
	v1 "vine-template-rpc/api/alarm/v1"
	"vine-template-rpc/internal/alarm/biz"
)

type AlarmService struct {
	v1.UnimplementedAlarmServer
	eb *biz.EngineBiz
	rb *biz.RuleBiz
}

func NewAlarmService(
	eb *biz.EngineBiz,
	rb *biz.RuleBiz,
) *AlarmService {
	return &AlarmService{
		eb: eb,
		rb: rb,
	}
}
