// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"network/app/sbi/task/internal/biz"
	"network/app/sbi/task/internal/conf"
	"network/app/sbi/task/internal/data"
	"network/app/sbi/task/internal/server"
	"network/app/sbi/task/internal/service"
)

// Injectors from wire.go:

func initApp(confData *conf.Data, confServer *conf.Server, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	taskRepo := data.NewTaskRepo(dataData, logger)
	taskUseCase := biz.NewTaskUseCase(taskRepo, logger)
	taskService := service.NewTaskService(taskUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, taskService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
