/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 10:56
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 10:56
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vine-template-rpc/internal/order/biz"
	"vine-template-rpc/internal/order/data/schema"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func (s orderRepo) Add(ctx context.Context, order *schema.Order) error {
	return s.data.gdb.Create(order).Error
}

func (s orderRepo) Update(ctx context.Context, order *schema.Order) error {
	return s.data.gdb.Save(&order).Error
}

func (s orderRepo) Delete(ctx context.Context, order *schema.Order) error {
	panic("还没实现我哦")
}

func (s orderRepo) Get(ctx context.Context, order *schema.Order) (*schema.Order, error) {
	panic("还没实现我哦")
}

func (s orderRepo) List(ctx context.Context, order *schema.Order) ([]*schema.Order, error) {
	panic("还没实现我哦")
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/order")),
	}
}