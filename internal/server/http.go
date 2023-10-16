package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	systemV1 "vine-cms-rpc/api/system/v1"
	"vine-cms-rpc/internal/conf"
	systemService "vine-cms-rpc/internal/system/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	system *systemService.SystemService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	return srv
}
