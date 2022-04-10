package server

import (
	"geek/week4/api/test"
	"geek/week4/internal/conf"
	"geek/week4/internal/service"
	"net/http"
	"strconv"
)

type HttpServer0 http.Server

func NewHttpServer0(config *conf.ServerConfig, user *service.UserService) *HttpServer0 {
	mux := http.NewServeMux()
	mux.HandleFunc("/", test.NewHelloHandler(user))
	//
	return &HttpServer0{
		Addr:    ":" + strconv.Itoa(config.Server0Port),
		Handler: mux,
	}
}
