package service

import (
	"github.com/google/wire"
	"vine-template-rpc/internal/order/biz"
	"vine-template-rpc/internal/order/data"
	"vine-template-rpc/internal/order/service"
)

var OrderProviderSet = wire.NewSet(
	// service
	service.NewOrderService,

	//biz
	biz.NewOrderBiz,

	// data common
	data.NewData,

	//data biz
	data.NewOrderRepo,
)
