//go:build wireinject
// +build wireinject

package main

import (
	"network/app/sbi/task/internal/biz"
	"network/app/sbi/task/internal/conf"
	"network/app/sbi/task/internal/data"
	"network/app/sbi/task/internal/server"
	"network/app/sbi/task/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Data, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
