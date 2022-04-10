//go:build wireinject
// +build wireinject

package main

import (
	"geek/week4/internal/biz"
	"geek/week4/internal/common"
	"geek/week4/internal/conf"
	"geek/week4/internal/data"
	"geek/week4/internal/server"
	"geek/week4/internal/service"
	"github.com/google/wire"
	"log"
)

func initApp(config *conf.Config, logger *log.Logger) *common.App {
	wire.Build(server.WireSet, data.WireSet, biz.WireSet, service.WireSet, newApp)
	return &common.App{}
}
