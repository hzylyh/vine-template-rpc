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
	"vine-template-rpc/internal/page"
)

func (os *OrderService) AddOrder(ctx context.Context, request *v1.AddOrderRequest) (*v1.AddOrderReply, error) {
	return os.ob.AddOrder(ctx, request)
}

func (os *OrderService) UpdateOrder(ctx context.Context, request *v1.UpdateOrderRequest) (*v1.UpdateOrderReply, error) {
	return os.ob.UpdateOrder(ctx, request)
}

func (os *OrderService) DeleteOrder(ctx context.Context, request *v1.DeleteOrderRequest) (*v1.DeleteOrderReply, error) {
	return os.ob.DeleteOrder(ctx, request)
}

func (os *OrderService) GetOrder(ctx context.Context, request *v1.GetOrderRequest) (*v1.GetOrderReply, error) {
	return os.ob.GetOrder(ctx, request)
}

func (os *OrderService) ListOrder(ctx context.Context, request *v1.ListOrderRequest) (*page.Page, error) {
	return os.ob.ListOrder(ctx, request)
}

func (os *OrderService) ReviewOrder(ctx context.Context, request *v1.ReviewOrderRequest) (*v1.ReviewOrderReply, error) {
	return os.ob.ReviewOrder(ctx, request)
}
