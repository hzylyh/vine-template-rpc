//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"vine-cms-rpc/internal/conf"
	"vine-cms-rpc/internal/server"
	systemService "vine-cms-rpc/internal/system"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		systemService.SystemProviderSet,
		newApp,
	),
	)
}
