//go:build wireinject
// +build wireinject

package main

import (
	"network/app/sbi/service/internal/biz"
	"network/app/sbi/service/internal/conf"
	"network/app/sbi/service/internal/data"
	"network/app/sbi/service/internal/server"
	"network/app/sbi/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Data, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
