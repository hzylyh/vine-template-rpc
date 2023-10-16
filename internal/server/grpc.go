package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	systemV1 "vine-template-rpc/api/system/v1"
	"vine-template-rpc/internal/conf"
	systemService "vine-template-rpc/internal/system/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	system *systemService.SystemService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
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
	return srv
}
