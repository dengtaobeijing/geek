// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"network/app/sbi/service/internal/biz"
	"network/app/sbi/service/internal/conf"
	"network/app/sbi/service/internal/data"
	"network/app/sbi/service/internal/server"
	"network/app/sbi/service/internal/service"
)

// Injectors from wire.go:

func initApp(confData *conf.Data, confServer *conf.Server, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	sbiRepo := data.NewSbiRepo(dataData, logger)
	sbiUseCase := biz.NewSbiUseCase(sbiRepo, logger)
	sbiService := service.NewSbiService(sbiUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, sbiService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
