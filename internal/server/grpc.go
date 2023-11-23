package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	alarmV1 "vine-template-rpc/api/alarm/v1"
	emonitorV1 "vine-template-rpc/api/emonitor/v1"
	systemV1 "vine-template-rpc/api/system/v1"
	alarmService "vine-template-rpc/internal/alarm/service"
	"vine-template-rpc/internal/conf"
	emonitorService "vine-template-rpc/internal/emonitor/service"
	"vine-template-rpc/internal/middleware"
	systemService "vine-template-rpc/internal/system/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	system *systemService.SystemService,
	emonitor *emonitorService.EMonitorService,
	alarm *alarmService.AlarmService,
	logger log.Logger,
	enforcer *casbin.Enforcer,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			middleware.AuthMiddleware(enforcer),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	systemV1.RegisterSystemServer(srv, system)
	emonitorV1.RegisterEmonitorServer(srv, emonitor)
	alarmV1.RegisterAlarmServer(srv, alarm)
	return srv
}
