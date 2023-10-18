//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"vine-template-rpc/internal/conf"
	"vine-template-rpc/internal/server"
	systemService "vine-template-rpc/internal/system"

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
