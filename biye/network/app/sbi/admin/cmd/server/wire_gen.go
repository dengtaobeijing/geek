// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"network/app/sbi/admin/internal/biz"
	"network/app/sbi/admin/internal/conf"
	"network/app/sbi/admin/internal/data"
	"network/app/sbi/admin/internal/server"
	"network/app/sbi/admin/internal/service"
)

// Injectors from wire.go:

func initApp(confData *conf.Data, confServer *conf.Server, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	sbiRepo := data.NewSbiRepo(dataData, logger)
	sbiUseCase := biz.NewSbiUseCase(sbiRepo, logger)
	ordersRepo := data.NewOrdersRepo(dataData, logger)
	ordersUseCase := biz.NewOrdersUseCase(ordersRepo, logger)
	adminService := service.NewAdminService(sbiUseCase, ordersUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, adminService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
