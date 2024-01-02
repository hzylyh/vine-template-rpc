/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 10:02
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 10:02
 */

package service

import (
	v1 "vine-template-rpc/api/order/v1"
	"vine-template-rpc/internal/order/biz"
)

type OrderService struct {
	v1.UnimplementedOrderServer
	ob *biz.OrderBiz
}

func NewOrderService(
	ob *biz.OrderBiz,
) *OrderService {
	return &OrderService{
		ob: ob,
	}
}
