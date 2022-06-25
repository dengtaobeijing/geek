//go:build wireinject
// +build wireinject

package main

import (
	"network/app/sbi/interface/internal/biz"
	"network/app/sbi/interface/internal/conf"
	"network/app/sbi/interface/internal/data"
	"network/app/sbi/interface/internal/server"
	"network/app/sbi/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Grpc, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
