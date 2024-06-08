//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/toheart/go-structure/internal/biz"
	"github.com/toheart/go-structure/internal/conf"
	"github.com/toheart/go-structure/internal/data"
	"github.com/toheart/go-structure/internal/server"
	"github.com/toheart/go-structure/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		service.ProviderServices,
		server.ProviderServer,
		data.ProviderData,
		biz.ProviderBiz,
		server.NewServerList, newApp))
}
