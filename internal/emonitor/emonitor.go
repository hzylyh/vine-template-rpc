package emonitor

import (
	"github.com/google/wire"
	"vine-template-rpc/internal/emonitor/biz"
	"vine-template-rpc/internal/emonitor/data"
	"vine-template-rpc/internal/emonitor/service"
)

// EMonitorProviderSet is service providers.
var EMonitorProviderSet = wire.NewSet(
	// service
	service.NewEMonitorService,

	// biz
	biz.NewSiteBiz,

	// data common
	data.NewData,
	data.NewEntClient,

	//data biz
	data.NewSiteRepo,
)