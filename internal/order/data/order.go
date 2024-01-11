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

func (s orderRepo) Get(ctx context.Context, query *schema.Order) (order *schema.Order, err error) {
	err = s.data.gdb.Where(query).First(&order).Error
	return
}

func (s orderRepo) List(ctx context.Context, order *schema.Order) (orders []*biz.Order, err error) {
	db := s.data.gdb.Table("tb_order").
		Select("tb_order.id, tb_order.name, tb_order.type, date_format(tb_order.created_at, '%Y-%m-%d %H:%i:%s') as created_at," +
			" tb_user.username, tb_order.user_id, tb_order.equipment_id, tb_order.priority, tb_order.describe, tb_order.status, tb_equipment.lon, tb_equipment.lat").
		Joins("left join tb_equipment on tb_order.equipment_id = tb_equipment.id").
		Joins("left join tb_user on tb_order.user_id = tb_user.id")
	if order.Status == "" {
		db = db.Where(order).Not("tb_order.status", "0")
	} else {
		db = db.Where(order)
	}
	err = db.Scan(&orders).Error
	return
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/order")),
	}
}
