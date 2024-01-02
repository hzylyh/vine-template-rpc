/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 10:39
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 10:39
 */

package service

import (
	"context"
	v1 "vine-template-rpc/api/order/v1"
)

func (os *OrderService) AddOrder(ctx context.Context, request *v1.AddOrderRequest) (*v1.AddOrderReply, error) {

	return os.ob.AddOrder(ctx, request)
}
