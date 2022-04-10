package server

import (
	"geek/week4/api/test"
	"geek/week4/internal/conf"
	"geek/week4/internal/service"
	"net/http"
	"strconv"
)

type HttpServer1 http.Server

func NewHttpServer1(config *conf.ServerConfig, user *service.UserService) *HttpServer1 {
	mux := http.NewServeMux()
	mux.HandleFunc("/", test.NewHelloHandler(user))
	//
	return &HttpServer1{
		Addr:    ":" + strconv.Itoa(config.Server1Port),
		Handler: mux,
	}
}
