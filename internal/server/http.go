package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	alarmV1 "vine-template-rpc/api/alarm/v1"
	emonitorV1 "vine-template-rpc/api/emonitor/v1"
	orderV1 "vine-template-rpc/api/order/v1"
	systemV1 "vine-template-rpc/api/system/v1"
	alarmService "vine-template-rpc/internal/alarm/service"
	"vine-template-rpc/internal/conf"
	emonitorService "vine-template-rpc/internal/emonitor/service"
	"vine-template-rpc/internal/middleware"
	orderService "vine-template-rpc/internal/order/service"
	systemService "vine-template-rpc/internal/system/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	system *systemService.SystemService,
	emonitor *emonitorService.EMonitorService,
	alarm *alarmService.AlarmService,
	order *orderService.OrderService,
	logger log.Logger,
	enforcer *casbin.Enforcer,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.AuthMiddleware(enforcer),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	systemV1.RegisterSystemHTTPServer(srv, system)
	emonitorV1.RegisterEmonitorHTTPServer(srv, emonitor)
	alarmV1.RegisterAlarmHTTPServer(srv, alarm)
	orderV1.RegisterOrderHTTPServer(srv, order)
	return srv
}
