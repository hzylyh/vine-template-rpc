/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 10:07
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 10:07
 */

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	orderV1 "vine-template-rpc/api/order/v1"
	"vine-template-rpc/internal/order/data/schema"
	"vine-template-rpc/internal/page"
	"vine-template-rpc/pkg/claim"
	"vine-template-rpc/pkg/pagehelper"
)

type OrderRepo interface {
	Add(ctx context.Context, site *schema.Order) error
	Update(ctx context.Context, site *schema.Order) error
	Delete(ctx context.Context, site *schema.Order) error
	Get(ctx context.Context, site *schema.Order) (*schema.Order, error)
	List(ctx context.Context, site *schema.Order) ([]*Order, error)
}

type OrderHistoryRepo interface {
	Add(ctx context.Context, history *schema.OrderHistory) error
	//Update(ctx context.Context, site *Site) error
	//Delete(ctx context.Context, site *Site) error
	//Get(ctx context.Context, site *Site) (*Site, error)
	//List(ctx context.Context, site *schema.Site) ([]*schema.Site, error)
}

type Order struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Priority  int32  `json:"priority,omitempty"`
	Lon       string `json:"lon,omitempty"`
	Lat       string `json:"lat,omitempty"`
	Status    string `json:"status,omitempty"`
	SiteName  string `json:"siteName,omitempty"`
	Describe  string `json:"describe,omitempty"`
}

type OrderBiz struct {
	repo OrderRepo
	log  *log.Helper
}

// AddOrder
func (ob *OrderBiz) AddOrder(ctx context.Context, request *orderV1.AddOrderRequest) (*orderV1.AddOrderReply, error) {
	claim := claim.GetUserInfoFromCtx(ctx)
	orderInfo := &schema.Order{
		Name:   request.Name,
		UserID: claim.UserId,
	}
	err := ob.repo.Add(ctx, orderInfo)
	if err != nil {
		return nil, err
	}
	return &orderV1.AddOrderReply{}, nil
}

func (ob *OrderBiz) UpdateOrder(ctx context.Context, request *orderV1.UpdateOrderRequest) (*orderV1.UpdateOrderReply, error) {
	return nil, nil
}

func (ob *OrderBiz) DeleteOrder(ctx context.Context, request *orderV1.DeleteOrderRequest) (*orderV1.DeleteOrderReply, error) {
	return nil, nil
}

func (ob *OrderBiz) GetOrder(ctx context.Context, request *orderV1.GetOrderRequest) (*orderV1.GetOrderReply, error) {
	return nil, nil
}

func (ob *OrderBiz) ListOrder(ctx context.Context, request *orderV1.ListOrderRequest) (*page.Page, error) {
	claim := claim.GetUserInfoFromCtx(ctx)
	orderList, err := ob.repo.List(ctx, &schema.Order{
		UserID: claim.UserId,
		Status: request.Status,
	})
	if err != nil {
		return nil, err
	}
	pageClient := pagehelper.NewMemPage(orderList)
	resWithPage := pageClient.Paginator(request.PageNum, request.PageSize)
	return resWithPage, nil
}

func (ob *OrderBiz) ReviewOrder(ctx context.Context, request *orderV1.ReviewOrderRequest) (*orderV1.ReviewOrderReply, error) {
	err := ob.repo.Update(ctx, &schema.Order{
		Model: gorm.Model{
			ID: uint(request.Id),
		},
		Status: request.Status,
	})
	if err != nil {
		return nil, err
	}
	return &orderV1.ReviewOrderReply{}, nil
}

func NewOrderBiz(
	repo OrderRepo,
	logger log.Logger,
) *OrderBiz {
	return &OrderBiz{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/rule")),
	}
}
